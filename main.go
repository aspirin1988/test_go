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

	Commands := map[string]string{
		"start"          : "/start",
		"help"           : "/help",
		"menu"           : "Ⓜ️Меню",
		"mainNews"       : "Главные новости",
		"news"           : "Новости",
		"economic"       : "Экономика",
		"accidents"      : "Происшествия",
		"sports"         : "Спорт в Казахстане",
		"tech"           : "Технологии",
		"life"           : "Жизнь",
		"culture"        : "Культура",
		"lastNews"       : "Последние новости",
		"articles"       : "Статьи",
		"back"           : "Назад",
		"showMor"        : "Показать еще",
		"showPrev"       : "Показать предыдущую",
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
			sendMessage(update)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "news":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "economic":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "accidents":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "sports":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "tech":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "life":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "culture":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "lastNews":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "articles":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "back":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			fmt.Println("CurrentCommand:", Command)
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
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "opinion":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "promises":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "conference":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
			fmt.Println("CurrentCommand:", Command)
		}
		break

	case "fotoarchive":
		NewMethod = func(update conf.Update) {
			setCommand(update, Command)
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

func sendMessage( user conf.Update)  {

	method := fmt.Sprintf(conf.APIEndpoint, conf.BOT_TOKEN, "sendMessage")
	form := url.Values{}
	form.Add("chat_id", strconv.Itoa(user.Message.From.ID))
	form.Add("text", user.Message.Text)

	row :=NewKeyboardButtonRow(NewKeyboardButton("Главные новости"),NewKeyboardButton("test"),NewKeyboardButton("test"))

	keyboard:= NewReplyKeyboard(row)

	fmt.Println(keyboard)
	json1,_ := json.Marshal(keyboard)
	form.Add("reply_markup", string(json1) )


	req, _ := http.NewRequest("POST", method ,  bytes.NewBufferString(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	defer req.Body.Close()

	//Отправка сообщения
	client := &http.Client{}
	resp, _ := client.Do(req)
	fmt.Println(resp.Status)
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

// NewKeyboardButtonRow creates a row of keyboard buttons.
func NewKeyboardButtonRow(buttons ...conf.KeyboardButton) []conf.KeyboardButton {
	var row []conf.KeyboardButton

	row = append(row, buttons...)

	return row
}

// NewReplyKeyboard creates a new regular keyboard with sane defaults.
func NewReplyKeyboard(rows ...[]conf.KeyboardButton) conf.ReplyKeyboardMarkup {
	var keyboard [][]conf.KeyboardButton

	keyboard = append(keyboard, rows...)

	return conf.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard:       keyboard,
	}
}


