package starter

import (
	"fmt"
	"net/http"
)

func SayHello(name string) string {
	return fmt.Sprintf("Hello %v. Welcome!", name)
}

func OddOrEven(number int) string {
	oddOrEven := "odd"
	if number%2 == 0 {
		oddOrEven = "even"
	}
	return fmt.Sprintf("%v is an %v number", number, oddOrEven)
}

func CheckHealth(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "health check passed")
}
