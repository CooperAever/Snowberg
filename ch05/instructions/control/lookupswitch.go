// switch-case have two ways to implement
// the second switch-case implementation is lookupswitch


package control
import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

type LOOKUP_SWITCH struct{
	defaultOffset int32
	npairs int32 	// key-value pair number
	// kind like map, key is case , value is jump offset
	// store as an array, [key1,value1,key2,value2,……]
	matchOffsets []int32 	

}

func (self *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader){
	reader.SkipPadding()	// there are 0~3 bytes padding need to skip
	self.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	self.matchOffsets = reader.ReadInt32s(self.npairs *2)
}

// pop an int value from operandStack, then use it find  matchOffsets
// if can find matched key,jump according to key-value
// if can't find, jump according to defaultOffset
func (self *LOOKUP_SWITCH) Execute(frame *rtda.Frame){
	key := frame.OperandStack().PopInt()
	for i := int32(0);i<self.npairs*2;i+=2{
		if self.matchOffsets[i] == key{
			offset := self.matchOffsets[i+1]
			base.Branch(frame,int(offset))
			return
		}
	}
	base.Branch(frame,int(self.defaultOffset))
}
