package gprompt

import (
	"bufio"
	"os"
	"strings"

	isatty "github.com/mattn/go-isatty"
	"github.com/spiegel-im-spiegel/gocli/rwi"
	errors "golang.org/x/xerrors"
)

//Prompt is a class for interactive mode in CUI shell
type Prompt struct {
	rw        *rwi.RWI
	function  func(string) (string, error)
	headerMsg string
	promptStr string
	scanner   *bufio.Scanner
}

//OptFunc is self-referential function for functional options pattern
type OptFunc func(*Prompt)

//New returns new Prompt instance
func New(rw *rwi.RWI, function func(string) (string, error), opts ...OptFunc) *Prompt {
	p := &Prompt{rw: rw, function: function, scanner: bufio.NewScanner(rw.Reader())}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

//WithPromptString returns function for setting Reader
func WithPromptString(s string) OptFunc {
	return func(p *Prompt) {
		p.promptStr = s
	}
}

//WithHeaderMessage returns function for setting Reader
func WithHeaderMessage(s string) OptFunc {
	return func(p *Prompt) {
		p.headerMsg = s
	}
}

//IsTerminal returns true if running in terminal
func (p *Prompt) IsTerminal() bool {
	if file, ok := p.rw.Reader().(*os.File); !ok {
		return false
	} else if !isatty.IsTerminal(file.Fd()) && !isatty.IsCygwinTerminal(file.Fd()) {
		return false
	}
	if file, ok := p.rw.Writer().(*os.File); !ok {
		return false
	} else if !isatty.IsTerminal(file.Fd()) && !isatty.IsCygwinTerminal(file.Fd()) {
		return false
	}
	return true
}

//Run function starts interactive mode.
func (p *Prompt) Run() error {
	if p == nil {
		return ErrTerminate
	}
	if len(p.headerMsg) > 0 {
		if err := p.rw.Outputln(p.headerMsg); err != nil {
			return err
		}
	}

	for {
		s, ok := p.get()
		if !ok {
			break
		}
		if res, err := p.function(s); err != nil {
			_ = p.rw.Outputln(res)
			if !errors.Is(err, ErrTerminate) {
				return err
			}
			return nil
		} else if err := p.rw.Outputln(res); err != nil {
			return err
		}
	}

	if err := p.scanner.Err(); err != nil {
		return err
	}
	return nil
}

//Once function starts interactive mode (one round).
func (p *Prompt) Once() error {
	if p == nil {
		return ErrTerminate
	}
	if len(p.headerMsg) > 0 {
		if err := p.rw.Outputln(p.headerMsg); err != nil {
			return err
		}
	}

	s, ok := p.get()
	if !ok {
		return ErrTerminate
	}
	if err := p.scanner.Err(); err != nil {
		return err
	}

	if res, err := p.function(s); err != nil {
		_ = p.rw.Outputln(res)
		if !errors.Is(err, ErrTerminate) {
			return err
		}
		return nil
	} else if err := p.rw.Outputln(res); err != nil {
		return err
	}
	return nil
}

func (p *Prompt) get() (string, bool) {
	if len(p.promptStr) > 0 {
		_ = p.rw.Output(p.promptStr)
	}
	if !p.scanner.Scan() {
		return "", false
	}
	return strings.Trim(p.scanner.Text(), "\t "), true
}

/* MIT License
 *
 * Copyright 2019 Spiegel
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
