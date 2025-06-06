package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	conf "github.com/tarkour/product-service/internal/config"
)

func ConnectDB() (*pgx.Conn, error) {

	cfg, err := conf.ReadConfig()
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
