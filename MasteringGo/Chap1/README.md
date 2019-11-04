## Error handling in Go

Errors and error handling are two very important Go topics. Go likes error
messages so much that it has a separate data type for errors, named error!
this also means that you can easily create your own error mesages if you find
that what Go gives you is inadequate. You will most likely need to create and
handle your own errors when you are developing your own go packages.

Note that having an error condition is one thing, while deciding how to react
to an error condition is a totally different thing. Put simply, not all error
conditions are created equal, which means that some error conditions might
require that you immediately stop the execution of a programs, whereas other
error situations might require printing a warning message for the user and
continuing with the execution of the program. It is up to the developer to use
common sense and decide what to do with each `error` value that the program
might get.

### The `defer` keyword

The `defer` keyword postpones the execution of a function untill the
surrounding function returns. It is widely used in file inpu9t and output
operations because it saves you from having to remember when to close an opened
file: the `defer` keyword allows you to to put the function call that closes
an opened file near to the function call that opened it.

It is very important to remember that __deferred functions__ are executed in
`Last In First Out` order after the return of the surrounding function. Put
simply, this means that if you `defer` function `f1()` first, function `f2()`
second (conceptually, it means `defer f1...; defer f2 ...` then when the
surrounding function is about to return, function `f2() -> f1()` executes in that order. 


