package tgbotapi

import (
	"fmt"
	"time"
)

type (
	ApiError struct {
		Message string
	}

	GetUpdatesFailed struct {
		Message string
		Timeout time.Duration
	}

	HandlerMismatch struct {
		Got      UpdateType
		Expected UpdateType
	}

	InvalidType struct {
		Got      any
		Expected any
	}
)

func (e ApiError) Error() string {
	return fmt.Sprintf("ERROR: api error: %s", e.Message)
}

func (e GetUpdatesFailed) Error() string {
	return fmt.Sprintf("ERROR: getUpdates failed: %s. Retrying in %.1f seconds...", e.Message, e.Timeout.Seconds())
}

func (e HandlerMismatch) Error() string {
	return fmt.Sprintf("WARNING: handler mismatch: got update of type %d, expected %d", e.Got, e.Expected)
}

func (e InvalidType) Error() string {
	return fmt.Sprintf("ERROR: invalid value type: got %s, expected %s", e.Got, e.Expected)
}
