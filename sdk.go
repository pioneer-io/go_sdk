package sdk

import (
	"log"
	ga "github.com/jpillora/go-ogle-analytics"

	"github.com/pioneer-io/go_sdk/pkg/models"
)

func InitMember(scoutServer, sdkKey string) *models.Member {
	return &models.Member{
		ScoutServer: scoutServer,
		SDKKey:      sdkKey,
	}
}

func InitAnalytics(identifier string) *models.Analytics {
	analyticsClient, err := ga.NewClient(identifier)

	if err != nil {
		log.Fatal("Error connecting to google analytics. ", err)
	}

	return &models.Analytics{
		GoogleTrackingId: identifier,
		Client:           analyticsClient,
		EventType:        "Pioneer Analytics Log Event",
	}
}
