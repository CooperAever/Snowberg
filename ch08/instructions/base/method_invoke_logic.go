package base

import "fmt"
import "jvmgo/ch08/rtda"
import "jvmgo/ch08/rtda/heap"

// after locating the invoke method , JVM need create a new Frame for this method
// and push this frame into JVMstack,then pass argument

func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	// create a new Frame and push into JVMstack
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)
	// pass argument

	// first count argument will occupy how many space in LocalVars
	// space not equal to argument size:1.maybe long and double,2.will contain a hidden argument-> this
	argSlotCount := int(method.ArgSlotCount())
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			// pop variable from invokerFrame operandStack and put into newFrame LocalVars in proper order	
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}

	// hack!
	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		} else {
			panic(fmt.Sprintf("native method: %v.%v%v\n",
				method.Class().Name(), method.Name(), method.Descriptor()))
		}
	}
}