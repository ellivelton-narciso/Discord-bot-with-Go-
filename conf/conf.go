package conf

import (
	"eln.bot/variables"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadConfig() error {
	file, err := ioutil.ReadFile("conf/config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err = json.Unmarshal(file, &variables.Config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	variables.Token = variables.Config.Token
	variables.BotPrefix = variables.Config.BotPrefix

	return nil
}
