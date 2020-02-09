package classfile


// Code is a variable length attr(not fixed length)
// store byte code info
type CodeAttribute struct{
	cp ConstantPool
	maxStack uint16 	// max depth of operator stack (ch04)
	maxLocals uint16  	// the size of local variable table (ch04)
	code []byte 		//byte code (ch05)
	exceptionTable []*ExceptionTableEntry 		// exception handler table (ch10)
	attributes []AttributeInfo 			// attribute table
}


type ExceptionTableEntry struct{
	startPc uint16
	endPc uint16
	handlerPc uint16
	catchType uint16
}

func (self *CodeAttribute) readInfo(reader *ClassReader){
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader,self.cp)
}

func (self *CodeAttribute) MaxStack() uint {
	return uint(self.maxStack)
}
func (self *CodeAttribute) MaxLocals() uint {
	return uint(self.maxLocals)
}
func (self *CodeAttribute) Code() []byte {
	return self.code
}
func (self *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return self.exceptionTable
}


func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry{
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([] *ExceptionTableEntry,exceptionTableLength)
	for i := range exceptionTable{
		exceptionTable[i] = &ExceptionTableEntry{
			startPc : reader.readUint16(),
			endPc : reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType:reader.readUint16(),
		}
	}
	return exceptionTable
}

func (self *ExceptionTableEntry) StartPc() uint16 {
	return self.startPc
}
func (self *ExceptionTableEntry) EndPc() uint16 {
	return self.endPc
}
func (self *ExceptionTableEntry) HandlerPc() uint16 {
	return self.handlerPc
}
func (self *ExceptionTableEntry) CatchType() uint16 {
	return self.catchType
}