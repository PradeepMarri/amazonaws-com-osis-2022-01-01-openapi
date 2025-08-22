package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/amazon-opensearch-ingestion/mcp-server/config"
	"github.com/amazon-opensearch-ingestion/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func ListpipelinesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["nextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("nextToken=%v", val))
		}
		if val, ok := args["MaxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("MaxResults=%v", val))
		}
		if val, ok := args["NextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("NextToken=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/2022-01-01/osis/listPipelines%s", cfg.BaseURL, queryString)
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
		var result models.ListPipelinesResponse
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

func CreateListpipelinesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_2022-01-01_osis_listPipelines",
		mcp.WithDescription("Lists all OpenSearch Ingestion pipelines in the current Amazon Web Services account and Region. For more information, see <a href="https://docs.aws.amazon.com/opensearch-service/latest/developerguide/list-pipeline.html">Viewing Amazon OpenSearch Ingestion pipelines</a>."),
		mcp.WithNumber("maxResults", mcp.Description("An optional parameter that specifies the maximum number of results to return. You can use <code>nextToken</code> to get the next page of results.")),
		mcp.WithString("nextToken", mcp.Description("If your initial <code>ListPipelines</code> operation returns a <code>nextToken</code>, you can include the returned <code>nextToken</code> in subsequent <code>ListPipelines</code> operations, which returns results in the next page.")),
		mcp.WithString("MaxResults", mcp.Description("Pagination limit")),
		mcp.WithString("NextToken", mcp.Description("Pagination token")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ListpipelinesHandler(cfg),
	}
}
