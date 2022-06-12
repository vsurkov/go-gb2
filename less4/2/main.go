package main

import (
	"fmt"
	"github.com/vsurkov/lgr/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//Написать программу, которая при получении в канал сигнала SIGTERM останавливается не позднее,
//чем за одну секунду (установить таймаут).

func main() {

	lgr.Logger(fmt.Sprintf("init complete"))

	<-makeSigtermChan()
	lgr.Logger(fmt.Sprintf("SIGTERM was handled, system stops"))

	time.Sleep(1 * time.Second)
	lgr.Logger(fmt.Sprintf("the system has been stopped correctly"))
}

func makeSigtermChan() chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	return ch
}
