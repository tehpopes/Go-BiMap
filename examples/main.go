package main

import (
	"bimap"
)

func main() {
	m := bimap.NewBiMap()

	m.Insert(1, "Vivek")
	m.Insert(2, "Sai")
	m.Insert(3, "Mogili")

	m.Print()
}