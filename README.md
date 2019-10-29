# [gprompt] - Simple CUI Prompt by Golang

[![Build Status](https://travis-ci.org/spiegel-im-spiegel/gprompt.svg?branch=master)](https://travis-ci.org/spiegel-im-spiegel/gprompt)
[![GitHub license](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/gprompt/master/LICENSE)
[![GitHub release](http://img.shields.io/github/release/spiegel-im-spiegel/gprompt.svg)](https://github.com/spiegel-im-spiegel/gprompt/releases/latest)

## Declare [gprompt] module

See [go.mod](https://github.com/spiegel-im-spiegel/gprompt/blob/master/go.mod) file. 

## Usage of [gprompt] package

```go
package main

import (
	"fmt"
	"os"

	"github.com/spiegel-im-spiegel/gocli/rwi"
	"github.com/spiegel-im-spiegel/gprompt"
)

func main() {
	p := gprompt.New(
		rwi.New(
			rwi.WithReader(os.Stdin),
			rwi.WithWriter(os.Stdout),
			rwi.WithErrorWriter(os.Stderr),
		),
		func(s string) (string, error) {
			if s == "q" || s == "quit" {
				return "quit prompt", gprompt.ErrTerminate
			}
			runes := []rune(s)
			for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
				runes[i], runes[j] = runes[j], runes[i]
			}
			return string(runes), nil
		},
		gprompt.WithPromptString("Sample> "),
		gprompt.WithHeaderMessage("Input 'q' or 'quit' to stop"),
	)
	if !p.IsTerminal() {
		fmt.Fprintln(os.Stderr, gprompt.ErrNotTerminal)
		return
	}
	if err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
```

[gprompt]: https://github.com/spiegel-im-spiegel/gprompt "spiegel-im-spiegel/gprompt: Simple CUI Prompt by Golang"
