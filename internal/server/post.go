package server

import (
	"net/http"
	"rgb/internal/store"
	"time"

	"github.com/gin-gonic/gin"
)

func createPost(ctx *gin.Context) {
	post := new(store.Post)
	if err := ctx.Bind(post); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := store.AddPost(user, post); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Post created successfully.",
		"data": post,
	})
}

func indexPosts(ctx *gin.Context) {
	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := store.FetchUserPosts(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Posts fetched successfully.",
		"data": user.Posts,
	})
}

func updatePost(ctx *gin.Context) {
	jsonPost := new(store.Post)
	if err := ctx.Bind(jsonPost); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := currentUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	dbPost, err := store.FetchPost(jsonPost.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user.ID != dbPost.UserID {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Not authorized."})
		return
	}
	jsonPost.ModifiedAt = time.Now()
	if err := store.UpdatePost(jsonPost); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Post updated successfully.",
		"data": jsonPost,
	})
}
