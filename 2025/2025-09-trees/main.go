package main

import (
	"fmt"
	// "iafisher.com/small-projects/trees/avl"
	"iafisher.com/small-projects/trees/twothree"
)

func main() {
	if true {
		s := "poiuytrewqlkjhgfdsamnbvcxz"

		root := twothree.New("0", "")
		for _, c := range s {
			fmt.Printf("insert %c\n", c)
			root = root.Insert(string(c), "")
			// root.Check()
		}
		fmt.Println(root.String())
	}
	// fmt.Println("insert 6")
	// root = root.Insert("6", "")
	// root.Print()
	// fmt.Println("insert 8")
	// root = root.Insert("8", "")
	// root.Print()
	// fmt.Println("insert 3")
	// root = root.Insert("3", "")
	// root.Print()
	// fmt.Println("insert 2")
	// root = root.Insert("2", "")
	// root.Print()
	// fmt.Println("insert 4")
	// root = root.Insert("4", "")
	// root.Print()
	// fmt.Println("insert 7")
	// root = root.Insert("7", "")
	// root.Print()
	// root.Check()

	if false {
		root := twothree.New("9", "")
		root = root.Insert("5", "")
		root = root.Insert("8", "")
		root = root.Insert("3", "")
		root = root.Insert("2", "")
		root = root.Insert("4", "")
		root = root.Insert("7", "")
		fmt.Println(root.String())
	}
}
