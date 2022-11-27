package tp_tcp_plugin

import (
	"github.com/sllt/ergo/etf"
	"github.com/sllt/ergo/gen"
	"github.com/sllt/tp-tcp-plugin/global"
	"github.com/sllt/tp-tcp-plugin/model"
	"strconv"
	"strings"
)

type rawServer struct {
	gen.TCP
}

func (rs *rawServer) InitTCP(process *gen.TCPProcess, args ...etf.Term) (gen.TCPOptions, error) {
	hostArr := strings.Split(global.Config.CustomTcpProtocolAddr, ":")
	port, _ := strconv.Atoi(hostArr[1])
	options := gen.TCPOptions{
		Host: hostArr[0],
		Port: uint16(port),
		Handler: &rawTCPHandler{
			onlineDevice: make(map[string]*model.Device, 0),
		},
	}

	return options, nil
}
