package main

import (
	"flag"
	"fmt"
)

type config struct {
	addr      string
	staticDir string
	debug     bool
	port      int
}

func main() {
	var cfg config

	// 1. Bind command-line flags to memory addresses of the struct fields
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.BoolVar(&cfg.debug, "debug", false, "Enable debug mode")
	flag.IntVar(&cfg.port, "port", 8080, "Port to listen on")

	// 2. Parse the command-line flags AFTER defining them
	flag.Parse()

	// 3. Access the values directly from the struct
	fmt.Println("Configuration Loaded:")
	fmt.Printf("  Address:    %s\n", cfg.addr)
	fmt.Printf("  Static Dir: %s\n", cfg.staticDir)
	fmt.Printf("  Debug Mode: %t\n", cfg.debug)
	fmt.Printf("  Port:       %d\n", cfg.port)
}
