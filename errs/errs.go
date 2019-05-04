package errs

import "fmt"

//Num is error number for gpgpdumo
type Num int

const (
	ErrNullPointer Num = iota + 1
	ErrTerminate
	ErrNotTerminal
)

var errMessage = map[Num]string{
	ErrNullPointer: "Null reference instance",
	ErrTerminate:   "Terminate prompt",
	ErrNotTerminal: "Not terminal (or pipe?)",
}

func (n Num) Error() string {
	if s, ok := errMessage[n]; ok {
		return s
	}
	return fmt.Sprintf("Unknown error (%d)", int(n))
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
