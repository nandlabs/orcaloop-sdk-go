package service

import (
	"fmt"
	"net/http"
	"strings"

	"oss.nandlabs.io/golly/ioutils"
	"oss.nandlabs.io/golly/rest"
	"oss.nandlabs.io/orcaloop-sdk/data"
	"oss.nandlabs.io/orcaloop-sdk/events"
	"oss.nandlabs.io/orcaloop-sdk/handlers"
	"oss.nandlabs.io/orcaloop-sdk/models"
)

const (
	ActionsEndPoint  = "/api/actions"
	InstanceEndpoint = "/api/instances/:instanceId/actions/:actionId"
)

type OrcaloopClient struct {
	client  *rest.Client
	baseurl string
}

func (oc *OrcaloopClient) Register(actionHandler handlers.ActionHandler) (err error) {
	var res *rest.Response

	req := oc.client.NewRequest(oc.baseurl+ActionsEndPoint, http.MethodPost)
	req.SetContentType(ioutils.MimeApplicationJSON)
	req.SetBody(actionHandler.Spec())
	res, err = oc.client.Execute(req)
	if err != nil {
		return
	}
	switch res.StatusCode() {
	case http.StatusCreated:
		err = nil
	case http.StatusInternalServerError:
		var errorResponse *models.Error
		err = res.Decode(&errorResponse)
		if err == nil {
			err = fmt.Errorf("registration of action failed with error code %s and message %s", errorResponse.Code, errorResponse.Message)
		}
	default:
		err = res.GetError()
	}

	if err == nil {
		handlers.ActionRegistry.Register(actionHandler.Spec().Id, actionHandler)
	}
	return
}

func (oc *OrcaloopClient) Respond(actionSpec models.ActionSpec, pipeline *data.Pipeline) (err error) {
	var res *rest.Response
	var stepId, instanceId, workflowId string
	var status models.Status
	instanceId = pipeline.Id()
	endpoint := oc.baseurl + InstanceEndpoint
	endpoint = strings.ReplaceAll(endpoint, ":instanceId", instanceId)
	endpoint = strings.ReplaceAll(endpoint, ":actionId", actionSpec.Id)
	req := oc.client.NewRequest(endpoint, http.MethodPost)
	req.SetContentType(ioutils.MimeApplicationJSON)
	outputData := data.NewPipeline(instanceId)
	// Set the error response
	if pipeline.Has(data.StepIdKey) {
		stepId = pipeline.GetStepId()
		outputData.Set(data.StepIdKey, stepId)
	} else {
		err = fmt.Errorf("missing step id in pipeline")
		return
	}

	if pipeline.Has(data.WorkflowIdKey) {
		workflowId = pipeline.GetWorkflowId()
		outputData.Set(data.WorkflowIdKey, workflowId)
	} else {
		err = fmt.Errorf("missing workflow id in pipeline")
		return
	}

	if pipeline.Has(data.ErrorKey) {
		// Set the error response
		outputData.Set(data.ErrorKey, pipeline.GetError())
		status = models.StatusFailed

	} else {

		// Create a new map to hold the output
		for _, returnField := range actionSpec.Returns {
			var val any
			val, _ = pipeline.Get(returnField.Name)
			outputData.Set(returnField.Name, val)
		}
		status = models.StatusCompleted
	}
	event := &events.StepChangeEvent{
		InstanceId: instanceId,
		StepId:     stepId,
		Status:     status,
		Data:       outputData.Map(),
	}

	req.SetBody(event)
	res, err = oc.client.Execute(req)
	if err != nil {
		return
	}
	switch res.StatusCode() {
	case http.StatusOK:
		err = nil
	case http.StatusInternalServerError:
		var errorResponse *models.Error
		err = res.Decode(&errorResponse)
		if err == nil {
			err = fmt.Errorf("execution of action failed with error code %s and message %s", errorResponse.Code, errorResponse.Message)
		} else {
			err = fmt.Errorf("execution of action failed with error code %d and message %s", res.StatusCode(), res.Status())
		}
	default:
		err = res.GetError()
	}

	return
}
