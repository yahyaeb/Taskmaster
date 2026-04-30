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

func (pm *ProcessManager) SpawnProcess(name string, instance int) {
    program := pm.config.Programs[name]
    instanceName := fmt.Sprintf("%s:%d", name, instance)
    
    cmd := exec.Command(program.Cmd)
    if err := cmd.Start(); err != nil {
        fmt.Printf("Failed to start '%s': %v\n", instanceName, err)
        return
    }
    
    fmt.Printf("Started '%s' with PID %d\n", instanceName, cmd.Process.Pid)
    
    // add to map
    pm.mu.Lock()
    pm.processes[instanceName] = &Process{
        Name:   instanceName,
        Config: program,
        Cmd:    cmd,
        State:  StateRunning,
    }
    pm.mu.Unlock()
}

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

	fmt.Printf("================================================================\n")
	fmt.Printf("TASKMASTER STARTING...\n")
	fmt.Printf("================================================================\n")
	programCount := 0
	for name, program := range config.Programs {
		fmt.Printf("\n[%s]\n", name)
		if program.AutoStart {
			fmt.Printf("Starting '%s'...\n", program.Cmd)
			fmt.Printf("  command:          %s\n", program.Cmd)
			fmt.Printf("  numprocs:         %d\n", program.NumProcs)
			fmt.Printf("  autostart:          %t\n", program.AutoStart)
			fmt.Printf("  autorestart:		%s\n", program.AutoRestart)
			cmd := exec.Command(program.Cmd)
			if err := cmd.Start(); err != nil {
				fmt.Printf("Failed to start '%s': %v\n", program.Cmd, err)
			} else {
				fmt.Printf("Started successfully with PID %d\n", cmd.Process.Pid)
				programCount++
				// only watch it if it actually started
				go func(c *exec.Cmd, n string) {
					c.Wait()
					fmt.Printf("%s exited\n", n)
				}(cmd, name)
}
		}
		
	}
	fmt.Printf("================================================================\n")
	fmt.Printf("%d programs started.\n", programCount)
	fmt.Printf("================================================================\n")
	// select {}



//-----------------------
	pm := &ProcessManager{
		processes: make(map[string]*Process),
		config:    config,
	}




	// Start all autostart programs

	fmt.Printf("ProcessManager ready, tracking %d processes.\n", len(pm.processes))

}
