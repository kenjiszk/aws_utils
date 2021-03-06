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

func railsConsoleValidator() error {
        if os.Getenv("EC2_SERVER_NAME") == "" {
                return errors.New("EC2_SERVER_NAME is not set.")
        }
	if os.Getenv("ECS_SERVICE_NAME") == "" {
		return errors.New("ECS_SERVICE_NAME is not set.")
	}
	return nil
}
