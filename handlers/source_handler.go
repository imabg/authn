package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/imabg/authn/store"
	"github.com/imabg/authn/types"
	"github.com/imabg/authn/utils"
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
		utils.Send400Response(c, "Bad request", err.Error())
		return
	}
	var source types.Source
	source.Name = body.Name
	source.Description = body.Description
	source.ID = utils.GenerateUUID()
	id, err := s.store.Create(&source)
	if err != nil {
		utils.Send500Response(c, "Internal server error", err.Error())
		return
	}
	utils.Send201Response(c, "Source created successfully", id)
}

func (s *SourceHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	source, err := s.store.GetByID(id)
	if err != nil {
		utils.Send500Response(c, "Internal server error", err.Error())
		return
	}
	utils.Send200Response(c, "Source fetched successfully", &source)
}
