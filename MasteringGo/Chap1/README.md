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


