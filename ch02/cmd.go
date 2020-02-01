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
	// add -Xjre to instruct jre catalogue location
	XjreOption string
	args []string
}


func parseCmd() *Cmd{
	cmd := &Cmd{}		// equal to cmd := new(Cmd)
	flag.Usage = printUsage		// set printUsage() value to flag.Usage
	flag.BoolVar(&cmd.helpFlag,"help",false,"print help message") 	// bind information to argument 1 , with content arguments 2(flagname),3(value),4(usage)
	flag.BoolVar(&cmd.helpFlag,"?",false,"print help message")
	flag.BoolVar(&cmd.versionFlag,"version",false,"print version and exit")
	flag.StringVar(&cmd.cpOption,"classpath","","classpath")
	flag.StringVar(&cmd.cpOption,"cp","","classpath")
	flag.StringVar(&cmd.XjreOption,"Xjre","","path to jre")
	flag.Parse()	// parse cmd , and the result can be accessed by flag.Args()
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

