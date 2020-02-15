package lang

import "unsafe"
import "jvmgo/ch08/native"
import "jvmgo/ch08/rtda"

const jlObject = "java/lang/Object"

func init() {
	native.Register(jlObject, "getClass", "()Ljava/lang/Class;", getClass)
	native.Register(jlObject, "hashCode", "()I", hashCode)
	native.Register(jlObject, "clone", "()Ljava/lang/Object;", clone)
}

//public final native Class<?> getClass();
func getClass(frame *rtda.Frame){
	this := frame.LocalVars().GetThis() 	//getRef(0)
	class := this.Class().JClass()			// get classObject
	frame.OperandStack().PushRef(class)		// push classOnject into operandStack
}

// public native int hashCode();
// ()I
func hashCode(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()	// get object ref
	hash := int32(uintptr(unsafe.Pointer(this))) //transform ref into uintptr, then into int32 and push into operandStack
	frame.OperandStack().PushInt(hash)
}

// protected native Object clone() throws CloneNotSupportedException;
// ()Ljava/lang/Object;
func clone(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()

	cloneable := this.Class().Loader().LoadClass("java/lang/Cloneable")
	if !this.Class().IsImplements(cloneable) { 	// if class not implment clone interface
		panic("java.lang.CloneNotSupportedException")
	}

	frame.OperandStack().PushRef(this.Clone())
}


