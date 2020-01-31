package main
import "fmt"

// main function is the extrance of program
func main(){
	cmd := parseCmd()
	if cmd.versionFlag{
		fmt.Println("Snowberg version 0.0.1")
	}else if cmd.helpFlag || cmd.class == ""{
		printUsage()
	}else{
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd){
	fmt.Printf("classpath:%s class:%s args:%v\n",cmd.cpOption,cmd.class,cmd.args)
}

