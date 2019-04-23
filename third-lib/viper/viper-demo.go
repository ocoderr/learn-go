package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type favorite struct {
	Sports      []string
	Music       []string
	LuckyNumber int
}

type information struct {
	Name   string
	Age    int
	Alise  []string
	Image  string
	Public bool
}

type YamlConfig struct {
	TimeStamp   string
	Author      string
	PassWd      string
	Information information
	Favorite    favorite
}

func main() {
	viper := viper.New()
	viper.SetConfigFile("config.yaml")
	//viper.AddConfigPath("/Users/yanliang/go/src/github.com/ocoderr/learn-go/third-lib/viper/")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	log.Printf("age: %s, name: %s \n", viper.Get("information.age"), viper.Get("information.name"))

	var cf YamlConfig
	if err := viper.Unmarshal(&cf); err != nil {
		panic(err)
	}
	fmt.Printf("%v", cf)

}
