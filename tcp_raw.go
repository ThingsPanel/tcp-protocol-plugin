package tp_tcp_plugin

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/sllt/ergo"
	"github.com/sllt/ergo/etf"
	"github.com/sllt/ergo/gen"
	"github.com/sllt/ergo/node"
)

type rawApp struct {
	gen.Application
}

func (ra *rawApp) Load(args ...etf.Term) (gen.ApplicationSpec, error) {
	return gen.ApplicationSpec{
		Name:        "rawTCPApp",
		Description: "Raw TCP APP",
		Version:     "v.1.0",
		Children: []gen.ApplicationChildSpec{
			{
				Child: &rawServer{},
				Name:  "raw",
			},
		},
	}, nil
}

func (ra *rawApp) Start(process gen.Process, args ...etf.Term) {
	fmt.Println("Raw TCP Application started!")
}

func StartRawServer() {
	opts := node.Options{
		Applications: []gen.ApplicationBehavior{
			&rawApp{},
		},
	}

	log.Info("started raw tcp server...")
	rawNode, err := ergo.StartNode("raw@127.0.0.1", "secret", opts)
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Info(err)
		return
	}

	rawNode.Wait()
}
