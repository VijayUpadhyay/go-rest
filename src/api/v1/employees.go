package v1

import (
	"github.com/avast/retry-go"
	"github.com/cenkalti/backoff"
	"github.com/sirupsen/logrus"
)

func GetEmployee() {
	logrus.Debugf("Inside GetEmployee")

}

func CreateEmployee() {
	logrus.Debugf("Inside CreateEmployee")
}

func BackOffOnError() {
	// An operation that may fail.
	operation := func() error {
		return nil // or an error
	}

	err := backoff.Retry(operation, backoff.NewExponentialBackOff())
	if err != nil {
		// Handle error.
		return
	}

	// Operation is successful.
}

func AvastRetry() {
	retry.Attempts(3)
}
