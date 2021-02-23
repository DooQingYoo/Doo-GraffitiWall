package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conf struct {
	Port int `json:"port"`
	Path string `json:"path"`
}

var Config = new(Conf)

func init() {
	file, err := os.Open("config" + string(os.PathSeparator) + "config.json")
	if err == nil {
		stat, _ := file.Stat()
		size := stat.Size()
		b := make([]byte, size)
		_, err = file.Read(b)
		if err == nil {
			err = json.Unmarshal(b, Config)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println(err.Error())
	}
}
