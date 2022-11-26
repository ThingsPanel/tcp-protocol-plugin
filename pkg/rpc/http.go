package rpc

import (
	"github.com/sllt/tp-tcp-plugin/model"
	"time"
)

func GetDeviceBufferConfig(accessToken string) (*model.DeviceConnConfig, error) {

	// request delay 100ms
	time.Sleep(100 * time.Millisecond)

	return &model.DeviceConnConfig{
		Token:              accessToken,
		InBoundByteLength:  1,
		OutBoundByteLength: 1,
	}, nil
}
