package resp

type DeviceConfig struct {
	InBoundByteLength  int `json:"InBoundByteLength"`
	OutBoundByteLength int `json:"OutBoundByteLength"`
}

type GetFormConfigResp struct {
	ID           string        `json:"Id"`
	AccessToken  string        `json:"AccessToken"`
	ProtocolType string        `json:"ProtocolType"`
	DeviceType   string        `json:"DeviceType"`
	DeviceConfig *DeviceConfig `json:"DeviceConfig"`
}

type GetFormConfigRespWithBody struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    *GetFormConfigResp `json:"data"`
}
