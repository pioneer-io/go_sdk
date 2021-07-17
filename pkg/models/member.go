package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/donovanhide/eventsource"
)

type Member struct {
	SDKKey      string
	ScoutServer string
	SSEClient   *eventsource.Stream
	Ruleset     map[string]*FlagData
	HasRuleset  bool
}

func parseJSONtoSlice(data string) []Flag {
	var flags []Flag
	json.Unmarshal([]byte(data), &flags)

	return flags
}

func mapRuleset(flags []Flag) map[string]*FlagData {
	ruleset := make(map[string]*FlagData)

	for _, flag := range flags {
		ruleset[flag.Title] = &FlagData{flag.Is_Active, flag.Rollout}
	}

	return ruleset
}

func (client *Member) Get(flagKey string) bool {
	if flag, ok := client.Ruleset[flagKey]; ok {
		return flag.Is_Active
	} else {
		log.Fatal("The flag '", flagKey, "' is not in the ruleset")
		return false
	}
}

func (client *Member) GetWithDefault(flagKey string, defaultVal bool) bool {
	if flag, ok := client.Ruleset[flagKey]; ok {
		return flag.Is_Active
	} else {
		fmt.Println("The flag '", flagKey, "' is not in the ruleset. Returning the default value you provided, ", defaultVal)
		return defaultVal
	}
}

func SumContext(context string) int {
	// iterate over runes in context to get code point value
	sum := 0
	for _, rune := range context {
		sum += int(rune)
	}
	// sum % 100 because max rollout is 100
	return sum % 100
}

func (client *Member) GetWithContext(flagKey, context string) bool {
	if flag, ok := client.Ruleset[flagKey]; ok {
		intContext := SumContext(context)
		return intContext <= flag.Rollout
	} else {
		log.Fatal("The flag '", flagKey, "' is not in the ruleset.")
		return false
	}
}

func (client *Member) GetWithContextWithDefault(flagKey, context string, defaultVal bool) bool {
	if _, ok := client.Ruleset[flagKey]; ok {
		return client.GetWithContext(flagKey, context)
	} else {
		fmt.Println("The flag '", flagKey, "' is not in the ruleset. Returning the default value you provided, ", defaultVal)
		return defaultVal
	}
}

func (client *Member) Connect() (*Member, error) {
	req, err := http.NewRequest("GET", client.ScoutServer, nil)
	req.Header.Add("Authorization", client.SDKKey)

	// sseClient is an eventsource pkg *Stream object
	sseClient, err := eventsource.SubscribeWithRequest("", req)
	maxTries := 10 // don't try more than 10 times
	connectionTries := 1

	for err != nil && connectionTries < maxTries {
		time.Sleep(2 * time.Second) // wait two seconds then try again
		fmt.Println("eventsource connection failed. Trying agian.")
		sseClient, err = eventsource.SubscribeWithRequest("", req)
		connectionTries += 1
	}

	if err != nil {
		log.Fatal("ERROR: ", err)
		return nil, err
	}

	fmt.Println("Successful connection")
	client.SSEClient = sseClient

	return client, nil
}

func (client *Member) Listen() {
	go client.HandleIncomingData()

	for !client.HasRuleset {
		time.Sleep(time.Second)
	}
}

func (client *Member) HandleIncomingData() {
	for {
		newRuleset := <-client.SSEClient.Events
		fmt.Println("Event detected")
		data := newRuleset.Data()

		var parsedEvent Event
		json.Unmarshal([]byte(data), &parsedEvent)

		if parsedEvent.EventType == "CREATE_CONNECTION" {
			fmt.Println("Initial SSE connection made.")

		} else if parsedEvent.EventType == "ALL_FEATURES" {
			ruleset := mapRuleset(parseJSONtoSlice(string(parsedEvent.Payload)))
			client.Ruleset = ruleset
			client.HasRuleset = true
		}
	}
}
