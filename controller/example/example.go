package example_controller

import (
	example_service "foo/service/example"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExampleController struct{}

func New() *ExampleController {
	return &ExampleController{}
}

type Params struct {
	Key   string `form:"key" binding:"required"`
	Value string `form:"value"`
}

func (ec *ExampleController) Put(ctx *gin.Context) {
	var params Params
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	// service
	service := example_service.New()
	if err := service.Put(params.Key, params.Value); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": nil, "msg": "SUCCESS::调用成功"})
}

func (ec *ExampleController) Get(ctx *gin.Context) {
	var params Params
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	// service
	service := example_service.New()
	val, err := service.Get(params.Key)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": map[string]string{"key": params.Key, "value": val}, "msg": "SUCCESS::调用成功"})
}
