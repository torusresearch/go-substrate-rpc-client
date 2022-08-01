package contract

type CallRequest struct {
	Origin              string `json:"origin"`
	Dest                string `json:"dest"`
	Value               int64  `json:"value"`
	GasLimit            int64  `json:"gasLimit"`
	StorageDepositLimit *int64 `json:"storageDepositLimit"`
	InputData           string `json:"inputData"`
}

type CallResponse struct {
	GasConsumed    float64        `mapstructure:"gasConsumed"`
	GasRequired    float64        `mapstructure:"gasRequired"`
	StorageDeposit StorageDeposit `mapstructure:"storageDeposit"`
	DebugMessage   string         `mapstructure:"debugMessage"`
	Result         Result         `mapstructure:"result"`
}

type StorageDeposit struct {
	Charge string `mapstructure:"charge"`
}

type Result struct {
	Err     *Err     `mapstructure:"Err"`
	Success *Success `mapstructure:"ok"`
}

type Err struct {
	Module Module `mapstructure:"Module"`
}

type Module struct {
	Index   float64   `mapstructure:"index"`
	Error   []float64 `mapstructure:"error"`
	Message string    `mapstructure:"message"`
}

type Success struct {
	Flags uint64 `mapstructure:"flags"`
	Data  string `mapstructure:"data"`
}
