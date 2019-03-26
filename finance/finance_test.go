package finance

import (
	"fmt"
	"testing"
)

func TestFinace(t *testing.T) {
	//ctx := context.Background()

	rate := Rate(48, -407423, 14773500, 0, 0, 0.01)

	yrate := rate * 12
	fmt.Println(rate)
	fmt.Println(yrate)

}
