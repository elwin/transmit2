package main

import (
	"fmt"
	"io"
	"os"

	"github.com/elwin/transmit2/mode"

	ftp "github.com/elwin/transmit2/client"
)

func main() {

	if err := run(); err != nil {
		fmt.Print(err)
	}

}

func run() error {
	conn, err := ftp.Dial("localhost:2121", ftp.DialWithDebugOutput(os.Stdout))
	if err != nil {
		return err
	}

	err = conn.Login("admin", "123456")
	if err != nil {
		return err
	}

	err = conn.Mode(mode.ExtendedBlockMode)
	if err != nil {
		return err
	}

	/*
		err = conn.Stor("stor.txt", strings.NewReader("Hello World!"))
		if err != nil {
			return err
		}
	*/

	res, err := conn.Retr("retr.txt")

	f, err := os.Create("/Users/elwin/ftp/result.txt")
	if err != nil {
		return err
	}

	n, err := io.Copy(f, res)
	if err != nil {
		return err
	}

	fmt.Printf("Read %d bytes\n", n)

	res.Close()

	entries, err := conn.List("/")
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fmt.Printf("- %s (%d)\n", entry.Name, entry.Size)
	}

	return conn.Quit()
}
