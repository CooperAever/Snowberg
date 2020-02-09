// constantpool store 2 type infomation
// 1.literal :literal including int,float and string ;
// 2.symbolic reference : symbolic reference including class symbolic reference、field symbolic reference、method symbolic reference and interface symbolic reference
package heap
import "fmt"
import "jvmgo/ch06/classfile"

type Constant interface{}
type ConstantPool struct{
	class *Class 
	consts []Constant
}

// transformer constantpool in document into constantpool in running time
// namely []classfile.ConstantInfo ===> []heap.ConstantInfo
func newConstantPool(class *Class,cfCp classfile.ConstantPool) *ConstantPool{
	cpCount := len(cfCp)
	consts := make([]Constant,cpCount)
	rtCp := &ConstantPool{class,consts}
	for i:=1;i<cpCount;i++{
		cpInfo := cfCp[i]
		// simplest is int or float,pick up constant value and put into consts
		switch cpInfo.(type){
		//literal 
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value() //int32
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value() //float32	
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value() //int64
			i++  //long occupied 2 position,need increment one more index	
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value() //float64
			i++ //double occupied 2 position,need increment one more index
		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = stringInfo.String() //string

		//symbolic reference
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtCp,classInfo) 
		case *classfile.ConstantFieldrefInfo:
			fieldrefInfo := cpInfo.(*classfile.ConstantFieldrefInfo)
			consts[i] = newFieldRef(rtCp,fieldrefInfo)
		case *classfile.ConstantMethodrefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantMethodrefInfo)
			consts[i] = newMethodRef(rtCp,methodrefInfo) 	
		case *classfile.ConstantInterfaceMethodrefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
			consts[i] = newInterfaceMethodRef(rtCp,methodrefInfo) 
		}
	}
	return rtCp
}

func (self *ConstantPool) GetConstant(index uint) Constant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}