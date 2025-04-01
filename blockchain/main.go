package main

import (
	"log"
	"time"

	"go_blockchain/core"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := core.Block{Index: 0, Timestamp: t.String(), BPM: 0, Hash: "", PrevHash: ""}
		spew.Dump(genesisBlock)
		core.Blockchain = append(core.Blockchain, genesisBlock)
	}()
	log.Fatal(core.Run())
}
