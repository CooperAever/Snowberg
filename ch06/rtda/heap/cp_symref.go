package heap
//symbolic reference
type SymRef struct{
	cp *ConstantPool //store running time constant pool pointer
	className string 
	class *Class //cache class struct pointer,so only need to parse once
}

// resolve class reference
func (self *SymRef) ResolvedClass() *Class{
	if self.class == nil{
		self.resolveClassRef()
	}
	return self.class
}

// basiclly,if class D try use symbolic reference N ref class C,
// want to resolve N,first use D.loader to load C,and check if D qualified to access C
// if not ,throw Illegal AccessError
func (self *SymRef) resolveClassRef() {
	d := self.cp.class 
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d){ 	//defined in class.go
		panic("java.lang.IllegalAccessError")
	}
	self.class = c 
}



