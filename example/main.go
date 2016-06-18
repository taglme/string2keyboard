package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/taglme/string2keyboard"
)

func main() {

	fmt.Print("Enter text to write as keyboard input: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\n")
	fmt.Println("Select text input field to write. Wait 5 sec to start writting...")
	time.Sleep(5 * time.Second)
	string2keyboard.KeyboardWrite(text)

}
