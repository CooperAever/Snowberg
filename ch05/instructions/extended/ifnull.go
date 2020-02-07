package extended
import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"
type IFNULL struct{base.BranchInstruction} //branch is reference is null
type IFNONNULL struct{base.BranchInstruction} //branch if reference not null

// decide wether jump based on reference == null
// ifnull and ifnonnull instructions pop stack top reference
func(self *IFNULL) Execute(frame *rtda.Frame){
	ref := frame.OperandStack().PopRef()
	if ref == nil{
		base.Branch(frame,self.Offset)
	}
}

func(self *IFNONNULL) Execute(frame *rtda.Frame){
	ref := frame.OperandStack().PopRef()
	if ref != nil{
		base.Branch(frame,self.Offset)
	}
}

