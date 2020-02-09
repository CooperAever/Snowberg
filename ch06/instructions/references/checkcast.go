package references
import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"

// check whether object is of given type
type CHECK_CAST struct{
	base.Index16Instruction
}

// checkcast is like instanceof,but instanceof will change operand
// checkcast will not change operand
// if checkcast is false, directly throw a ClassCaseException
func (self *CHECK_CAST) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil{
		return
	}
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class){
		panic("java.lang.ClassCaseException")
	}
}

