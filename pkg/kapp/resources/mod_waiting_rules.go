package resources

import "log"

var globalWaitingRules []WaitingRuleMod

func SetGlobalWaitingRules(rules []WaitingRuleMod) {
	globalWaitingRules = rules
}

type WaitingRuleMod struct {
	SupportsObservedGeneration bool              `json:"supportsObservedGeneration"`
	SuccessfulConditions       []string          `json:"successfulConditions"`
	FailureConditions          []string          `json:"failureConditions"`
	ResourceMatchers           []ResourceMatcher `json:"resourceMatchers"`
}

// Find waiting rule for specified resource
func GetWaitingRule(res Resource) WaitingRuleMod {
	rules := globalWaitingRules
	mod := WaitingRuleMod{}
	log.Printf("Finding Waiting Rule: length: %v\n", len(rules))
	log.Printf("%v", rules)
	log.Printf("res: %v", res.Description())
	for _, rule := range rules {
		for _, matcher := range rule.ResourceMatchers {
			if matcher.Matches(res) {
				log.Printf("Match found: %v", rule)
				mod.SupportsObservedGeneration = rule.SupportsObservedGeneration
				mod.SuccessfulConditions = append(mod.SuccessfulConditions, rule.SuccessfulConditions...)
				mod.FailureConditions = append(mod.FailureConditions, rule.FailureConditions...)
			}
		}
	}
	log.Printf("mod after Get: %v", mod)
	return mod
}
