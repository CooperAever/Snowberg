package classfile

// sourceFile used for indicating source file name
// attribute_lenght must be 2 size 
// sourcefile_index is a pointer point to constantPool

type SourceFileAttribute struct{
	cp ConstantPool
	sourceFileIndex uint16
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader){
	self.sourceFileIndex = reader.readUint16()
}

func (self *SourceFileAttribute) FileName() string{
	return self.cp.getUtf8(self.sourceFileIndex)
}
