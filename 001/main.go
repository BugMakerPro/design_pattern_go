package main

import "fmt"

func main() {
	var n Book = new(OffNovelBook)
	n.GetPrice()
}

type Book interface {
	GetPrice()
}

type NovelBook struct {

}

func (n *NovelBook)GetPrice()  {
	fmt.Println("NovelBook")
}

type OffNovelBook struct {
	NovelBook
}

func (n *OffNovelBook)GetPrice()  {
	fmt.Println("OffNovelBook")
}


type Game interface {
	initPlay()
	start()
	end()
}

type GameStarter struct {

}

func (s GameStarter)Start(g Game)  {
	g.initPlay()
	g.start()
	g.end()
}