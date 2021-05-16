// +build client

package common

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

var ErrNotPiped = errors.New("not piped")

func ReadPipe() (res string, err error) {
	var info os.FileInfo
	if info, err = os.Stdin.Stat(); err != nil {
		return
	}
	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		err = ErrNotPiped
		return
	}
	reader := bufio.NewReader(os.Stdin)
	var output []rune
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}
	res = strings.TrimSpace(string(output))
	return
}
