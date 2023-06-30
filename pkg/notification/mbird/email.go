package mbird

import (
	"context"
	"fmt"
	"os"

	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/models"
	"github.com/alexandredsa/2fa-poc-api/pkg/applog"
	"github.com/alexandredsa/2fa-poc-api/pkg/appredis"
	"github.com/alexandredsa/2fa-poc-api/pkg/notification"
	messagebird "github.com/messagebird/go-rest-api"
	"github.com/messagebird/go-rest-api/conversation"
)

type EmailParams struct {
	channelID string
}

type EmailNotifier struct {
	client        *messagebird.Client
	logger        *applog.Logger
	params        *EmailParams
	codeGenerator notification.Generator
	redisClient   appredis.RedisClientWrapper
}

// NewEmailNotifier creates a new instance of EmailNotifier
// with the provided MessageBird client and Redis client.
func NewEmailNotifier(
	client *messagebird.Client,
	generator notification.Generator,
	redisClient appredis.RedisClientWrapper,
) *EmailNotifier {
	return &EmailNotifier{
		client:        client,
		logger:        applog.NewLogger("EmailNotifier"),
		params:        &EmailParams{channelID: os.Getenv("NOTIFICATION_EMAIL_CHANNEL_ID")},
		codeGenerator: generator,
		redisClient:   redisClient,
	}
}

const (
	emailCodeLength = 4
)

func (e *EmailNotifier) formatKey(k string) string {
	return fmt.Sprintf("email_notifier_%s", k)
}

func (e *EmailNotifier) SendVerificationCode(ctx context.Context, user models.User) error {
	code := e.codeGenerator.NewCode(emailCodeLength)

	_, err := conversation.Start(e.client, &conversation.StartRequest{
		ChannelID: e.params.channelID,
		Content: &conversation.MessageContent{
			Text: fmt.Sprintf("Confirm your e-mail address "+
				"using the following code\n<b>%s</b>", code),
		},
		To:   user.Email,
		Type: conversation.MessageTypeText,
	})
	if err != nil {
		return err
	}

	key := e.formatKey(user.ID)
	// Set a key in Redis with userID as the key and code as the value
	err = e.redisClient.Set(ctx, key, code, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (e *EmailNotifier) VerifyCode(ctx context.Context, user models.User, code string) (bool, error) {
	key := e.formatKey(user.ID)
	// Retrieve the stored code from Redis using the userID as the key
	storedCode, err := e.redisClient.Get(ctx, key).Result()
	if err != nil {
		return false, err
	}

	// Compare the provided code with the stored code
	if code == storedCode {
		// Code matches, delete the stored code from Redis
		err := e.redisClient.Del(ctx, user.ID).Err()
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}
