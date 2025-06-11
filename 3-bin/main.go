package make

import (
	"fmt"
)

func main() {
	id := "1"
	private := true
	name := ""
	myBin := newBin(id, private, name)
	fmt.Println(myBin)
}
