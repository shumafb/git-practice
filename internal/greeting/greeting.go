package greeting

import (
	"fmt"
	"time"
)

func Hello(name string) {
	current := time.Now().Format("15:04")
	fmt.Printf("[%s] Привет, %s!\n", current, name)
}

func Bye(name string) {
	fmt.Println("--")
	fmt.Printf("Пока, %s!\n", name)
}
