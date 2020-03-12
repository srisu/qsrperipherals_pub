package service

import (
	"fmt"
	"gofrugalqsr/peripherals/peripheralaccess/models"
	"gofrugalqsr/peripherals/peripheralaccess/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)



func GetConfig() models.Values {
str,err := ioutil.ReadFile(fmt.Sprintf("%s/peripheralconfig.yml", utils.GetCwd()))
if err != nil{
	utils.Log.Fatalf("Error while reading config %v",err.Error())
}
	var cfg models.Values
	err = yaml.Unmarshal(str, &cfg)
	if err != nil {
		utils.Log.Fatalf("Error while parsing config %v",err.Error())
	}
	return cfg
}
