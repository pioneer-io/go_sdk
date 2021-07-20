package models

import (
	"log"
	"strconv"
	ga "github.com/jpillora/go-ogle-analytics"
)

type Analytics struct {
	GoogleTrackingId string
	Client *ga.Client
	EventType string
}

func (analytics *Analytics) LogAnalyticsEvent(descriptor string, ruleset map[string]*FlagData) (error) {
	for flagKey, data := range ruleset {
		flagData := flagKey + " -- " + strconv.FormatBool(data.Is_Active)
		analyticEvent := ga.NewEvent(analytics.EventType, descriptor).Label(flagData)

		err := analytics.Client.Send(analyticEvent)

		if err != nil {
			log.Fatal("There was an error sending the analytics event to Google ", err)
		}
	}

	return nil
}