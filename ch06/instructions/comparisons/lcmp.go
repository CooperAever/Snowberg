// comparisons can be divided into 2 part:
// 1. push compare result into operandStack
// 2. according to compare result to jump,like if_else
package comparisons
import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

// compare long
type LCMP struct{base.NoOperandsInstruction}

func(self *LCMP) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	}else if v1 < v2{
		stack.PushInt(-1)
	}else{
		stack.PushInt(0)
	}
}
