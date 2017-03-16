package conf


const (
	BOT_TOKEN string = "371139318:AAHYqiYcvAp4bzJrWJbh5K2BiRcH3KKpYiU"

	APIEndpoint string = "https://api.telegram.org/bot%s/%s"

)

var Menu = map[string][][]string{
	"main_menu": {
		{"Главные новости", "Последние новости"},
		{"Новости"},
		{"Статьи", "Битва мнений", "Блоги и мнения"},
		{"Обещания", "Интервью", "Фотогалерея"},
	},
	"news": {
		{"⏮Предыдущая", "Следующая⏭", },
		{"Подписаться на рассылку" },
		{"Экономика", "Спорт", "Технологии", },
		{"Происшествия", "Жизнь", "Культура", },
		{"Назад" },
	},
	"single": {
		{"⏮Предыдущая", "Следующая⏭"},
		{"Подписаться на рассылку"},
		{"Назад"},
	},

}

var Back = map[string]string{
	"articles":       "main_menu",
	"battleopinions": "main_menu",
	"opinion":        "main_menu",
	"promises":       "main_menu",
	"conference":     "main_menu",
	"fotoarchive":    "main_menu",
	"news":           "main_menu",
	"economic":       "news",
	"accidents":      "news",
	"sports":         "news",
	"tech":           "news",
	"life":           "news",
	"culture":        "news",
}

func init() {

}



