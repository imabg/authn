package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imabg/authn/store"
	"github.com/imabg/authn/types"
	"github.com/imabg/authn/utils"
	"time"
)

type LoginHandler struct {
	store       store.ILoginStore
	tokenMaster utils.Maker
	config      utils.Config
}

func NewLoginHandler(lStore store.ILoginStore, token utils.Maker, config utils.Config) *LoginHandler {
	return &LoginHandler{
		store:       lStore,
		tokenMaster: token,
		config:      config,
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
	hashPwd, err := l.store.GetUserCred(user.ID)
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
	duration, err := time.ParseDuration(l.config.AccessTokenDuration)
	token, tokenPayload, err := l.tokenMaster.CreateToken(duration, user.ID, user.SourceID)
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
