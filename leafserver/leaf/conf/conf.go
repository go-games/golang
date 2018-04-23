package conf

import "strconv"

var (
	LenStackBuf = 4096

	// log
	LogLevel string
	LogPath  string
	LogFlag  int

	// console
	ConsolePort   int
	ConsolePrompt string = "Leaf# "
	ProfilePath   string

	// cluster
	ListenAddr      string
	ConnAddrs       []string
	PendingWriteNum int

	//db
	dbusername string = "root"
	dbpassword string = "name"
	dbipport   string = "127.0.0.1:3306"
	dbname     string = "app"

	//redis
)
var appConfig map[string]string

func init() {
	appConfig = make(map[string]string)
	//TODO 可以从文件里面读取数据
	appConfig["dbusername"] = dbusername
	appConfig["dbpassword"] = dbpassword
	appConfig["dbipport"] = dbipport
	appConfig["dbname"] = dbname
}
func GetString(name string, defaultValue string) string {
	if value, ok := appConfig[name]; ok {
		return value
	} else {
		return defaultValue
	}
}
func GetInt(name string, defaultValue int) int {
	if value, ok := appConfig[name]; ok {
		if ret, err := strconv.Atoi(value); err != nil {
			panic(err)
		} else {
			return ret
		}
	}
	return defaultValue
}
