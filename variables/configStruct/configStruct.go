package configStruct

var Config *configStruct

type configStruct struct {
	Token     string `json: "Token"`
	BotPrefix string `json: "BotPrefix"`
}
