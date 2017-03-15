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

	Command,Isset :=isCommand(update.Message.Text)

	if Isset{
		Command:=getMethod(Command)
		Command()
	}


	fmt.Println(Command,Isset)
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
			result = true
			text = k
		}
	}

	return text,result

}

func getMethod (Command string)func(){

	NewMethod:= func() {}
	switch Command {
	case "Start":
		NewMethod = func() {
			fmt.Println("Start")
		}
	default:
		NewMethod = func() {
			fmt.Println("Default")
		}

	}

	return NewMethod

}



