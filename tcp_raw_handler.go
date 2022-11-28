package tp_tcp_plugin

import (
	log "github.com/sirupsen/logrus"
	"github.com/sllt/ergo/gen"
	"github.com/sllt/tp-tcp-plugin/global"
	"github.com/sllt/tp-tcp-plugin/model"
	"github.com/sllt/tp-tcp-plugin/pkg/rpc"
)

type rawTCPHandler struct {
	gen.TCPHandler
	onlineDevice map[string]*model.Device
}

func (th *rawTCPHandler) HandleConnect(process *gen.TCPHandlerProcess, conn *gen.TCPConnection) gen.TCPHandlerStatus {
	log.Info("device connect:", conn.Addr.String())
	th.onlineDevice[conn.Addr.String()] = nil
	return gen.TCPHandlerStatusOK
}
func (th *rawTCPHandler) HandleDisconnect(process *gen.TCPHandlerProcess, conn *gen.TCPConnection) {
	delete(th.onlineDevice, conn.Addr.String())
	log.Info("device disconnect:", conn.Addr.String())
}

func (th *rawTCPHandler) HandlePacket(process *gen.TCPHandlerProcess, packet []byte, conn *gen.TCPConnection) (int, int, gen.TCPHandlerStatus) {
	log.Info(th.onlineDevice)
	if th.onlineDevice[conn.Addr.String()] == nil {
		info, err := rpc.GetDeviceBufferConfig(string(packet))
		if err != nil {
			return 0, 0, gen.TCPHandlerStatusClose
		}
		device := &model.Device{
			AccessToken: info.Token,
			ConnectType: "custom",
			ClientConn:  conn,
			ConnConfig:  info,
			Online:      true,
		}
		th.onlineDevice[conn.Addr.String()] = device
		global.Devices[info.Token] = device
		log.Info("device authed success:", string(packet))
		return 0, 0, gen.TCPHandlerStatusOK
	}

	device := th.onlineDevice[conn.Addr.String()]

	err := global.DefaultMqttClient.SendRawData(device.AccessToken, packet)
	if err != nil {
		conn.Socket.Write([]byte("error"))
	}
	log.Info("send data to mqtt success:", string(packet))
	return 0, 0, gen.TCPHandlerStatusOK
}
