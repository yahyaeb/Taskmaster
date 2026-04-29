package main

import (
	// Config & setup
	"os"

	_ "github.com/chzyer/readline" // shell line editing + history
	"gopkg.in/yaml.v3"
	_ "gopkg.in/yaml.v3" // YAML config parsing

	_ "os"        // file ops, env vars, umask, working directory
	"os/exec"     // spawning child processes
	_ "os/signal" // catching SIGCHLD, SIGHUP, SIGINT

	// Process control
	_ "syscall" // signals (SIGTERM, SIGKILL, USR1...), umask, waitpid

	// Concurrency
	_ "sync" // Mutex — protecting shared process state
	_ "time" // starttime, stoptime timers

	// Logging
	"fmt"   // formatting output
	_ "log" // logging events to file

	// Control shell
	_ "strconv" // string to int conversions
	_ "strings" // parsing user commands
)

func main() {
	// Step 1 - did user provide a config file?
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./taskmaster config.yml")
		os.Exit(1)
	}

	// Step 2 - read the file
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		os.Exit(1)
	}

	// Step 3 - parse it
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("Error parsing config: %v\n", err)
		os.Exit(1)
	}

	// Step 4 - confirm it worked
	// fmt.Printf("Loaded %d programs:\n", len(config.Programs))
	// for name, program := range config.Programs {
	// 	fmt.Printf("\n[%s]\n", name)
	// 	fmt.Printf("  cmd:          %s\n", program.Cmd)
	// 	// 	// 	fmt.Printf("  numprocs:     %d\n", program.NumProcs)
	// 	// 	// 	fmt.Printf("  umask:        %d\n", program.Umask)
	// 	// 	// 	fmt.Printf("  workingdir:   %s\n", program.WorkingDir)
	// 	// 	// 	fmt.Printf("  autostart:    %t\n", program.AutoStart)
	// 	// 	// 	fmt.Printf("  autorestart:  %s\n", program.AutoRestart)
	// 	// 	// 	fmt.Printf("  startretries: %d\n", program.StartRetries)
	// 	// 	// 	fmt.Printf("  starttime:    %d\n", program.StartTime)
	// 	// 	// 	fmt.Printf("  stopsignal:   %s\n", program.StopSignal)
	// 	// 	// 	fmt.Printf("  stoptime:     %d\n", program.StopTime)
	// 	// 	// 	fmt.Printf("  stdout:       %s\n", program.Stdout)
	// 	// 	// 	fmt.Printf("  stderr:       %s\n", program.Stderr)
	// }

	// fmt.Printf("%+v\n", config)
	// cmd := exec.Command(config.Programs["myprogram"].Cmd)

	for name, program := range config.Programs {
		fmt.Printf("\n[%s]\n", name)
		if program.AutoStart {
			fmt.Printf("Starting '%s'...\n", program.Cmd)
			cmd := exec.Command(program.Cmd)
			if err := cmd.Start(); err != nil {
				fmt.Printf("Failed to start '%s': %v\n", program.Cmd, err)
			}
		}
	}

	select {}
}
