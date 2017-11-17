package main

import (
	"fmt"
	"encoding/json"
	"os"
	"os/exec"
	"syscall"
	"strconv"
	"log"
	"io"
	"io/ioutil"
	"strings"
	"path/filepath"
	"bytes"
	)

var currentId int
var processName string
var currentWorkingDir string
var arg1 string
var fullcommand string
var state string
var runningPid int
var pidRunning string

var args string

var cmd string
var todos Todos
var apps Apps


// Give us some seed data
func init() {
	}


func RepoAppControl(a App) App {
      currentId += 1
      a.Id = currentId
      processName = a.Command
      arg1 = a.Arg1
      currentWorkingDir = a.Cwd
      state = a.State
      apps = append(apps, a)
      b, _ := json.Marshal(a)
      s := string(b)

     if (state == "started"){



      fmt.Println("Application started: ", s )
      fmt.Println("Application started: ", processName )
      fmt.Println("With args: ", arg1 )
      fmt.Println("Current Working Dir:",  currentWorkingDir )
      fmt.Println("State:", state )

     // fullcommand = processName + " " + arg1


      cmd := exec.Command(processName, arg1)
      cmd.Stdout = os.Stdout 
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
      //args :=  []string{"-i"}a
      err := cmd.Start()
      if err !=nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Println("Error: not Started" , processName)
		os.Exit(1)
	}
	runningPid = cmd.Process.Pid
	fmt.Println("Application Successfully Started!:", cmd )
        fmt.Println("Running:" , cmd.Process.Pid)
	//isProcessRunning(runningPid)

	//if processErr := cmd.Wait(); processErr != nil {
	// if exiterr, ok := err.(*exec.ExitError); ok {
          //  if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
            //    log.Printf("Exit Status: %d", status.ExitStatus())
          //  }
      //  } else {
       //     log.Fatalf("cmd.Wait: %v", err)
      //  }
//}
	cmd.Wait()
	exitStatus := cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	fmt.Println("Error:", err)
        fmt.Println("Status:", exitStatus)
	if exitStatus == -1 {
	for{
	if  (state != "stopped" ) {
	fmt.Println("Process : ", processName, "stopped outside of controller: Restarting")
	cmd := exec.Command(processName, arg1)
        cmd.Stdout = os.Stdout
        cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
      //args :=  []string{"-i"}a
      	err := cmd.Start()
      	if err !=nil {
                fmt.Fprintln(os.Stderr, err)
                fmt.Println("Error: not Started" , processName)
                os.Exit(1)
        }
        runningPid = cmd.Process.Pid
        fmt.Println("Application Successfully Restarted!:", cmd )
        fmt.Println("Running:" , cmd.Process.Pid)
	cmd.Wait() 
}
	}
	}
       // close(done)

} else if (state == "stopped" ) {
	isProcessRunning(runningPid)
	fmt.Println("Stopping:" , processName)
	fmt.Println("Stopping PID: ", runningPid)
        //syscall.Kill(runningPid, syscall.SIGKILL)
	args = processName
	err2 := filepath.Walk("/proc", findAndKillProcess)
        if err2 != nil {
                if err2 == io.EOF {
            // Not an error, just a signal when we are done
                err2 = nil
        } else {
            log.Fatal(err2)

}
}


}
      return a
}







func findProcess(path string, info os.FileInfo, err error) error {
    if err != nil {
        return nil
    }
    processName := path
    fmt.Printf("findProcess:: PROCESSNAME1::%s\n", path)

    if strings.Count(path, "/") == 3 {
        if strings.Contains(path, "/status") {

            pid, err := strconv.Atoi(path[6:strings.LastIndex(path, "/")])
                 fmt.Printf("TEST::%d", pid)
            if err != nil {
                log.Println(err)
                return nil
            }

            f, err := ioutil.ReadFile(path)
            if err != nil {
                log.Println(err)
                return nil
            }

            name := string(f[6:bytes.IndexByte(f, '\n')])
                        fmt.Printf("findProcess:: PROCESSNAME::%s NAME::%s PATH::%s PID::%d \n",processName, name, path, pid)
            if name == args {
                fmt.Printf("PID: %d, Name: %s is running.\n", pid, name)
                proc, err := os.FindProcess(pid)
                if err != nil {
                    log.Println(err)
                }
                fmt.Printf("PROC %d\n", proc)
                pgid, err := syscall.Getpgid(pid)
                fmt.Printf("PGID %d\n", pgid)
                if err == nil {
		fmt.Printf("PID: %d, Name: %s is running.\n", pid, name)

                }
                return io.EOF
            }

        }
    }
 return nil
}



func stopProcess(pid string) string{
	pidToKill, err := strconv.Atoi(pid)
	fmt.Sprintf("%d", pidToKill)
	if err !=nil {
		fmt.Sprintf("Failed to kill: %s ", err)
	}
	//fmt.Sprintf("%d", pidToKill)
	syscall.Kill(pidToKill, syscall.SIGTERM)
	return (pid) 
}

func isProcessRunning(pid int) bool {
        process, err := os.FindProcess(pid)
        if err != nil {
            fmt.Printf("Failed to find process: %s\n", err)
	    return false
        } else {
            err := process.Signal(syscall.Signal(0))
            fmt.Printf("process.Signal on pid %d returned: %v\n", pid, err)
	    return true
        }
}

func findAndKillProcess(path string, info os.FileInfo, err error) error {
    if err != nil {
        return nil
    }
    processName := path

    if strings.Count(path, "/") == 3 {
        if strings.Contains(path, "/status") {

            pid, err := strconv.Atoi(path[6:strings.LastIndex(path, "/")])
		 fmt.Printf("TEST::%d", pid)
            if err != nil {
                log.Println(err)
                return nil
            }

            f, err := ioutil.ReadFile(path)
            if err != nil {
                log.Println(err)
                return nil
            }

            name := string(f[6:bytes.IndexByte(f, '\n')])
			fmt.Printf("PROCESSNAME::%s NAME::%s PATH::%s PID::%d \n",processName, name, path, pid)
            if name == args {
                fmt.Printf("PID: %d, Name: %s will be killed.\n", pid, name)
                proc, err := os.FindProcess(pid)
                if err != nil {
                    log.Println(err)
                }
		fmt.Printf("PROC %d\n", proc)
		pgid, err := syscall.Getpgid(pid)
		fmt.Printf("PGID %d\n", pgid)
		if err == nil {
		syscall.Kill(-pid, syscall.SIGKILL)
		}
                return io.EOF
            }

        }
    }

    return nil
}
