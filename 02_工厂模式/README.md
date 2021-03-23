###工厂模式 ###

---
#### 概念
> 简单工厂模式(Simple Factory Pattern)
> 工厂方法模式(Factory Method)：定义了一个创建对象的接口，但由子类决定要实例化的类是哪一个。工厂方法让类把实例化推迟到子类  
> 抽象工厂模式：提供一个接口，用于创建相关或依赖对象的家族，而不需要明确指定具体类   
---
#### 要解决的问题/用途
- 当创建逻辑比较复杂，是一个“大工程”的时候，我们就考虑使用*工厂模式*，  
  封装对象的创建过程，将对象的创建和使用相分离
  代码中存在```if-else```分支判断，动态地根据不同的类型创建不同的对象。  
  针对这种情况，我们就考虑使用*工厂模式*，将这一大坨```if-else```创建对象的代码抽离出来，放到工厂类中.  
  还有一种情况，尽管我们不需要根据不同的类型创建不同的对象，但是，单个对象本身的创建过程比较复杂，  
  比如要组合其他类对象，做各种初始化操作。在这种情况下，我们也可以考虑使用*工厂模式*，  
  将对象的创建过程封装到工厂类中
---
#### 如何实现
#####  1,简单工厂模式(Simple Factory Pattern)
 ```go
/*
简单工厂模式1(Simple Factory Pattern)
*/

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


```
还有一种简单工厂模式的实现,和单例模式相结合
```go
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


```
> 简单工厂模式总结：尽管简单工厂模式的代码实现中，有多处if分支判断逻辑，违背开闭原则，  
  但权衡扩展性和可读性，这样的代码实现在大多数情况下（比如，不需要频繁地添加parser，也没有太多的parser）  
  是没有问题的

#####  2,工厂方法（Factory Method）
```go
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

```
考虑到工厂类只包含方法，不包含成员变量，所以完全可以复用，于是有了工厂方法模式的另外一种实现：
```go
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

```
> 工厂方法模式总结:对于规则配置文件解析这个应用场景来说，工厂模式需要额外创建诸多Factory类，  
  也会增加代码的复杂性，而且，每个Factory类只是做简单的new操作，功能非常单薄（只有一行代码），  
  也没必要设计成独立的类，所以，在这个应用场景下，简单工厂模式简单好用，比工方法厂模式更加合适

#### 抽象工厂（Abstract Factory）
> 一个factory接口创建多个parser
 ```go
/*
抽象工厂（Abstract Factory）
*/
type IconfigParserFactory interface {
	CreateRuleParser()
	CreateSystemParser()
}

```