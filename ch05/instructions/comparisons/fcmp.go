// fcmpg and fcmpl used for comparing float variable
package comparisons
import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type FCMPG struct{base.NoOperandsInstruction}
type FCMPL struct{base.NoOperandsInstruction}

// float comparison like lcmp, but float compare may has one more result
// 4.NaN(not a number),means can not compare
func _fcmp(frame *rtda.Frame,gFlag bool){
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2{
		stack.PushInt(1)
	}else if v1 == v2{
		stack.PushInt(0)
	}else if v1 < v2{
		stack.PushInt(-1)
	}else if gFlag{
		stack.PushInt(1)
	}else{
		stack.PushInt(-1)
	}
}

// FCMPG and FCMPL differ from defination of result 4
// when two float variable is at least one NaN,fcmpg result is 1,
// fcmpl is -1
func (self* FCMPG)Execute(frame *rtda.Frame){
	_fcmp(frame,true)
}

func (self* FCMPL)Execute(frame *rtda.Frame){
	_fcmp(frame,false)
}