// math instruction can be divided into add、sub、mul、div、rem、neg.
package math

import "math"
import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

type DREM struct{base.NoOperandsInstruction}
type FREM struct{base.NoOperandsInstruction}
type IREM struct{base.NoOperandsInstruction}
type LREM struct{base.NoOperandsInstruction}


//get reminder of operandStack top two number
func(self *IREM) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0{
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1%v2
	stack.PushInt(result)

}

func(self *DREM) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	// because double type have infinity value,so dived 0 would invoke ArithmeticException
	result := math.Mod(v1,v2)
	stack.PushDouble(result)
}


func (self *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := float32(math.Mod(float64(v1), float64(v2))) // todo
	stack.PushFloat(result)
}

func (self *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 % v2
	stack.PushLong(result)
}
