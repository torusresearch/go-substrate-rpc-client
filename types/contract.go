package types

type ContractABI struct {
	V3 V3 `json:"V3"`
}
type V3 struct {
	Spec Spec `json:"spec"`
}
type Spec struct {
	Messages []Message `json:"messages"`
}
type Message struct {
	Label    string `json:"label"`
	Selector string `json:"selector"`
}
