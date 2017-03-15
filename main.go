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


	NewMethod:= func(update conf.Update) {}
	switch Command {
	case "Start":
		NewMethod = func(update conf.Update) {
			setCommand(update,Command)
			fmt.Println("CurrentCommand:",Command)
		}
	default:
		NewMethod = func(update conf.Update) {
			var val = GetCommand(update)
			fmt.Println("LastCommand:",val)

		}

	}

	return NewMethod

}

func setCommand(UserID conf.Update, Command string)  {
	mc := memcache.New("127.0.0.1:11211")
	mc.Set(&memcache.Item{Key:strconv.Itoa(UserID.Message.From.ID),Value:[]byte(Command)})
}

func GetCommand(UserID conf.Update)(string)  {
	mc := memcache.New("127.0.0.1:11211")
	val, _ := mc.Get(strconv.Itoa(UserID.Message.From.ID))
	var result = string(val.Value)
	return result
}



