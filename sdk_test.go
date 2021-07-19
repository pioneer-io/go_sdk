package sdk

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/pioneer-io/go_sdk/pkg/models"
)

func dummyRuleset() map[string]*models.FlagData {
	testFlag1 := &models.FlagData{Is_Active: true, Rollout: 50}
	testFlag2 := &models.FlagData{Is_Active: false, Rollout: 34}

	ruleset := make(map[string]*models.FlagData)
	ruleset["test_flag_1"] = testFlag1
	ruleset["test_flag_2"] = testFlag2
	return ruleset
}

func testClient(t *testing.T) *models.Member {
	client := InitMember("http://localhost:3030/features", "JazzyElksRule")
	client.Ruleset = dummyRuleset()

	return client
}

func TestGet(t *testing.T) {
	assert := assert.New(t)
	client := testClient(t)

	assert.Equal(2, len(client.Ruleset), "Ruleset should contain 2 flags")
	assert.Equal(true, client.Get("test_flag_1"), "Toggled on flag should return true")
	assert.Equal(false, client.Get("test_flag_2"), "Toggled off flag should return false")
}

func TestGetWithDefault(t *testing.T) {
	assert := assert.New(t)
	client := testClient(t)

	assert.Equal(true, client.GetWithDefault("non_existent", true), "Should return default if flag not present")
	assert.Equal(false, client.GetWithDefault("test_flag_2", true), "should return Is_Active value if flag is present, ignoring default")
}

func TestGetWithContext(t *testing.T) {
	assert := assert.New(t)
	client := testClient(t)
	id1 := "it-is-a-dummy-uuid" // sum % 100 = 13
	id2 := "ITSZ A DUMMY" // sum % 100 = 54

	assert.Equal(false, client.GetWithContext("test_flag_2", id1), "toggled off flag should return false regardless")
	assert.Equal(true, client.GetWithContext("test_flag_1", id1), "toggled on flag should return true if context falls within rollout")
	assert.Equal(false, client.GetWithContext("test_flag_1", id2), "toggled on flag should return false if context falls above rollout")
}

func TestGetWIthContextWithDefault(t *testing.T) {
	assert := assert.New(t)
	client := testClient(t)
	id1 := "it-is-a-dummy-uuid" // sum % 100 = 13
	id2 := "ITSZ A DUMMY" // sum % 100 = 54

	assert.Equal(false, client.GetWithContextWithDefault("non_existent", id1, false), "default value should be returned if flag not in ruleset")
	assert.Equal(true, client.GetWithContextWithDefault("test_flag_1", id1, false), "default value should be ignored if flag in ruleset. return true if context falls within rollout")
	assert.Equal(false, client.GetWithContextWithDefault("test_flag_1", id2, true), "default value should be ignored if flag in ruleset. return false if context falls outside rollout")
}