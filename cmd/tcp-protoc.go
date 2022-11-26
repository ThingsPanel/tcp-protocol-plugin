package main

import (
	plugin "github.com/sllt/tp-tcp-plugin"
	"github.com/sllt/tp-tcp-plugin/conf"
	"github.com/sllt/tp-tcp-plugin/global"
	"github.com/sllt/tp-tcp-plugin/pkg/api"
)

func main() {
	conf.LogInit()
	global.DefaultMqttClient.Init()
	// subscribe data from mqtt
	global.DefaultMqttClient.Subscribe()

	// started api server
	go api.NewCustomServer(global.Config.Api.CustomAddr).Start()
	go api.NewSelfApiServer(global.Config.Api.SelfAddr).Start()
	// started tcp server
	go plugin.StartCustomProtocolServer()
	plugin.Start(global.Config.TcpProtocolAddr)
}
