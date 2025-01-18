package v1

import (
	"net/http"

	"oss.nandlabs.io/golly/rest/server"
	"oss.nandlabs.io/orcaloop-sdk/data"
	"oss.nandlabs.io/orcaloop-sdk/handlers"
	"oss.nandlabs.io/orcaloop-sdk/models"
)

const (
	ActionIDParam = "actionId"
)

var transformError = func(code int, message string) *models.Error {
	return &models.Error{
		Code:    http.StatusText(code),
		Message: message,
	}
}

func ExecuteAction(ctx server.Context) {
	var actionHandler handlers.ActionHandler
	actionId, error := ctx.GetParam("actionId", server.PathParam)
	if error != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.WriteJSON(transformError(http.StatusBadRequest, error.Error()))
		return
	}

	actionHandler = handlers.ActionRegistry.Get(actionId)
	if actionHandler == nil {
		ctx.SetStatusCode(http.StatusNotFound)
		ctx.WriteJSON(transformError(http.StatusNotFound, "Action not found"))
		return
	}
	input := make(map[string]any)
	err := ctx.Read(&input)
	if err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.WriteJSON(transformError(http.StatusBadRequest, err.Error()))
		return
	}
	instanceId, ok := input[data.InstanceIdKey].(string)
	if !ok {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.WriteJSON(transformError(http.StatusBadRequest, "Instance Id is required"))
		return
	}

	pipeline := data.NewPipelineFrom(instanceId, input)
	err = actionHandler.Handle(pipeline)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.WriteJSON(transformError(http.StatusInternalServerError, err.Error()))
		return
	} else {
		ctx.SetStatusCode(http.StatusAccepted)
		return
	}
}
