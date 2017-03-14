package main

import (
	"./conf"
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

func main() {

	http.HandleFunc("/", parseMessage)
	http.ListenAndServe(":88", nil)

}

func parseMessage(rw http.ResponseWriter, request *http.Request){

	bytes, _ := ioutil.ReadAll(request.Body)

	var update conf.Update
	json.Unmarshal(bytes, &update)
	for k, v := range conf.Command {
		fmt.Println(k,v)
	}
	fmt.Println(update)
}
