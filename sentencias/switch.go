package sentencias

import (
	"fmt"
	"runtime"
)

func SentenciaSwitch() {
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "OS X":
		fmt.Println("Esto es OS X")
	default:
		fmt.Printf("%s \n", os)
	}
}