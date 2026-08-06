package main

import "fmt"

type MyErr struct{}

func (*MyErr) Error() string { return "boom" }

func mayFail(fail bool) *MyErr {
	if fail {
		return &MyErr{}
	}
	return nil
}

func doWork() error {
	var p *MyErr = mayFail(false) // nil *MyErr
	return p                      // wrapped into a non-nil error interface
}

func main() {
	err := doWork()
	fmt.Println("err == nil:", err == nil)
	if err != nil {
		fmt.Println("BUG REPRODUCED: interface holding typed nil pointer is != nil")
	} else {
		fmt.Println("not reproduced")
	}
}
