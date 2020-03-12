package service

import

(
	"fmt"
	"go.bug.st/serial"
	"gofrugalqsr/peripherals/peripheralaccess/models"
	"gofrugalqsr/peripherals/peripheralaccess/utils"
	"log"
)

func Getdatafromscale(in models.Weighingscale) (string,error) {
	//mode := &serial.Mode{
	//	BaudRate: in.Baud,
	//	DataBits: in.DataBits,
	//	Parity:   serial.Parity(in.Parity),
	//	StopBits: serial.StopBits(in.StopBits)}
	////port, err := serial.Open(in.Port, mode)
	//port, err := serial.Open(in.Port, mode)
	//if err != nil {
	//	service.Log.Printf("Error accessing port %v",err)
	//	return "",err
	//}
	//buff := make([]byte, 100)
	////for {
	//	n, err := port.Read(buff)
	//	if err != nil {
	//		service.Log.Fatalf("Error Reading data from port %v",err)
	//		//break
	//	}
	//	if n == 0 {
	//		fmt.Println("\nEOF")
	//		//break
	//	}
	//	fmt.Printf("%v", string(buff[:n]))
	//	return string(buff[:n]),nil
	////}
	mode := &serial.Mode{
		BaudRate: 9600,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit}
	port, err := serial.Open("/dev/ttyUSB0", mode)
	if err != nil {
		utils.Log.Printf(err.Error())
		return "",nil
	}
	buff := make([]byte, 100)
	// for {
	n, err := port.Read(buff)
	if err != nil {
		utils.Log.Printf(err.Error())
		return "",nil
	}
	if n == 0 {
		fmt.Println("\nEOF")
		// break
	}
	fmt.Printf("%v", string(buff[:n]))
	// }
	return string(buff[:n]),nil
}

func ListPorts () []string {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	utils.Log.Print(ports)
	return ports
}