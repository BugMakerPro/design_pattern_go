package main

import "errors"

/*
简单工厂模式1(Simple Factory Pattern)
*/

var rcpFactory1 *RuleConfigParserFactory1

func Load1(configFormat string) (RuleConfig, error) {
	parser, err := rcpFactory1.createParser(configFormat)
	if err != nil {
		return nil, err
	}
	var text string
	// load form config file
	return parser.Parse(text)

}

// 简单工厂模式1
type RuleConfigParserFactory1 struct {
}

func (r *RuleConfigParserFactory1) createParser(configFormat string) (IRuleConfigParser, error) {
	var parser IRuleConfigParser
	if configFormat == "json" {
		parser = &JsonRuleConfigParser{}
	} else if configFormat == "xml" {
		parser = &XmlRuleConfigParser{}
	} else if configFormat == "yml" {
		parser = &YamlRuleConfigParser{}
	} else {
		return nil, errors.New("configFormat invalid")
	}
	return parser, nil
}

type IRuleConfigParser interface {
	Parse(text string) (RuleConfig, error)
}

type JsonRuleConfigParser struct {
}

func (j *JsonRuleConfigParser) Parse(text string) (RuleConfig, error) {
	panic("implement me")
}

type YamlRuleConfigParser struct {
}

func (y *YamlRuleConfigParser) Parse(text string) (RuleConfig, error) {
	panic("implement me")
}

type XmlRuleConfigParser struct {
}

func (x *XmlRuleConfigParser) Parse(text string) (RuleConfig, error) {
	panic("implement me")
}

type RuleConfig interface {
}
