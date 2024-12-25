package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type ConnStr struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Username string `yaml:"username"`
	Dbname   string `yaml:"dbname"`
}

func Get_connStr(FilePath string) string {
	f := &ConnStr{}
	source, err := os.ReadFile(FilePath)
	if err != nil {
		log.Println(err)
	}

	err = yaml.Unmarshal([]byte(source), &f)
	if err != nil {
		log.Printf("error: %v", err)
	}

	fmt.Println(f.Dbname)
	fmt.Println(f.Host)
	fmt.Println(f.Password)
	fmt.Println(f.Port)
	fmt.Println(f.Username)

	connStr := "postgres://" + f.Username + ":" + f.Password + "@" + f.Host + ":" + f.Port + "/" + f.Dbname + "?sslmode=disabled"

	return connStr
}
