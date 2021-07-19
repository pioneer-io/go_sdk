package sdk

import (
	"fmt"
	"testing"

	"github.com/pioneer-io/go_sdk/pkg/models"
)

// type FlagData struct {
// 	Is_Active bool
// 	Rollout   int
// }

func dummyRuleset() map[string]*models.FlagData {
	testFlag1 := &models.FlagData{Is_Active: true, Rollout: 50}
	testFlag2 := &models.FlagData{Is_Active: false, Rollout: 34}

	ruleset := make(map[string]*models.FlagData)
	ruleset["test_flag_1"] = testFlag1
	ruleset["test_flag_2"] = testFlag2
	return ruleset
}

func testClient() *models.Member {
	client := InitMember("http://localhost:3030/features", "JazzyElksRule")
	client.Ruleset = dummyRuleset()
	return client
}

func TestGet(t *testing.T) {
	client := testClient()

	
}
