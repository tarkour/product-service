package main

import (
	"context"
	"fmt"
	"os"

	db "githib.com/tarkour/product-service/database"
)

func main() {

	conn, err := db.ConnectDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: ", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	fmt.Println("Database connected.")

}
