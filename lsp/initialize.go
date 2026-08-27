package lsp

type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
	// ... there's ton more that goes here
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"results"`
}

type InitializeResult struct {
	Capablities ServerCapablities `json:"capablities"`
	ServerInfo  ServerInfo        `json:"serverInfo"`
}

type ServerCapalities struct {
	TextDocumentSync int `json:"textDocumentSync"`

	HoverProvider      bool           `json:"hoverProvider"`
	DefinitionProvider bool           `json:"definitionProvider "`
	CodeActionProvider bool           `json:"codeAction"`
	CompletionProvider map[string]any `json:"completionProvider"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResult{
			Capablities: ServerCapalities{
				TextDocumentSync:   1,
				HoverProvider:      true,
				DefinitionProvider: true,
				CodeActionProvider: true,
				CompletionProvider: map[string]any{},
			},
			ServerInfo: ServerInfo{
				Name:    "educational lsp",
				Version: "0.0.0.0.0.0-beta1.final",
			},
		},
	}
}
