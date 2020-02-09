package classfile
// ConstantValue is fixed length atrribute, to represent constant value type
// attribute_length must be 2 size ; constantvalue_index in a cp index,
type ConstantValueAttribute struct{
	constantValueIndex uint16 	// point to constant type in cp, e.g.CONSTANT_Integer_info or CONSTANT_Long_info
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader){
	self.constantValueIndex = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantValueIndex() uint16{
	return self.constantValueIndex
}