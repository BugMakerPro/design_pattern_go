package main

import (
	"errors"
	"strings"
)

/*
简单工厂模式2(Simple Factory Pattern)
*/

var rcpFactory2 = &RuleConfigParserFactory2{
	cacheParsers: map[string]IRuleConfigParser{
		"json": &JsonRuleConfigParser{},
		"xml":  &XmlRuleConfigParser{},
		"yaml": &YamlRuleConfigParser{},
	},
}

func Load2(configFormat string) (RuleConfig, error) {
	parser, err := rcpFactory2.createParser(configFormat)
	if err != nil {
		return nil, err
	}
	var text string
	// load form config file
	return parser.Parse(text)

}

// 简单工厂模式2
type RuleConfigParserFactory2 struct {
	cacheParsers map[string]IRuleConfigParser
}

func (r *RuleConfigParserFactory2) createParser(configFormat string) (IRuleConfigParser, error) {
	configFormat = strings.ToLower(configFormat)
	parser, ok := r.cacheParsers[configFormat]
	if !ok {
		return nil, errors.New("configFormat invalid")
	}
	return parser, nil
}
