package main

import "fmt"

func main() {
	ok, err := say("Hello World")
	if err != nil {
		panic(err.Error())
	}
	switch ok {
	case true:
		fmt.Println("Deu certo")
	default:
		fmt.Println("Deu errado")
	}
}

func say(what string) (bool, error) {
	if what == "" {
		return false, fmt.Errorf("empty string")
	}
	fmt.Println(what)
	return true, nil
}
