package usecase

import (
	"github.com/keikohi/gorchitect/domain"
	"github.com/keikohi/gorchitect/domain/repository"
)

func Analyzing(projectPath string, writer repository.Writer) error {

	filepaths, err := domain.LoadGoFiles(projectPath)
	if err != nil {
		return err
	}
	nf := domain.NewTreeFactory(filepaths)
	rootNode := nf.Build()

	analyzer := domain.NewAnalyzer(filepaths, projectPath)
	codes := analyzer.Analyse()
	dr := codes.DependecyRelation()

	writer.Write(rootNode, &dr)
	return nil
}
