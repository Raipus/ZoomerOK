package security

import (
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/gin-gonic/gin"
)

type SMTPInterface interface {
	SendConfirmEmail(login, email string, cache caching.CachingInterface) error
	SendChangePassword(login, email string, cache caching.CachingInterface) error
	SendEmail(email string, message []byte) error
}

var ProductionSMTPInterface SMTPInterface = &RealSMTP{smtp: initSMTP()}

type RealSMTP struct {
	smtp smtp.Auth
}

func initSMTP() smtp.Auth {
	if gin.Mode() == gin.ReleaseMode {
		log.Println("Smtp initialized")
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

func (Smtp *RealSMTP) SendConfirmEmail(login, email string, cache caching.CachingInterface) error {
	code := GenerateLink()
	// var confirmationLink string = config.Config.FrontendLink + "/confirm_email/" + code
	/*confirmEmailMessage := []byte(fmt.Sprintf(
		"Здравствуйте, %s!\n\n"+
			"Спасибо за регистрацию на нашем сайте. Чтобы активировать вашу учетную запись, пожалуйста, подтвердите свой адрес электронной почты, перейдя по следующей ссылке:\n\n"+
			"%s\n\n"+
			"Если вы не регистрировались на нашем сайте, просто проигнорируйте это сообщение.\n\n"+
			"С уважением,\n"+
			"ZoomerOk",
		login, confirmationLink,
	))*/
	cache.SetCacheConfirmationLink(login, code)
	log.Println("code:", code)

	return fmt.Errorf("ewr")
	/*var subject string = "Подтверждение электронной почты ZoomerOk"
	headers := []byte("From: " + config.Config.SmtpUsername + "\n" +
		"To: " + email + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-Version: 1.0\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"\n" +
		"\n")
	message := append(headers, confirmEmailMessage...)
	return Smtp.SendEmail(email, message)*/
}

func (Smtp *RealSMTP) SendChangePassword(login, email string, cache caching.CachingInterface) error {
	code := GenerateLink()
	var resetLink string = config.Config.FrontendLink + "/confirm_password/" + code

	changePasswordMessage := []byte(fmt.Sprintf(
		"Здравствуйте, %s!\n\n"+
			"Мы получили запрос на сброс пароля для вашей учетной записи. Чтобы установить новый пароль, перейдите по следующей ссылке:\n\n"+
			"%s\n\n"+
			"Если вы не запрашивали сброс пароля, просто проигнорируйте это сообщение.\n\n"+
			"С уважением,\n"+
			"ZoomerOk",
		login, resetLink,
	))
	cache.SetCacheResetLink(login, code)
	log.Println(code)

	var subject string = "Подтверждение смены пароля ZoomerOk"
	headers := []byte("From: " + config.Config.SmtpUsername + "\n" +
		"To: " + email + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-Version: 1.0\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"\n" +
		"\n")
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
