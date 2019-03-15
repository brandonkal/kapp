package clusterapply

import (
	ctldiff "github.com/k14s/kapp/pkg/kapp/diff"
	ctlres "github.com/k14s/kapp/pkg/kapp/resources"
)

type ClusterChangeFactory struct {
	opts                ClusterChangeOpts
	identifiedResources ctlres.IdentifiedResources
	changeFactory       ctldiff.ChangeFactory
	ui                  UI
}

func NewClusterChangeFactory(
	opts ClusterChangeOpts,
	identifiedResources ctlres.IdentifiedResources,
	changeFactory ctldiff.ChangeFactory,
	ui UI,
) ClusterChangeFactory {
	return ClusterChangeFactory{opts, identifiedResources, changeFactory, ui}
}

func (f ClusterChangeFactory) NewClusterChange(change ctldiff.Change) ClusterChange {
	return NewClusterChange(change, f.opts, f.identifiedResources, f.changeFactory, f.ui)
}
