package interfaces

import (
	"agourdel.com/kgpkernel/domains/foodomain/models"
)

type IFooRepositories interface {
	GetVirgins() []models.VirginModel
}
