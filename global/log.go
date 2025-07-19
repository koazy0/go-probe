package global

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"regexp"
	"time"
)

var Log *zap.Logger

const (
	BuleColor   = "\033[34m"
	YellowColor = "\033[33m"
	RedColor    = "\033[31m"
	ResetColor  = "\033[0m"
)

// logEncoder 时间分片和level分片同时做
// 组合一个zap的Encoder继承接口方法，与此同时额外增加字段，做到日志双写+分片
type logEncoder struct {
	zapcore.Encoder
	errFile     *os.File
	file        *os.File
	currentDate string
}

func (e *logEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	// 先调用原始的 EncodeEntry 方法生成日志行
	buff, err := e.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}
	data := buff.String()
	buff.Reset()
	buff.AppendString("[gateway] " + data)
	data = buff.String()

	// logdata去除了转义字符
	re := regexp.MustCompile(`\033\[[\d;]*m`)
	logdata := re.ReplaceAllString(data, "")

	// 时间分片，以天为间隔进行分片
	now := time.Now().Format("2006-01-02")

	//更新时间到最新
	if e.currentDate != now {
		os.MkdirAll(fmt.Sprintf("logs/%s", now), 0666)
		// 时间不同，先创建目录
		name := fmt.Sprintf("logs/%s/out.log", now)
		file, _ := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		e.file = file
		e.currentDate = now
	}

	switch entry.Level {
	//按等级进行分片，单独把error拉出来写入
	case zapcore.ErrorLevel:
		if e.errFile == nil {
			name := fmt.Sprintf("logs/%s/err.log", now)
			file, _ := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
			e.errFile = file
		}
		e.errFile.WriteString(logdata)
	}

	if e.currentDate == now {
		e.file.WriteString(logdata)
	}
	return buff, nil
}

func InitLog() {
	// 使用 zap 的 NewDevelopmentConfig 快速配置
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // 替换时间格式化方式
	cfg.EncoderConfig.EncodeLevel = myEncodeLevel
	// 创建自定义的 Encoder
	encoder := &logEncoder{
		Encoder: zapcore.NewConsoleEncoder(cfg.EncoderConfig), // 使用 Console 编码器
	}
	// 创建 Core
	core := zapcore.NewCore(
		encoder,                    // 使用自定义的 Encoder
		zapcore.AddSync(os.Stdout), // 输出到控制台
		zapcore.InfoLevel,          // 设置日志级别
	)
	// 创建 Logger
	logger := zap.New(core, zap.AddCaller())

	zap.ReplaceGlobals(logger)
	Log = logger

	//定期删除10天之前的日志
	go func() {
		for {
			cutoff := time.Now().Add(-10 * 24 * time.Hour).Format("2006-01-02")
			entries, _ := os.ReadDir("./logs")
			for _, entry := range entries {
				if entry.Name() < cutoff {
					os.RemoveAll(path.Join("./logs/" + entry.Name()))
				}
			}
			<-time.After(12 * time.Hour)
		}
	}()
}

func myEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch level {
	case zapcore.InfoLevel:
		enc.AppendString(BuleColor + "INFO" + ResetColor)
	case zapcore.WarnLevel:
		enc.AppendString(YellowColor + "WARN" + ResetColor)
	case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		enc.AppendString(RedColor + "ERROR" + ResetColor)
	default:
		enc.AppendString(level.String())
	}
}
