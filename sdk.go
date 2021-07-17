package sdk

import (
	"github.com/pioneer-io/go_sdk/pkg/models"
)

func InitMember(scoutServer, sdkKey string) *models.Member {
	return &models.Member{
		ScoutServer: scoutServer,
		SDKKey:      sdkKey,
	}
}