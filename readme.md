# string2keyboard
Emulate keyboard input from string

## Overview
Package is used to emulate keyboard input. 
It is generates keystrokes on the keyboard and sends them to the active text input field.
So cursor should be in some text input field before calling *KeyboardWrite* method.
Escape sequences are available for some special characters ('Tab', 'Enter' etc.)

## Install

```golang
go get github.com/taglme/string2keyboard
```

## Import

```golang
import (

    "github.com/taglme/string2keyboard"

)
```

## Usage

```golang
message := "Hello world!"
err := string2keyboard.KeyboardWrite(message)

```
## Special characters
It is possible to pass some special characters using escape sequence. Use backslash character to start escape sequence.
List of available special characters: 

* \n - line feed or newline
* \b - backspace
* \t - horizontal tab
* \\\ - backslash
* \\" - double quote




## Basic example
Check usage in example folder.

## Special characters example

```golang
//Newline example
message := "123\n456"
err := string2keyboard.KeyboardWrite(message)

//Output
123
456

//Tab example
message := "123\t456"
err := string2keyboard.KeyboardWrite(message)

//Output
123   456

//Backslash example
message := "123\\456"
err := string2keyboard.KeyboardWrite(message)

//Output
123\456

```