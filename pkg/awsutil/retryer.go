package awsutil

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/client"
)

type CustomRetryer struct {
	client.DefaultRetryer
}

func NewCustomRetryer() CustomRetryer {
	return CustomRetryer{
		DefaultRetryer: client.DefaultRetryer{
			NumMaxRetries:    10,
			MinRetryDelay:    1000 * time.Millisecond,
			MinThrottleDelay: 2000 * time.Millisecond,
			MaxRetryDelay:    500 * time.Second,
			MaxThrottleDelay: 1000 * time.Second,
		},
	}
}
