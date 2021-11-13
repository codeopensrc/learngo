package main

import (
    "os"
    "os/exec"
    "fmt"
    "strings"
)

func RunCommand(myCommand string, myArgs []string) (string, error) {
    //Declare byte array and error type
    var cmdOutput []byte
    var err error

    //Execute external command and get output
    cmdOutput, err = exec.Command(myCommand, myArgs[:]...).Output()

    if err != nil {
        return "", err
    }
    return string(cmdOutput), nil
}

// Look at 'flag' built-in when coming back to this
//https://golangdocs.com/command-line-arguments-in-golang

func main() {
    //Get first arg
    programpath := os.Args[0]
    //Get all user provided args (start at 1 till end)
    userargs := os.Args[1:]

    //Default commmand
    cmdToRun := "date"

    //User provides at least a command
    if len(userargs) > 0 {
        if len(userargs) > 1 {
            //Command with args
            cmdToRun, userargs = userargs[0], userargs[1:]
        } else {
            //Otherwise only user command
            cmdToRun = userargs[0]
            userargs = nil
        }
    }
    output, err := RunCommand(cmdToRun, userargs)

    if err != nil {
        fmt.Printf("Found error running: '%s %s'.\nErrMsg: '%s'\n", cmdToRun, userargs, err)
    }

    fmt.Println("cmdToRun: " + cmdToRun)
    if len(userargs) > 0 {
        //Join []string (array or slice) into single string
        fmt.Println("User args: [ " + strings.Join(userargs[0:], ", ") + " ]")
    }

    fmt.Println("output: " + output)
    fmt.Println("Program path: " + programpath)
}
