package target

import (
	"time"

	"github.com/foxcpp/maddy/log"
	"github.com/foxcpp/maddy/module"
)

func DeliveryLogger(l log.Logger, msgMeta *module.MsgMetadata) log.Logger {
	out := l.Out
	if out == nil {
		out = log.DefaultLogger.Out
	}

	return log.Logger{
		Out: log.FuncOutput(func(t time.Time, debug bool, str string) {
			out.Write(t, debug, str+" (msg ID = "+msgMeta.ID+")")
		}, out.Close),
		Name:  l.Name,
		Debug: l.Debug,
	}
}
