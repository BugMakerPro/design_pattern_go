package main

/*
抽象工厂（Abstract Factory）
*/
type IconfigParserFactory interface {
	CreateRuleParser()
	CreateSystemParser()
}