package main

import (
	"flag"
	"fmt"
	"log"
	"tesks-service/internal/api"
	"tesks-service/internal/config"
	"tesks-service/internal/storage"

	"github.com/BurntSushi/toml"
)

func init() {
	//Скажем, что наше приложение будет на этапе запуска получать путь до конфиг файла из внешнего мира
	flag.StringVar(&config.ConfigPath, "path", "configs/api.toml", "path to config file in .toml format")
}

func main() {
	fmt.Println("Privet")

	//server instance initialization
	cfg := config.NewConfig()
	_, err := toml.DecodeFile(config.ConfigPath, cfg) // Десериализиуете содержимое .toml файла
	if err != nil {
		log.Println("can not find configs file. using default values:", err)
	}

	storage.Open()

	server := api.New(cfg)

	//api server start
	log.Fatal(server.Start())
}
