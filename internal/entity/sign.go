package entity

type Sign struct {
	Text string `json:"text,omitempty"`
	Key  string `json:"key,omitempty"`
	Hex  string `json:"hex,omitempty"`
}
