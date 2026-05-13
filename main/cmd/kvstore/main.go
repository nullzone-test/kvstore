package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nullzone-test/kvstore/internal/engine"
)

var version = "dev"

func main() {
	dataDir := flag.String("data", "./data", "data directory path")
	listenAddr := flag.String("addr", ":8080", "HTTP listen address")
	walSync := flag.Bool("wal-sync", true, "sync WAL on every write")
	flag.Parse()

	cfg := engine.Config{
		DataDir:     *dataDir,
		WALSync:     *walSync,
		MemTableSize: 4 * 1024 * 1024, // 4MB
		BloomFPRate: 0.01,
	}

	db, err := engine.Open(cfg)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	fmt.Printf("kvstore %s listening on %s (data: %s)\n", version, *listenAddr, *dataDir)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	fmt.Println("\nshutting down...")
}
