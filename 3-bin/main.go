package make

import (
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.DateTime
	name      string
}

func newBin(id string, private bool, name string) *Bin {
	return &Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}

func (bin Bin) BinList() {
	binList := []Bin{}
	for _, value := range binList {
		fmt.Println(value)
	}
}

func main() {
	id := "1"
	private := true
	name := ""
	myBin := newBin(id, private, name)
	fmt.Println(myBin)
}
