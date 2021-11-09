package g

type GlobalConfig struct {
	Debug         bool              `json:"debug"`
	LogLevel   		string			`json:"loglevel"`
	LogFile 		string			`json:"logfile"`
}