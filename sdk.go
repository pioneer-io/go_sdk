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

func InitAnalytics(identifier string) *models.Analytics {

	return &models.Analytics{ // google analytics collector in go sdk for featurehub
		GoogleTrackingId: identifier,
		EventType:        "Pioneer Analytics",
	}
}
