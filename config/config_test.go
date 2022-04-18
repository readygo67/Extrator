package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConfig(t *testing.T) {
	log.SetFormatter(&log.TextFormatter{
		DisableColors:          true,
		DisableTimestamp:       true,
		DisableSorting:         true,
		DisableLevelTruncation: true})
	config, err := New("../config.yml")
	require.NoError(t, err)
	log.Println(config.DB)
	log.Println(config.RPCURL)
}
