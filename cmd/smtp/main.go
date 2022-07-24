package main

import (
	"smtp/pkg/logger"
	smtpSrv "smtp/services/smtp"
)

func main() {
	logger.Init()
	smtpSrv.Start()
}
