package main

import (
	"BharadwajaD/DistSys/base"

	"sync"
	"flag"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)


func main(){

	host := flag.String("host", "127.0.0.1", "ip address")
	port := flag.Int("port", 42069, "port")
	debug := flag.Bool("debug", false, "debug flag")
	flag.Parse()

    zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
    	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	node := base.NewNode(*host, *port)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func (){
		defer wg.Done()
		node.Start()
	}()

	testMessageSending(node, *port)
	wg.Wait()
}

func testMessageSending(node *base.Node, port int) {
	if port == 42069 {
		reply , err := node.Send("127.0.0.1:42070", base.NewRPCMessage("Hello test"))
		if err != nil {
			log.Fatal().Msg(err.Error())
		}
		log.Info().Msg(reply.Content)
	}
}
