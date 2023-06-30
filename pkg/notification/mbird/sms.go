package mbird

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/models"
	"github.com/alexandredsa/2fa-poc-api/pkg/applog"
	"github.com/alexandredsa/2fa-poc-api/pkg/appredis"
	messagebird "github.com/messagebird/go-rest-api"
	"github.com/messagebird/go-rest-api/verify"
)

type SMSNotifier struct {
	client      *messagebird.Client
	logger      *applog.Logger
	redisClient appredis.RedisClientWrapper
}

// NewSMSNotifier creates a new instance of SMSNotifier
// with the provided MessageBird client.
func NewSMSNotifier(
	client *messagebird.Client,
	redisClient appredis.RedisClientWrapper,
) *SMSNotifier {
	return &SMSNotifier{
		client:      client,
		logger:      applog.NewLogger("SMSNotifier"),
		redisClient: redisClient,
	}
}

func (s *SMSNotifier) SendVerificationCode(ctx context.Context, user models.User) error {
	v, err := verify.Create(s.client, user.Email, &verify.Params{
		Originator: os.Getenv("NOTIFICATION_SMS_ORIGINATOR"),
		Timeout:    s.getTimeoutFromEnv(),
	})
	if err != nil {
		return err
	}

	s.logger.Infof("Verification: %v\n", v)

	key := s.formatKey(user.ID)
	err = s.redisClient.Set(ctx, key, v.ID, 0).Err()

	return err
}

func (s *SMSNotifier) VerifyCode(ctx context.Context, user models.User, code string) (bool, error) {
	key := s.formatKey(user.ID)
	vID, err := s.redisClient.Get(ctx, key).Result()
	if err != nil {
		return false, err
	}

	v, err := verify.VerifyToken(
		s.client,
		vID,
		code,
	)

	s.logger.Infof("Verification: %v\n", v)

	return v.Status == string(StatusVerified), err
}

func (s *SMSNotifier) formatKey(k string) string {
	return fmt.Sprintf("sms_notifier_%s", k)
}

const (
	smsDefaultTimeout int = 120
)

func (s *SMSNotifier) getTimeoutFromEnv() int {
	timeoutStr := os.Getenv("SMS_TIMEOUT")
	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		s.logger.Infof(
			"Invalid or missing timeout value from environment variables."+
				"Using default timeout: %d", smsDefaultTimeout)

		return smsDefaultTimeout
	}
	return timeout
}
