package security

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/gin-gonic/gin"
)

type SMTPInterface interface {
	SendConfirmEmail(username, email string, cache caching.CachingInterface) error
	SendChangePassword(username, email string, cache caching.CachingInterface) error
	SendEmail(email string, message []byte) error
}

var ProductionSMTPInterface SMTPInterface = &RealSMTP{smtp: initSMTP()}

type RealSMTP struct {
	smtp smtp.Auth
}

func initSMTP() smtp.Auth {
	if gin.Mode() == gin.ReleaseMode {
		return smtp.PlainAuth(
			"",
			config.Config.SmtpUsername,
			config.Config.SmtpPassword,
			config.Config.SmtpHost,
		)
	} else {
		return nil
	}
}

func (Smtp *RealSMTP) SendConfirmEmail(username, email string, cache caching.CachingInterface) error {
	var confirmationLink string = GenerateLink()
	confirmEmailMessage := []byte(fmt.Sprintf(
		"Здравствуйте, %s!\n\n"+
			"Спасибо за регистрацию на нашем сайте. Чтобы активировать вашу учетную запись, пожалуйста, подтвердите свой адрес электронной почты, перейдя по следующей ссылке:\n\n"+
			"%s\n\n"+
			"Если вы не регистрировались на нашем сайте, просто проигнорируйте это сообщение.\n\n"+
			"С уважением,\n"+
			"ZoomerOk",
		username, confirmationLink,
	))
	cache.SetCacheConfirmationLink(username, confirmationLink)

	var subject string = "Подтверждение электронной почты ZoomerOk"
	headers := []byte("From: " + config.Config.SmtpUsername + "\n" +
		"To: " + email + "\n" +
		"Subject: " + subject + "\n\n")
	message := append(headers, confirmEmailMessage...)
	return Smtp.SendEmail(email, message)
}

func (Smtp *RealSMTP) SendChangePassword(username, email string, cache caching.CachingInterface) error {
	var resetLink string = GenerateLink()

	changePasswordMessage := []byte(fmt.Sprintf(
		"Здравствуйте, %s!\n\n"+
			"Мы получили запрос на сброс пароля для вашей учетной записи. Чтобы установить новый пароль, перейдите по следующей ссылке:\n\n"+
			"%s\n\n"+
			"Если вы не запрашивали сброс пароля, просто проигнорируйте это сообщение.\n\n"+
			"С уважением,\n"+
			"ZoomerOk",
		username, resetLink,
	))
	cache.SetCacheResetLink(username, resetLink)

	var subject string = "Подтверждение смены пароля ZoomerOk"
	headers := []byte("From: " + config.Config.SmtpUsername + "\n" +
		"To: " + email + "\n" +
		"Subject: " + subject + "\n\n")
	message := append(headers, changePasswordMessage...)
	return Smtp.SendEmail(email, message)
}

func GenerateLink() string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, config.Config.GenerateLinkLength)
	for i := range b {
		b[i] = config.Config.GenerateLinkCharset[seededRand.Intn(len(config.Config.GenerateLinkCharset))]
	}
	return string(b)
}

func (Smtp *RealSMTP) SendEmail(email string, message []byte) error {
	err := smtp.SendMail(config.Config.SmtpHost+":"+config.Config.SmtpPort, Smtp.smtp, config.Config.SmtpUsername, []string{email}, message)
	if err != nil {
		return err
	}

	return nil
}
