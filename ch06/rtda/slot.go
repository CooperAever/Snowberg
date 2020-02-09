package rtda
import "jvmgo/ch06/rtda/heap"
// LocalVars can be accessed by index,so its best fit structure is array
// every element of array should be able to store a int or a reference,two element can store a long or double.
// should not use int[] as array,because when it store a element of uintptr,
// it will be collected by GC when no point pointer to it except array.

type Slot struct{
	num int32 		//num store integer
	ref *heap.Object		// ref store reference
}

