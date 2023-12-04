package website

import (
	"fmt"
	"plezk/lib/common"
	"plezk/lib/domain"
)

func View(dom domain.Domain) {
	common.Cls()
	fmt.Println(dom.Info().GetInfoString())

}
