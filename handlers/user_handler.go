package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/imabg/authn/store"
	"github.com/imabg/authn/types"
	"github.com/imabg/authn/utils"
	"net/http"
)

type UserHandler struct {
	store store.UserStoreInterface
}

func NewUserHandler(uStore store.UserStoreInterface) *UserHandler {
	return &UserHandler{
		store: uStore,
	}
}

func (u *UserHandler) Create(c *gin.Context) {
	body := types.UserDTO{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})
		return
	}
	var user types.User
	user.ID = utils.GenerateUUID()
	user.Email = body.Email
	user.Phone = body.Phone
	user.SourceID = body.SourceID
	user.CountryCode = body.CountryCode
	dId, err := utils.GenerateDisplayID()
	user.DisplayID = dId
	id, err := u.store.Create(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"id":      id,
	})
}
