package config

type Config struct {
	NodeUrl string `json:"node_url"`

	GasPrice string `json:"gas_price"`
	GasLimit string `json:"gas_limit"`

	AuthorityAddress string `json:"authority_address"`
	ChargeAddress    string `json:"charge_address"`
	Ddc721Address    string `json:"ddc_721_address"`
	Ddc1155Address   string `json:"ddc_1155_address"`
}
