package rpc

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/sllt/tp-tcp-plugin/model"
	"github.com/sllt/tp-tcp-plugin/pkg/rpc/req"
	"github.com/sllt/tp-tcp-plugin/pkg/rpc/resp"
	"strings"
)

var (
	URL    = "http://dev.thingspanel.cn:9999"
	client = resty.New()
)

func GetDeviceBufferConfig(accessToken string) (*model.DeviceConnConfig, error) {
	accessToken = strings.Trim(accessToken, "\n")
	result := &resp.GetFormConfigRespWithBody{}
	address := URL + "/api/plugin/device/config"
	_, err := client.R().
		SetBody(req.GetFormConfigReq{AccessToken: accessToken}).
		SetResult(result).
		Post(address)
	if err != nil {
		return nil, err
	}

	if result.Data != nil && result.Data.AccessToken == "" {
		return nil, errors.New("device not found")
	} else if result.Data != nil && result.Data.DeviceConfig != nil {
		return &model.DeviceConnConfig{
			DeviceType:         result.Data.DeviceType,
			InBoundByteLength:  result.Data.DeviceConfig.InBoundByteLength,
			OutBoundByteLength: result.Data.DeviceConfig.OutBoundByteLength,
		}, nil
	} else {
		// return default value
		return &model.DeviceConnConfig{
			Token:              accessToken,
			DeviceType:         result.Data.DeviceType,
			InBoundByteLength:  1,
			OutBoundByteLength: 1,
		}, nil
	}
}
