package rtda
// use linked list data structure to inplement JVM stack
// so stack can use memory, and pop frame can be collected by GC


type Stack struct{
	maxSize uint 
	size uint
	_top *Frame
}

func newStack(maxSize uint) *Stack{
	return &Stack{
		maxSize:maxSize,
	}
}

func (self *Stack) push(frame *Frame){
	if self.size >= self.maxSize{
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil{
		frame.lower = self._top
	}
	self._top = frame
	self.size++
}

func (self *Stack) pop() *Frame{
	if self._top == nil{
		panic("jvm stack is empty!")
	}
	top := self._top
	self._top = top.lower
	self.size --
	return top
}

func (self *Stack) top() *Frame{
	if self._top == nil{
		panic("jvm stack is empty!")
	}
	return self._top
}