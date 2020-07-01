package cli

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

type Cli struct {
	F    *os.File
	In   chan []byte
	Out  chan []byte
	Err  chan []byte
	Done chan struct{}
}

func NewCli() (*Cli, error) {
	cmd := exec.Command("sh")
	cmd.Env = []string{"TERM=xterm"}

	f, err := pty.Start(cmd)
	if err != nil {
		return nil, err
	}

	in := make(chan []byte, 1024)
	out := make(chan []byte, 1024)
	done := make(chan struct{})

	cli := &Cli{
		F:    f,
		In:   in,
		Out:  out,
		Done: done,
	}

	//send to pty
	go func(c *Cli) {
		defer func() {
			log.Print("PTY input was closed")
			close(c.In)
		}()
		for {
			select {
			case com := <-c.In:
				c.Do(com)
			case <-c.Done:
				return
			}
		}
	}(cli)

	// read from pty
	go func(c *Cli) {
		defer func() {
			log.Print("PTY output was closed")
			close(c.Out)
		}()
		liner := bufio.NewReaderSize(cli.F, 1024)
		for {
			line, _, err := liner.ReadLine()
			if err != nil && err != io.EOF {
				log.Fatal(err)
			}
			c.Out <- line
			if err != nil && err == io.EOF {
				log.Fatal(err)
				break
			}
		}
	}(cli)

	return cli, nil
}

func (c *Cli) Do(com []byte) {
	log.Printf("Command '%s' was receved", string(com))
	_, err := c.F.Write([]byte(string(com) + "\n"))
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Cli) Close() {
	log.Print("PTY closing...")
	c.F.Close()
	close(c.Done)
}
