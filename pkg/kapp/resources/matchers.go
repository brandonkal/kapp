package resources

type ResourceMatcher interface {
	Matches(Resource) bool
}

type APIGroupKindMatcher struct {
	APIGroup string
	Kind     string
}

var _ ResourceMatcher = APIGroupKindMatcher{}

func (m APIGroupKindMatcher) Matches(res Resource) bool {
	return res.APIGroup() == m.APIGroup && res.Kind() == m.Kind
}

type APIVersionKindMatcher struct {
	APIVersion string
	Kind       string
}

var _ ResourceMatcher = APIVersionKindMatcher{}

func (m APIVersionKindMatcher) Matches(res Resource) bool {
	return res.APIVersion() == m.APIVersion && res.Kind() == m.Kind
}

type KindNamespaceNameMatcher struct {
	Kind, Namespace, Name string
}

var _ ResourceMatcher = KindNamespaceNameMatcher{}

func (m KindNamespaceNameMatcher) Matches(res Resource) bool {
	return res.Kind() == m.Kind && res.Namespace() == m.Namespace && res.Name() == m.Name
}

type AllMatcher struct{}

var _ ResourceMatcher = AllMatcher{}

func (AllMatcher) Matches(Resource) bool { return true }

type AnyMatcher struct {
	Matchers []ResourceMatcher
}

var _ ResourceMatcher = AnyMatcher{}

func (m AnyMatcher) Matches(res Resource) bool {
	for _, m := range m.Matchers {
		if m.Matches(res) {
			return true
		}
	}
	return false
}

type NotMatcher struct {
	Matcher ResourceMatcher
}

var _ ResourceMatcher = NotMatcher{}

func (m NotMatcher) Matches(res Resource) bool {
	return !m.Matcher.Matches(res)
}

type AndMatcher struct {
	Matchers []ResourceMatcher
}

var _ ResourceMatcher = AndMatcher{}

func (m AndMatcher) Matches(res Resource) bool {
	for _, m := range m.Matchers {
		if !m.Matches(res) {
			return false
		}
	}
	return true
}

type HasAnnotationMatcher struct {
	Keys []string
}

var _ ResourceMatcher = HasAnnotationMatcher{}

func (m HasAnnotationMatcher) Matches(res Resource) bool {
	anns := res.Annotations()
	for _, key := range m.Keys {
		if _, found := anns[key]; !found {
			return false
		}
	}
	return true
}

type HasNamespaceMatcher struct {
	Names []string
}

var _ ResourceMatcher = HasNamespaceMatcher{}

func (m HasNamespaceMatcher) Matches(res Resource) bool {
	resNs := res.Namespace()
	if len(resNs) == 0 {
		return false // cluster resource
	}
	if len(m.Names) == 0 {
		return true // matches any name, but not cluster
	}
	for _, name := range m.Names {
		if name == resNs {
			return true
		}
	}
	return false
}

var (
	builtinAPIGroups = map[string]struct{}{
		"":                             {},
		"admissionregistration.k8s.io": {},
		"apiextensions.k8s.io":         {},
		"apps":                         {},
		"authentication.k8s.io":        {},
		"authorization.k8s.io":         {},
		"autoscaling":                  {},
		"batch":                        {},
		"certificates.k8s.io":          {},
		"coordination.k8s.io":          {},
		"extensions":                   {},
		"metrics.k8s.io":               {},
		"migration.k8s.io":             {},
		"networking.k8s.io":            {},
		"node.k8s.io":                  {},
		"policy":                       {},
		"rbac.authorization.k8s.io":    {},
		"scheduling.k8s.io":            {},
		"storage.k8s.io":               {},
	}
)

type CustomResourceMatcher struct{}

var _ ResourceMatcher = CustomResourceMatcher{}

func (m CustomResourceMatcher) Matches(res Resource) bool {
	_, found := builtinAPIGroups[res.APIGroup()]
	return !found
}
