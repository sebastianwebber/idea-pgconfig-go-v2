package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/pgconfig/api/doc"
)

func main() {
	a, err := doc.Get("effective_cache_size", 11.0)

	if err != nil {
		panic(err)
	}

	spew.Dump(a)

}
