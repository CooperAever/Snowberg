package rtda

// every method invoke correspond to a frame
type Frame struct{
	lower *Frame 	//used for implement stack,only need upper
	localVars LocalVars 		//store pointer point to local variable 
	operandStack *OperandStack 	//store operandStack pointer
	thread *Thread 	//used for branch()
	nextPC int  	//used for branch()
}

// maxLocals and maxStack are pre-computed, store in Code attribute in class document
func NewFrame(thread *Thread,maxLocals,maxStack uint) *Frame{
	return &Frame{
		thread:thread,
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

func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
