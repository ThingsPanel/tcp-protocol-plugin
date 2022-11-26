package api

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/sllt/tp-tcp-plugin/pkg/api/resp"
	"net/http"
)

//go:embed custom_form_config.json
var customFormConfig string

//go:embed self_form_config.json
var selfFromConfig string

func CustomGetFormConfig(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "application/json", []byte(customFormConfig))
}

func CustomDeviceConfigUpdate(ctx *gin.Context) {
	resp.Success(ctx)
}

func CustomDeviceConfigCreate(ctx *gin.Context) {
	resp.Success(ctx)
}

func CustomDeviceConfigDelete(ctx *gin.Context) {
	resp.Success(ctx)
}

func SelfGetFormConfig(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "application/json", []byte(selfFromConfig))
}

func SelfDeviceConfigUpdate(ctx *gin.Context) {
	resp.Success(ctx)
}

func SelfDeviceConfigCreate(ctx *gin.Context) {
	resp.Success(ctx)
}

func SelfDeviceConfigDelete(ctx *gin.Context) {
	resp.Success(ctx)
}
