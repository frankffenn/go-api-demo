package routers

import (
	"github.com/frankffenn/trading-assistants/comm"
	"github.com/frankffenn/trading-assistants/errors"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(data comm.JsonObj) gin.H {
	return gin.H{
		"success": true,
		"data":    data,
	}
}

func ResponseFailWithError(err *errors.Error) gin.H {
	return gin.H{
		"success": false,
		"code":    err.Code,
		"message": err.Message,
	}
}
