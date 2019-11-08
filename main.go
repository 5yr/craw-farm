package main

import (
	"fmt"

	"github.com/5yr/nnew/pkg/process"
	"github.com/5yr/nnew/pkg/website"
)

func init() {
	website.Setup()
}

func main() {
	t, err := process.LoadTask("/Users/jc-yiran/yiran-git/github/nnew/examples/test1.toml")
	if err != nil {
		fmt.Println(err)
	}
	t.ExecSequence()
}
