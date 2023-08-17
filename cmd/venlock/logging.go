package main

import (
	"fmt"
	"log"
	"os"
)

type logging struct {
	logger *log.Logger
}

func newLogger() *logging {
	return &logging{
		logger: log.New(os.Stderr, "", 0),
	}
}

func (l logging) print(str string) {
	l.logger.Print(str)
}

func (l logging) printHeader(str string) {
	l.logger.Printf("\n> %s", str)
}

func (l logging) printResult(str string) {
	l.logger.Printf("  %s", str)
}

func (l logging) printfResult(format string, args interface{}) {
	l.printResult(fmt.Sprintf(format, args))
}

func (l logging) printFatal(str string) {
	l.logger.Printf("  %s", str)
	panic(str)
}
