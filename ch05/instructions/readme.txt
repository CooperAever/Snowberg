package instructions
// JVM , namly a virtual machine, which accept bytecode as a readable machine code
// every class or interface will be compile as a class document
// Bytecode contain JVM instruction,every instruction begin with a single byte operand -> opcode 
// Single byte operand means there are most 256(2^8) bytecode instructions
// now already have 205, from 0(0x00) to 202(0xCA) and from 254(0xFE) to 255(0xFF) , all these are called JVM instruction set.

// one instruction contain mnemonic + operand,like 0xB20002 , mnemonic B2 means getstatic , operand is argument, means second constant in constantPool
// mentioned in ch04,Localvars and operandStack only store value not type
// that means mnemonic should indicate type,like iadd is add int value
// mnemonic first letter means type,casting relation is down below:
// reference -> a ; byte/boolean ->b ; char ->c ; double->d ; 
// float ->f ; int -> i ; long ->l ; short -> s;

// instructions can be divided 11 types by usage:
// constant ; loads ; stores ; stack ; math ; conversions ; 
// comparisons ; control ; references ; extended ; reserved ;

// we will implement 9 types , their share code put in base dir