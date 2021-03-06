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
