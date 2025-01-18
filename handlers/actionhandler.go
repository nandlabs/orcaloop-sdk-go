package handlers

import (
	"oss.nandlabs.io/golly/managers"
	"oss.nandlabs.io/orcaloop-sdk/data"
	"oss.nandlabs.io/orcaloop-sdk/models"
)

var ActionRegistry managers.ItemManager[ActionHandler] = managers.NewItemManager[ActionHandler]()

type ActionHandler interface {
	Handle(pipeline *data.Pipeline) error
	Spec() *models.ActionSpec
}
