package utils

import (
	"fmt"
	"github.com/jordan-wright/email"
	"iotvep/internal/app/config"
	"iotvep/pkg/logger"
	"net/smtp"
)

func SendEmail(to string, code string) bool {
	// 构造邮件的内容
	e := &email.Email{
		From:    fmt.Sprintf("%v <%v>", config.CONFIG.Email.Name, config.CONFIG.Email.Addr),
		To:      []string{to},
		Subject: "邮箱验证",
		Text:    []byte("【物联网虚拟仿真平台】注册验证码: " + code + ", 有效期 10 分钟"),
	}
	// 发送邮件
	err := e.Send(config.CONFIG.Email.Smtp, smtp.PlainAuth("", config.CONFIG.Email.Addr, config.CONFIG.Email.Password, config.CONFIG.Email.Host))
	if err != nil {
		logger.Error("验证码发送失败！", err)
		return false
	}
	return true
}
