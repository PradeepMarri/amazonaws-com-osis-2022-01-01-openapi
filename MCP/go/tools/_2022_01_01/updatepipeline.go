package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/amazon-opensearch-ingestion/mcp-server/config"
	"github.com/amazon-opensearch-ingestion/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func UpdatepipelineHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		PipelineNameVal, ok := args["PipelineName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: PipelineName"), nil
		}
		PipelineName, ok := PipelineNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: PipelineName"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/2022-01-01/osis/updatePipeline/%s", cfg.BaseURL, PipelineName)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.UpdatePipelineResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateUpdatepipelineTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_2022-01-01_osis_updatePipeline_PipelineName",
		mcp.WithDescription("Updates an OpenSearch Ingestion pipeline. For more information, see <a href="https://docs.aws.amazon.com/opensearch-service/latest/developerguide/update-pipeline.html">Updating Amazon OpenSearch Ingestion pipelines</a>."),
		mcp.WithString("PipelineName", mcp.Required(), mcp.Description("The name of the pipeline to update.")),
		mcp.WithNumber("MaxUnits", mcp.Description("Input parameter: The maximum pipeline capacity, in Ingestion Compute Units (ICUs)")),
		mcp.WithNumber("MinUnits", mcp.Description("Input parameter: The minimum pipeline capacity, in Ingestion Compute Units (ICUs).")),
		mcp.WithString("PipelineConfigurationBody", mcp.Description("Input parameter: The pipeline configuration in YAML format. The command accepts the pipeline configuration as a string or within a .yaml file. If you provide the configuration as a string, each new line must be escaped with <code>\\n</code>.")),
		mcp.WithObject("LogPublishingOptions", mcp.Description("Input parameter: Container for the values required to configure logging for the pipeline. If you don't specify these values, OpenSearch Ingestion will not publish logs from your application to CloudWatch Logs.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    UpdatepipelineHandler(cfg),
	}
}
