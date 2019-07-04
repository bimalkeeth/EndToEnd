package main

import fmt "fmt"

type MyError struct {
	ShortMessage    string
	DetailedMessage string
}

func (e *MyError) Error() string {
	return e.ShortMessage + "\n" + e.DetailedMessage
}

func main() {

	fmt.Println(doSomething())

}

func doSomething() error {
	return &MyError{ShortMessage: "Error", DetailedMessage: "Long"}
}
