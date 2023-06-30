package notifier

import (
	"context"
	"os"

	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/models"
	"github.com/alexandredsa/2fa-poc-api/pkg/appredis"
	"github.com/alexandredsa/2fa-poc-api/pkg/notification"
	"github.com/alexandredsa/2fa-poc-api/pkg/notification/mbird"
	messagebird "github.com/messagebird/go-rest-api"
)

type Component string

const (
	EmailComponent Component = "email"
	SMSComponent   Component = "sms"
)

type Notifier interface {
	SendVerificationCode(ctx context.Context, user models.User) error
	VerifyCode(ctx context.Context, user models.User, code string) (bool, error)
}

type NotifierFactory struct {
	redisClient appredis.RedisClientWrapper
}

func NewNotifierFactory(redisClient appredis.RedisClientWrapper) NotifierFactory {
	return NotifierFactory{redisClient: redisClient}
}

func (n *NotifierFactory) NewNotifier(component string) Notifier {
	client := messagebird.New(os.Getenv("MESSAGEBIRD_ACCESS_KEY"))
	generator := notification.NewGenerator()

	switch component {
	case string(EmailComponent):
		return mbird.NewEmailNotifier(client, generator, n.redisClient)
	case string(SMSComponent):
		return mbird.NewSMSNotifier(client, n.redisClient)
	}

	return nil
}
