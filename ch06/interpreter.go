package main
import "fmt"
import "jvmgo/ch06/classfile"
import "jvmgo/ch06/instructions"
import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"

// interpreter() method has argument MemberInfo ref,
// through MemberInfo.CodeAttribute() can get code attribute
func interpreter(methodInfo *heap.Method){
	thread := rtda.NewThread()
	frame := thread.NewFrame(method) //defined in rtda/thread.go
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread,method.Code())
}

func catchErr(frame *rtda.Frame){
	// recover use to catch panic() and excute before program terminated
	if r := recover(); r!=nil{
		fmt.Printf("localVars:%v \n",frame.LocalVars())
		fmt.Printf("OperandStack:%v \n",frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread,bytecode []byte){
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		// 1.read pc
		pc := frame.NextPC()
		thread.SetPC(pc)

		// 2.decode instruction
		reader.Reset(bytecode,pc)
		opcode := reader.ReadUint8() 	//read opcode to define instruction type
		inst := instructions.NewInstruction(opcode) 	// in "jvmgo/ch05/instructions/factory.go"
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		// 3.execute
		fmt.Printf("pc:%2d  inst:%T %v\n",pc,inst,inst)
		inst.Execute(frame)
	}
}







