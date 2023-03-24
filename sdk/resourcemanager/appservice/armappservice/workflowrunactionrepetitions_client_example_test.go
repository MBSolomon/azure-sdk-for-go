//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/10c740b7224771c889cfb92f128168f5a0568c26/specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/WorkflowRunActionRepetitions_List.json
func ExampleWorkflowRunActionRepetitionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkflowRunActionRepetitionsClient().NewListPager("testResourceGroup", "test-name", "testFlow", "08586776228332053161046300351", "testAction", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.WorkflowRunActionRepetitionDefinitionCollection = armappservice.WorkflowRunActionRepetitionDefinitionCollection{
		// 	Value: []*armappservice.WorkflowRunActionRepetitionDefinition{
		// 		{
		// 			Name: to.Ptr("000000"),
		// 			Type: to.Ptr("/workflows/runs/actions/repetitions"),
		// 			ID: to.Ptr("/workflows/testFlow/runs/08586776228332053161046300351/actions/testAction/repetitions/000000"),
		// 			Properties: &armappservice.WorkflowRunActionRepetitionProperties{
		// 				Code: to.Ptr("OK"),
		// 				Correlation: &armappservice.RunActionCorrelation{
		// 					ClientTrackingID: to.Ptr("08586775357427610445444523191"),
		// 					ActionTrackingID: to.Ptr("0d8152bb-e198-44a9-bde8-5138eea16dd4"),
		// 				},
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-18T17:05:57.2264835Z"); return t}()),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-18T17:05:57.217991Z"); return t}()),
		// 				Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
		// 				InputsLink: &armappservice.ContentLink{
		// 					ContentHash: &armappservice.ContentHash{
		// 						Algorithm: to.Ptr("md5"),
		// 						Value: to.Ptr("8q1zMKS5ZyHBrPF+qF1xXw=="),
		// 					},
		// 					ContentSize: to.Ptr[int64](8),
		// 					ContentVersion: to.Ptr("8q1zMKS5ZyHBrPF+qF1xXw=="),
		// 					URI: to.Ptr("https://test-site.azurewebsites.net:443/runtime/webhooks/workflow/scaleUnits/prod-00/workflows/2dfde2fce2584c49bea17ef0b04c95f4/runs/08586776228332053161046300351/actions/testAction/repetitions/000000/contents/ActionInputs?api-version=2016-10-01&se=2018-04-18T21%3A00%3A00.0000000Z&sp=%2Fruns%2F08586776228332053161046300351%2Factions%2FtestAction%2Frepetitions%2F000000%2Fcontents%2FActionInputs%2Fread&sv=1.0&sig=vw4BDdYp4Ap5RXdM7tY_wl9C38DeAHfnixLBEOpideA"),
		// 				},
		// 				OutputsLink: &armappservice.ContentLink{
		// 					ContentHash: &armappservice.ContentHash{
		// 						Algorithm: to.Ptr("md5"),
		// 						Value: to.Ptr("8q1zMKS5ZyHBrPF+qF1xXw=="),
		// 					},
		// 					ContentSize: to.Ptr[int64](8),
		// 					ContentVersion: to.Ptr("8q1zMKS5ZyHBrPF+qF1xXw=="),
		// 					URI: to.Ptr("https://test-site.azurewebsites.net:443/runtime/webhooks/workflow/scaleUnits/prod-00/workflows/2dfde2fce2584c49bea17ef0b04c95f4/runs/08586776228332053161046300351/actions/testAction/repetitions/000000/contents/ActionOutputs?api-version=2016-10-01&se=2018-04-18T21%3A00%3A00.0000000Z&sp=%2Fruns%2F08586776228332053161046300351%2Factions%2FtestAction%2Frepetitions%2F000000%2Fcontents%2FActionOutputs%2Fread&sv=1.0&sig=y8Wq7jbu85tmlMo_1zpRyqNJuoCaQCFQtZ3bgSovLY0"),
		// 				},
		// 				TrackingID: to.Ptr("0d8152bb-e198-44a9-bde8-5138eea16dd4"),
		// 				RepetitionIndexes: []*armappservice.RepetitionIndex{
		// 					{
		// 						ItemIndex: to.Ptr[int32](0),
		// 						ScopeName: to.Ptr("For_each"),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("000001"),
		// 			Type: to.Ptr("/workflows/runs/actions/repetitions"),
		// 			ID: to.Ptr("/workflows/testFlow/runs/08586776228332053161046300351/actions/testAction/repetitions/000001"),
		// 			Properties: &armappservice.WorkflowRunActionRepetitionProperties{
		// 				Code: to.Ptr("OK"),
		// 				Correlation: &armappservice.RunActionCorrelation{
		// 					ClientTrackingID: to.Ptr("08586775357427610445444523191"),
		// 					ActionTrackingID: to.Ptr("f84f23eb-b331-4772-9f39-cc307fa83bc3"),
		// 				},
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-18T17:05:57.1015421Z"); return t}()),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-18T17:05:57.0929911Z"); return t}()),
		// 				Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
		// 				InputsLink: &armappservice.ContentLink{
		// 					ContentHash: &armappservice.ContentHash{
		// 						Algorithm: to.Ptr("md5"),
		// 						Value: to.Ptr("OA3i83YHGYVch+N8BQJIRQ=="),
		// 					},
		// 					ContentSize: to.Ptr[int64](6),
		// 					ContentVersion: to.Ptr("OA3i83YHGYVch+N8BQJIRQ=="),
		// 					URI: to.Ptr("https://test-site.azurewebsites.net:443/runtime/webhooks/workflow/scaleUnits/prod-00/workflows/2dfde2fce2584c49bea17ef0b04c95f4/runs/08586776228332053161046300351/actions/testAction/repetitions/000001/contents/ActionInputs?api-version=2016-10-01&se=2018-04-18T21%3A00%3A00.0000000Z&sp=%2Fruns%2F08586776228332053161046300351%2Factions%2FtestAction%2Frepetitions%2F000001%2Fcontents%2FActionInputs%2Fread&sv=1.0&sig=dLmnt50joimEMK4k9rR6njHQh94iSFJ9rrDxFbkEg5M"),
		// 				},
		// 				OutputsLink: &armappservice.ContentLink{
		// 					ContentHash: &armappservice.ContentHash{
		// 						Algorithm: to.Ptr("md5"),
		// 						Value: to.Ptr("OA3i83YHGYVch+N8BQJIRQ=="),
		// 					},
		// 					ContentSize: to.Ptr[int64](6),
		// 					ContentVersion: to.Ptr("OA3i83YHGYVch+N8BQJIRQ=="),
		// 					URI: to.Ptr("https://test-site.azurewebsites.net:443/runtime/webhooks/workflow/scaleUnits/prod-00/workflows/2dfde2fce2584c49bea17ef0b04c95f4/runs/08586776228332053161046300351/actions/testAction/repetitions/000001/contents/ActionOutputs?api-version=2016-10-01&se=2018-04-18T21%3A00%3A00.0000000Z&sp=%2Fruns%2F08586776228332053161046300351%2Factions%2FtestAction%2Frepetitions%2F000001%2Fcontents%2FActionOutputs%2Fread&sv=1.0&sig=B3-X5sqIAv1Lb31GOD34ZgIRUXGuiM2QllWiNwXFYAw"),
		// 				},
		// 				TrackingID: to.Ptr("f84f23eb-b331-4772-9f39-cc307fa83bc3"),
		// 				RepetitionIndexes: []*armappservice.RepetitionIndex{
		// 					{
		// 						ItemIndex: to.Ptr[int32](1),
		// 						ScopeName: to.Ptr("For_each"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/10c740b7224771c889cfb92f128168f5a0568c26/specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/WorkflowRunActionRepetitions_Get.json
func ExampleWorkflowRunActionRepetitionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowRunActionRepetitionsClient().Get(ctx, "testResourceGroup", "test-name", "testFlow", "08586776228332053161046300351", "testAction", "000001", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowRunActionRepetitionDefinition = armappservice.WorkflowRunActionRepetitionDefinition{
	// 	Name: to.Ptr("000001"),
	// 	Type: to.Ptr("/workflows/runs/actions/repetitions"),
	// 	ID: to.Ptr("/workflows/testFlow/runs/08586776228332053161046300351/actions/testAction/repetitions/000001"),
	// 	Properties: &armappservice.WorkflowRunActionRepetitionProperties{
	// 		Code: to.Ptr("OK"),
	// 		Correlation: &armappservice.RunActionCorrelation{
	// 			ClientTrackingID: to.Ptr("08586775357427610445444523191"),
	// 			ActionTrackingID: to.Ptr("f84f23eb-b331-4772-9f39-cc307fa83bc3"),
	// 		},
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-18T17:05:57.1015421Z"); return t}()),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-18T17:05:57.0929911Z"); return t}()),
	// 		Status: to.Ptr(armappservice.WorkflowStatusSucceeded),
	// 		InputsLink: &armappservice.ContentLink{
	// 			ContentHash: &armappservice.ContentHash{
	// 				Algorithm: to.Ptr("md5"),
	// 				Value: to.Ptr("OA3i83YHGYVch+N8BQJIRQ=="),
	// 			},
	// 			ContentSize: to.Ptr[int64](6),
	// 			ContentVersion: to.Ptr("OA3i83YHGYVch+N8BQJIRQ=="),
	// 			URI: to.Ptr("https://test-site.azurewebsites.net:443/runtime/webhooks/workflow/scaleUnits/prod-00/workflows/2dfde2fce2584c49bea17ef0b04c95f4/runs/08586776228332053161046300351/actions/testAction/repetitions/000001/contents/ActionInputs?api-version=2016-10-01&se=2018-04-18T21%3A00%3A00.0000000Z&sp=%2Fruns%2F08586776228332053161046300351%2Factions%2FtestAction%2Frepetitions%2F000001%2Fcontents%2FActionInputs%2Fread&sv=1.0&sig=dLmnt50joimEMK4k9rR6njHQh94iSFJ9rrDxFbkEg5M"),
	// 		},
	// 		OutputsLink: &armappservice.ContentLink{
	// 			ContentHash: &armappservice.ContentHash{
	// 				Algorithm: to.Ptr("md5"),
	// 				Value: to.Ptr("OA3i83YHGYVch+N8BQJIRQ=="),
	// 			},
	// 			ContentSize: to.Ptr[int64](6),
	// 			ContentVersion: to.Ptr("OA3i83YHGYVch+N8BQJIRQ=="),
	// 			URI: to.Ptr("https://test-site.azurewebsites.net:443/runtime/webhooks/workflow/scaleUnits/prod-00/workflows/2dfde2fce2584c49bea17ef0b04c95f4/runs/08586776228332053161046300351/actions/testAction/repetitions/000001/contents/ActionOutputs?api-version=2016-10-01&se=2018-04-18T21%3A00%3A00.0000000Z&sp=%2Fruns%2F08586776228332053161046300351%2Factions%2FtestAction%2Frepetitions%2F000001%2Fcontents%2FActionOutputs%2Fread&sv=1.0&sig=B3-X5sqIAv1Lb31GOD34ZgIRUXGuiM2QllWiNwXFYAw"),
	// 		},
	// 		TrackingID: to.Ptr("f84f23eb-b331-4772-9f39-cc307fa83bc3"),
	// 		RepetitionIndexes: []*armappservice.RepetitionIndex{
	// 			{
	// 				ItemIndex: to.Ptr[int32](1),
	// 				ScopeName: to.Ptr("For_each"),
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/10c740b7224771c889cfb92f128168f5a0568c26/specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/WorkflowRunActionRepetitions_ListExpressionTraces.json
func ExampleWorkflowRunActionRepetitionsClient_NewListExpressionTracesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkflowRunActionRepetitionsClient().NewListExpressionTracesPager("testResourceGroup", "test-name", "testFlow", "08586776228332053161046300351", "testAction", "000001", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Inputs {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ExpressionTraces = armappservice.ExpressionTraces{
		// 	Inputs: []*armappservice.ExpressionRoot{
		// 		{
		// 			Text: to.Ptr("items('For_each')?['OccuringLocation']?['Environment']"),
		// 			Value: "PROD",
		// 			Path: to.Ptr(""),
		// 	}},
		// }
	}
}
