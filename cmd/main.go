package main

import (
	smtpSrv "smtp/app"
	"smtp/pkg/logger"
)

func main() {
	logger.Init()
	smtpSrv.Start()
}
