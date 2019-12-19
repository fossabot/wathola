package config

import "time"

// Instance holds configuration values
var Instance = Config{
	Receiver: ReceiverConfig{
		Teardown: ReceiverTeardownConfig{
			Duration: 5 * time.Second,
		},
	},
	Sender: SenderConfig{
		Address:  "http://localhost:8080",
		Interval: 20 * time.Millisecond,
	},
}