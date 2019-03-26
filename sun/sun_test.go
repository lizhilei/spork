package sun

import (
	"fmt"
	"testing"
)

func TestFinace(t *testing.T) {
	//ctx := context.Background()

	r := SunRise(39.9818378462,116.3069070835)

	s := SunSet(39.9818378462,116.3069070835)
	fmt.Println(r,s)

}
