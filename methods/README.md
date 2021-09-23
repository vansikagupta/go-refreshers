While technically Go isn’t an Object Oriented Programming language, types and methods allow for an object-oriented style of programming.

A method is a function that has a defined receiver, in OOP terms, a method is a function on an instance of an object.

Go does not have classes. However, we can define methods on struct types. In fact on any types expect for the basic types.

Methods are associated with named types either through value receivers or pointers receivers. Pointer receivers are preferred more as they avoid copying the value on each method call (more efficient if the value type is a large struct).
The other reason why you might want to use a pointer is so that the method can modify the value that its receiver points to.

The adjacent .go file has examples that explain the above points