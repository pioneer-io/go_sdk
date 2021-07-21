package interfaces

import (
	"github.com/pioneer-io/go_sdk/pkg/models"
)

type Member interface {
	Connect() (*models.Member, error)
	HandleIncomingData()
	Listen()
	Get(flagKey string) (string)
	GetWithContext(flagKey, context string) (bool, error)
	GetWithDefault(flagKey string, defaultVal bool) (bool)
	GetWithContextWithDefault(flagKey, context string, defaultVal bool) (bool)
}
