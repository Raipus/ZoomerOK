package security

import (
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/caching"
	"github.com/Raipus/ZoomerOK/account/pkg/config"
)

func SendConfirmEmail(username, confirmationLink string) []byte {
	confirmEmailMessage := []byte(fmt.Sprintf(
		"Здравствуйте, %s!\n\n"+
			"Спасибо за регистрацию на нашем сайте. Чтобы активировать вашу учетную запись, пожалуйста, подтвердите свой адрес электронной почты, перейдя по следующей ссылке:\n\n"+
			"%s\n\n"+
			"Если вы не регистрировались на нашем сайте, просто проигнорируйте это сообщение.\n\n"+
			"С уважением,\n"+
			"ZoomerOk",
		username, confirmationLink,
	))
	caching.SetCacheConfirmationLink(username, confirmationLink)
	return confirmEmailMessage
}

func SendChangePassword(username, resetLink string) []byte {
	changePasswordMessage := []byte(fmt.Sprintf(
		"Здравствуйте, %s!\n\n"+
			"Мы получили запрос на сброс пароля для вашей учетной записи. Чтобы установить новый пароль, перейдите по следующей ссылке:\n\n"+
			"%s\n\n"+
			"Если вы не запрашивали сброс пароля, просто проигнорируйте это сообщение.\n\n"+
			"С уважением,\n"+
			"ZoomerOk",
		username, resetLink,
	))
	caching.SetCacheResetLink(username, resetLink)
	return changePasswordMessage
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

func sendEmail(message string) error {
	err := smtp.SendMail(smtpHost+":"+smtpPort, nil)

	if err != nil {
		return fmt.Errorf("Ошибка при отправке сообщения: %s\n", err)
	}

	log.Println("Сообщение успешно отправлено!")
	return nil
}
