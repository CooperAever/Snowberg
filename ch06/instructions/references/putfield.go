package references
import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"

// Set field in object
type PUT_FIELD struct{
	base.Index16Instruction
}

// putfield instruction set instance variable,need 3 operands
// first two operand are constantpool index and value
// third operand is object ref,pop from operandStack

func (self *PUT_FIELD) Execute(frame *rtda.Frame){
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heapFieldRef)
	field := fieldRef.ResolveField()

	if field.IsStatic(){
		panic("java.lang.IncopatibleClassChangeError")
	}
	if field.IsFinal(){
		if currentClass != field.Class() || currentMethod.Name() != "<init>"{
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	stack := frame.OperandStack()
	switch descriptor[0]{
	case 'Z','B','C','S','I':
		val := stack.PopInt()
		ref := stack.PopRef()
		if ref == nil{
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetInt(slotId,val)
	case 'F':
		val := stack.PopFloat()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetFloat(slotId, val)
	case 'J':
		val := stack.PopLong()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetLong(slotId, val)
	case 'D':
		val := stack.PopDouble()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetDouble(slotId, val)
	case 'L', '[':
		val := stack.PopRef()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetRef(slotId, val)
	default:
		// todo
	}
}