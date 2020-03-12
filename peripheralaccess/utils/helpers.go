package utils

import (
	"os"
	"strconv"
)

func GetCwd() string{
	val,err := os.Getwd()
	if err != nil{
		return err.Error()
	}
	return val
}

func SavePID(pid int,PIDFile string) {

	file, err := os.Create(PIDFile)
	if err != nil {
		Log.Fatalf("Unable to create pid file : %v\n", err)
	}

	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(pid))

	if err != nil {
		Log.Fatalf("Unable to create pid file : %v\n", err)
	}

	file.Sync() // flush to disk

}