package tp_tcp_plugin

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/sllt/tp-tcp-plugin/global"
	"github.com/sllt/tp-tcp-plugin/model"
	"github.com/sllt/tp-tcp-plugin/pkg/rpc"
	"github.com/sllt/tp-tcp-plugin/pkg/server"
	"io"
	"net"
)

func handleCustomPacket(conn net.Conn) {
	defer conn.Close()
	ctx := context.Background()
	for {
		device, ok := ctx.Value("device").(*model.Device)
		if ok {
			goto authed
		}
		{
			buf := make([]byte, 40)
			n, err := conn.Read(buf)
			if err != nil {
				log.Info(err)
			}
			log.Info(string(buf[0:n]))
			// authed
			info, err := rpc.GetDeviceBufferConfig(string(buf[0:n]))
			device := &model.Device{
				AccessToken: info.Token,
				ConnectType: "custom",
				ClientConn:  conn,
				ConnConfig:  info,
				Online:      true,
			}
			ctx = context.WithValue(ctx, "deviceConnConfig", info)
			ctx = context.WithValue(ctx, "device", device)
			global.Devices[info.Token] = device
			if err != nil {
				conn.Write([]byte("error"))
				conn.Close()
				return
			}
			conn.Write([]byte("ok"))
			log.Info("device authed success:", string(buf[0:n]))
			continue
		}
	authed:
		buf := make([]byte, device.ConnConfig.InBoundByteLength*1024)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Info("client closed")
				device.Online = false
				delete(global.Devices, device.AccessToken)
				return
			}
			log.Println(err)
		}
		log.Info("received:", string(buf[0:n]))

		//conn.Write([]byte("received"))
		err = global.DefaultMqttClient.SendRawData(device.AccessToken, buf[0:n])
		if err != nil {
			conn.Write([]byte("error"))
		}
	}
}

func StartCustomProtocolServer() {
	s := server.NewServer(global.Config.CustomTcpProtocolAddr)
	s.AddConnectionHandler(handleCustomPacket)
	log.Println("custom protocol server started...")
	s.Start()
}
