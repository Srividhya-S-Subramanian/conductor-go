package workflow_e2e

import (
	"testing"

	"github.com/conductor-sdk/conductor-go/examples/shipment_example"
	"github.com/conductor-sdk/conductor-go/pkg/workflow/def/workflow"
	"github.com/conductor-sdk/conductor-go/tests/e2e/e2e_properties"
)

func TestTaskRegistration(t *testing.T) {
	err := e2e_properties.ValidateTaskRegistration(
		*shipment_example.TaskCalculateTaxAndTotal.ToTaskDef(),
		*shipment_example.TaskChargePayment.ToTaskDef(),
		*shipment_example.TaskGroundShippingLabel.ToTaskDef(),
		*shipment_example.SameDayShippingLabel.ToTaskDef(),
		*shipment_example.AirShippingLabel.ToTaskDef(),
		*shipment_example.UnsupportedShippingLabel.ToTaskDef(),
		*shipment_example.TaskShippingLabel.ToTaskDef(),
		*shipment_example.TaskSendEmail.ToTaskDef(),
		*shipment_example.TaskGetOrderDetails.ToTaskDef(),
		*shipment_example.TaskGetUserDetails.ToTaskDef(),
		*shipment_example.TaskGetInParallel.ToTaskDef(),
		*shipment_example.TaskGenerateDynamicFork.ToTaskDef(),
		*shipment_example.TaskProcessOrder.ToTaskDef(),
		*shipment_example.TaskUpdateState.ToTaskDef(),
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWorkflowRegistration(t *testing.T) {
	workflows := []*workflow.ConductorWorkflow{
		shipment_example.NewOrderWorkflow(e2e_properties.WorkflowExecutor),
		shipment_example.NewShipmentWorkflow(e2e_properties.WorkflowExecutor),
	}
	for _, workflow := range workflows {
		err := e2e_properties.ValidateWorkflowRegistration(workflow)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestOrderWorkflow(t *testing.T) {

}

func TestShipmentWorkflow(t *testing.T) {

}