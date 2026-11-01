package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for idx, el := range dataset {
		err := dp.Parse(el)
		if err != nil {
			log.Print(err, " элемент по индексу ", idx)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Print(err, " элемент по индексу ", idx)
			continue
		}
		fmt.Println(info)
	}
}
