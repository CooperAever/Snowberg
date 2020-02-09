package constants

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

//bipush instruction get a byte from operand, transform to int and push into stack
type BIPUSH struct{
	val int8
}

//sipush instruction get a short from operand, transform to int and push into stack
type SIPUSH struct{
	val int16
}

func(self *BIPUSH) FetchOperands(reader *base.BytecodeReader){
	self.val = reader.ReadInt8()
}

func(self *BIPUSH) Execute(frame *rtda.Frame){
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

func(self *SIPUSH) FetchOperands(reader *base.BytecodeReader){
	self.val = reader.ReadInt16()
}

func(self *SIPUSH) Execute(frame *rtda.Frame){
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}