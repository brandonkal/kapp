package resources

type WaitingRuleMod struct {
	SupportsObservedGeneration bool
	SuccessfulConditions       []string
	FailureConditions          []string
	ResourceMatchers           []ResourceMatcher
}
