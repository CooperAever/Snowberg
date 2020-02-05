// Deprecated and Synthetic are two simplest attributes, only used fo 
// mark , not contain any data(attribute_length == 0)
package classfile
// Deprecated attribute mark class、interface、field and method which are
// not suggested use.  @deprecated
type DeprecatedAttribute struct{
	MarkerAttribute
}

// Synthetic attribute mark nest class or nest interface
type SyntheticAttribute struct{
	MarkerAttribute
}

type MarkerAttribute struct{}

// Because these two attributes do not have data , so readInfo() is empty
func (self *MarkerAttribute) readInfo(reader *ClassReader){
	// read nothing
}