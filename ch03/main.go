package main
import "fmt"
import "strings"
import "jvmgo/ch03/classpath"

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
	cp := classpath.Parse(cmd.XjreOption,cmd.cpOption)
	fmt.Printf("classpath:%v class:%s args:%v\n",cp,cmd.class,cmd.args)
	className := strings.Replace(cmd.class,".","/",-1)		//func Replace(s, old, new string, n int) string
	classData,_,err := cp.ReadClass(className)
	if err != nil{
		fmt.Printf("Could not find or load main class %s\n",cmd.class)
		return
	}
	fmt.Printf("class data:%v\n",classData)
}

