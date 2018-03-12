package util

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"simple_blockchain/common/http_base"
)

// 返回请求结果
func RenderGinJsonResult(c *gin.Context, result interface{}) {
	rv := reflect.ValueOf(result)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.FieldByName("Success").Bool() {
		c.JSON(200, result)
	} else {
		c.JSON(400, result)
	}
}

// 根据是否有err获得http返回结果
func GetBaseResp(err error, successInfo string) *http_base.BaseResp {
	if err != nil {
		return &http_base.BaseResp{
			Success: false,
			Info: err.Error(),
		}
	} else {
		return &http_base.BaseResp{
			Success: true,
			Info: successInfo,
		}
	}
}
