`crown` is a cron program for poor people having access to local accounts
only on unix machines.  Start it using `nohup`, i.e.

	nohup crown &


# Configuration

`crown` is controlled by the configuration file
`${XDG_CONFIGDIR}/crown/config.yaml`, e.g.

	crontabs:
	   - label: basic
	     spec:  "* * * * *"
	     command: go/bin/basic
	     args: []

Usually, `${XDG_CONFIGDIR}` is set to `${HOME}/.config/` in UNIX
operating systems.  In this file we enlist all crontabs.

Each crontab contains a label mentioned by `label` parameter which is
used in log files to referring to this crontab.

`spec` parameter specifies the time `crown`
should trigger its jobs.  Their syntax is similar
to time specifications for crontabs for unix cron.  Cf.
[here](https://pkg.go.dev/github.com/robfig/cron/v3#hdr-CRON_Expression_Format)
for details.

The `command` parameter contains a file name for an executable which
`crown` should execute.  Relative file names relate to home directory.
Here `go/bin/basic` relates to a binary `${HOME}/go/bin/basic`.

The `args` parameter contains the command line parameters for the binary
given in the `command` parameter.


# Files

`${HOME}/run/crown.pid` contains the pid under which `crown` is running.

`${HOME}/log/crown.log` contains the log messages and the `stderr` and
`stdout` outputs of all jobs triggered by `crown`.
