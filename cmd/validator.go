package cmd

import (
	"errors"
	"os"
)

func awsEnvValidator() error {
	if os.Getenv("AWS_DEFAULT_REGION") == "" {
		return errors.New("AWS_DEFAULT_REGION is not set.")
	}
	if os.Getenv("AWS_ACCESS_KEY_ID") == "" {
		return errors.New("AWS_ACCESS_KEY_ID is not set.")
	}
	if os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
		return errors.New("AWS_SECRET_ACCESS_KEY is not set.")
	}
	return nil
}
