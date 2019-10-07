package main

import "fmt"

// type, indica que se define un nuevo tipo de tipo "Usuario" y es struct
type User struct {
	age            int
	name, lastName string
}

func main() {
	alexander := User{age: 23, name: "Alexander", lastName: "Armúa"}
	fmt.Println(alexander)

	// new retorna un puntero a un tipo de la estructura
	josesito := new(User)
	fmt.Println(josesito)

	// Tomo el valor de donde apunta y veo el "age"
	fmt.Println((*josesito).age)
	// x2 pero con syntactic sugar
	fmt.Println(josesito.age)

	josesito.age = 55
	// (*josesito).age = 27

	fmt.Println(josesito.age)
}
