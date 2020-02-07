package main
import "fmt"
import "jvmgo/ch05/classfile"
import "jvmgo/ch05/instructions"
import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// interpreter() method has argument MemberInfo ref,
// through MemberInfo.CodeAttribute() can get code attribute
func interpreter(methodInfo *classfile.MemberInfo){
	// get code info
	codeAttr := methodInfo.CodeAttribute()
	// we can further get localVars and operandStack and method's bytecode
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()
	// create a thread instance , and create a frame then push to jvm stack
	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals,maxStack)
	thread.PushFrame(frame)
	defer catchErr(frame) 		// not implement return instruction, will occur an error and be catched
	loop(thread,bytecode)
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







