package contract

const method = "contracts_call"

func (c *contract) Call(result interface{}, callRequest CallRequest) error {
	return c.client.Call(result, method, callRequest)
}
