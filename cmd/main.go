package main

import (
	"smtp/pkg/logger"
	smtpSrv "smtp/app"
)

func main() {
	logger.Init()
	smtpSrv.Start()
}
