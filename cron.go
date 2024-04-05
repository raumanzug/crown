package main

import (
	"context"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
	"syscall"
)

func getExec(ctx context.Context, pLogFile *os.File, pCrontab *crontab_s) func() {
	deeperCtx, _ := context.WithCancel(ctx)

	return func() {
		spawn(deeperCtx, pCrontab.Label, pLogFile, pCrontab.Command, pCrontab.Args...)
	}
}

func startCron(ctx context.Context, pLogFile *os.File, pConfigData *crown_conf_s) (pCron *cron.Cron, err error) {
	pCron = cron.New()

	for _, crontab := range pConfigData.Crontabs {
		pCron.AddFunc(crontab.Spec, getExec(ctx, pLogFile, &crontab))
	}

	pCron.Start()

	return
}

func cronLoop(ctx context.Context, pLogFile *os.File) (err error) {
	var configData crown_conf_s
	var pCron *cron.Cron

	err = loadConfigFile(configFile, &configData)
	if err != nil {
		return
	}
	pCron, err = startCron(ctx, pLogFile, &configData)
	if err != nil {
		return
	}

	hupCh := make(chan os.Signal, 1)
	stopCh := make(chan os.Signal, 1)

	signal.Notify(hupCh, syscall.SIGHUP)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	for isLoop := true; isLoop; {
		select {
		case <-hupCh:
			pCron.Stop()
			err = loadConfigFile(configFile, &configData)
			if err != nil {
				return
			}
			pCron, err = startCron(ctx, pLogFile, &configData)
			if err != nil {
				return
			}
		case <-stopCh:
			pCron.Stop()
			isLoop = false
		}
	}

	return
}
