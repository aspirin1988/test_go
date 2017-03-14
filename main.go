package main

import (
	"./conf"
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

func main() {

	//var JsonStr string = "{\"update_id\":144847163,\"message\":{\"message_id\":198,\"from\":{\"id\":169105432,\"first_name\":\"Demidov\",\"last_name\":\"Sergey\",\"username\":\"Aspirin_Sergey\"},\"chat\":{\"id\":169105432,\"first_name\":\"Demidov\",\"last_name\":\"Sergey\",\"username\":\"Aspirin_Sergey\",\"type\":\"private\"},\"date\":1487839051,\"text\":\"/console\",\"entities\":[{\"type\":\"bot_command\",\"offset\":0,\"length\":8}]}}"
	//fmt.Println(conf.BOT_TOKEN)
	//var bytes =[]byte(JsonStr)
	//
	//
	//var update conf.Update
	//json.Unmarshal(bytes, &update)

	http.HandleFunc("/", parseMessage)
	http.ListenAndServe(":88", nil)

}

func parseMessage(rw http.ResponseWriter, request *http.Request){

	fmt.Println("test")

	bytes, _ := ioutil.ReadAll(request.Body)

	var update conf.Update
	json.Unmarshal(bytes, &update)

	fmt.Println(update)
}
