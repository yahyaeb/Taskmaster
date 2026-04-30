package main

import (
    "os/exec"
    "sync"
)

type ProgramConfig struct {
	Cmd          string            `yaml:"cmd"`
	NumProcs     int               `yaml:"numprocs"`
	Umask        int               `yaml:"umask"`
	WorkingDir   string            `yaml:"workingdir"`
	AutoStart    bool              `yaml:"autostart"`
	AutoRestart  string            `yaml:"autorestart"`
	ExitCodes    []int             `yaml:"exitcodes"`
	StartRetries int               `yaml:"startretries"`
	StartTime    int               `yaml:"starttime"`
	StopSignal   string            `yaml:"stopsignal"`
	StopTime     int               `yaml:"stoptime"`
	Stdout       string            `yaml:"stdout"`
	Stderr       string            `yaml:"stderr"`
	Env          map[string]string `yaml:"env"`
}

type Config struct {
	Programs map[string]ProgramConfig `yaml:"programs"`
}

type ProcessState string

const (
    StateStarting ProcessState = "STARTING"
    StateRunning  ProcessState = "RUNNING"
    StateStopped  ProcessState = "STOPPED"
    StateFatal    ProcessState = "FATAL"
)

type Process struct {
    Name    string
    Config  ProgramConfig
    Cmd     *exec.Cmd
    State   ProcessState
}

type ProcessManager struct {
    mu       sync.Mutex
    processes map[string]*Process
    config   Config
}