package logx

import (
	"encoding/json"
	"strconv"
	"strings"
)

type LogxConfig struct {
	FilePath        string
	DevMode         bool //if true all log of debug level will be outputted or will be ignored,the default value is true
	ToDifferentFile bool

	Maxbuffer     string
	MaxbufferInt  int
	DisableBuffer bool
}

var jsonConfig *LogxConfig

//priority of configure from file is higher than flags
func LoadConfJson(conf []byte) *LogxConfig {
	jsonConfig = new(LogxConfig)
	jsonConfig.DevMode = true
	err := json.Unmarshal(conf, &jsonConfig)
	if err != nil {
		panic(err.Error())
	}
	jsonConfig.MaxbufferInt = unitParse(jsonConfig.Maxbuffer)
	if jsonConfig.MaxbufferInt > maxDefaultBufferSize {
		jsonConfig.MaxbufferInt = maxDefaultBufferSize
	}
	return jsonConfig
}

//unitParse implements convert string amount to amount of byte, e.g 1kb-> 1024
const (
	unitKb   = 'K'
	unitMb   = 'M'
	unitByte = 'B'
)

func unitParse(one string) int {
	i, err := strconv.Atoi(one)
	if err == nil {
		return i
	}
	one = strings.ToUpper(strings.TrimSpace(one))
	if len(one) < 1 {
		return 0
	}
	oneLen := len(one)
	if one[oneLen-2] == unitKb {
		i = 1 << 10
		one = one[:oneLen-2]

	} else if one[oneLen-2] == unitMb {
		i = 1 << 20
		one = one[:oneLen-2]
	} else if one[oneLen-4] == unitByte {
		i = 1
		one = one[:oneLen-4]
	}
	one = strings.TrimSpace(one)
	amount, err := strconv.Atoi(one)
	if err != nil {
		return 0
	}
	return amount * i
}
