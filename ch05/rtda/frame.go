package rtda

// every method invoke correspond to a frame
type Frame struct{
	lower *Frame 	//used for implement stack,only need upper
	localVars LocalVars 		//store pointer point to local variable 
	operandStack *OperandStack 	//store operandStack pointer
}

// maxLocals and maxStack are pre-computed, store in Code attribute in class document
func NewFrame(maxLocals,maxStack uint) *Frame{
	return &Frame{
		localVars: newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

// getters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
