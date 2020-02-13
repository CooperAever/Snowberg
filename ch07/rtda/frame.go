package rtda

import "jvmgo/ch07/rtda/heap"

// stack frame
type Frame struct {
	lower        *Frame // stack is implemented as linked list
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	method       *heap.Method
	nextPC       int // the next instruction after the call
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

// getters & setters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) Method() *heap.Method {
	return self.method
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
// if class not initialized, need first initialize
// but and this time, instruction execution is half way, means nextPC point to next instruction
// need change nextPc to current instruction
func (self *Frame) RevertNextPC() {
	self.nextPC = self.thread.pc
}