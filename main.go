// main executable.
package main

import (
	"fmt"
	"os"

	"github.com/bluenviron/mavp2p/mavp2p"
)

func main() {
	p, err := mavp2p.NewProgramFromCLI(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERR: %s\n", err)
		os.Exit(1)
	}
	defer p.Close()

	p.Wait()
}
