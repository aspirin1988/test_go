package main

import (
	"./conf"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {

	http.HandleFunc("/", parseMessage)
	http.ListenAndServe(":88", nil)

}

func parseMessage(rw http.ResponseWriter, request *http.Request){

	bytes, _ := ioutil.ReadAll(request.Body)

	var update conf.Update
	json.Unmarshal(bytes, &update)

	isCommand(update.Message.Text)

	fmt.Println(update)
}

func isCommand(text string)(string, bool)  {

	var result bool = false
	Commands:= make(map[string]string)
	Commands["Start"]="/start"
	Commands["MainNews"]="Главные новости"
	Commands["LastNews"]="Последние новости"
	Commands["News"]="Новости"
	Commands["Article"]="Статьи"
	Commands["OpinionBatle"]="Битва мнений"
	Commands["Opinions"]="Блоги и мнения"

	for k, v := range Commands {

		if v==text {
			fmt.Println(k, v)
			result = true
			text = k
		}
	}

	return text,result

}



