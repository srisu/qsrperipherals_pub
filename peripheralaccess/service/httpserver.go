package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gofrugalqsr/peripherals/peripheralaccess/deviceFactoryGenerator"
	"gofrugalqsr/peripherals/peripheralaccess/models"
	"gofrugalqsr/peripherals/peripheralaccess/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	//"strings"
)
var configs models.Configs
func init() {
	configs =  GetConfig().Configs
}
var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan models.SocketResponse)          // broadcast channel

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Define our message object


func  StartServer() {
	// Create a simple file server
	fs := http.FileServer(http.Dir(fmt.Sprintf("%s/views/dist/views", utils.GetCwd())))
	http.Handle("/", fs)
	http.HandleFunc("/api/configurations",configurations)

	// Configure websocket router
	http.HandleFunc("/ws", handleConnections)

	// Start listening for incoming chat messages
	go handleMessages()

	// Start the server on localhost port 54749 and log any errors
	utils.Log.Println("http server started on port 54749")
	err := http.ListenAndServe(":54749", nil)
	if err != nil {
		utils.Log.Fatal("Error while initiating http server: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.Log.Printf("Exception while upgrading the client to socket %v",err.Error())
	}
	//Make sure we close the connection when the function returns
	//defer ws.Close()

	// Register our new client
	clients[ws] = true

	//for {

		var msg models.SocketRequest
		var response models.SocketResponse

	wserr := ws.ReadJSON(&msg)
	if err != nil {
		utils.Log.Printf("Error while Reading message from client %v",wserr)
		delete(clients, ws)
		//break
	}
		utils.Log.Printf(msg.DeviceType)
	utils.Log.Printf(msg.ActionCode)
	utils.Log.Printf(msg.ReadInterval)
		device,err := deviceFactoryGenerator.GetDeviceFactory(msg.DeviceType)
		if err != nil {
			utils.Log.Printf("Error while getting devicefactory : %v", err)
			//break
			response.Message = err.Error()
			response.Status = "FAIL"
		}


		if strings.EqualFold(msg.ActionCode,"READ") {
			deviceVal, err := device.GetDataFromDevice(configs)
			if err != nil {
				utils.Log.Printf("Error reading data from device %v", err)
				response.Message = err.Error()
				response.Status = "FAIL"
			} else {
				response.Value = deviceVal.Value
				response.Status = "PASS"
			}

			broadcast <- response
			//}
		} else {
			response.Message = "Specify  a valid actionCode"
			response.Status = "FAIL"
		}
		}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		output := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(output)
			if err != nil {
				utils.Log.Printf("Error sending Data to socket: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func configurations(w http.ResponseWriter, r *http.Request){
	configList := GetConfig()
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(configList)
	case "POST":
		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()
		var t = models.Values{}
		w.Header().Set("Content-Type", "application/json")
		err := d.Decode(&t)
		if err != nil{
			utils.Log.Printf("Error parsing request %v",err)
			w.WriteHeader(400)
			response := models.ApiResponse{Status: "ERROR",Message: "Unable to parse the request body"}
			json.NewEncoder(w).Encode(response)
		} else {
			data , err := yaml.Marshal(&t)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			err = ioutil.WriteFile(fmt.Sprintf("%s/peripheralconfig.yml", utils.GetCwd()), data, 0644)
			if err != nil {
				w.WriteHeader(500)
				response := models.ApiResponse{Status: "ERROR", Message: "Cannot save configurations to file"}
				json.NewEncoder(w).Encode(response)
			}
			response := models.ApiResponse{Status: "OK", Message: "Configurations Updated Successfully"}
			json.NewEncoder(w).Encode(response)
		}
	default:
		w.WriteHeader(405)
		w.Write([]byte("Only GET and POST are allowed"))
	}

}