package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, d := range dataset {
		err := dp.Parse(d)
		if err != nil {
			log.Println(err)
			continue
		}
		result, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(result)
	}
	
}
