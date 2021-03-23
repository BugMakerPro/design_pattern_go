package main

import "sync"

/*
懒汉模式适用场景：
懒汉式相对于饿汉式的优势是支持延迟加载。
*/

var lz *LazySingleton

type LazySingleton struct {
}

// 第一种：非线程安全
//func GetLzInstance() *LazySingleton  {
//	if lz == nil{
//		lz = &LazySingleton{}
//	}
//	return lz
//}
//

// 第二种：加锁，并发性能最差

//var mu sync.Mutex

//func GetLzInstance() *LazySingleton {
//	mu.Lock()
//	defer mu.Unlock()
//	if lz == nil {
//		lz = &LazySingleton{}
//	}
//	return lz
//}

// 第三种：锁+双重检查，并发性较第二种稍好一点

//var mu sync.Mutex
//
//func GetLzInstance() *LazySingleton {
//	if lz == nil {
//		mu.Lock()
//		defer mu.Unlock()
//		if lz == nil {
//			lz = &LazySingleton{}
//		}
//	}
//	return lz
//}

// 第四种：懒汉式最优实现,通过golang特有的sync.Once来实现

var once sync.Once

func GetLzInstance() *LazySingleton {
	once.Do(func() {
		lz = &LazySingleton{}
	})
	return lz
}
