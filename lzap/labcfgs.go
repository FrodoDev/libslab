package lzap

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LabCfgs explore different kinds of configurations
func LabCfgs() {
	logs := makeCfgLoggers()
	fs := buildFields()
	for _, log := range logs {
		log.Debug("Lab cfg", fs...)
		log.Error("Lab cfg", fs...)
	}
}

func LabSample() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = myTimeEncoder
	cfg.Sampling.Initial = 5
	cfg.Sampling.Thereafter = 3
	cfg.Sampling.Hook = mySampleHook
	log := zap.Must(cfg.Build())

	n := 70
	for i := 1; i <= n; i++ {
		log.Debug("testsample", zap.Int("index", i))
		log.Info("testsample", zap.Int("index", i))
		log.Warn("testsample", zap.Int("index", i))
		time.Sleep(time.Millisecond * 50)
	}
}

func buildFields() (fields []zap.Field) {
	contents := make(map[string]interface{})
	contents["foo"] = "bar"
	contents["No"] = 123
	// contents["boolean"] = true
	// contents["float64"] = 3.14159
	// contents["other"] = []int32{1, 2, 3}
	contents["duration"] = 1*time.Hour + 14*time.Minute + 3*time.Second + 456*time.Millisecond
	// contents["interface"] = ?

	for k, v := range contents {
		switch v := v.(type) {
		case int:
			fields = append(fields, zap.Int(k, v))
		case string:
			fields = append(fields, zap.String(k, v))
		case bool:
			fields = append(fields, zap.Bool(k, v))
		case float64:
			fields = append(fields, zap.Float64(k, v))
		case time.Duration:
			fields = append(fields, zap.Duration(k, v))
		default:
			fields = append(fields, zap.Any(k, v))
		}
	}
	return fields
}

// makeCfgLoggers constructs a set of loggers, all derived from the default production
// configuration but with subtle individual variations. These variations are achieved
// by modifying fields in zap.Config and zapcore.EncoderConfig one by one.
func makeCfgLoggers() (logs []*zap.Logger) {
	cfg := zap.NewProductionConfig()
	cfg.Level.SetLevel(zap.DebugLevel)
	log := zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `begin with zap.NewProductionConfig(); cfg.Level.SetLevel(zap.DebugLevel)`)
	logs = append(logs, log)

	cfg = zap.NewProductionConfig()
	cfg.Level.SetLevel(zap.InfoLevel)
	cfg.EncoderConfig.NameKey = "caseName"
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `"cfg.Level.SetLevel(zap.InfoLevel); cfg.EncoderConfig.NameKey = "caseName"`)
	logs = append(logs, log)

	cfg.DisableCaller = true
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + "cfg.DisableCaller = true")
	logs = append(logs, log)

	cfg.DisableCaller = false
	cfg.DisableStacktrace = true
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + "cfg.DisableCaller = false; cfg.DisableStacktrace = true")
	logs = append(logs, log)

	cfg.Encoding = "console"
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.Encoding = "console"`) // case5
	logs = append(logs, log)

	cfg.EncoderConfig.ConsoleSeparator = "[ ]"
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.ConsoleSeparator = "[ ]"`) // case6
	logs = append(logs, log)

	cfg.Encoding = "json"
	cfg.OutputPaths = []string{"./log1.log", "./log2.log"}
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.Encoding = "json"; cfg.OutputPaths = []string{"./log1.log", "./log2.log"}`)
	logs = append(logs, log)

	cfg.OutputPaths = []string{"stderr"}
	cfg.ErrorOutputPaths = []string{"./errlog1.log"}
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.OutputPaths = []string{"stderr"}; cfg.ErrorOutputPaths = []string{"./errlog1.log"`)
	logs = append(logs, log)

	cfg.ErrorOutputPaths = []string{"stderr"}
	cfg.InitialFields = make(map[string]interface{})
	cfg.InitialFields["initField1"] = "hello"
	cfg.InitialFields["initField2"] = "zap"
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.ErrorOutputPaths = []string{"stderr"}; cfg.InitialFields = make(map[string]interface{}); cfg.InitialFields["initField1"] = "hello"; cfg.InitialFields["initField2"] = "zap"`)
	logs = append(logs, log)

	cfg.InitialFields = nil
	cfg.EncoderConfig.MessageKey = "userMsg"
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.InitialFields = nil; cfg.EncoderConfig.MessageKey = "userMsg"`)
	logs = append(logs, log)

	cfg.EncoderConfig.MessageKey = "msg"
	cfg.EncoderConfig.LevelKey = "lv"
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.MessageKey = "msg"; cfg.EncoderConfig.LevelKey = "lv"`)
	logs = append(logs, log)

	cfg.EncoderConfig.TimeKey = "logtime"
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.TimeKey = "logtime"`)
	logs = append(logs, log)

	cfg.EncoderConfig.TimeKey = "ts"
	cfg.EncoderConfig.CallerKey = "from"
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.TimeKey = "ts"; cfg.EncoderConfig.CallerKey = "from"`)
	logs = append(logs, log)

	cfg.EncoderConfig.FunctionKey = "funName"
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.FunctionKey = "funName"`)
	logs = append(logs, log)

	cfg.EncoderConfig.FunctionKey = ""
	cfg.DisableStacktrace = false
	cfg.EncoderConfig.StacktraceKey = "stack"
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.FunctionKey = ""; cfg.DisableStacktrace = false; cfg.EncoderConfig.StacktraceKey = "stack"`)
	logs = append(logs, log)

	cfg.DisableStacktrace = true
	cfg.EncoderConfig.StacktraceKey = ""
	cfg.EncoderConfig.SkipLineEnding = true
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.DisableStacktrace = true; cfg.EncoderConfig.StacktraceKey = "" cfg.EncoderConfig.SkipLineEnding = true`)
	logs = append(logs, log)

	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `SkipLineEnding second part`)
	logs = append(logs, log)

	cfg.EncoderConfig.SkipLineEnding = false
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.SkipLineEnding = false`)
	logs = append(logs, log)

	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder`)
	logs = append(logs, log)

	cfg.EncoderConfig.EncodeTime = zapcore.EpochTimeEncoder
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.EncodeTime = zapcore.EpochTimeEncoder`)
	logs = append(logs, log)

	cfg.EncoderConfig.EncodeTime = zapcore.EpochMillisTimeEncoder
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.EncodeTime = zapcore.EpochMillisTimeEncoder`)
	logs = append(logs, log)

	cfg.EncoderConfig.EncodeTime = zapcore.EpochNanosTimeEncoder
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.EncodeTime = zapcore.EpochNanosTimeEncoder`)
	logs = append(logs, log)

	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder`)
	logs = append(logs, log)

	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder`)
	logs = append(logs, log)

	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder`)
	logs = append(logs, log)

	cfg.EncoderConfig.EncodeTime = myTimeEncoder
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.EncodeTime = myTimeEncoder`) // case26
	logs = append(logs, log)

	cfg.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.EncodeTime = zapcore.EpochTimeEncoder; cfg.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder`)
	logs = append(logs, log)

	cfg.EncoderConfig.EncodeDuration = zapcore.NanosDurationEncoder
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.EncodeDuration = zapcore.NanosDurationEncoder`)
	logs = append(logs, log)

	cfg.EncoderConfig.EncodeDuration = zapcore.MillisDurationEncoder
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.EncodeDuration = zapcore.MillisDurationEncoder`)
	logs = append(logs, log)

	cfg.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder`)
	logs = append(logs, log)

	cfg.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder`)
	logs = append(logs, log)

	cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder`)
	logs = append(logs, log)

	cfg.EncoderConfig.EncodeName = myNameEncoder
	log = zap.Must(cfg.Build())
	log = log.Named(fmt.Sprintf("Case%d: ", len(logs)+1) + `cfg.EncoderConfig.EncodeName = myNameEncoder`)
	logs = append(logs, log)

	return logs
}

func myTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.999999999"))
}

func myNameEncoder(loggerName string, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("prefix " + loggerName)
}

func mySampleHook(en zapcore.Entry, sd zapcore.SamplingDecision) {
	switch sd {
	case zapcore.LogDropped:
		fmt.Println("call mySampleHook with LogDropped")
	default:
		fmt.Println("call mySampleHook with LogSampled")
	}
}
