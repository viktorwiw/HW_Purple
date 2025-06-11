package bins

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
