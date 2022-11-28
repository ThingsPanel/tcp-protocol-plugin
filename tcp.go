package tp_tcp_plugin

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/sllt/ergo"
	"github.com/sllt/ergo/etf"
	"github.com/sllt/ergo/gen"
	"github.com/sllt/ergo/node"
)

type tcpApp struct {
	gen.Application
}

func (ta *tcpApp) Load(args ...etf.Term) (gen.ApplicationSpec, error) {
	return gen.ApplicationSpec{
		Name:        "tcpApp",
		Description: "ThingsPanel Common TCP APP",
		Version:     "v.1.0",
		Children: []gen.ApplicationChildSpec{
			{
				Child: &tcpServer{},
				Name:  "raw",
			},
		},
	}, nil
}

func (ta *tcpApp) Start(process gen.Process, args ...etf.Term) {
	fmt.Println("Common TCP Application started!")
}

func StartTcpServer() {
	opts := node.Options{
		Applications: []gen.ApplicationBehavior{
			&tcpApp{},
		},
	}

	log.Info("started tcp server...")
	tcpNode, err := ergo.StartNode("tcp@127.0.0.1", "secret", opts)
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Info(err)
		return
	}

	tcpNode.Wait()
}
