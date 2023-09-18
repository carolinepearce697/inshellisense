// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["wc"] = model.Subcommand{
		Name:        []string{"wc"},
		Description: `World, line, character, and byte count`,
		Args: []model.Arg{{
			Templates:   []model.Template{model.TemplateFilepaths},
			Name:        "file",
			Description: `File to count in`,
			IsOptional:  true,
			IsVariadic:  true,
		}},
		Options: []model.Option{{
			Name:        []string{"-c"},
			Description: `Output the number of bytes to the standard input`,
		}, {
			Name:        []string{"-l"},
			Description: `Output the number of lines to the standard input`,
		}, {
			Name:        []string{"-m"},
			Description: `Output the number of characters to the standard input`,
		}, {
			Name:        []string{"-w"},
			Description: `Output the number of words to the standard input`,
		}},
	}
}