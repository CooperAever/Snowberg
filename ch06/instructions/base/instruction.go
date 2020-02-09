package base
import "jvmgo/ch06/rtda"

type Instruction interface{
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

// the instruction with no operand 
type NoOperandsInstruction struct{} 

func(self *NoOperandsInstruction)FetchOperands(reader *BytecodeReader){
	// nothing to do
}

// the instruction used for jump ,namely plus offset
type BranchInstruction struct{
	Offset int
}

func(self *BranchInstruction) FetchOperands(reader *BytecodeReader){
	self.Offset = int(reader.ReadInt16())
}

// the instruction give localVars index
type Index8Instruction struct{
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader){
	self.Index = uint(reader.ReadUint8())
}

// the instruction give localVars index, some index give by two bytes
type Index16Instruction struct{
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader){
	self.Index = uint(reader.ReadUint16())
}

