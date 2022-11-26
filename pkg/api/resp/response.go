package resp

import "github.com/gin-gonic/gin"

type resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Success(ctx *gin.Context) {
	ctx.JSON(200, resp{
		Code:    0,
		Message: "success",
	})
}
