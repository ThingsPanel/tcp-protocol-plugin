package global

import (
	"github.com/sllt/tp-tcp-plugin/conf"
	"github.com/sllt/tp-tcp-plugin/model"
)

var (
	Devices map[string]*model.Device
	Config  *conf.Config
)
