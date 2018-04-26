package cmd

import (
	"github.com/hlts2/gson"
)

// REPL is REPL interface
type REPL interface {
	Run(args []string) error
}

// repl repsents REPL base structor
type repl struct {
	Gson *gson.Gson
}

// NewREPL returns repl instance
func NewREPL(gson *gson.Gson) REPL {
	return &repl{
		Gson: gson,
	}
}

// Run execute REPL of json-viewer
func (r *repl) Run(args []string) error {
	return nil
}
