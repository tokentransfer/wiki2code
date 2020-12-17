package main

import (
	"fmt"
	"strings"
)

func PrintGolang(list []*Category) {
	fmt.Println("type TransactionResult int16")
	fmt.Println()
	for _, c := range list {
		fmt.Println("const (")
		fmt.Println("\t//", c.Name)
		fmt.Println("\t//", c.Desc)

		for _, item := range c.Items {
			fmt.Printf("\t%s TransactionResult = %d  // %s\n", strings.Title(item.Type), item.Status, item.Desc)
		}
		fmt.Println(")")
		fmt.Println()
	}

	fmt.Println(`var resultNames = map[TransactionResult]struct {
		Token string
		Human string
	}{`)
	for _, c := range list {

		for _, item := range c.Items {
			fmt.Printf("\t%s:{\"%s\", `%s`},\n", strings.Title(item.Type), item.Type, item.Desc)
		}
	}
	fmt.Println("}")
}
