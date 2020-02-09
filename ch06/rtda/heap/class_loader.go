// JVM class loader is very complicate
package heap
import "fmt"
import "jvmgo/ch06/classfile"
import "jvmgo/ch06/classpath"

// classloader use classpath find and read class document
// cp store classpath path
// classMap store classInfo which have been loaded
// and its key is full qualified name
type ClassLoader struct{
	cp *classpath.Classpath 
	classMap map[string]*Class //load class
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader{
	return &ClassLoader{
		cp:cp,
		classMap: make(map[string]*Class),
	}
}

// load class info into method area
func(self *ClassLoader) LoadClass(name string) *Class{
	if class,ok := self.classMap[name];ok{
		return class //class has been loaded
	}
	return self.loadNonArrayClass(name) //load class
}

// there is big difference between arrayClass and NonArrayClass
// arrayClass data not from class document ,but created by JVM in running-time
// not implement loadNonArrayClass,we will discuss arrayClass load in ch08
func(self *ClassLoader) loadNonArrayClass(name string) *Class{
	// find class document and read into memory
	data,entry := self.readClass(name)
	// parse class document and  create class info which can be used by JVM
	// and put into method area
	class := self.defineClass(data)
	// link
	link(class)
	fmt.Printf("[Loaded %s from %s ]\n",name,entry)
	return class
}

func(self *ClassLoader) readClass(name string) ([]byte,classpath.Entry){
	data,entry,err := self.cp.ReadClass(name)
	if err != nil{
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data,entry 	// class data nad class filepath
}

func(self *ClassLoader) defineClass(data[]byte) *Class{
	class := parseClass(data) 	// class document ===> class struct
	class.loader = self
	resolveSuperCLass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class{
	cf,err := classfile.Parse(data)
	if err != nil{
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

func resolveSuperCLass(class *Class){
	if class.name != "java/lang/Object"{
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func resolveInterfaces(class *Class){
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0{
		class.interfaces = make([] *Class,interfaceCount)
		for i,interfaceName := range class.interfaceNames{
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class){
	verify(class)
	prepare(class)
}


// JVM standard will strictly verify class info before execute
// and textbook didn't give implement detail

func verify(class *Class){
	// todo
}

// prepare give class variable a space and default value
func prepare(class *Class){
	//to do 
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

// calculate instance field and mark
func calcInstanceFieldSlotIds(class *Class){
	slotId := uint(0)
	if class.superClass != nil{
		slotId = class.superClass.instanceSlotCount
	}
	for _,field := range class.fields{
		if !field.IsStatic(){
			field.slotId = slotId
			slotId++
			if field.IsLongOrDouble(){
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

// calculate static field and mark
func calcStaticFieldSlotIds(class *Class){
	slotId := uint(0)
	for _,field := range class.fields{
		if field.IsStatic(){
			field.slotId = slotId
			slotId++
			if field.IsLongOrDouble(){
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

// give class variable space and initialize them
// If static field is Final,this value sotre in class document
func allocAndInitStaticVars(class *Class){
	class.staticVars = newSlots(class.staticSlotCount)
	for _,field := range class.fields{
		if field.IsStatic() && field.IsFinal(){
			initStaticFinalVar(class,field)
		}
	}
	// golang can automatically initialize
}

func initStaticFinalVar(class *Class,field *Field){
	vars := class.staticVars
	cp := class.constantPool 
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()
	if cpIndex > 0{
		switch field.Descriptor(){
		case "Z","B","C","S","I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId,val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId,val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId,val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId,val)
		case "Ljava/lang/String": 		//string 
			panic("todo")
		}
	}
}
