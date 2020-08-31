package logger

import "go.uber.org/zap"

var (
	OptimizedLog *zap.Logger        = nil
	Log          *zap.SugaredLogger = nil
)

func Init() (err error) {
	OptimizedLog, err = zap.NewProduction()
	if err != nil {
		return
	}

	Log = OptimizedLog.Sugar()
	return
}

func IsInitialized() bool {
	return OptimizedLog != nil && Log != nil
}
