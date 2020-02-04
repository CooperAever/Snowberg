package classfile
import "fmt"

// ClassFile struct define the information contained by JVM
// in go lang, all type、struct、field、variable、function and method with big initial is public, can be access by other package
// on the contrary,small initial can only be access inside the package
// with the intention of increasing code readibility, all struct is public, means with big initial.
type ClassFile struct{
	//magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags uint16 		// a 16bits bitmask, indicate class document has class or interface defination, is public or private, etc. 
	thisClass uint16   		// give class name
	superClass uint16 		// give superclass name index in constantPool , every class has superclass except java.lang.Object(index==0).
	interfaces []uint16  	// give interface index in constantPool
	fields [] *MemberInfo	
	methods [] *MemberInfo
	attributes [] AttributeInfo
}

// Parse []byte data to ClassFile struct
func Parse(classData []byte)(cf *ClassFile,err error){
	defer func(){
		if r:=recover() ; r!= nil{
			var ok bool
			err,ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v",r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return 
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader,self.constantPool)
	self.methods = readMembers(reader,self.constantPool)
	self.attributes = readAttributes(reader,self.constantPool)
}


// magic number : a lot of documents declare their format with a fixed begin, like PDF document begin with "%PDF" (4bytes)
// and ZIP begin with "PK" (2bytes). The class document correspond to "0XCAFEBABE" magic word
func (self *ClassFile) readAndCheckMagic(reader *ClassReader){
	magic := reader.readUint32()
	if magic != 0XCAFEBABE{
		panic("java.lang.ClassFormatError:magic!")
	}
}

// after magic number is class document's minorVersion number and majorVersion number
// if a class has majorVersion M and minorVersion m, the full version is M.m
// after J2SE 1.2 , minorVersion is useless , and set to 0 defaultly
// when we use java 8, it can support class document which Version number is between 45.0~52.0(only 45 has minor version)
func (self *ClassFile) readAndCheckVersion(reader *ClassReader){
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion{
	case 45:
		return
	case 46,47,48,49,50,51,52:
		if self.minorVersion == 0{
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}



// getter
func (self *ClassFile) MinorVersion() uint16{
	return self.minorVersion
}

// getter
func (self *ClassFile) MajorVersion() uint16{
	return self.majorVersion
}

// getter
func (self *ClassFile) ConstantPool() ConstantPool{
	return self.constantPool
}

// getter
func (self *ClassFile) AccessFlags() uint16{
	return self.accessFlags
}

// getter
func (self *ClassFile) Fields() [] *MemberInfo{
	return self.fields
}

// getter
func (self *ClassFile) Methods() [] *MemberInfo{
	return self.methods
}

// search class name from constant pool
func (self *ClassFile) ClassName() string{
	return self.constantPool.getClassName(self.thisClass)
}

// search super class name from constant pool
func (self *ClassFile) SuperClassName() string{
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" 	//only java.lang.Object hasn't super class
}

func (self *ClassFile) InterfaceNames() []string{
	interfaceNames := make([]string,len(self.interfaces))
	for i,cpIndex := range self.interfaces{
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}






