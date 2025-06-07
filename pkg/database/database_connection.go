package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/tarkour/product-service/pkg/config"
)

func ConnectDB(path string) (*pgx.Conn, error) {

	cfg, err := config.LoadConfig(path)
	if err != nil {
		log.Fatal("Config error: ", err)
	}

	conn, err := pgx.Connect(context.Background(), cfg.Database.GetConnectionString())
	if err != nil {
		return nil, fmt.Errorf("database connection error: %v", err)
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("database test ping error: %v", err)
	}

	return conn, nil

}
