// new instruction used for creating class instance
// putstatic and getstatic instructions used for getting and storing static value
// putfield and getfield used for getting and storing instance value
// instanceof and checkcase used for juding object's type
// ldc serial instructions push running-time constant into operandStack

package references
import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"

// create new Object
type NEW struct{base.Index16Instruction}

// new instruction's operandNum is a uint16 index,find a symbolic reference in constant pool
// and create object, then push object reference into stack
func (self *NEW) Execute(frame *rtda.Frame){
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	// interface and abstract can not be instance
	if class.IsInterface() || class.IsAbstract(){
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}



