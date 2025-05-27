package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	db "githib.com/tarkour/product-service/pkg/database"
	"github.com/spf13/viper"
)

func main() {

	config := viper.New()
	config.Set("database.safe_mode", true)

	conn := db.ConnectDB(config)
	queryExec := db.NewQueryExecutor(
		conn,
		config.GetBool("database.safe_mode"),
		slog.Default(), // Логгер из pkg/slog_response
	)

	conn, err := db.ConnectDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: ", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	fmt.Println("Database connected.")

}
