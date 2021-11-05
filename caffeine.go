package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/rehacktive/caffeine/database"
	"github.com/rehacktive/caffeine/service"
)

var (
	isProduction = false
	addr         = ":8000"
	domain       = ""
)

func main() {
	flag.StringVar(&addr, "ip_port", ":8000", "ip:port to expose")
	flag.BoolVar(&isProduction, "production", false, "if true, we start HTTPS server")
	flag.StringVar(&domain, "domain", "", "domain to use in production")
	flag.Parse()

	server := service.Server{
		Address:      addr,
		IsProduction: isProduction,
	}
	go server.Init(&database.MemDatabase{})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	log.Println("bye")
}
