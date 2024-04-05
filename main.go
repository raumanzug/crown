package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	err := initConstants()
	if err != nil {
		panic(err)
	}

	pRunFile, err := os.OpenFile(runFile, os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(pRunFile, "%d\n", os.Getpid())
	defer func() {
		err = pRunFile.Close()
		if err != nil {
			panic(err)
		}
		err = os.Remove(runFile)
		if err != nil {
			panic(err)
		}
	}()

	ctx, _ := context.WithCancel(context.Background())

	pLogFile, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = pLogFile.Close()
		if err != nil {
			panic(err)
		}
	}()

	err = cronLoop(ctx, pLogFile)
	if err != nil {
		panic(err)
	}

}
