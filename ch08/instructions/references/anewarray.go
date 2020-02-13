package references
import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"
import "jvmgo/ch08/rtda/heap"

// Create new array of reference
type ANEW_ARRAY struct{base.Index16Instruction}

// anewarray need two operands, first is uint16 index, point to array type
// second is array length, pop from operandStack
func (self *ANEW_ARRAY) Execute(frame *rtda.Frame){
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	componentClass := classRef.ResolvedClass()
	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	arrClass := componentClass.ArrayClass() 	// return arrayclass correspond to right type
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}