// define 6 if<cond> instructions
package comparisons
import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

//Branch if int comparison with zero succeeds

type IFEQ struct{base.BranchInstruction}
type IFNE struct{base.BranchInstruction}
type IFLT struct{base.BranchInstruction}
type IFLE struct{base.BranchInstruction}
type IFGT struct{base.BranchInstruction}
type IFGE struct{base.BranchInstruction}

// if<cond> instructions pop operandstack top int value x,
// compare to 0, if meet the condition, jump
// ifeq: x == 0
// ifne: x != 0
// iflt: x < 0
// ifle: x <= 0
// ifgt: x > 0
// ifge: x >= 0

// Branch() defined in instructions\base\branch_logic.go
func (self *IFEQ) Execute(frame *rtda.Frame){
	val := frame.OperandStack().PopInt()
	if val == 0{
		base.Branch(frame,self.Offset)
	}
}

func (self *IFNE) Execute(frame *rtda.Frame){
	val := frame.OperandStack().PopInt()
	if val != 0{
		base.Branch(frame,self.Offset)
	}
}

func (self *IFLT) Execute(frame *rtda.Frame){
	val := frame.OperandStack().PopInt()
	if val < 0{
		base.Branch(frame,self.Offset)
	}
}

func (self *IFLE) Execute(frame *rtda.Frame){
	val := frame.OperandStack().PopInt()
	if val <= 0{
		base.Branch(frame,self.Offset)
	}
}

func (self *IFGT) Execute(frame *rtda.Frame){
	val := frame.OperandStack().PopInt()
	if val > 0{
		base.Branch(frame,self.Offset)
	}
}

func (self *IFGE) Execute(frame *rtda.Frame){
	val := frame.OperandStack().PopInt()
	if val >= 0{
		base.Branch(frame,self.Offset)
	}
}

