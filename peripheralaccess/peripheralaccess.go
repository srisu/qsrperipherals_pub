package main

import (
	"fmt"
	"gofrugalqsr/peripherals/peripheralaccess/service"
	"gofrugalqsr/peripherals/peripheralaccess/utils"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

var PIDFile = fmt.Sprintf("%s/peripherals.lock",utils.GetCwd())

func main() {
	if len(os.Args) == 1 {
		// check if daemon already running.
		if _, err := os.Stat(PIDFile); err == nil {
			fmt.Println("Already running or lock file exists.")
			os.Exit(1)
		}
		cmd := exec.Command(os.Args[0], "main")
		cmd.Start()
		fmt.Println("Daemon process ID is : ", cmd.Process.Pid)
		utils.SavePID(cmd.Process.Pid,PIDFile)
		os.Exit(0)
	}


	if strings.ToLower(os.Args[1]) == "main" { // this only called by the process and not directly

		service.StartServer() // invokes the http server in the specified port
		// Make arrangement to remove PID file upon receiving the SIGTERM from kill command
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)
		go func() { // starting a separate thread to notify the exit
			signalType := <-ch
			signal.Stop(ch)
			utils.Log.Println("Exit signal received. Exiting...",signalType)
			// remove PID file
			os.Remove(PIDFile)
			os.Exit(0)
		}()
	}

	// upon receiving the stop command read the Process ID stored in lockfile ; kill the process using the Process ID and exit
	//If Process ID does not exist, prompt error and quit
	if len(os.Args) > 1 { //only when a stop argument is passed
		if strings.ToLower(os.Args[1]) == "stop" {
			if _, err := os.Stat(PIDFile); err == nil {
				data, err := ioutil.ReadFile(PIDFile)
				if err != nil {
					fmt.Println("Not running")
					os.Exit(1)
				}
				ProcessID, err := strconv.Atoi(string(data))
				if err != nil {
					fmt.Println("Unable to read and parse process id found in ", PIDFile)
					os.Exit(1)
				}
				process, err := os.FindProcess(ProcessID)
				if err != nil {
					fmt.Printf("Unable to find process ID [%v] with error %v \n", ProcessID, err)
					os.Exit(1)
				}
				// remove PID file
				os.Remove(PIDFile)

				fmt.Printf("Killing process ID [%v] now.\n", ProcessID)
				// kill process and exit immediately
				err = process.Kill()

				if err != nil {
					fmt.Printf("Unable to kill process ID [%v] with error %v \n", ProcessID, err)
					os.Exit(1)
				} else {
					fmt.Printf("Killed process ID [%v]\n", ProcessID)
					utils.Log.Printf("Killed the process [%v]",ProcessID)
					os.Exit(0)
				}

			} else {
				fmt.Println("Not running.")
				os.Exit(1)
			}
		} else {
			fmt.Printf("Unknown command : %v\n", os.Args[1])
			fmt.Printf("Usage : %s [stop]\n", os.Args[0]) // return the program name back to %s
			os.Exit(1)
		}
	}

}
