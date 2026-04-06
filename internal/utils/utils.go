package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetTerminalInput(message string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s", message)
	textInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return textInput, nil
}
