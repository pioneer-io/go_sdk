package main

import (
	"fmt"
	"time"

	sdk "github.com/pioneer-io/go_sdk"
	// "github.com/pioneer-io/go_sdk/pkg/models"
	// "gopkg.in/segmentio/analytics-go.v3"
)

/*
run a simple program like this to test out the SDK
*/

func main() {
	// Initialize an SDK client
	client := sdk.InitMember("http://localhost:3030/features", "JazzyElksRule")

	// connect SDK client to Scout to listen for SSE updates
	client.Connect()
	client.Listen()

	// initialize a google analytics integration
	analytics := sdk.InitAnalytics("UA-XXXXXX-X")
	// log a google analytics  event
	fmt.Println(analytics.LogAnalyticsEvent("pioneer", "log", "1"))


	testFlag1 := &models.FlagData{Is_Active: true, Rollout: 50}
	ruleset := make(map[string]*models.FlagData)
	ruleset["test_flag_1"] = testFlag1


	// this example supposes that you have an existing flag
	// called 'test this flag' that is toggled off

	fmt.Println(client.Get("test this flag")) // false

	time.Sleep(12 * time.Second)
	// wait so you can toggle "test this flag" von ia UI
	// and create "a_new_flag" and toggle it on

	// after we're finished sleeping, we can log the updated ruleset
	fmt.Println(client.Get("a_new_flag"))     // true
	fmt.Println(client.Get("test this flag")) // true

}