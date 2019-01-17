package main

import (
	"appletMessagesServer/src/g"
	"appletMessagesServer/src/work"

	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var Conf g.Config

func main() {
	version := flag.Bool("v", false, "show version")
	flag.Parse()
	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}
	g.GetConf(&Conf)
	g.Init(Conf)

	work.Start()
	// 停止 处理
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		log.Println("all service stopping...")
		work.Stop()
		log.Println("all service stop ok ")
		os.Exit(0)
	}()
	select {}
}
