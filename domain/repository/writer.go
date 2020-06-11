package repository

import (
	"github.com/keikohi/gorchitect/domain"
)

type Writer interface {
	Write(rootDir domain.Node, dr *domain.DependecyRelation)
}
