package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ParadoxZero/wnote/internal/store"
)

func main() {
    fmt.Println("Hello, World!")
		s, err := store.CreateSqliteStore("db.sqlite3")
		if err != nil {
			log.Fatalln("Error creating Store: ", err)
			os.Exit(1)
		}
		defer s.Close()
}
