// exception object not a normal object, if check java.lang.Exception and RuntimeException
// can find their construction function invoke java.lang.Throwable construction, and Throwable construction
// again invoke fillInStackTrace().
// means if want to throw exception , JVM need implement fillInStackTrace(int) method
package lang

import "fmt"
import "jvmgo/ch08/native"
import "jvmgo/ch08/rtda"
import "jvmgo/ch08/rtda/heap"

const jlThrowable = "java/lang/Throwable"

// StackTraceElement struct used for tracing jvm frame info
type StackTraceElement struct {
	fileName   string 		// file name
	className  string
	methodName string
	lineNumber int 			// current execute which line
}

func (self *StackTraceElement) String() string {
	return fmt.Sprintf("%s.%s(%s:%d)",
		self.className, self.methodName, self.fileName, self.lineNumber)
}

func init() {
	native.Register(jlThrowable, "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

// private native Throwable fillInStackTrace(int dummy);
// (I)Ljava/lang/Throwable;
func fillInStackTrace(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)

	stes := createStackTraceElements(this, frame.Thread())
	this.SetExtra(stes)
}

// because stack top two frame is fillInStackTrace(int) and fillInStackTrace(),so need to skip 2 frame
// and next these two frames, there are some exceptionClass's construction method, also need skip, skip number correspond to class extends number
func createStackTraceElements(tObj *heap.Object, thread *rtda.Thread) []*StackTraceElement {
	skip := distanceToObject(tObj.Class()) + 2 	// compute skip number
	frames := thread.GetFrames()[skip:] 	// get needed JVM Frames
	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}

func distanceToObject(class *heap.Class) int {
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}
	return distance
}

func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		fileName:   class.SourceFile(),
		className:  class.JavaName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}