package website

import (
	"github.com/ProductionPanic/go-input"
	"plezk/lib/common"
	"plezk/lib/domain"
)

func Delete(dom domain.Domain) int {
	for {
		common.Cls()
		answer, err := input.GetText("Are you sure you want to delete this domain? (y/n)")
		if err != nil {
			panic(err)
		}
		if answer == "y" {
			dom.Delete()
			return -1
		}
		if answer == "n" {
			break
		}
	}
	return 0
}
