package main

import (
	"os"

	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func run() {
	items := [...]string{"C", "C++", "C#", "Python", "Ruby", "Go", "Javascript", "Typescript"}

	for _, item := range items {
		wf.NewItem(item).
			Var("itemName", item).
			Valid(true)
	}

	args := os.Args
	if len(args) > 1 {
		wf.Filter(args[1])
	}
	wf.WarnEmpty("No programming language.", "Try different programming language.")

	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
