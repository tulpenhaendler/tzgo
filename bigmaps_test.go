package tzgo

import (
	"context"
	"fmt"
	"github.com/tulpenhaendler/tzgo/rpc"
	"github.com/tulpenhaendler/tzgo/tezos"
	"testing"
)

func TestAdd(t *testing.T) {
	addr := tezos.MustParseAddress("KT1J1yVJEDkaYzJ7WuEomryEGWBNMGcUw1WA") // Mooncakes
	c, e := rpc.NewClient("https://archive_vin14.mantodev.com/", nil)
	if e != nil {
		t.Error(e)
	}
	script, e := c.GetContractScript(context.Background(), addr)
	if e != nil {
		t.Error(e)
	}
	bs := script.Bigmaps()
	fmt.Println(bs)
	fmt.Println("HI")
}
