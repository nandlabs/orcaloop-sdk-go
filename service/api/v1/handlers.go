package v1

import (
	"net/http"

	"oss.nandlabs.io/golly/rest"
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

func ExecuteAction(ctx rest.ServerContext) {
	var actionHandler handlers.ActionHandler
	actionId, error := ctx.GetParam("actionId", rest.PathParam)
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
	stepId, ok := input[data.StepIdKey].(string)
	if !ok {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.WriteJSON(transformError(http.StatusBadRequest, "Step Id is required"))
		return
	}

	pipeline := data.NewPipelineFrom(instanceId, input)
	err = actionHandler.Handle(pipeline)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.WriteJSON(transformError(http.StatusInternalServerError, err.Error()))
		return
	} else {
		responseData := make(map[string]any)
		for _, returnParam := range actionHandler.Spec().Returns {
			returnval, err := pipeline.Get(returnParam.Name)
			if err != nil {
				ctx.SetStatusCode(http.StatusInternalServerError)
				ctx.WriteJSON(transformError(http.StatusInternalServerError, err.Error()))
				return
			}
			responseData[returnParam.Name] = returnval
		}

		ctx.SetStatusCode(http.StatusOK)
		response := &models.StepChangeEvent{
			InstanceId: instanceId,
			StepId:     stepId,
			Status:     models.StatusCompleted,
			Data:       data.NewPipelineFrom(instanceId, responseData),
		}
		ctx.WriteJSON(response)

		return
	}
}
