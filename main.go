package main

import (
	"flag"
	"fmt"
	"github.com/signmem/log-template/g"
	"log"
	"os"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")

	flag.Parse()

	if *version {
		version := g.Version
		fmt.Printf("%s", version)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)
	g.InitLog()
	
	log.Printf("[INFO] aaaaa bbbb cdcc")
	if g.Config().Debug {
		log.Printf("[DEBUG] yes")
	}



}