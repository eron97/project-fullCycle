package configs

import (
	"fmt"
	"os"
)

var (
	cfg *config
)

type config struct {
	db_driver       string
	db_host         string
	db_port         string
	db_user         string
	db_password     string
	db_name         string
	web_server_port string
	jwt_secret      string
	jwt_exxperesIn  int
}

func NewConfig() *config {
	return cfg
}

func init() {

	env := os.Getenv("DB_DRIVER")

	fmt.Println(env)

}
