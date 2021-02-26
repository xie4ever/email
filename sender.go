package email

import "github.com/go-redis/redis/v7"

type config struct {
	host        string
	port        string
	address     string
	username    string
	password    string
	ignoreError bool // 如果为false，任何错误都会导致整体发送失败。如果为true，将把可以发的邮件尽可能发完。
	client      *redis.Client
}

// Sender ...
type Sender struct {
	config
}

// NewSender ...
func NewSender(setters ...Setter) (*Sender, error) {
	var err error
	s := &Sender{}
	for _, setter := range setters {
		if err := setter(&s.config); err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	return s, err
}
