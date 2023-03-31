package rpc

import (
	"errors"
	"strings"

	"github.com/go-resty/resty/v2"
	global "github.com/sllt/tp-tcp-plugin/global"
	"github.com/sllt/tp-tcp-plugin/model"
	"github.com/sllt/tp-tcp-plugin/pkg/rpc/req"
	"github.com/sllt/tp-tcp-plugin/pkg/rpc/resp"
)

var (
	URL    = global.Config.Tp.HttpAddr
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
		if result.Data.DeviceConfig.InBoundByteLength != 0 {
			return &model.DeviceConnConfig{
				Token:              result.Data.AccessToken,
				DeviceType:         result.Data.DeviceType,
				InBoundByteLength:  result.Data.DeviceConfig.InBoundByteLength,
				OutBoundByteLength: result.Data.DeviceConfig.OutBoundByteLength,
			}, nil
		} else {
			// the server return an empty config
			// use default config
			return &model.DeviceConnConfig{
				Token:              accessToken,
				DeviceType:         result.Data.DeviceType,
				InBoundByteLength:  1,
				OutBoundByteLength: 1,
			}, nil
		}
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
