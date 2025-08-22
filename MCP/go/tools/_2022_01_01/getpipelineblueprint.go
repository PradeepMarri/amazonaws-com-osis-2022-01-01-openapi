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

func GetpipelineblueprintHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		BlueprintNameVal, ok := args["BlueprintName"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: BlueprintName"), nil
		}
		BlueprintName, ok := BlueprintNameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: BlueprintName"), nil
		}
		url := fmt.Sprintf("%s/2022-01-01/osis/getPipelineBlueprint/%s", cfg.BaseURL, BlueprintName)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.GetPipelineBlueprintResponse
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

func CreateGetpipelineblueprintTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_2022-01-01_osis_getPipelineBlueprint_BlueprintName",
		mcp.WithDescription("Retrieves information about a specific blueprint for OpenSearch Ingestion. Blueprints are templates for the configuration needed for a <code>CreatePipeline</code> request. For more information, see <a href="https://docs.aws.amazon.com/opensearch-service/latest/developerguide/creating-pipeline.html#pipeline-blueprint">Using blueprints to create a pipeline</a>."),
		mcp.WithString("BlueprintName", mcp.Required(), mcp.Description("The name of the blueprint to retrieve.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetpipelineblueprintHandler(cfg),
	}
}
