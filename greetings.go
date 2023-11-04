package greetings

import (
	"fmt"
)

func Hello(name string) string {
	mensaje := fmt.Sprintf("Hola %v, Bienvenido", name)
	return mensaje
}
