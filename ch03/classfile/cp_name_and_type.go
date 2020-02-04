// field name or method name given by name_index 
// field descriptor or method descriptor given by descriptor_index
// name_index and descriptor_index all point to CONSTANT_Utf8_info constant

// basic type descriptor : B->byte、S->short、C->char、I->int、J->long、F->float、D->double
// ref type descriptor : L + class full name + ';'
// array type descriptor : '[' + array element type descriptor
// field descriptor : field type descriptor
// method descriptor : (argument type descriptor;argument type descriptor...) + return type argument type descriptor
// why CONSTANT_NameAndType_info need contain name and descriptor ? to distinguish different override methods

package classfile
type ConstantNameAndTypeInfo struct{
	nameIndex uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader){
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}