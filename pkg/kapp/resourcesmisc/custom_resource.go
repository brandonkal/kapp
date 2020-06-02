package resourcesmisc

import (
	"fmt"

	ctlres "github.com/k14s/kapp/pkg/kapp/resources"
	appsv1 "k8s.io/api/apps/v1"
)

type CustomResource struct {
	resource ctlres.Resource
}

func NewCustomResource(resource ctlres.Resource) *CustomResource {
	for _, rule := range resource. {
		mods = append(mods, rule...)
	}
	matcher := ctlres.APIVersionKindMatcher{
		APIVersion: "apps/v1",
		Kind:       "DaemonSet",
	}
	if matcher.Matches(resource) {
		return &CustomResource{resource}
	}
	return nil
}

func (s CustomResource) IsDoneApplying() DoneApplyState {
	dset := appsv1.DaemonSet{}
	err := s.resource.AsUncheckedTypedObj(&dset)
	if err != nil {
		return DoneApplyState{Done: true, Successful: false, Message: fmt.Sprintf("Error: Failed obj conversion: %s", err)}
	}
	status := s.resource.Status()
	if dset.Generation != status["ObservedGeneration"] {
		return DoneApplyState{Done: false, Message: fmt.Sprintf(
			"Waiting for generation %d to be observed", dset.Generation)}
	}

	if dset.Status.NumberUnavailable > 0 {
		return DoneApplyState{Done: false, Message: fmt.Sprintf(
			"Waiting for %d unavailable pods", dset.Status.NumberUnavailable)}
	}

	return DoneApplyState{Done: true, Successful: true}
}
