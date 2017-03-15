package main

import (
	"./conf"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"strconv"
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
		Command(update)
	}


	//fmt.Println(Command,Isset)
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

func getMethod (Command string)func(update conf.Update){

	mc := memcache.New("127.0.0.1:11211")

	NewMethod:= func(update conf.Update) {}
	switch Command {
	case "Start":
		NewMethod = func(update conf.Update) {
			mc.Set(&memcache.Item{Key:strconv.Itoa(update.Message.From.ID),Value:[]byte(Command)})
			fmt.Println("CurrentCommand:",Command)
		}
	default:
		NewMethod = func(update conf.Update) {
			val, _ := mc.Get(strconv.Itoa(update.Message.From.ID))
			fmt.Println("LastCommand:",string(val.Value))

		}

	}

	return NewMethod

}



