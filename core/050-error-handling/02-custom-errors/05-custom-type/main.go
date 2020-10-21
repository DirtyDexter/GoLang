package main

import (
	"fmt"
	"log"
)

type myError struct {
	err error
}

func (e *myError) Error() string {
	return fmt.Sprintf("my error: %v", e.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		er := fmt.Errorf("square root of negative number: %v", f)
		return 0, &myError{er}
	}
	// implementation
	return 42, nil
}

// see use of structs with error type in standard library:
// http://www.goinggo.net/2014/11/error-handling-in-go-part-ii.html
//
// http://golang.org/pkg/net/#OpError
// http://golang.org/src/pkg/net/dial.go
// http://golang.org/src/pkg/net/net.go
//
// http://golang.org/src/pkg/encoding/json/decode.go
