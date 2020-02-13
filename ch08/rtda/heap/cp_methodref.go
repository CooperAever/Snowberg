package heap

import "jvmgo/ch08/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

// jvms8 5.4.3.3
func (self *MethodRef) resolveMethodRef() {
	//class := self.Class()
	// class d try to access class c's method through methodRef
	d := self.cp.class
	// parse methodRef find class c
	c := self.ResolvedClass()
	// if c is an interface, throw error
	if c.IsInterface(){
		panic("java.lang.IncompatibleClassChangeError")
	}
	// find method through methodName and methodDescriptor
	method := lookupMethod(c,self.name,self.descriptor)
	if method == nil{
		panic("java.lang.NoSuchMethodError")
	}
	// defined in classMember struct
	if !method.isAccessibleTo(d){
		panic("java.lang.IllegalAccessError")
	}
	self.method = method 
}

func lookupMethod(class *Class,name,descriptor string) *Method{
	// first find in extends, defined in heap\method_lookup.go
	method := LookupMethodInClass(class,name,descriptor)
	if method == nil{
		// then find in interface
		method = lookupMethodInInterfaces(class.interfaces,name,descriptor)
	}
	return method
}
