package resources

type WaitingRuleMod struct {
	SupportsObservedGeneration bool              `json:"supportsObservedGeneration"`
	SuccessfulConditions       []string          `json:"successfulConditions"`
	FailureConditions          []string          `json:"failureConditions"`
	ResourceMatchers           []ResourceMatcher `json:"resourceMatchers"`
}
