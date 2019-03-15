package app

import (
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ChangeImpl struct {
	name   string
	nsName string

	coreClient kubernetes.Interface
	meta       ChangeMeta
}

var _ Change = &ChangeImpl{}

func (c *ChangeImpl) Name() string     { return c.name }
func (c *ChangeImpl) Meta() ChangeMeta { return c.meta }

func (c *ChangeImpl) Fail() error {
	return c.update(func(meta *ChangeMeta) {
		falseBool := false

		meta.Successful = &falseBool
		meta.FinishedAt = time.Now().UTC()
	})
}

func (c *ChangeImpl) Succeed() error {
	return c.update(func(meta *ChangeMeta) {
		trueBool := true

		meta.Successful = &trueBool
		meta.FinishedAt = time.Now().UTC()
	})
}

func (c *ChangeImpl) update(doFunc func(*ChangeMeta)) error {
	change, err := c.coreClient.CoreV1().ConfigMaps(c.nsName).Get(c.name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("Getting app change: %s", err)
	}

	meta := NewChangeMetaFromData(change.Data)
	doFunc(&meta)

	c.meta = meta
	change.Data = meta.AsData()

	_, err = c.coreClient.CoreV1().ConfigMaps(c.nsName).Update(change)
	if err != nil {
		return fmt.Errorf("Updating app change: %s", err)
	}

	return nil
}

type NoopChange struct{}

var _ Change = NoopChange{}

func (NoopChange) Name() string     { return "" }
func (NoopChange) Meta() ChangeMeta { return ChangeMeta{} }
func (NoopChange) Fail() error      { return nil }
func (NoopChange) Succeed() error   { return nil }
