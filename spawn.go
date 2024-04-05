package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

const loggerFlags = log.LstdFlags | log.LUTC | log.Lmicroseconds

func getAbsolutePathName(inPathName string) (outPathName string, err error) {
	var pathName string

	if filepath.IsAbs(inPathName) {
		pathName = inPathName
	} else {
		pathName = filepath.Join(homeDir, inPathName)
	}

	outPathName = filepath.Clean(pathName)
	return
}

func streamLog(pLogger *log.Logger, pWg *sync.WaitGroup, pFatalLogger *log.Logger, rdr io.Reader) {
	defer pWg.Done()

	scanner := bufio.NewScanner(rdr)

	for scanner.Scan() {
		pLogger.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		pFatalLogger.Panicln(err)
	}
}

func spawn(ctx context.Context, label string, pLogFile *os.File, cmd string, args ...string) {
	var err error

	pLoggerFat := log.New(pLogFile, "(FF) "+label+" ", loggerFlags)

	absoluteCommand, err := getAbsolutePathName(cmd)
	if err != nil {
		pLoggerFat.Panicln(err)
		return
	}
	pCmd := exec.CommandContext(ctx, absoluteCommand, args...)

	stderr, err := pCmd.StderrPipe()
	if err != nil {
		pLoggerFat.Panicln(err)
		return
	}
	stdout, err := pCmd.StdoutPipe()
	if err != nil {
		pLoggerFat.Panicln(err)
		return
	}

	pLoggerErr := log.New(pLogFile, "(EE) "+label+" ", loggerFlags)
	pLoggerOut := log.New(pLogFile, "(II) "+label+" ", loggerFlags)

	var wg sync.WaitGroup
	wg.Add(1)
	go streamLog(pLoggerErr, &wg, pLoggerFat, stderr)
	wg.Add(1)
	go streamLog(pLoggerOut, &wg, pLoggerFat, stdout)

	err = pCmd.Start()
	if err != nil {
		pLoggerFat.Println(err)
		return
	}

	wg.Wait()

	err = pCmd.Wait()
	if err != nil {
		pLoggerFat.Panicln(err)
		return
	}

}
