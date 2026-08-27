package lsp

type PuslbishDiagnosticNotification struct {
	Notifications
	Params PushbulletDiagnosticParams `json:"params"`
}

type PushbulletDiagnosticParams struct {
	Uri         string       `json:"uri"`
	Diagnostics []Diagnostic `json:"diagnostics"`
}

type Diagnostic struct {
	Range Range `json:"range"`
	// Severity defines how the message should be interpreted.
	Severity DiagnosticSeverity `json:"severity"`
	// Code is a number that identifies this particular occurrence of the problem.
	Code string `json:"code"`
	// Source is a human-readable string describing the source of this
	// diagnostic, e.g. 'typescript' or 'super lint'.
	Source string `json:"source"`
	// Message is the diagnostic message. It should be phrased in a human-readable
	// way.
	Message string `json:"message"`
}
