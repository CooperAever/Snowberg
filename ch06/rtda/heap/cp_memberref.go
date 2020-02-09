//define MemberRef store the info field symbol and method symbol both have
package heap
import "jvmgo/ch06/classfile"

type MemberRef struct{
	SymRef
	name string 
	descriptor string
}

func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberInfo){
	self.className = refInfo.ClassName()
	self.name,self.descriptor = refInfo.NameAndDescriptor()
}