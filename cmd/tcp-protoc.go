package main

import (
	log "github.com/sirupsen/logrus"
	plugin "github.com/sllt/tp-tcp-plugin"
	"github.com/sllt/tp-tcp-plugin/conf"
	"github.com/sllt/tp-tcp-plugin/global"
	"github.com/sllt/tp-tcp-plugin/model"
	"github.com/sllt/tp-tcp-plugin/pkg/api"
	"os"
)

func main() {
	conf.LogInit()
	global.Devices = make(map[string]*model.Device, 0)
	err := global.DefaultMqttClient.Init()
	if err != nil {
		log.Info(err)
		os.Exit(1)
	}
	// subscribe data from mqtt
	global.DefaultMqttClient.Subscribe()

	// started api server
	go api.NewCustomServer(global.Config.Api.CustomAddr).Start()
	go api.NewSelfApiServer(global.Config.Api.SelfAddr).Start()
	// started tcp server
	go plugin.StartRawServer()
	plugin.Start(global.Config.TcpProtocolAddr)
}
