package nvidiasmi

import (
	"github.com/c3sr/config"
	logger "github.com/c3sr/logger"
)

var (
	log         = logger.New().WithField("pkg", "nvidia-smi")
	initialized = make(chan struct{}, 1)
)

func Wait() {
	<-initialized
}

func init() {
	config.BeforeInit(func() {
		Init()
	})
	config.AfterInit(func() {
		log = logger.New().WithField("pkg", "nvidia-smi")
	})
}
