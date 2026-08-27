package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`

	// We will just specify the type of the params in all the request types
	// Params...
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id,omitempty"`

	// Result
	// Error
}

type Notifications struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}
