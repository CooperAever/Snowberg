package main
import (
	"flag"	// used for handle cmd instrustions
	"fmt"
	"os" 
)

type Cmd struct{
	helpFlag bool
	versionFlag bool
	cpOption string
	class string
	args []string
}


func parseCmd() *Cmd{
	cmd := &Cmd{}		// equal to cmd := new(Cmd)
	flag.Usage = printUsage		// set printUsage() value to flag.Usage
	flag.BoolVar(&cmd.helpFlag,"help",false,"print help message") 	// return a pointer to argument 1 with content arguments 2,3,4
	flag.BoolVar(&cmd.helpFlag,"?",false,"print help message")
	flag.BoolVar(&cmd.versionFlag,"version",false,"print version and exit")
	flag.StringVar(&cmd.cpOption,"classpath","","classpath")
	flag.StringVar(&cmd.cpOption,"cp","","classpath")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0{
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

// print method of using cmd to command monitor
func printUsage(){
	fmt.Printf("Usage: %s  [-options] class [args...] \n",os.Args[0])
}

