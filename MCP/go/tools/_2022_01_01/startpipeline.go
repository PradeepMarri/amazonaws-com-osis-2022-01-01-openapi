package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/amazon-opensearch-ingestion/mcp-server/config"
	"github.com/amazon-opensearch-ingestion/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func StartpipelineHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/2022-01-01/osis/startPipeline/%s", cfg.BaseURL, PipelineName)
		req, err := http.NewRequest("PUT", url, nil)
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
		var result models.StartPipelineResponse
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

func CreateStartpipelineTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_2022-01-01_osis_startPipeline_PipelineName",
		mcp.WithDescription("Starts an OpenSearch Ingestion pipeline. For more information, see <a href="https://docs.aws.amazon.com/opensearch-service/latest/developerguide/pipeline--stop-start.html#pipeline--start">Starting an OpenSearch Ingestion pipeline</a>."),
		mcp.WithString("PipelineName", mcp.Required(), mcp.Description("The name of the pipeline to start.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    StartpipelineHandler(cfg),
	}
}
