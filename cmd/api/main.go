package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/luthfiupil/dojang/internal/config"
	"github.com/luthfiupil/dojang/internal/db"
	"github.com/luthfiupil/dojang/internal/server"
)

func main() {
	cfg := config.LoadConfig("config.yaml")
	pool := db.ConnectDB(cfg.Database.Url)
	defer pool.Close()

	srv := server.NewServer(pool)

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Listening on %s", addr)
	if err := http.ListenAndServe(addr, srv.Router); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
