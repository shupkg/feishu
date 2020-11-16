package log

type logger interface {
	Error(msg ...interface{})
	Errorf(msg string, args ...interface{})
	Info(msg ...interface{})
	Infof(msg string, args ...interface{})
}

var log logger

func Init(l logger) {
	log = l
}

func Error(msg ...interface{}) {
	if log != nil {
		log.Error(msg...)
	}
}

func ErrorF(msg string, args ...interface{}) {
	if log != nil {
		log.Errorf(msg, args...)
	}
}

func Info(msg ...interface{}) {
	if log != nil {
		log.Info(msg...)
	}
}

func InfoF(msg string, args ...interface{}) {
	if log != nil {
		log.Infof(msg, args...)
	}
}
