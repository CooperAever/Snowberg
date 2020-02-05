

public static float circumference(float r){
	float pi = 3.14f;
	float area = 2 *pi *r;
	return area;
}

// the code above will be compiled by javac to bytecode following:
// 00 ldc #4
// 02 fstore_1   pop from OperandStack and store in LocalVars
// 03 fconst_2   push into OperandStack
// 04 fload_1	 push LocalVars[1] into OperandStack top
// 05 fmul		 float multiply
// 06 fload_0
// 07 fmul
// 08 fstore_2
// 09 fload_2
// 10 return 	 pop OperandStack top

// before excute, jvm compute LocalVars size = 3,OperandStack size = 2
// assume argument r = 1.6f

// ### value in every phrase
// LocalVars: 1 , 2, 3 means LocalVars array [0] = 1,[1] = 2,[2] = 3
// OperandStack: 1,2  means bottom is 2 , and top is 1

// before excute 		{pc: ; LocalVars:1.6 ; OperandStack: }
// ldc #4 				{pc: 00 ; LocalVars:1.6 ; OperandStack:3.14 }
// fstore_1				{pc: 02 ; LocalVars:1.6,3.14 ; OperandStack: }
// fconst_2				{pc: 03 ; LocalVars:1.6,3.14 ; OperandStack:2.0}
// fload_1				{pc: 04 ; LocalVars:1.6,3.14  ; OperandStack:3.14,2.0}
// fmul					{pc: 05 ; LocalVars:1.6,3.14  ; OperandStack:6.28}
// fload_0				{pc: 06 ; LocalVars:1.6,3.14  ; OperandStack:1.6,6.28}
// fmul					{pc: 07 ; LocalVars:1.6,3.14  ; OperandStack:10.048}
// fstore_2				{pc: 08 ; LocalVars:1.6,3.14,10.048 ; OperandStack: }
// fload_2				{pc: 08 ; LocalVars:1.6,3.14,10.048 ; OperandStack: 10.048}
// return 				{pc: 08 ; LocalVars:1.6,3.14,10.048 ; OperandStack: }





