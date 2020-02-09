// conversion instruction correspond to java force conversion
// conversion instruction can be divide 3 part:
// 1. i2x -> int force cast to other
// 2. l2x -> long force cast to other
// 3. f2x -> float force cast to other
// 4. d2x -> double force cast to other

package conversions
import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

type D2F struct{base.NoOperandsInstruction}
type D2I struct{base.NoOperandsInstruction}
type D2L struct{base.NoOperandsInstruction}

func(self *D2I) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

func (self *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

func (self *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}