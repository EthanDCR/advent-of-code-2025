package main

import (
	"fmt"
)

// this is how short code works type shit

func sayHello(name string) (message string) {
	return "hello there, " + name + "!"
}

///// bunch random stuff loremipsm nutsack ...

///// bunch random stuff loremipsm nutsack ...

///// bunch random stuff loremipsm nutsack ...

///// bunch random stuff loremipsm nutsack ...

///// bunch random stuff loremipsm nutsack ...

///// bunch random stuff loremipsm nutsack ...

///// bunch random stuff loremipsm nutsack ...
///// bunch random stuff loremipsm nutsack ...

///// bunch random stuff loremipsm nutsack ...

func main() {
	// we use the code thats written above right here, without needing to write it all into the main file
	// this calls the sayHello function which is defined above (could be in another file)
	fmt.Println(sayHello("owen"))
}
