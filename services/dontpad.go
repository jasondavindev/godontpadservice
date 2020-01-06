package services

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jasondavindev/godontpadservice/utils"
)

type Commands []string

func GetCommands() Commands {
	doc, err := goquery.NewDocument("http://dontpad.com/executable_golang")

	if err != nil {
		log.Fatal(err)
	}

	cmds := []string{}

	doc.Find("#text").Each(func(_ int, e *goquery.Selection) {
		for _, str := range strings.Split(e.Text(), "\n") {
			cmds = append(cmds, str)
		}
	})

	return cmds
}

func (c Commands) Exec() {
	for _, cmd := range c {
		fmt.Println(time.Now().Format("2006-01-02T15:04:05-0700"), cmd)
		utils.GetPlataformCommand(cmd).Run()
	}
}
