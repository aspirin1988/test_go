package main

import (
	"./conf"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"strconv"
	"net/url"
	"bytes"
	"time"
	"database/sql"
)

import _ "github.com/go-sql-driver/mysql"

var db *sql.DB
var err error


func init() {
	db, err = sql.Open("mysql","tengrinews:vfzq2sE8hKzNqkkLyar6@/tengrinews")
	if err != nil {
		panic(err.Error())
	}else {
		fmt.Println("OK!")
	}
	// sql.DB should be long lived "defer" closes it once this function ends
	defer db.Close()

	// Test the connection to the database
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	res, err :=db.Query("SELECT count(id) FROM news")
	if err != nil {
		panic(err)
	}
	//for res.Next() {
		println(res)
	//}

}

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

	Commands := map[string]string{
		"start"          : "/start",
		"help"           : "/help",
		"menu"           : "Ⓜ️Меню",
		"mainNews"       : "Главные новости",
		"news"           : "Новости",
		"economic"       : "Экономика",
		"accidents"      : "Происшествия",
		"sports"         : "Спорт",
		"tech"           : "Технологии",
		"life"           : "Жизнь",
		"culture"        : "Культура",
		"lastNews"       : "Последние новости",
		"articles"       : "Статьи",
		"back"           : "Назад",
		"showMor"        : "Следующая⏭",
		"showPrev"       : "⏮Предыдущая",
		"clearCache"     : "Очистить",
		"battleopinions" : "Битва мнений",
		"opinion"        : "Блоги и мнения",
		"promises"       : "Обещания",
		"conference"     : "Интервью",
		"fotoarchive"    : "Фотогалерея",
		"userlist"       : "Список пользователей",
		"subscribe"      : "Подписаться на рассылку",
		"unsubscribe"    : "Подписаться на рассылку✔️",
	}


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
	case "start":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			var args = map[string]interface{}{"user":update,"menu":"main_menu"}
			go sendMessage1(args)
			fmt.Println("CurrentCommand:", Command)

		}
		break

	case "help":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "menu":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "mainNews":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			var args = map[string]interface{}{"user":update}
			sendMessage(args)

			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "lastNews":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			var args = map[string]interface{}{"user":update}
			sendMessage(args)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "news":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			var args = map[string]interface{}{"user":update,"menu":"news"}
			sendMessage(args)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "economic":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			var args = map[string]interface{}{"user":update,"menu":"single"}
			sendMessage(args)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "accidents":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			var args = map[string]interface{}{"user":update,"menu":"single"}
			sendMessage(args)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "sports":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			var args = map[string]interface{}{"user":update,"menu":"single"}
			sendMessage(args)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "tech":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			var args = map[string]interface{}{"user":update,"menu":"single"}
			sendMessage(args)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "life":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			var args = map[string]interface{}{"user":update,"menu":"single"}
			sendMessage(args)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "culture":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			var args = map[string]interface{}{"user":update,"menu":"single"}
			sendMessage(args)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "articles":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			var args = map[string]interface{}{"user":update,"menu":"single"}
			sendMessage(args)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "back":
		NewMethod = func(update conf.Update) {
			LastCommand := GetCommand(update)
			var args = map[string]interface{}{"user":update,"menu":conf.Back[LastCommand]}
			if conf.Back[LastCommand]!="menu"{
				setCommand(update,conf.Back[LastCommand])
			}
			sendMessage(args)
			fmt.Println("CurrentCommand:", conf.Back[LastCommand])
			fmt.Println("CurrentCommand:", LastCommand)
		}
		break

	case "showMor":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "showPrev":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "clearCache":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "battleopinions":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			var args = map[string]interface{}{"user":update,"menu":"single"}
			sendMessage(args)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "opinion":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			var args = map[string]interface{}{"user":update,"menu":"single"}
			sendMessage(args)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "promises":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			var args = map[string]interface{}{"user":update,"menu":"single"}
			sendMessage(args)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "conference":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			var args = map[string]interface{}{"user":update,"menu":"single"}
			sendMessage(args)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "fotoarchive":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			var args = map[string]interface{}{"user":update,"menu":"single"}
			sendMessage(args)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "userlist":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "subscribe":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "unsubscribe":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	default:
		NewMethod = func(update conf.Update) {
			var val= GetCommand(update)
			fmt.Println("LastCommand:", val)
		}

	}

	return NewMethod

}

func sendMessage(args map[string]interface{})  {

	start := int (time.Now().UnixNano())

	user := args["user"].(conf.Update)

	method := fmt.Sprintf(conf.APIEndpoint, conf.BOT_TOKEN, "sendMessage")
	form := url.Values{}
	form.Add("chat_id", strconv.Itoa(user.Message.From.ID))
	form.Add("text", user.Message.Text)

	if args["menu"]!=nil {
		MenuName := args["menu"].(string)
		menu, _ := json.Marshal(getMenu(MenuName))
		form.Add("reply_markup", string(menu) )
	}


	req, _ := http.NewRequest("GET", method ,  bytes.NewBufferString(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Body.Close()

	//Отправка сообщения
	client := &http.Client{}
	resp, _ := client.Do(req)
	fmt.Println(resp.Status)

	end := int (time.Now().UnixNano())
	fmt.Println(end-start)

}

func sendMessage1(args map[string]interface{})  {
	start := int (time.Now().UnixNano())

	user := args["user"].(conf.Update)

	method := fmt.Sprintf(conf.APIEndpoint, conf.BOT_TOKEN, "sendMessage")
	form := url.Values{}
	form.Add("chat_id", strconv.Itoa(user.Message.From.ID))
	form.Add("text", user.Message.Text)

	if args["menu"]!=nil {
		MenuName := args["menu"].(string)
		menu, _ := json.Marshal(getMenu(MenuName))
		form.Add("reply_markup", string(menu) )
	}


	req, _ := http.NewRequest("GET", method ,  bytes.NewBufferString(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Body.Close()

	//Отправка сообщения
	client := &http.Client{}
	resp, _ := client.Do(req)
	fmt.Println(resp.Status)

	end := int (time.Now().UnixNano())
	fmt.Println(end-start)

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


func NewKeyboardButton(text string) conf.KeyboardButton {
	return conf.KeyboardButton{
		Text: text,
	}
}

func getMenu(MenuName string) conf.ReplyKeyboardMarkup  {

	menu :=conf.Menu[MenuName]
	var keyboard [][]conf.KeyboardButton
	for _,row := range menu {
		var rows[]conf.KeyboardButton
		for _,cel := range row {
			rows = append(rows, NewKeyboardButton(cel))

		}
		keyboard = append(keyboard, rows)
	}
	return conf.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard:       keyboard,
	}
}

