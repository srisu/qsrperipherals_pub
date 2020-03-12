package utils

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"log"
	"os"
)

var Log *log.Logger

func init() {
e, err := os.OpenFile(fmt.Sprintf("%s/qsrperipherals.log", GetCwd()), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

if err != nil {
fmt.Printf("error opening file: %v", err)
os.Exit(1)
}
Log = log.New(e, "", log.Ldate|log.Ltime)
Log.SetOutput(&lumberjack.Logger{
Filename:   fmt.Sprintf("%s/qsrperipherals.log", GetCwd()),
MaxSize:    1,  // megabytes after which new file is created
MaxBackups: 3,  // number of backups
MaxAge:     28, //days
})
}