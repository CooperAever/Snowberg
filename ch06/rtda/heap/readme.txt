// ch02 implement classfile,find class document and load into memory
// ch03 parse class document into a classfile struct
// before we parse thread private info, ch06 will further parse classfile struct, mainly thread-share info


// methodArea is a rtda logic area, shared by multi-thread
// methodArea store class info from class document and class variable also sotre in methodArea
// when JVM first use a class,it will search classpath and find class document,read and parse and put correspond info into methodArea

// As for where is method area,and is fixed size or variable size ,
// and if there is garbage collection,JVM standard has no clear stipulation


// which infos need to put into methodArea:
// 1.class
// 2.Field
// 3.Method
// 4.ConstantPool

