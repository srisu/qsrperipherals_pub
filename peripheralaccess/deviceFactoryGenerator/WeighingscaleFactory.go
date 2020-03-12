package deviceFactoryGenerator

import (
	"errors"
	"go.bug.st/serial"
	"gofrugalqsr/peripherals/peripheralaccess/models"
	"gofrugalqsr/peripherals/peripheralaccess/utils"
	"strings"
)

type WeighingscaleFactory struct {
}


func  (WeighingscaleFactory WeighingscaleFactory)  ListPorts () []string {
	ports, err := serial.GetPortsList()
	if err != nil {
		utils.Log.Fatal(err)
	}
	utils.Log.Print(ports)
	return ports
}

func (WeighingscaleFactory WeighingscaleFactory)GetDataFromDevice(configs models.Configs) ( out models.DeviceResult,err error){
	val,err := Getdatafromscale(configs.Weighingscale)
	utils.Log.Printf("%v",err)
	if err != nil {
		return
	}
	return val,nil
}

func (WeighingscaleFactory WeighingscaleFactory)SendDataToDevice(data models.DeviceResult, config models.Configs) (n int, err error){
	return 0,errors.New("is nill");
}

func Getdatafromscale(in models.Weighingscale) (WeighingScaleResult models.DeviceResult,err error) {
	utils.Log.Printf("inside getdatafromsacle")
	mode := &serial.Mode{
		BaudRate: in.Baud,
		DataBits: in.DataBits,
		Parity:   serial.Parity(in.Parity),
		StopBits: serial.StopBits(in.StopBits)}
	port, err := serial.Open(in.Port, mode)
	if err != nil {
		return WeighingScaleResult,err
	}
	defer port.Close()
	buff := make([]byte, 100)
	//for {
	//var val string
	//for i := 0; i < 5; i++ {
	//	n, err := port.Read(buff)
	//	if err != nil {
	//		service.Log.Fatal(err)
	//		//break
	//	}
	//	if n == 0 {
	//		fmt.Println("\nEOF")
	//		//break
	//	}
	//	service.Log.Printf("%v", string(buff[:n]))
	//	val =  string(buff[:n])
	//	//}
	//}
	//val = strings.Replace(val, "\n", "", -1)
	//val = val[2:13]
	//return val,nil
	var val,unit string
	for i := 0; i < 5; i++{
		n, err := port.Read(buff)
		if err != nil {
			utils.Log.Printf("error from read port is %v",err)
			return WeighingScaleResult,err
		}
		if n == 0 {
			continue
		}
		val = string(buff[:n])
	}
	utils.Log.Println(val)
	orgval := strings.Replace(val, "\n", "", -1)
	val = orgval[in.Position.Start:in.Position.Start+in.Position.Length]
	unit = orgval[in.Position.Start+in.Position.Length:in.Position.Start+in.Position.Length+4]
	utils.Log.Println(unit)
	WeighingScaleResult.Unit  = unit
	WeighingScaleResult.Value = val
	return
}


