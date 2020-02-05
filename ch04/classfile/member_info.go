package classfile

// fields and methods have same structure,so use the same struct construct -- MemberInfo
type MemberInfo struct{
	cp ConstantPool 	// a pointer to constantPool
	accessFlags uint16 		// access flag, a bit mask
	nameIndex uint16 		// name index in constantPool
	desriptorIndex uint16  	// desriptorIndex in constantPool 
	attributes []AttributeInfo 		// attribute table
}

//use readMembers to read fields table or methods table
func readMembers(reader *ClassReader,cp ConstantPool) [] *MemberInfo{
	memberCount := reader.readUint16()
	members := make([] *MemberInfo,memberCount)
	for i:= range members{
		members[i] = readMember(reader,cp)
	}
	return members
}

func readMember(reader *ClassReader,cp ConstantPool) *MemberInfo{
	return &MemberInfo{
		cp : cp,
		accessFlags : reader.readUint16(),
		nameIndex : reader.readUint16(),
		desriptorIndex: reader.readUint16(),
		attributes : readAttributes(reader,cp),
	}
}

// getter
func (self *MemberInfo) AccessFlags() uint16{
	return self.accessFlags
}

// getter
func (self *MemberInfo) Name() string{
	return self.cp.getUtf8(self.nameIndex)
}

// getter
func (self *MemberInfo) Descriptor() string{
	return self.cp.getUtf8(self.desriptorIndex)
}


