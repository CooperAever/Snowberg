package heap
import "jvmgo/ch06/classfile"

type FieldRef struct{
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool,refInfo *classfile.ConstantFieldrefInfo) *FieldRef{
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref 
}

func (self *FieldRef) ResolvedField() *Field{
	if self.field == nil{
		self.resolveFieldRef()
	}
	return self.field
}

// if class D try to access class C's field, need first reslove C,
// and find filed by look up field name and descriptor
func (self *FieldRef) resolveFieldRef(){
	d := self.cp.class 
	c := self.ResolvedClass()
	field := lookupField(c,self.name,self.descriptor)
	if field == nil{
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessiableTo(d){
		panic("java.lang.IllegalAccessError")
	}
	self.field = field
}

func lookupField(c *Class, name,descriptor string) *Field{
	for _,field := range c.fields{
		if field.name == name && field.descriptor == descriptor{
			return field
		}
	}

	for _,iface := range c.interfaces{
		if field := lookupField(iface,name,descriptor);field != nil{
			return field
		}
	}

	if c.superClass != nil{
		return lookupField(c.superClass,name,descriptor)
	}
	return nil
}