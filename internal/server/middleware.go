package server

import (
	"errors"
	"net/http"
	"rgb/internal/store"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Authorization middleware extracts token from Authorization header.
// It first checks if header exists, if it’s in valid format, and then calls verifyJWT() function.
// If JWT verification passes, user’s ID is returned.
// User with that ID is fetched from database and set as current user for this context.
func authorization(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing."})
		return
	}
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format is not valid."})
		return
	}
	if headerParts[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing bearer part."})
		return
	}
	userID, err := verifyJWT(headerParts[1])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	user, err := store.FetchUser(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.Set("user", user)
	ctx.Next()
}

// First we check if user is set for this context. If not, error is returned.
// Since ctx.Get() returns interface, we must check if value is of type *store.User. If not, error is returned.
// When both checks are passed, current user is returned from context.
func currentUser(ctx *gin.Context) (*store.User, error) {
	var err error
	_user, exists := ctx.Get("user")
	if !exists {
		err = errors.New("current context user not set")
		log.Error().Err(err).Msg("")
		return nil, err
	}
	user, ok := _user.(*store.User)
	if !ok {
		err = errors.New("context user is not valid type")
		log.Error().Err(err).Msg("")
		return nil, err
	}
	return user, nil
}
