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
	if len(dataset) == 0 {
		return
	}
	for _, d := range dataset {
		err := dp.Parse(d)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	result, err := dp.ActionInfo()
	if err != nil {
		log.Println(err)
		return
	}
	
	fmt.Println(result)
}
