package db

import (
	"context"
	"fmt"
	"log"

	"githib.com/tarkour/product-service/internal/config"
	"github.com/jackc/pgx/v5"
)

func ConnectDB() (*pgx.Conn, error) {

	cfg, err := config.LoadConfig("./database")
	if err != nil {
		log.Fatal("Config error: ", err)
	}

	conn, err := pgx.Connect(context.Background(), cfg.Database.GetConnectionString())
	if err != nil {
		return nil, fmt.Errorf("Database connection error: %v", err)
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Database test ping error: %v", err)
	}

	return conn, nil

}
