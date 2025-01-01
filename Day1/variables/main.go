package main

import "fmt"

func main() {
	var name string = "golang" // strict typing
	var name2 = "golang2"      // golang infer the data type
	name3 := "krishna"         // shorthand syntax hai ye
	fmt.Println("Ye hai new lang", name)
	fmt.Println("Ye hai new lang", name2)
	fmt.Println("Ye mera naam hai", name3)

	var nami string
	nami = "ye definition hai" // yaha define karna hi pada
	fmt.Println(nami)
}
