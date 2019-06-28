package api

/*
* 基础类目 下一个版本将支持 单个struct 自动解析。
 */

import (
	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

//
func Newctx(c *gin.Context) *Context {
	return &Context{c}
}

//获取版本号
func (c *Context) GetVersion() string {
	return c.Param("version")
}

//写入json对象
func (c *Context) WriteJson(obj interface{}) {
	c.JSON(200, obj)
}

//获取用户信息
