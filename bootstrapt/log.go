package bootstrapt

import (
	"flag"
	log "github.com/xiaomi-tc/log15"
)

var (
	logPath, logLevel string
)

func InitLog() {
	logPath = "edu.log"
	flag.StringVar(&logLevel, "l", "debug", "log level: debug, info, error")
	h, err := log.NetFileHandler(logPath, "edu", log.LogfmtFormat(), log.WithDstAddr("127.0.0.1:9999"))
	if err != nil {
		log.Error("log.NetFileHandler", "error", err)
		return
	}
	log.Root().SetHandler(h)
	switch logLevel {
	case "debug":
		log.SetOutLevel(log.LvlDebug)
	case "info":
		log.SetOutLevel(log.LvlInfo)
	case "error":
		log.SetOutLevel(log.LvlError)
	default:
		log.SetOutLevel(log.LvlDebug)
	}
}
