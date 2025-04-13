package sentencias

import (
	"fmt"
	"runtime"
)

func SentenciaIf() {
	if os := runtime.GOOS; os == "linux" || os == "OS X" {
		fmt.Println("Esto no es Windows, es ", os)
	} else {
		fmt.Println("Esto es Windows")
	}
}