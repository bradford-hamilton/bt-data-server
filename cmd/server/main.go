package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bradford-hamilton/bt-data-server/internal/server"
	"github.com/bradford-hamilton/bt-data-server/internal/storage"
)

func main() {
	db, err := storage.NewDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := server.New(db)
	port := os.Getenv("BT_DATA_SERVER_SERVER_PORT")
	if port == "" {
		port = "4000"
	}

	fmt.Printf("serving application on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, s.Mux))
}
