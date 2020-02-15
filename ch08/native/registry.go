// before implement native method, first implement a native method register
// 
package native
import "jvmgo/ch08/rtda"

// define NativeMethod as a function
// argument frame is nativeMethod work place, also the bridge connect JVM and java class
type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}

// className, methodName, methodDescriptor all together can  define a method
func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

// java.lang.Object and other class use a registerNatives() native method to register other native method
func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}
	return nil
}

func emptyNativeMethod(frame *rtda.Frame) {
	// do nothing because have already implement all register native method
}
