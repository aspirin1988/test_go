package conf


const (
	BOT_TOKEN string = "225366019:AAGDMEeCVLARyx8A5g5BAUHqjMAPjR7AFOQ"
)

var Command map[string]string


func init() {
	Command["Start"]="/start"
	Command["MainNews"]="Главные новости"
	Command["LastNews"]="Последние новости"
	Command["News"]="Новости"
	Command["Article"]="Статьи"
	Command["OpinionBatle"]="Битва мнений"
	Command["Opinions"]="Блоги и мнения"
}