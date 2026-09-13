package lucy

import "encoding/json"

// BuildRequest is the stable host / CLI / wasm JSON input schema.
type BuildRequest struct {
	Samples []Sample         `json:"samples"`
	Options *DensityOptions  `json:"options,omitempty"`
}

// BuildResponse is the stable host / CLI / wasm JSON output schema.
type BuildResponse struct {
	Version string          `json:"version"`
	Options DensityOptions  `json:"options"`
	Board   LPD             `json:"board"`
}

// BuildFromJSON unmarshals a BuildRequest, runs BuildLPDWithOptions, returns a response.
func BuildFromJSON(raw []byte) (BuildResponse, error) {
	var req BuildRequest
	if err := json.Unmarshal(raw, &req); err != nil {
		return BuildResponse{}, err
	}
	opts := DensityOptions{}
	if req.Options != nil {
		opts = *req.Options
	}
	opts = opts.normalized()
	board := BuildLPDWithOptions(req.Samples, opts)
	return BuildResponse{
		Version: Version,
		Options: opts,
		Board:   board,
	}, nil
}

// MarshalBoard is a convenience for CLI/wasm stdout.
func MarshalBoard(resp BuildResponse) ([]byte, error) {
	return json.Marshal(resp)
}
