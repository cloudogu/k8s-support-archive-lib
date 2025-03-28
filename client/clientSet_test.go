package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"k8s.io/client-go/rest"
)

func TestNewSupportArchiveClientSet(t *testing.T) {
	t.Run("should create new support archive client set", func(t *testing.T) {
		// given
		config := &rest.Config{}

		// when
		clientSet, err := NewSupportArchiveClientSet(config)

		// then
		require.NoError(t, err)
		assert.NotEmpty(t, clientSet)
	})
	t.Run("should fail to create new support archive client for wrong config", func(t *testing.T) {
		// given
		config := &rest.Config{
			Host: "foo:/error",
		}

		// when
		clientSet, err := NewSupportArchiveClientSet(config)

		// then
		require.Error(t, err)
		require.Nil(t, clientSet)
		assert.ErrorContains(t, err, "host must be a URL or a host:port pair")
	})
}

func Test_clientSet_SupportArchiveV1(t *testing.T) {
	t.Run("should return V1Alpha1Client", func(t *testing.T) {
		// given
		config := &rest.Config{}
		client, err := NewSupportArchiveClientSet(config)
		require.NoError(t, err)

		// when
		componentClient := client.SupportArchiveV1()

		// then
		assert.NotEmpty(t, componentClient)
	})
}
