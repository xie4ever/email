package email

import (
	"errors"
	"fmt"
	"time"

	"github.com/badoux/checkmail"
	"github.com/go-redis/redis/v7"
)

const (
	emailAddressKey        = "email:address:%s"
	addressValid           = 1
	addressNotValid        = 2
	addressValidDuration   = 30 * 24 * time.Hour
	addressInValidDuration = 24 * time.Hour
)

func getKey(address string) string {
	return fmt.Sprintf(emailAddressKey, address)
}

func setValidAddress(client *redis.Client, key string) error {
	return client.Set(key, addressValid, addressValidDuration).Err()
}

func setInValidAddress(client *redis.Client, key string) error {
	return client.Set(key, addressNotValid, addressInValidDuration).Err()
}

func (s *Sender) isValidAddress(address string) error {
	if err := checkmail.ValidateFormat(address); err != nil {
		return err
	}

	if s.client == nil {
		return checkmail.ValidateHost(address)
	}

	key := getKey(address)
	val, err := s.client.Get(key).Int()
	switch {
	case err == nil:
		if val == addressValid {
			return nil
		} else {
			return checkmail.ErrUnresolvableHost
		}
	case errors.Is(err, redis.Nil):
		if err := checkmail.ValidateHost(address); err != nil {
			_ = setInValidAddress(s.client, key)
			return err
		}
		_ = setValidAddress(s.client, key)
		return nil
	default:
		return err
	}
}
