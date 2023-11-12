package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imabg/authn/store"
	"github.com/imabg/authn/types"
	"github.com/imabg/authn/utils"
	"os"
)

type LoginHandler struct {
	store  store.ILoginStore
	cStore store.ICredentialStore
}

func NewLoginHandler(lStore store.ILoginStore, cStore store.ICredentialStore) *LoginHandler {
	return &LoginHandler{
		store:  lStore,
		cStore: cStore,
	}
}

func (l *LoginHandler) LoginViaEmail(c *gin.Context) {
	body := types.LoginViaEmailDTO{}
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.Send400Response(c, "Bad request", err.Error())
		return
	}
	user, err := l.store.CheckAccess(body.Email)
	if err != nil {
		utils.Send400Response(c, "Permission related issue", err.Error())
		return
	}
	hashPwd, err := l.cStore.GetUserCred(user.ID)
	pwdConfig := utils.NewArgon2ID()
	ok, err := pwdConfig.Verify(body.Password, hashPwd)
	if err != nil {
		utils.Send500Response(c, "Internal server error", err.Error())
		return
	}
	if !ok {
		utils.Send400Response(c, "Authentication failed", "Password does not match")
		return
	}
	tokenMaster, err := utils.NewPasetoMaker(os.Getenv("TOKEN_SECRET_KEY"))
	token, tokenPayload, err := tokenMaster.CreateToken(user.ID, user.SourceID)
	if err != nil {
		utils.Send500Response(c, "Internal server error", err.Error())
		return
	}
	//TODO: make platform dynamic
	loginLog := types.LoginLog{
		Ip:          c.ClientIP(),
		Platform:    types.WebPlatform,
		UserAgent:   c.GetHeader("User-Agent"),
		AccessToken: token,
		IsActive:    true,
		UserId:      user.ID,
	}
	if err := l.store.Log(&loginLog); err != nil {
		utils.Send500Response(c, "Internal server error", err.Error())
		return
	}
	fmt.Println(tokenPayload.ExpiredAt)
	utils.Send200Response(c, "Login success", token)
}
