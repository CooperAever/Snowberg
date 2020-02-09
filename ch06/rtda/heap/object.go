package heap
type Object struct{
	//ch06 todo
	class *Class 
	fields Slots
}

// need to know 1.how many space do static variable and instance variable need？
// and  2. which filed correspond to which slot？

// for question 1: assume a class have m static field and n instance field
// so they need space are m' and n'. 
// And class can be inherite,so need iterative compute superclass's instance 
// And long and double filed need 2 position, so m' >= m and n'>=n

// for question 2:use number to identify , and static field and instance field need identify separately
// and for instance,need identify from top to bottom, that means from java.lang.Object
// and need consider long and double's different

func newObject(class *Class) *Object{
	return &Object{
		class:class,
		fields: newSlots(class.instanceSlotCount),
	}
}

func (self *Object) IsInstanceOf(class *Class)bool{
	return class.isAssignableFrom(self.class) // defined in class_hierarchy.go
}


