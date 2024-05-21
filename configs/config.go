package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-chi/jwtauth"
	"github.com/lpernett/godotenv"
)

var (
	cfg *config
)

type config struct {
	dbDriver      string
	dbHost        string
	dbPort        string
	dbUser        string
	dbPassword    string
	dbName        string
	webServerPort string
	jwtSecret     string
	jwtExpiresIn  int
	tokenAuth     *jwtauth.JWTAuth
}

func NewConfig() *config {
	return cfg
}

func (c *config) GettokenAuth() *jwtauth.JWTAuth {
	return c.tokenAuth
}

func (c *config) GetjwtExpiresIn() int {
	return c.jwtExpiresIn
}

func init() {

	cfg = &config{}

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	cfg.dbDriver = os.Getenv("DB_DRIVER")
	cfg.dbHost = os.Getenv("DB_HOST")
	cfg.dbPort = os.Getenv("DB_PORT")
	cfg.dbUser = os.Getenv("DB_USER")
	cfg.dbPassword = os.Getenv("DB_PASSWORD")
	cfg.dbName = os.Getenv("DB_NAME")
	cfg.webServerPort = os.Getenv("WEB_SERVER_PORT")
	cfg.jwtSecret = os.Getenv("JWT_SECRET")
	jwtExpiresIn, _ := strconv.Atoi(os.Getenv("JWT_EXPIRESIN"))
	cfg.jwtExpiresIn = jwtExpiresIn
	cfg.tokenAuth = jwtauth.New("HS256", []byte(cfg.jwtSecret), nil)
}
