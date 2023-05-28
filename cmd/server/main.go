package main

import "github.com/mdiaas/goapi/configs"

func main() {
	configs, _ := configs.LoadConfig(".")
	println(configs.DBHost)
}
