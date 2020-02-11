package references
import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"
import "jvmgo/ch07/rtda/class"

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

	base.InvokeMethod(frame, resolvedMethod)
}



