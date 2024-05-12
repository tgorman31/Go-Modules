package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func User_Input(msg string) (out string) {
	fmt.Println(msg)

	reader := bufio.NewReader(os.Stdin)

	out, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("Error reading input: %e", err)
	}

	out = strings.TrimSpace(out)

	return out
}
