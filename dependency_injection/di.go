package dependencyinjection

import (
	"fmt"
	"io"
	"net/http"
)

// Greet's given string. Using `io.Writer` as its interface for multiple usages
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}
