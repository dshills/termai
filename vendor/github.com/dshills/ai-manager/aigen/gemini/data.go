package gemini

type Request struct {
	Contents []Content `json:"contents"`
}

type Response struct {
	Candidates []Candidate
}

type Candidate struct {
	Content      Content `json:"content"`
	FinishReason string  `json:"finish_reason"`
	TokenCount   int     `json:"token_count"`
	Index        int     `json:"index"`
}

type Content struct {
	Role  string `json:"role"`
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}
