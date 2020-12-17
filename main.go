package main

import (
	"fmt"
	"os"
)

func PrintSummary(list []*Category) {
	fmt.Printf("%s\t%s\n", "Count", "Name")
	for _, c := range list {
		fmt.Printf("%d\t%s\n", len(c.Items), c.Name)
	}
}

func main() {
	format := ""
	file := "./test.md"
	switch len(os.Args) {
	case 2:
		file = os.Args[1]
	case 3:
		format = os.Args[1]
		file = os.Args[2]
	default:
		fmt.Println("usage: ./wiki2code [*|json|golang] ../ipfslib.wiki/chain错误信息整理.md")
		return
	}

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	list, err := Parse(f)
	if err != nil {
		panic(err)
	}
	switch format {
	case "json":
		PrintJSON(list)
	case "golang":
		PrintGolang(list)
	default:
		PrintSummary(list)
	}
}
