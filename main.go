package main

import (
	"eln.bot/conf"
	"eln.bot/connect"
	"fmt"
)

func main() {
	err := conf.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	connect.Start()

	<-make(chan struct{})
	return
	fmt.Println("Bot está funcionando!")

}
