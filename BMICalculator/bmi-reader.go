package main

import (
	"bufio"
	"os"
)

/*
	we cannot declare it as const  it will give error because

in go lang const value is not changed but here we are calling
functions and these functions will return values on runtime
so this value will be changed afterwards , so const means in
go lang that the value defined before execution should not
change.
*/
var Reader = bufio.NewReader(os.Stdin)
