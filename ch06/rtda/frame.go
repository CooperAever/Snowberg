package rtda

// every method invoke correspond to a frame
type Frame struct{
	lower *Frame 	//used for implement stack,only need upper
	localVars LocalVars 		//store pointer point to local variable 
	operandStack *OperandStack 	//store operandStack pointer
	thread *Thread 	//used for branch()
	nextPC int  	//used for branch()
	method *heap.Method
}

// maxLocals and maxStack are pre-computed, store in Code attribute in class document
func NewFrame(thread *Thread,method *heap.Method) *Frame{
	return &Frame{
		thread:thread,
		localVars: newLocalVars(method.maxLocals),
		operandStack: newOperandStack(method.maxStack),
		method : method,
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

func (self *Frame) Method() {
	return self.method
}