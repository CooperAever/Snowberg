// stack instructions used for operator stack
// pop and pop2 used to poping stack top
// dup used to duplicate top value
// swap used to swap top variable
package stack
import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

type POP struct{base.NoOperandsInstruction}

type POP2 struct{base.NoOperandsInstruction}

// pop fit for 1 size variable,like int、float,etc
func(self* POP) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PopSlot()
}

// pop2 fit for 2 size variable,like double and long
func(self* POP2) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}



