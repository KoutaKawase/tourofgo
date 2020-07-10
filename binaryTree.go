package main

// import (
// 	"fmt"

// 	"golang.org/x/tour/tree"
// )

// // Walk walks the tree t sending all values
// // from the tree to the channel ch.
// func Walk(t *tree.Tree, ch chan int) {
// 	walk(t, ch)
// 	close(ch)
// }

// func walk(t *tree.Tree, ch chan int) {
// 	if t.Left != nil {
// 		walk(t.Left, ch)
// 	}
// 	ch <- t.Value
// 	if t.Right != nil {
// 		walk(t.Right, ch)
// 	}
// }

// // Same determines whether the trees
// // t1 and t2 contain the same values.
// func Same(t1, t2 *tree.Tree) bool {
// 	c1, c2 := make(chan int), make(chan int)
// 	go Walk(t1, c1)
// 	go Walk(t2, c2)
// 	for {
// 		v1, ok1 := <-c1
// 		v2, ok2 := <-c2

// 		switch {
// 		case !ok1, !ok2:
// 			return ok1 == ok2
// 		case v1 != v2:
// 			return false
// 		}
// 	}
// }

// func main() {
// 	c := make(chan int)
// 	go Walk(tree.New(2), c)
// 	for i := range c {
// 		fmt.Println(i)
// 	}
// 	fmt.Println(Same(tree.New(1), tree.New(1)))
// 	fmt.Println(Same(tree.New(1), tree.New(2)))
// }
