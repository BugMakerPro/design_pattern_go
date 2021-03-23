package main

import (
	"errors"
	"strings"
)

/*
 工厂方法2（Factory Method）
 工厂方法1中,工厂类对象的创建逻辑又耦合进了load()函数中，
 跟我们最初的简单工厂模式1非常相似，引入工厂方法非但没有解决问题，反倒让设计变得更加复杂了
 解决办法：
 我们可以为工厂类再创建一个简单工厂，也就是工厂的工厂，用来创建工厂类对象。
 RuleConfigParserFactory2，
 getFactory()返回的是缓存好的单例工厂对象
*/
var rcpFactory2 = &RuleConfigParserFactory2{
	m: map[string]RuleConfigParseFactory{
		"json": &JsonRuleConfigParserFactory{},
		"xml":  &XmlRuleConfigParserFactory{},
		"yaml": &YamlRuleConfigParserFactory{},
	},
}

func Load2(configFormat string) (RuleConfig, error) {
	factory, err := rcpFactory2.getFactory(configFormat)
	if err != nil {
		return nil, err
	}
	parser := factory.CreateParser()
	var text string
	// load form config file
	return parser.Parse(text)

}

// 简单工厂模式2
type RuleConfigParserFactory2 struct {
	m map[string]RuleConfigParseFactory
}

func (r *RuleConfigParserFactory2) getFactory(configFormat string) (RuleConfigParseFactory, error) {
	configFormat = strings.ToLower(configFormat)
	factory, ok := r.m[configFormat]
	if !ok {
		return nil, errors.New("configFormat invalid")
	}
	return factory, nil
}
