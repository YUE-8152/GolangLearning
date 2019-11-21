package user

import (
	"GolangLearning/model/domain/po"
	"GolangLearning/model/domain/user"
	commutils "GolangLearning/model/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type User struct {
}

func (this *User) Create(ctx *gin.Context) {
	var u po.User
	u.Name = ctx.Request.PostFormValue("name")
	u.Age, _ = strconv.Atoi(ctx.Request.PostFormValue("age"))
	u.Password = ctx.Request.PostFormValue("password")
	if err := user.Create(u); err != nil {
		commutils.FailedResponse(ctx, nil, err.Error())
		return
	}
	commutils.SuccessResponse(ctx, "")
}

func (this *User) QueryById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	user, err := user.QueryById(id)
	if err != nil {
		commutils.FailedResponse(ctx, nil, err.Error())
		return
	}
	commutils.SuccessResponse(ctx, user)
}

func (this *User) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Request.PostFormValue("id"))
	age, _ := strconv.Atoi(ctx.Request.PostFormValue("age"))
	u := po.User{
		Id:       id,
		Name:     ctx.Request.PostFormValue("name"),
		Age:      age,
		Password: ctx.Request.PostFormValue("password"),
	}
	err := user.Update(u)
	if err != nil {
		commutils.FailedResponse(ctx, nil, err.Error())
		return
	}
	commutils.SuccessResponse(ctx, "")
}

func (this *User) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	err := user.Delete(id)
	if err != nil {
		commutils.FailedResponse(ctx, nil, err.Error())
		return
	}
	commutils.SuccessResponse(ctx, "")
}

func (this *User) QueryUser(ctx *gin.Context) {
	res, err := user.QueryUser()
	if err != nil {
		commutils.FailedResponse(ctx, nil, err.Error())
		return
	}
	commutils.SuccessResponse(ctx, res)
}
