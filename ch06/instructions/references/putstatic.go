package references
import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"

// Set static field in class
type PUT_STATIC struct{
	base.Index16Instruction
}

// give a static variable value,need two operand num
// first one is uint16 index,from byte code, can use it find which static variable need to set
// second one is value,pop from operandStack

func (self *PUT_STATIC) Execute(frame *rtda.Frame){
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()

	if !field.IsStatic(){
		panic("java.lang.IncompatobleClassChangeError")
	}
	if field.IsFinal(){
		// if its final field, only can set when initialize
		// and class initial method is "<clinit>"
		if currentClass != class || currentMethod.Name() != "<clinit>"{
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()
	switch descriptor[0]{
	case 'Z','B','C','S','I':
		slots.SetInt(slotId,stack.PopInt())
	case 'F':
		slots.SetFloat(slotId,stack.PopFloat())
	case 'J':
		slots.SetLong(slotId,stack.PopLong())
	case 'D':
		slots.SetDouble(slotId,stack.PopDouble())
	case 'L','[':
		slots.SetRef(slotId,stack.PopRef())
	}
}

