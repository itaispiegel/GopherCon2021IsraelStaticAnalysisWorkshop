package main

import (
	"os"
	"strconv"
)

var a int

func main() {
	a, _ = strconv.Atoi(os.Args[1])
	if a > 0 {
		a = 3
	} else if a == 0 {
		a = 4
	} else {
		a = 5
	}
}

// 1. In the source code, the condition can go to the if, the else if or the else blocks. Here, each block can have up to 2
// successors. Notice that the first block can only jump to 1 (a=3) or 3 (a==0) and 3 evaluates the condition again and can go to
// 4 (a=4) or 5(a=5)

// 2. `a` is declared in the global scope (and not under the function scope). Because it's in the global scope, it might
// be used in other places, so the compiler isn't sure if it could have side effects, (not considered dead) so it can't
// optimize away it using dead code elimination.
