package main

import (
	"BharadwajaD/DistSys/base"

	"flag"
    "github.com/rs/zerolog"
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
	node.Start()
}
