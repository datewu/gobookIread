package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime/debug"
)

// MyError contain all of well-formed error
type MyError struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]interface{}
}

func wrapError(err error, messagef string, msgArgs ...interface{}) MyError {
	return MyError{
		err,
		fmt.Sprintf(messagef, msgArgs...),
		string(debug.Stack()),
		make(map[string]interface{}),
	}
}

func (err MyError) Error() string {
	return err.Message
}

type lowLevelErr struct {
	error
}

func isGloballyExec(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, lowLevelErr{wrapError(err, err.Error())}
	}
	return info.Mode().Perm()&0100 == 0100, nil
}

type intermediateErr struct {
	error
}

func runJob(id string) error {
	const jobBinPath = "/bad/job/binary"
	isExecutable, err := isGloballyExec(jobBinPath)
	if err != nil {
		return intermediateErr{
			wrapError(err, "cannot run job %q: requisite binaries not available", id),
		}
	} else if isExecutable == false {
		return wrapError(nil, "cannot run job %q: requisite binaries are not executable", id)
	}
	return exec.Command(jobBinPath, "--id="+id).Run()
}

func handleError(key int, err error, message string) {
	log.SetPrefix(fmt.Sprintf("[logID: %v]: ", key))
	log.Printf("%#v", err)
	fmt.Printf("[%v] %v", key, message)
}

func runErrorSample() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	err := runJob("8")
	if err != nil {
		msg := "There was an unexpected issue; please report this as a bug."
		if _, ok := err.(intermediateErr); ok {
			msg = err.Error()
		}
		handleError(8, err, msg)
	}

}
