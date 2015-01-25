package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type ContestantProcess struct {
	dir string
	cmd *exec.Cmd
}

func NewContestantProcess(contestant *Contestant) (*ContestantProcess, error) {
	var err error
	cp := &ContestantProcess{}

	// Create directory and import AI code
	cp.dir, err = ioutil.TempDir("", "sandbox")
	if err != nil {
		return nil, err
	}

	ai := filepath.Join(cp.dir, "main.go")
	if err := ioutil.WriteFile(ai, []byte(contestant.Code), 0400); err != nil {
		return nil, fmt.Errorf("error creating temp file %q: %v", ai, err)
	}

	exe := filepath.Join(cp.dir, "a.out")
	cmd := exec.Command("go", "build", "-o", exe, ai)
	cmd.Env = []string{"GOOS=nacl", "GOARCH=amd64p32", "GOPATH=/go"}
	if out, err := cmd.CombinedOutput(); err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			// Error compiling AI
			return nil, fmt.Errorf("Error compiling AI: %s", string(out))
		}
		return nil, fmt.Errorf("error building go source: %v", err)
	}

	return cp, nil
}

func (cp *ContestantProcess) Run() (time.Duration, error) {
	// Prepare AI to receive requests
	exe := filepath.Join(cp.dir, "a.out")
	cp.cmd = exec.Command("sel_ldr_x86_64", "-l", "/dev/null", "-S", "-e", exe)

	start := time.Now()
	err := cp.cmd.Start()
	if err != nil {
		return 0, fmt.Errorf("error starting AI: %v", err)
	}

	resc := make(chan time.Duration, 1)
	errc := make(chan error, 1)

	go func() {
		err := cp.cmd.Wait()
		if err != nil {
			errc <- err
			return
		}
		diff := time.Now().Sub(start)
		resc <- diff
		return
	}()

	t := time.NewTimer(time.Second)

	select {
	case err := <-errc:
		t.Stop()
		return 0, err
	case res := <-resc:
		t.Stop()
		log.Printf("executed piece in %v", res)
		return res, nil
	case <-t.C:
		cp.cmd.Process.Kill()
		return 0, fmt.Errorf("timeout")
	}
}

func (cp *ContestantProcess) Close() {
	os.RemoveAll(cp.dir)
	cp.cmd.Process.Kill()
}
