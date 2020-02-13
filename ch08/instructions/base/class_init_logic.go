package base

import "jvmgo/ch08/rtda"
import "jvmgo/ch08/rtda/heap"

// jvms 5.5

func InitClass(thread *rtda.Thread, class *heap.Class) {
	class.StartInit() // set initStarted to true
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}

// prepare execute class initial method
func scheduleClinit(thread *rtda.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		// exec <clinit>
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}

func initSuperClass(thread *rtda.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}