package rtda

// there are two types of run-time data area:1.multi-thread share rtda 2.private rtda

// 1 include class data(stored in Method area)  and class instance(stored in heap)
// class data include fileds and methods info、methods' bytecode、constant pool,etc

// 2 used for help execute java bytecode,include program counter and local variable and operand Stack(stored in JVM stack)
// private means every thread have their own private rtda

// direct use go's heap and GC, so this version JVM may not need GC

type Thread struct{
	pc int 
	stack *Stack 	// a jvm pointer
}

// jvm stack can be sequential or not sequential, can be variable len or fixed len
// if JVM has size limit , may occur StackOverflowError
// if JVM can resize, but no more memory , may occur OutOfMemoryError
func NewThread() *Thread{
	return &Thread{
		stack : newStack(1024),		//arguments means max can contain 1024 frame
	}
}

// getter
func (self *Thread) PC() int {
	return self.pc
}

// setter
func (self *Thread) SetPC(pc int){
	self.pc = pc 
}


func (self *Thread) PushFrame(frame *Frame){
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame{
	return self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame{
	return self.stack.top()
}

func (self *Thread) NewFrame(maxLocals,maxStack uint) *Frame{
	return NewFrame(self,maxLocals,maxStack)
}