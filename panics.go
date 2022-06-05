package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

type ErrorWithData struct {
	Time  string
	Trace string
	Err   error
	Msg   any
}

func (e *ErrorWithData) Error() string {
	currentTime := time.Now()
	return fmt.Sprintf("%d-%d-%d %d:%d:%d - %s\n Panic message: %s\n Trace: %s",
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Hour(),
		currentTime.Second(),
		e.Err,
		e.Msg,
		e.Trace)
}

func simplePanic() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("%s", v)
		}
	}()
	generatePanic()
}

func customPanicError() (err error) {
	defer func() {
		if v := recover(); v != nil {
			err = &ErrorWithData{
				Time:  time.Now().String(),
				Err:   fmt.Errorf("Recoreved panic"),
				Trace: string(debug.Stack()),
				Msg:   v,
			}
		}
	}()
	generatePanic()
	return nil
}

func generatePanic() {
	for i := 1; i < 10; i++ {
		if i%3 == 0 {
			panic("Some unwanted panic")
		}
	}
}
