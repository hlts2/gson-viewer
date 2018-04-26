package cmd

import (
	"fmt"
	"os"

	prompt "github.com/c-bata/go-prompt"
	"github.com/hlts2/gson"
)

// REPL is REPL interface
type REPL interface {
	Run()
}

// repl repsents REPL base structor
type repl struct {
	Gson   *gson.Gson
	Prompt *prompt.Prompt
}

// NewREPL returns repl instance
func NewREPL(gson *gson.Gson) REPL {
	repl := &repl{
		Gson: gson,
	}

	repl.Prompt = prompt.New(
		repl.executer,
		func(in prompt.Document) []prompt.Suggest {
			s := []prompt.Suggest{}
			return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
		},
		prompt.OptionPrefix(">>> "),
	)

	return repl
}

func (r *repl) executer(in string) {
	result, err := r.Gson.GetByPath(in)
	if err != nil {
		if in == "show" {
			str, _ := r.Gson.Indent("", "  ")
			fmt.Println(str)
			return
		}

		fmt.Fprintf(os.Stderr, "json value dose not exist: %s\n", in)
		return
	}

	fmt.Println(result.Indent("", "  "))
}

var welcomText = `
   _                            _
  (_)___  ___  _ __      __   _(_) _____      _____ _ __
  | / __|/ _ \| '_ \ ____\ \ / / |/ _ \ \ /\ / / _ \ '__|
  | \__ \ (_) | | | |_____\ V /| |  __/\ V  V /  __/ |
 _/ |___/\___/|_| |_|      \_/ |_|\___| \_/\_/ \___|_|
|__/

interactive command-line JSON viewer
`

// Run execute REPL
func (r *repl) Run() {
	fmt.Println(welcomText)
	r.Prompt.Run()
}
