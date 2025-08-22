package main

import (
	"github.com/amazon-opensearch-ingestion/mcp-server/config"
	"github.com/amazon-opensearch-ingestion/mcp-server/models"
	tools__2022_01_01 "github.com/amazon-opensearch-ingestion/mcp-server/tools/_2022_01_01"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools__2022_01_01.CreateTagresourceTool(cfg),
		tools__2022_01_01.CreateListpipelineblueprintsTool(cfg),
		tools__2022_01_01.CreateUntagresourceTool(cfg),
		tools__2022_01_01.CreateValidatepipelineTool(cfg),
		tools__2022_01_01.CreateCreatepipelineTool(cfg),
		tools__2022_01_01.CreateGetpipelineTool(cfg),
		tools__2022_01_01.CreateGetpipelinechangeprogressTool(cfg),
		tools__2022_01_01.CreateStartpipelineTool(cfg),
		tools__2022_01_01.CreateDeletepipelineTool(cfg),
		tools__2022_01_01.CreateGetpipelineblueprintTool(cfg),
		tools__2022_01_01.CreateUpdatepipelineTool(cfg),
		tools__2022_01_01.CreateListpipelinesTool(cfg),
		tools__2022_01_01.CreateListtagsforresourceTool(cfg),
		tools__2022_01_01.CreateStoppipelineTool(cfg),
	}
}
