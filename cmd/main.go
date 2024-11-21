package main

import (
	"BharadwajaD/DistSys/pkg/base"
	"flag"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)


func main(){

	host := flag.String("host", "127.0.0.1", "ip address")
	port := flag.Int("port", 42069, "port")
	debug := flag.Bool("debug", false, "debug flag")
	nm := flag.Bool("nm", false, "use this flag to create node manager")
	flag.Parse()

	log.Info().Msgf("ip %s:%d, debug: %d, nm :%d", *host, *port, *debug, *nm)

    zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
    	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	connector := base.NewConnector(*host, *port, base.NODE_MANAGER_ADDR)
	var node_manager *base.NodeManager
	if *nm {
		node_manager = base.NewNodeManager(*host, *port)
		connector.RegisterService(node_manager)
	}


	wg := sync.WaitGroup{}
	wg.Add(1)
	go func (){
		defer wg.Done()
		connector.Start()
	}()

	wg.Wait()
}
