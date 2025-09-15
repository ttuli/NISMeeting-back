package Gzap

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type MultiFileLogger struct {
	mu      sync.RWMutex
	loggers map[string]*zap.SugaredLogger
	keyToFile map[string]string
}

func NewMultiFileLogger() *MultiFileLogger {
	return &MultiFileLogger{
		loggers:   make(map[string]*zap.SugaredLogger),
		keyToFile: make(map[string]string),
	}
}

func (m *MultiFileLogger) RegisterKey(key, filename string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	m.keyToFile[key] = filename
	
	if _, exists := m.loggers[key]; exists {
		m.loggers[key].Sync()
		delete(m.loggers, key)
	}
}

func (m *MultiFileLogger) getOrCreateLogger(key string) *zap.SugaredLogger {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	if logger, exists := m.loggers[key]; exists {
		return logger
	}
	
	filename, exists := m.keyToFile[key]
	if !exists {
		filename = key + ".log"
	}
	
	logger := m.createLogger(filename)
	m.loggers[key] = logger
	
	return logger
}

func (m *MultiFileLogger) createLogger(filename string) *zap.SugaredLogger {

	writer := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    100,  // MB
		MaxBackups: 3,
		MaxAge:     28,   // days
		Compress:   true,
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(writer),
		zapcore.InfoLevel,
	)
	
	logger := zap.New(core, zap.AddCaller())
	return logger.Sugar()
}

func (m *MultiFileLogger) Write(key string, msg string, args ...interface{}) {
	logger := m.getOrCreateLogger(key)
	logger.Infof(msg, args...)
}


func (m *MultiFileLogger) WriteError(key string, msg string, args ...interface{}) {
	logger := m.getOrCreateLogger(key)
	logger.Errorf(msg, args...)
}

func (m *MultiFileLogger) WriteWarn(key string, msg string, args ...interface{}) {
	logger := m.getOrCreateLogger(key)
	logger.Warnf(msg, args...)
}

func (m *MultiFileLogger) WriteDebug(key string, msg string, args ...interface{}) {
	logger := m.getOrCreateLogger(key)
	logger.Debugf(msg, args...)
}

func (m *MultiFileLogger) GetLogger(key string) *zap.SugaredLogger {
	return m.getOrCreateLogger(key)
}

func (m *MultiFileLogger) Close() {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	for _, logger := range m.loggers {
		logger.Sync()
	}
}

var globalLogger = NewMultiFileLogger()

// RegisterKey 全局注册函数
func RegisterKey(key, filename string) {
	globalLogger.RegisterKey(key, filename)
}

func Write(key string, msg string, args ...interface{}) {
	globalLogger.Write(key, msg, args...)
}

func WriteError(key string, msg string, args ...interface{}) {
	globalLogger.WriteError(key, msg, args...)
}

func WriteWarn(key string, msg string, args ...interface{}) {
	globalLogger.WriteWarn(key, msg, args...)
}

func WriteDebug(key string, msg string, args ...interface{}) {
	globalLogger.WriteDebug(key, msg, args...)
}

func GetLogger(key string) *zap.SugaredLogger {
	return globalLogger.GetLogger(key)
}

func Close() {
	globalLogger.Close()
}