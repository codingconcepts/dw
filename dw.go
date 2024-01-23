package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("first argument must be a URL to the database")
	}

	url := os.Args[1]

	attempts := 0
	for range time.NewTicker(time.Second).C {
		attempts++

		db, err := pgx.Connect(context.Background(), url)
		if err != nil {
			fmt.Printf("attempt %d failed\r", attempts)
			continue
		}

		defer db.Close(context.Background())
		return
	}
}
