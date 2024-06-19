package slack

import (
	"fmt"
	"time"

	"github.com/slack-go/slack"
)

func retryOnRateLimit(operation func() (interface{}, error)) (interface{}, error) {
	for {
		result, err := operation()
		if rateLimitedError, ok := err.(*slack.RateLimitedError); ok {
			fmt.Printf("Rate limited. Retrying after %v seconds...\n", rateLimitedError.RetryAfter)
			time.Sleep(rateLimitedError.RetryAfter)
			continue
		} else if err != nil {
			return nil, err
		}
		return result, nil
	}
}