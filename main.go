package remote_module

import (
	"errors"
	"fmt"
)

//This is branch V2

func Greet(name string, lang string) error {

	switch lang {
	case "eng":
		fmt.Println("Hi ", name)
		return nil
	case "rus":
		fmt.Println("Привет ", name)
		return nil
	default:
		fmt.Println("Unknown language")
		return errors.New("unknown language")
	}

}
