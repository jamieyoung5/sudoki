//go:build darwin

package term

import "bufio"

func PrepareTerminal() error {
	return nil
}

func ReadKeySequence(reader *bufio.Reader) ([]byte, error) {
	return nil, nil
}

func Encode(sequence []byte) string {

}

func Clear() {

}
