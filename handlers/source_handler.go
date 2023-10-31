package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/imabg/authn/store"
	"github.com/imabg/authn/types"
	"net/http"
	"strconv"
)

type SourceHandler struct {
	store store.SourceStoreInterface
}

func NewSourceHandler(sStore store.SourceStoreInterface) *SourceHandler {
	return &SourceHandler{
		store: sStore,
	}
}

func (s *SourceHandler) Create(c *gin.Context) {
	body := types.SourceRTO{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})
		return
	}
	var source types.Source
	source.Name = body.Name
	source.Description = body.Description
	id, err := s.store.Create(&source)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Source created successfully",
		"id":      id,
	})
}

func (s *SourceHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})
		return
	}
	source, err := s.store.GetByID(i)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Source fetched successfully",
		"data":    &source,
	})
}
