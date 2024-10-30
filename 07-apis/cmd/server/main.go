package main

import "github.com/rafaelcmd/go-expert/07-apis/configs"

func main() {
	config := configs.LoadConfig(".")
	println(config.DBDriver)
}
