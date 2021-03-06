// switch-case have two ways to implement
// the first switch-case implementation is code case stituation into a table
// use index to jump
package control
import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

//Access jump table by index and jump
type TABLE_SWITCH struct{
	defaultOffset int32
	low int32 		// case low bound
	high int32 		// case high bound
	jumpOffsets []int32 	//  a index table
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader){
	reader.SkipPadding()	// there are 0~3 bytes padding need to skip
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low +1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame){
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= self.low && index <= self.high{
		offset = int(self.jumpOffsets[index-self.low])
	}else{
		offset = int(self.defaultOffset)
	}
	base.Branch(frame,offset)
}

