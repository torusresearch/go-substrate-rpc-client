package contract

type CallRequest struct {
	Origin              string `json:"origin"`
	Dest                string `json:"dest"`
	Value               int64  `json:"value"`
	GasLimit            int64  `json:"gasLimit"`
	StorageDepositLimit *int64 `json:"storageDepositLimit"`
	InputData           string `json:"inputData"`
}
