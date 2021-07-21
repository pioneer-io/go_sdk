package models

import (
	"net/http"
	"net/url"
	uuid "github.com/gofrs/uuid"
)

type Analytics struct {
	GoogleTrackingId string
	EventType string
}

func (analytics *Analytics) LogAnalyticsEvent(category, action, value string) error {

	gaPropertyID := analytics.GoogleTrackingId
	v := url.Values{
		"v":   {"1"},
		"tid": {gaPropertyID},
		// Anonymously identifies a particular user. See the parameter guide for
		// details:
		// https://developers.google.com/analytics/devguides/collection/protocol/v1/parameters#cid
		//
		// Depending on your application, this might want to be associated with the
		// user in a cookie.
		"cid": {uuid.Must(uuid.NewV4()).String()},
		"t":   {"event"},
		"ec":  {category},
		"ea":  {action},
		"ev": {value},
		"ua":  {"pioneer go sdk"},
	}
	_, err := http.PostForm("https://www.google-analytics.com/collect", v)
	return err
}