package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gitlab.eoitek.net/EOI/ckman/common"
	"gitlab.eoitek.net/EOI/ckman/config"
	"gitlab.eoitek.net/EOI/ckman/model"
	"io/ioutil"
	"path"
	"time"
)

type UserController struct {
	config *config.CKManConfig
}

func NewUserController(config *config.CKManConfig) *UserController {
	ck := &UserController{}
	ck.config = config
	return ck
}

// @Summary 登陆
// @Description 登陆
// @version 1.0
// @Param req body model.LoginReq true "request body"
// @Failure 200 {string} json "{"code":400,"msg":"请求参数错误","data":""}"
// @Failure 200 {string} json "{"code":5030,"msg":"用户不存在","data":""}"
// @Failure 200 {string} json "{"code":5031,"msg":"获取用户密码失败","data":""}"
// @Failure 200 {string} json "{"code":5032,"msg":"用户密码验证失败","data":""}"
// @Success 200 {string} json "{"code":200,"msg":"ok","data":{"username":"ckman","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"}}"
// @Router /login [post]
func (d *UserController) Login(c *gin.Context) {
	var req model.LoginReq

	if err := model.DecodeRequestBody(c.Request, &req); err != nil {
		model.WrapMsg(c, model.INVALID_PARAMS, model.GetMsg(model.INVALID_PARAMS), err.Error())
		return
	}

	if req.Username != common.DefaultUserName {
		model.WrapMsg(c, model.USER_VERIFY_FAIL, model.GetMsg(model.USER_VERIFY_FAIL), nil)
		return
	}

	passwordFile := path.Join(common.GetWorkDirectory(), "conf/password")
	data, err := ioutil.ReadFile(passwordFile)
	if err != nil {
		model.WrapMsg(c, model.GET_USER_PASSWORD_FAIL, model.GetMsg(model.GET_USER_PASSWORD_FAIL), err.Error())
		return
	}

	if pass := common.ComparePassword(string(data), req.Password); !pass {
		model.WrapMsg(c, model.PASSWORD_VERIFY_FAIL, model.GetMsg(model.PASSWORD_VERIFY_FAIL), nil)
		return
	}

	j := common.NewJWT()
	claims := common.CustomClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Second * time.Duration(d.config.Server.SessionTimeout)).Unix(),
		},
		Name:     common.DefaultUserName,
		ClientIP: c.ClientIP(),
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		model.WrapMsg(c, model.CREAT_TOKEN_FAIL, model.GetMsg(model.CREAT_TOKEN_FAIL), err.Error())
		return
	}

	rsp := model.LoginRsp{
		Username: req.Username,
		Token:    token,
	}

	model.WrapMsg(c, model.SUCCESS, model.GetMsg(model.SUCCESS), rsp)
}
