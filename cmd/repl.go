package cmd

import (
	"fmt"
	"os"
	"unsafe"

	prompt "github.com/c-bata/go-prompt"
	"github.com/hlts2/gson"
	"github.com/hlts2/gson-viewer/pkg/gson-viewer"
	"github.com/hokaccha/go-prettyjson"
	"github.com/mattn/go-colorable"
)

// REPL is REPL interface
type REPL interface {
	Run() error
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
	normalizedIn := gsonviewer.NormalizeInputText(&in)

	result, err := r.Gson.GetByPath(normalizedIn)
	if err != nil {
		if in == "show" {
			d, _ := prettyjson.Marshal(r.Gson.Interface())
			fmt.Fprintln(colorable.NewColorableStdout(), *(*string)(unsafe.Pointer(&d)))
			return
		}

		fmt.Fprintf(os.Stderr, "json value dose not exist: %s\n", in)
		return
	}

	d, _ := prettyjson.Marshal(result.Interface())
	fmt.Fprintln(colorable.NewColorableStdout(), *(*string)(unsafe.Pointer(&d)))
}

var welcomText = `
                                      _
   ____ __________  ____       _   __(_)__ _      _____  _____
  / __ // ___/ __ \/ __ \_____| | / / / _ \ | /| / / _ \/ ___/
 / /_/ (__  ) /_/ / / / /_____/ |/ / /  __/ |/ |/ /  __/ /
 \__, /____/\____/_/ /_/      |___/_/\___/|__/|__/\___/_/
/____/

interactive command-line JSON viewer
`

// Run execute REPL
func (r *repl) Run() error {
	fmt.Println(welcomText)
	r.Prompt.Run()

	return nil
}
