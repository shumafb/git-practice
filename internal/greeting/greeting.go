package greeting

import "fmt"

func Hello(name string) {
	fmt.Printf("Привет, %s!\n", name)
}

func Bye(name string) {
	fmt.Printf("Пока, %s!\n", name)
}
