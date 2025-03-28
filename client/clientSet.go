package client

import (
	"k8s.io/client-go/rest"

	v1 "github.com/cloudogu/k8s-support-archive-lib/client/v1"
)

// SupportArchiveEcosystemInterface exposes the clients for all the custom resources of this library.
type SupportArchiveEcosystemInterface interface {
	SupportArchiveV1() v1.SupportArchiveV1Interface
}

type clientSet struct {
	clientV1 v1.SupportArchiveV1Interface
}

// NewSupportArchiveClientSet creates a new instance of the support archive client set.
func NewSupportArchiveClientSet(config *rest.Config) (SupportArchiveEcosystemInterface, error) {
	clientV1, err := v1.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &clientSet{
		clientV1: clientV1,
	}, nil
}

// SupportArchiveV1 returns the support archive v1 client.
func (cswc *clientSet) SupportArchiveV1() v1.SupportArchiveV1Interface {
	return cswc.clientV1
}
