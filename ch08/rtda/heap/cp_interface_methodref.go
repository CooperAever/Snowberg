package heap

import "jvmgo/ch08/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if self.method == nil {
		self.resolveInterfaceMethodRef()
	}
	return self.method
}

// jvms8 5.4.3.4
func (self *InterfaceMethodRef) resolveInterfaceMethodRef() {
	//class := self.ResolveClass()
	// todo
	// class d try to access class c's method through interfaceMethodRe
	d := self.cp.class
	c := self.ResolvedClass()
	if !c.IsInterface(){
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupInterfaceMethod(c,self.name,self.descriptor)
	if method == nil{
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d){
		panic("java.lang.IllegalAccessError")
	}
	self.method = method
}


// if can find method in interface,then return method
// otherwise invoke lookupMethodInInterfaces() which is defined in method_lookup.go
// to find in interface's interface
func lookupInterfaceMethod(iface *Class, name, descriptor string) *Method {
	for _, method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}

	return lookupMethodInInterfaces(iface.interfaces, name, descriptor)
}

