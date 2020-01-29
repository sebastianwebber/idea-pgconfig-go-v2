package main

import "github.com/pgconfig/api/config"

import "fmt"

func main() {
	input := config.Input{
		TotalRAM: 1024 * config.KB,
	}
	output := input.Compute()

	if output != nil {
		for i := 0; i < len(*output); i++ {
			fmt.Printf("-- %s\n", (*output)[i].DocURL())
			fmt.Printf("%s\n", (*output)[i].ToSQL())
		}
	}

}
