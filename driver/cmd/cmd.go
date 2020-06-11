package cmd

import (
	"github.com/jessevdk/go-flags"
	"github.com/keikohi/gorchitect/registry"
	"github.com/keikohi/gorchitect/usecase"
)

type options struct {
	ProjectPath string `short:"p" long:"path" required:"true" description:"input project path"`
}

func getOptions() (*options, error) {
	options := new(options)
	_, err := flags.Parse(options)
	if err != nil {
		return nil, err
	}
	return options, err
}

func Run(registory registry.Writer) error {
	options, err := getOptions()
	if err != nil {
		return err
	}
	writer := registory.NewWriter()
	return usecase.Analyzing(options.ProjectPath, writer)
}
