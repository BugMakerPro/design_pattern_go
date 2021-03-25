package main
/*
饿汉模式适用场景：

 */

var es = &EagerSingleton{}

//func init() {
//	如果EagerSingleton有属性需要初始化,可以在init内进行
//}

func GetEsInstance() *EagerSingleton  {
	return es
}

type EagerSingleton struct {
	
}
