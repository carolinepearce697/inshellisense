// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["copypath"] = model.Subcommand{
		Name:        []string{"copypath"},
		Description: `Oh-My-Zsh plugin that copies the path of given directory or file to the clipboard`,
		Args: []model.Arg{{
			Templates:  []model.Template{model.TemplateFilepaths, model.TemplateFolders},
			Name:       "path",
			IsOptional: true,
		}},
	}
}