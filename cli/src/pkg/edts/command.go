package edts

import (
	"os"
	"strings"

	"github.com/lolPants/redts/cli/src/pkg/config"
)

func RunCommand(script string, args []string) {
	cfg := config.Load()
	if !cfg.EnsureRequiredFields() {
		os.Exit(1)
	}

	var parsedArgs strings.Builder
	for i, arg := range args {
		hasSpace := strings.ContainsRune(arg, ' ')

		if hasSpace {
			parsedArgs.WriteRune('"')
		}

		parsedArgs.WriteString(arg)

		if hasSpace {
			parsedArgs.WriteRune('"')
		}

		if i+1 != len(args) {
			parsedArgs.WriteRune(' ')
		}
	}

	err := callAPI(cfg, script, parsedArgs.String())
	if err != nil {
		panic(err)
	}
}
