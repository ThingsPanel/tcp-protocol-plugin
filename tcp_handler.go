package tp_tcp_plugin

import (
	log "github.com/sirupsen/logrus"
	"github.com/sllt/ergo/gen"
	"github.com/sllt/tp-tcp-plugin/global"
	"github.com/sllt/tp-tcp-plugin/model"
)

type tcpHandler struct {
	gen.TCPHandler
	onlineDevice map[string]*model.Device
}

func (th *tcpHandler) HandleConnect(process *gen.TCPHandlerProcess, conn *gen.TCPConnection) gen.TCPHandlerStatus {
	log.Info("device connect:", conn.Addr.String())
	th.onlineDevice[conn.Addr.String()] = nil
	return gen.TCPHandlerStatusOK
}
func (th *tcpHandler) HandleDisconnect(process *gen.TCPHandlerProcess, conn *gen.TCPConnection) {
	device := th.onlineDevice[conn.Addr.String()]
	if device.Conn.IsConnected() {
		device.Conn.Disconnect(0)
	}
	delete(global.Devices, device.AccessToken)
	delete(th.onlineDevice, conn.Addr.String())
	log.Info("device disconnect:", conn.Addr.String())
}

func (th *tcpHandler) HandlePacket(process *gen.TCPHandlerProcess, packet []byte, conn *gen.TCPConnection) (int, int, gen.TCPHandlerStatus) {
	if !(packet[0] == 'T' && packet[1] == 'P') {
		log.Println("invalid protocol")
		conn.Socket.Write([]byte("error"))
		return 0, 0, gen.TCPHandlerStatusClose
	}

	// heartbeat packet
	if packet[2] == 0x1 {
		// TODO: handle heartbeat packet
		return 0, 0, gen.TCPHandlerStatusOK
	}

	p := &model.Packet{}
	p.Parse(packet)

	device := &model.Device{}
	switch p.Cmd {
	case 0x0:
		// auth
		log.Println("get auth request...")
		device = &model.Device{
			AccessToken: string(p.Payload),
			ConnectType: "self",
			ClientConn:  conn,
			ConnConfig:  nil,
		}
		err := device.Auth(global.Config.Mqtt.Addr)
		log.Info(err)
		if err != nil {
			conn.Socket.Write([]byte("MQTT:" + err.Error()))
			return 0, 0, gen.TCPHandlerStatusClose
		}

		device.Online = true
		global.Devices[string(p.Payload)] = device
		th.onlineDevice[conn.Addr.String()] = device
	case 0x1:
		// publish
		log.Println("get publish request...")
		device = th.onlineDevice[conn.Addr.String()]
		if device == nil || device.Online == false {
			conn.Socket.Write([]byte("MQTT: device auth failed"))
			return 0, 0, gen.TCPHandlerStatusClose
		}
		device.Publish(string(p.Payload))
	case 0x2:
		// TODO publish events
	}

	return 0, 0, gen.TCPHandlerStatusOK
}
