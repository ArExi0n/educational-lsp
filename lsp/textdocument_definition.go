package lsp

type DefinitonRequest struct {
	Request
	Params DefinitionParams `json:"params"`
}

type DefinitionParams struct {
	TextDocumentPositionParams
}

type DefinitionResponse struct {
	Response
	Result Location `json:"result"`
}
