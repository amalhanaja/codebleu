package cli

import (
	"log"
	"os"
)

func Run() {
	if err := NewCliApp().Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
