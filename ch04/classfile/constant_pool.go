// constant pool occpied a large part of class document data
// store a lot of constant information, including number and string constant、
// struct and interface name 、 fields and methods name
package classfile

// there are three things need to pay attention
// 1.table header give the constantPool size,but this size equals to real size +1
// 2.the valid index is 1~n-1, and 0 is invalid.
// 3.CONSTANT_Long_info and CONSTANT_Double_info occupy 2 size,means some index is invalid 
type ConstantPool [] ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool{
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo,cpCount)
	for i:=1;i<cpCount;i++{
		// index begin from 1
		cp[i] = readConstantInfo(reader,cp)
		switch cp[i].(type){	//.(type) only use for switch,means compare variable type not value
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ //occupy two slot
		}
	}
	return cp
}


func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo{
	if cpInfo := self[index];cpInfo != nil{
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (self ConstantPool) getNameAndType(index uint16) (string,string){
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name,_type
}

func (self ConstantPool) getClassName(index uint16) string{
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

// find UTF-8 string from constant pool
func (self ConstantPool) getUtf8(index uint16) string{
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

