package common

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

var NotPiped = errors.New("not piped")

func ReadPipe() (res string, err error) {
	var info os.FileInfo
	if info, err = os.Stdin.Stat(); err != nil {
		return
	}
	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		err = NotPiped
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
	res = string(output)
	// remove new line at the end
	if strings.HasSuffix(res, "\n") {
		res = res[:len(res)-1]
	}
	return
}
