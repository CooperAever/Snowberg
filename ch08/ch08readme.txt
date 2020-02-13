// array is a special defination 
// 1.normal class is load from class document, array class is created by JVM in running time.
// 2.1-D array class is construct by newarray instruction, n-D by multianewarray
// 3.normal object store instance variable, use putfield and getfield access, array object store element, use index and <t>aload、 <t>astore instructions to access, t can be a(ref)、b(byte)、c(char)、d(double)、f(float)、i(int)、l(long)、s(short). additionally, use arraylength get array length.