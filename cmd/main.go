package main

import (
	"fmt"
	"github.com/eron97/project-fullCycle.git/configs"
)

func main() {
	cfg := configs.NewConfig()
	driver := cfg.GetDBDriver()
	fmt.Println(driver)
}
