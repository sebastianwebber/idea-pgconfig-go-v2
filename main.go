package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/pgconfig/api/config"
)

func main() {
	input := config.Input{
		TotalRAM: 1024 * config.KB,
	}
	output := input.Compute()

	fmt.Println("\n=== DEBUG OUTPUT ===============")
	spew.Dump(output)

	fmt.Println("\n=== SQL OUTPUT =================")
	printSQL(output)

	// fmt.Println("\n=== JSON OUTPUT ================")
	// printSQL(output)

}

func printSQL(output *[]config.Category) {

	for _, cat := range *output {
		fmt.Printf("\n... %s ................\n", cat.Name)

		for _, param := range cat.Parameters {
			fmt.Printf("-- %s\n", param.DocURL())
			fmt.Printf("%s\n\n", param.ToSQL())
		}
	}
}
