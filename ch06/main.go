package main


import "fmt"
import "strings"
import "jvmgo/ch06/classfile"
import "jvmgo/ch06/classpath"
import "jvmgo/ch06/rtda/heap"

// ch01 make cmd parse tools
// ch02 talk about how to search class document
// ch03 focus on parsing class document
// ch04 implement run-time data area
// ch05 first implement thread private run-time data area(Thread、Stack、Frame、OperandStack、LocalVars)
// ch06 parse codebyte and use a loop to execute instructions(1.count pc,2.decode instruction type,3.execute)

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
	classLoader := heap.NewClassLoader(cp)
	className := strings.Replace(cmd.class,".","/",-1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil{
		interpret(mainMethod)
	}else{
		fmt.Printf("Main method not found in class %s\n",cmd.class)
	}

}

func loadClass(className string,cp *classpath.Classpath) *classfile.ClassFile{
	classData,_,err := cp.ReadClass(className)
	if err != nil{
		panic(err)
	}
	cf,err := classfile.Parse(classData)
	if err != nil{
		panic(err)
	}
	return cf
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo{
	for _,m := range cf.Methods(){
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V"{
			return m
		}
	}
	return nil
}