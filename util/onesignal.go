package util

import "mf_backup_onetime/dto"

type OneSignal struct {
	Config *dto.OneSignalConfig
}

func InitOneSignal(config *dto.OneSignalConfig) *OneSignal {
	return &OneSignal{
		Config: config,
	}
}
