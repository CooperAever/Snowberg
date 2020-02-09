package heap
import "jvmgo/ch06/classfile"

type ClassMember struct{
	accessFlags uint16
	name string 
	descriptor string
	class *Class //store class struct pointer
}

func(self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo){
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}
func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}
func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}
func (self *ClassMember) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *ClassMember) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}

// getters
func (self *ClassMember) Name() string {
	return self.name
}
func (self *ClassMember) Descriptor() string {
	return self.descriptor
}
func (self *ClassMember) Class() *Class {
	return self.class
}


func(self *ClassMember) isAccessiableTo(d *Class) bool{
	// if field is public, any class can access
	if self.IsPublic(){
		return true
	}
	c := self.class 
	// if field is protected, only child class or under the same package can access
	if self.IsProtected() {
		return d == c || d.isSubClassOf(c) || c.getPackageName() == d.getPackageName()
	}
	// if field is not public、protected、privated, only under the same package can access

	if !self.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	// if field is private,only this field.Class can access
	return d == c 
}
