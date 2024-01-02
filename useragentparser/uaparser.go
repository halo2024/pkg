package useragentparser

import (
	"github.com/one-piece-official/Nimbus/useragentparser"
)

var uaParser useragentparser.UserAgentParser //nolint:gochecknoglobals

func Init() {
	uaParser = useragentparser.NewUserAgentParser()
}

func Get() useragentparser.UserAgentParser {
	return uaParser
}
