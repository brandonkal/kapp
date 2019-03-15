package config

import (
	ctlres "github.com/k14s/kapp/pkg/kapp/resources"
)

type Conf struct {
	configs []Config
}

func NewConfFromResources(resources []ctlres.Resource) ([]ctlres.Resource, Conf, error) {
	var rsWithoutConfigs []ctlres.Resource
	var configs []Config

	for _, res := range resources {
		if res.APIVersion() == configAPIVersion && res.Kind() == configKind {
			config, err := NewConfigFromResource(res)
			if err != nil {
				return nil, Conf{}, err
			}
			configs = append(configs, config)
		} else {
			rsWithoutConfigs = append(rsWithoutConfigs, res)
		}
	}

	return rsWithoutConfigs, Conf{configs}, nil
}

func (c Conf) RebaseMods() []ctlres.FieldCopyMod {
	var mods []ctlres.FieldCopyMod
	for _, config := range c.configs {
		for _, rule := range config.RebaseRules {
			mods = append(mods, rule.AsMods()...)
		}
	}
	return mods
}

func (c Conf) OwnershipLabelMods() func(kvs map[string]string) []ctlres.StringMapAppendMod {
	return func(kvs map[string]string) []ctlres.StringMapAppendMod {
		var mods []ctlres.StringMapAppendMod
		for _, config := range c.configs {
			for _, rule := range config.OwnershipLabelRules {
				mods = append(mods, rule.AsMods(kvs)...)
			}
		}
		return mods
	}
}

func (c Conf) LabelScopingMods() func(kvs map[string]string) []ctlres.StringMapAppendMod {
	return func(kvs map[string]string) []ctlres.StringMapAppendMod {
		var mods []ctlres.StringMapAppendMod
		for _, config := range c.configs {
			for _, rule := range config.LabelScopingRules {
				mods = append(mods, rule.AsMods(kvs)...)
			}
		}
		return mods
	}
}

func (c Conf) TemplateRules() []TemplateRule {
	var result []TemplateRule
	for _, config := range c.configs {
		result = append(result, config.TemplateRules...)
	}
	return result
}
