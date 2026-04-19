package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/nishthanabya/warp-sentinel/internal/context"
	"github.com/nishthanabya/warp-sentinel/internal/health"
)

func main() {
	color.Cyan("🛡️ WARP SENTINEL | Production Safety Gate")

	ctx := health.CheckStatus()
	if err := context.GenerateWarpContext(ctx); err != nil {
		color.Red("failed to generate warp context: %v", err)
		os.Exit(1)
	}

	if ctx.ErrorRate > 0.05 {
		color.Red("🛑 ACCESS DENIED: " + ctx.Message)
		color.Yellow("AI Agent execution restricted.")
		os.Exit(1)
	}

	color.Green("Access granted. Environment is within safety thresholds.")
}
