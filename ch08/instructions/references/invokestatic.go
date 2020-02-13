package references

import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"
import "jvmgo/ch08/rtda/heap"

// Invoke a class (static) method
type INVOKE_STATIC struct{base.Index16Instruction}

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	// assume parse methodRef and get method M , M must be static method
	// M can not initialized by class, If M not initialized , need first 
	// initialize and then invoke
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	class := resolvedMethod.Class()
	// check if class has been initialize
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	base.InvokeMethod(frame, resolvedMethod)
}


