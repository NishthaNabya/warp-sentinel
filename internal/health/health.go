package health

// SystemContext captures production safety signals for agent gating.
type SystemContext struct {
	Env          string  `json:"Env"`
	ErrorRate    float64 `json:"ErrorRate"`
	OnCall       string  `json:"OnCall"`
	SafetyStatus string  `json:"SafetyStatus"`
	Message      string  `json:"Message"`
}

// CheckStatus returns a mocked failing production state for development and tests.
func CheckStatus() SystemContext {
	return SystemContext{
		Env:          "production",
		ErrorRate:    0.08,
		OnCall:       "Nishtha",
		SafetyStatus: "RESTRICTED",
		Message:      "High 5xx errors in payments-service",
	}
}
