package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/imabg/authn/store"
	"github.com/imabg/authn/types"
	"github.com/imabg/authn/utils"
	"strings"
	"unicode"
)

type UserHandler struct {
	store       store.UserStoreInterface
	sourceStore store.SourceStoreInterface
}

func NewUserHandler(uStore store.UserStoreInterface, sStore store.SourceStoreInterface) *UserHandler {
	return &UserHandler{
		store:       uStore,
		sourceStore: sStore,
	}
}

func (u *UserHandler) CreateViaEmail(c *gin.Context) {
	body := types.UserEmailDTO{}
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.Send400Response(c, "Bad request", err.Error())
		return
	}
	sourceConfig, err := u.sourceStore.GetConfig(body.SourceID)
	if err != nil {
		utils.Send400Response(c, "Invalid source", err.Error())
		return
	}
	isPwdValid := checkPassword(body.Password, sourceConfig)
	if !isPwdValid {
		utils.Send400Response(c, "Bad request", "Password does not match source policy")
		return
	}
	var user types.User
	user.ID = utils.GenerateUUID()
	user.Email = body.Email
	user.SourceID = body.SourceID
	//TODO: encrypt the password and create a new entry in "credentials" and link with user_id
	id, err := u.store.Create(&user)
	if err != nil {
		utils.Send500Response(c, "Internal server error", err.Error())
		return
	}
	utils.Send201Response(c, "User created successfully", id)
}

func checkPassword(password string, config *types.Config) bool {
	if len(password) < config.PasswordLength {
		return false
	}

	var hasLower, hasUpper, hasDigit, hasSpecial bool

	for _, char := range password {
		switch {
		case unicode.IsLower(char) && config.PasswordLowerAllowed:
			hasLower = true
		case unicode.IsUpper(char) && config.PasswordUpperAllowed:
			hasUpper = true
		case unicode.IsDigit(char) && config.PasswordNumericAllowed:
			hasDigit = true
		case strings.ContainsRune("!@#\\$%\\^&\\*", char) && config.PasswordSpecialAllowed:
			hasSpecial = true
		}
	}

	return hasLower && hasUpper && hasDigit && hasSpecial
}
