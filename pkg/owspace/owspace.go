// Overwrite space

package owspace

import (
	"bytes"
	"io"
	"os"
)

var clearLine = []byte("\x1b[1A\x1b[2K")

type Writer struct {
	buffer bytes.Buffer
	space  io.Writer
	lines  int
}

func (w *Writer) Write(s string) {
	w.buffer.WriteString(s)
	data := w.buffer.Bytes()
	w.buffer.Reset()
	_, _ = w.space.Write(bytes.Repeat(clearLine, w.lines))
	_, _ = w.space.Write(data)
	w.lines = bytes.Count(data, []byte{'\n'})
}

func New(fn func(w *Writer)) {
	w := &Writer{space: io.Writer(os.Stdout)}
	fn(w)
}
