package main

import "errors"

/*
工厂方法1（Factory Method）,这样当我们新增一种parser的时候，只需要新增一个实现了IRuleConfigParserFactory接口的Factory类即可。
 所以，工厂方法模式比起简单工厂模式更加符合开闭原则
*/
var rcpFactory1 *RuleConfigParserFactory1

func Load1(configFormat string) (RuleConfig, error) {
	factory, err := rcpFactory1.getFactory(configFormat)
	if err != nil {
		return nil, err
	}
	parser := factory.CreateParser()
	var text string
	// load form config file
	return parser.Parse(text)

}

// 简单工厂模式1
type RuleConfigParserFactory1 struct {
}

func (r *RuleConfigParserFactory1) getFactory(configFormat string) (RuleConfigParseFactory, error) {
	var factory RuleConfigParseFactory
	if configFormat == "json" {
		factory = &JsonRuleConfigParserFactory{}
	} else if configFormat == "xml" {
		factory = &XmlRuleConfigParserFactory{}
	} else if configFormat == "yml" {
		factory = &YamlRuleConfigParserFactory{}
	} else {
		return nil, errors.New("configFormat invalid")
	}
	return factory, nil
}

type RuleConfigParseFactory interface {
	CreateParser() IRuleConfigParser
}

type JsonRuleConfigParserFactory struct {
}

func (j *JsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return &JsonRuleConfigParser{}
}

type YamlRuleConfigParserFactory struct {
}

func (j *YamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return &YamlRuleConfigParser{}
}

type XmlRuleConfigParserFactory struct {
}

func (j *XmlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return &XmlRuleConfigParser{}
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
