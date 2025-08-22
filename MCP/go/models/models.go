package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// CreatePipelineResponse represents the CreatePipelineResponse schema from the OpenAPI specification
type CreatePipelineResponse struct {
	Pipeline interface{} `json:"Pipeline,omitempty"`
}

// UpdatePipelineRequest represents the UpdatePipelineRequest schema from the OpenAPI specification
type UpdatePipelineRequest struct {
	Logpublishingoptions interface{} `json:"LogPublishingOptions,omitempty"`
	Maxunits interface{} `json:"MaxUnits,omitempty"`
	Minunits interface{} `json:"MinUnits,omitempty"`
	Pipelineconfigurationbody interface{} `json:"PipelineConfigurationBody,omitempty"`
}

// ValidationMessage represents the ValidationMessage schema from the OpenAPI specification
type ValidationMessage struct {
	Message interface{} `json:"Message,omitempty"`
}

// GetPipelineRequest represents the GetPipelineRequest schema from the OpenAPI specification
type GetPipelineRequest struct {
}

// LogPublishingOptions represents the LogPublishingOptions schema from the OpenAPI specification
type LogPublishingOptions struct {
	Isloggingenabled interface{} `json:"IsLoggingEnabled,omitempty"`
	Cloudwatchlogdestination interface{} `json:"CloudWatchLogDestination,omitempty"`
}

// ListPipelinesResponse represents the ListPipelinesResponse schema from the OpenAPI specification
type ListPipelinesResponse struct {
	Pipelines interface{} `json:"Pipelines,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// StopPipelineRequest represents the StopPipelineRequest schema from the OpenAPI specification
type StopPipelineRequest struct {
}

// GetPipelineResponse represents the GetPipelineResponse schema from the OpenAPI specification
type GetPipelineResponse struct {
	Pipeline interface{} `json:"Pipeline,omitempty"`
}

// ListPipelineBlueprintsResponse represents the ListPipelineBlueprintsResponse schema from the OpenAPI specification
type ListPipelineBlueprintsResponse struct {
	Blueprints interface{} `json:"Blueprints,omitempty"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// PipelineBlueprint represents the PipelineBlueprint schema from the OpenAPI specification
type PipelineBlueprint struct {
	Blueprintname interface{} `json:"BlueprintName,omitempty"`
	Pipelineconfigurationbody interface{} `json:"PipelineConfigurationBody,omitempty"`
}

// DeletePipelineRequest represents the DeletePipelineRequest schema from the OpenAPI specification
type DeletePipelineRequest struct {
}

// PipelineStatusReason represents the PipelineStatusReason schema from the OpenAPI specification
type PipelineStatusReason struct {
	Description interface{} `json:"Description,omitempty"`
}

// ChangeProgressStage represents the ChangeProgressStage schema from the OpenAPI specification
type ChangeProgressStage struct {
	Description interface{} `json:"Description,omitempty"`
	Lastupdatedat interface{} `json:"LastUpdatedAt,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"Tags,omitempty"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// GetPipelineChangeProgressRequest represents the GetPipelineChangeProgressRequest schema from the OpenAPI specification
type GetPipelineChangeProgressRequest struct {
}

// GetPipelineChangeProgressResponse represents the GetPipelineChangeProgressResponse schema from the OpenAPI specification
type GetPipelineChangeProgressResponse struct {
	Changeprogressstatuses interface{} `json:"ChangeProgressStatuses,omitempty"`
}

// DeletePipelineResponse represents the DeletePipelineResponse schema from the OpenAPI specification
type DeletePipelineResponse struct {
}

// UpdatePipelineResponse represents the UpdatePipelineResponse schema from the OpenAPI specification
type UpdatePipelineResponse struct {
	Pipeline interface{} `json:"Pipeline,omitempty"`
}

// StartPipelineRequest represents the StartPipelineRequest schema from the OpenAPI specification
type StartPipelineRequest struct {
}

// CreatePipelineRequest represents the CreatePipelineRequest schema from the OpenAPI specification
type CreatePipelineRequest struct {
	Tags interface{} `json:"Tags,omitempty"`
	Vpcoptions interface{} `json:"VpcOptions,omitempty"`
	Logpublishingoptions interface{} `json:"LogPublishingOptions,omitempty"`
	Maxunits interface{} `json:"MaxUnits"`
	Minunits interface{} `json:"MinUnits"`
	Pipelineconfigurationbody interface{} `json:"PipelineConfigurationBody"`
	Pipelinename interface{} `json:"PipelineName"`
}

// ChangeProgressStatus represents the ChangeProgressStatus schema from the OpenAPI specification
type ChangeProgressStatus struct {
	Changeprogressstages interface{} `json:"ChangeProgressStages,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Totalnumberofstages interface{} `json:"TotalNumberOfStages,omitempty"`
}

// ValidatePipelineRequest represents the ValidatePipelineRequest schema from the OpenAPI specification
type ValidatePipelineRequest struct {
	Pipelineconfigurationbody interface{} `json:"PipelineConfigurationBody"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
	Tagkeys interface{} `json:"TagKeys"`
}

// CloudWatchLogDestination represents the CloudWatchLogDestination schema from the OpenAPI specification
type CloudWatchLogDestination struct {
	Loggroup interface{} `json:"LogGroup"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"Tags"`
}

// VpcOptions represents the VpcOptions schema from the OpenAPI specification
type VpcOptions struct {
	Securitygroupids interface{} `json:"SecurityGroupIds,omitempty"`
	Subnetids interface{} `json:"SubnetIds"`
}

// ValidatePipelineResponse represents the ValidatePipelineResponse schema from the OpenAPI specification
type ValidatePipelineResponse struct {
	Isvalid interface{} `json:"isValid,omitempty"`
	Errors interface{} `json:"Errors,omitempty"`
}

// StartPipelineResponse represents the StartPipelineResponse schema from the OpenAPI specification
type StartPipelineResponse struct {
	Pipeline Pipeline `json:"Pipeline,omitempty"` // Information about an existing OpenSearch Ingestion pipeline.
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// ListPipelineBlueprintsRequest represents the ListPipelineBlueprintsRequest schema from the OpenAPI specification
type ListPipelineBlueprintsRequest struct {
}

// PipelineSummary represents the PipelineSummary schema from the OpenAPI specification
type PipelineSummary struct {
	Lastupdatedat interface{} `json:"LastUpdatedAt,omitempty"`
	Maxunits interface{} `json:"MaxUnits,omitempty"`
	Minunits interface{} `json:"MinUnits,omitempty"`
	Pipelinearn interface{} `json:"PipelineArn,omitempty"`
	Pipelinename interface{} `json:"PipelineName,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Statusreason PipelineStatusReason `json:"StatusReason,omitempty"` // Information about a pipeline's current status.
	Createdat interface{} `json:"CreatedAt,omitempty"`
}

// ListPipelinesRequest represents the ListPipelinesRequest schema from the OpenAPI specification
type ListPipelinesRequest struct {
}

// PipelineBlueprintSummary represents the PipelineBlueprintSummary schema from the OpenAPI specification
type PipelineBlueprintSummary struct {
	Blueprintname interface{} `json:"BlueprintName,omitempty"`
}

// GetPipelineBlueprintResponse represents the GetPipelineBlueprintResponse schema from the OpenAPI specification
type GetPipelineBlueprintResponse struct {
	Blueprint interface{} `json:"Blueprint,omitempty"`
}

// Pipeline represents the Pipeline schema from the OpenAPI specification
type Pipeline struct {
	Createdat interface{} `json:"CreatedAt,omitempty"`
	Maxunits interface{} `json:"MaxUnits,omitempty"`
	Vpcendpoints interface{} `json:"VpcEndpoints,omitempty"`
	Ingestendpointurls interface{} `json:"IngestEndpointUrls,omitempty"`
	Statusreason interface{} `json:"StatusReason,omitempty"`
	Lastupdatedat interface{} `json:"LastUpdatedAt,omitempty"`
	Logpublishingoptions interface{} `json:"LogPublishingOptions,omitempty"`
	Minunits interface{} `json:"MinUnits,omitempty"`
	Pipelinearn interface{} `json:"PipelineArn,omitempty"`
	Pipelineconfigurationbody interface{} `json:"PipelineConfigurationBody,omitempty"`
	Pipelinename interface{} `json:"PipelineName,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
}

// GetPipelineBlueprintRequest represents the GetPipelineBlueprintRequest schema from the OpenAPI specification
type GetPipelineBlueprintRequest struct {
}

// VpcEndpoint represents the VpcEndpoint schema from the OpenAPI specification
type VpcEndpoint struct {
	Vpcendpointid interface{} `json:"VpcEndpointId,omitempty"`
	Vpcid interface{} `json:"VpcId,omitempty"`
	Vpcoptions interface{} `json:"VpcOptions,omitempty"`
}

// StopPipelineResponse represents the StopPipelineResponse schema from the OpenAPI specification
type StopPipelineResponse struct {
	Pipeline Pipeline `json:"Pipeline,omitempty"` // Information about an existing OpenSearch Ingestion pipeline.
}
