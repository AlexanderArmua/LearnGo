package main

import "fmt"

type ApiError struct {
	Message    string
	StatusCode int
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("%s - SC: %d", e.Message, e.StatusCode)
}

func main() {
	var err error = &ApiError{Message: "Bad Request", StatusCode: 400}
	fmt.Printf("Error: %s", err.Error())

	fmt.Printf("StatusCoe: %d", err.StatusCode)
}
