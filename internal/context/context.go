package context

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nishthanabya/warp-sentinel/internal/health"
)

// GenerateWarpContext writes JSON context and agent rules for Warp.
func GenerateWarpContext(ctx health.SystemContext) error {
	data, err := json.MarshalIndent(ctx, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(".warp-context.json", data, 0o644); err != nil {
		return err
	}

	warpMD := fmt.Sprintf(`# Agent Rules
- SYSTEM_STATUS: %s
- DO NOT run destructive commands if ErrorRate > 0.05
- Current ErrorRate: %g
`, ctx.SafetyStatus, ctx.ErrorRate)

	return os.WriteFile("WARP.md", []byte(warpMD), 0o644)
}
