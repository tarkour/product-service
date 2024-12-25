package main

import (
	"fmt"

	"github.com/tarkour/product-service/config"
)

func main() {

	connstr := config.Get_connStr("./config/database_connection.yaml")

	fmt.Println(connstr)

}
