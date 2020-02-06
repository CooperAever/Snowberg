package main

import "fmt"
import "jvmgo/ch04/rtda"
// import "strings"
// import "jvmgo/ch04/classpath"
// import "jvmgo/ch04/classfile"

// ch01 make cmd parse tools
// ch02 talk about how to search class document
// ch03 focus on parsing class document
// ch04 implement run-time data area
// ch05 first implement thread private run-time data area(Thread、Stack、Frame、OperandStack、LocalVars)

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
	// cp := classpath.Parse(cmd.XjreOption,cmd.cpOption)
	// fmt.Printf("classpath:%v class:%s args:%v\n",cp,cmd.class,cmd.args)
	// className := strings.Replace(cmd.class,".","/",-1)		//func Replace(s, old, new string, n int) string
	// cf := loadClass(className,cp)
	// fmt.Println(cmd.class)
	// printClassInfo(cf)

	frame := rtda.NewFrame(100,100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())
}

func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
}

func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}