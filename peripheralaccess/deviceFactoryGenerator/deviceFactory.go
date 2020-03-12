package deviceFactoryGenerator

import (
	"fmt"
	"gofrugalqsr/peripherals/peripheralaccess/models"
	"gofrugalqsr/peripherals/peripheralaccess/utils"
	"strings"
)

type DeviceFactory interface {
	ListPorts() []string
	GetDataFromDevice(models.Configs) (models.DeviceResult,error)
	SendDataToDevice(models.DeviceResult, models.Configs) (int ,error)
}

func GetDeviceFactory(device string) (DeviceFactory, error) {
	utils.Log.Printf("device type is %s",device)
	if strings.EqualFold(device,"Weighingscale") {
		return WeighingscaleFactory{}, nil
	}
	if device == "Barcode" {
		return WeighingscaleFactory{}, nil
	}
	return nil, fmt.Errorf("Wrong Device type passed")
}

