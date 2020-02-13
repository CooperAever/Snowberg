// look method from invoking perspective,methods can be divided into 2 types
// static method and instance method. static method invoked by class, and instance method invoked by object ref.
// static method is bind while compiling, and instance method need wait until running-time to know which method to invoke

// look method from implement perspective,methods can be divided into 3 types
// 1.no implement(abstract method)
// 2.java method (use java language implement methods)
// 3.native method (implement by native language like C、C++)


// ch08 only implements java methods,
// invokestatic instruction used to invoke static methods
// invokespecial instruction used to invoke non-dynamic-bind instance method(construct method、private method、super. invoke method)
// invokeinterface instruction uesd to invoke interface type method,otherwise invokevirtual

// method invoke need n+1 operand , the first one operand is a uint16 index,
// used for find methodRef in constantpool(may not the final invoked method), the left n operands are method argument
