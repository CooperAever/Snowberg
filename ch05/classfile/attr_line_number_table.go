package classfile
// lineNumberTable store method linenumber info
// localVariableTable store method local variable
// these two and sourceFile attribute are all debug info,not running necessary

type LineNumberTableAttribute struct{
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct{
	startPc uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader){
	lineNumberTableLength := reader.readUint16()
	self.lineNumberTable = make([]*LineNumberTableEntry,lineNumberTableLength)
	for i:= range self.lineNumberTable{
		self.lineNumberTable[i] = &LineNumberTableEntry{
			startPc : reader.readUint16(),
			lineNumber : reader.readUint16(),
		}
	}
}