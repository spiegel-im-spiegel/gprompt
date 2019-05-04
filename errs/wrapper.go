package errs

import (
	"fmt"

	errors "golang.org/x/xerrors"
)

//wrapError is wrapper for error instance
type wrapError struct {
	msg   string
	cause error
	frame errors.Frame
}

//Wrap returns wraped error instance
func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}
	return &wrapError{msg: msg, cause: err, frame: errors.Caller(1)}
}

//Wrapf returns wraped error instance
func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &wrapError{msg: fmt.Sprintf(format, args...), cause: err, frame: errors.Caller(1)}
}

//Error method for error interface
func (we *wrapError) Error() string {
	return fmt.Sprintf("%v: %v", we.msg, we.cause)
}

//Unwrap method for errors.Wrapper interface
func (e *wrapError) Unwrap() error {
	return e.cause
}

//Format method for fmt.Formatter interface
func (we *wrapError) Format(s fmt.State, v rune) {
	errors.FormatError(we, s, v)
}

//FormatError method for errors.Formatter interface
func (we *wrapError) FormatError(p errors.Printer) error {
	p.Print(we.msg)
	we.frame.Format(p)
	return we.cause
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
