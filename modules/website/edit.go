package website

import (
	"fmt"
	"plezk/lib/domain"
)

func Edit(dom domain.Domain) {
	fmt.Println(dom.Info())
}
