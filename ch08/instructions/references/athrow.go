package references

import "reflect"
import "jvmgo/ch08/instructions/base"
import "jvmgo/ch08/rtda"
import "jvmgo/ch08/rtda/heap"

// Throw exception or error
type ATHROW struct{ base.NoOperandsInstruction }

// athrow instruction has a operand which is exceptionRef, pop from stack
func (self *ATHROW) Execute(frame *rtda.Frame) {
	ex := frame.OperandStack().PopRef()
	if ex == nil {
		panic("java.lang.NullPointerException")
	}

	thread := frame.Thread()
	if !findAndGotoExceptionHandler(thread, ex) {
		handleUncaughtException(thread, ex)
	}
}

// try find and goto exception hanlder code
// begin from currentFrame, traversal JVM and look for exception handler table
// assume now travel to Frame F, if can not find handler in F, Pop F and continue to traversal
// if can find in f, before jump to hanlder code, clear F'OperandStack and push exceptionRef
func findAndGotoExceptionHandler(thread *rtda.Thread, ex *heap.Object) bool {
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC() - 1

		handlerPC := frame.Method().FindExceptionHandler(ex.Class(), pc)
		if handlerPC > 0 {
			stack := frame.OperandStack()
			stack.Clear()
			stack.PushRef(ex)
			frame.SetNextPC(handlerPC)
			return true
		}

		thread.PopFrame()
		if thread.IsStackEmpty() {
			break
		}
	}
	return false
}

// todo
func handleUncaughtException(thread *rtda.Thread, ex *heap.Object) {
	thread.ClearStack()

	jMsg := ex.GetRefVar("detailMessage", "Ljava/lang/String;")
	goMsg := heap.GoString(jMsg)
	println(ex.Class().JavaName() + ": " + goMsg)

	stes := reflect.ValueOf(ex.Extra())
	for i := 0; i < stes.Len(); i++ {
		ste := stes.Index(i).Interface().(interface {
			String() string
		})
		println("\tat " + ste.String())
	}
}