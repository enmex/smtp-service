package main

import (
	smtpSrv "smtp/app"
	"smtp/app/config"
	"smtp/pkg/logger"
)

func main() {
	logger.Init()
	config.Init()
	smtpSrv.Start()
}
