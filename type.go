package trongrid

type Error struct {
	Error string `json:"error"`
}

type Meta struct {
	Links       *MetaLinks `json:"links"`
	Fingerprint string     `json:"fingerprint"`
	At          int64      `json:"at"`
	PageSize    int32      `json:"page_size"`
}

type MetaLinks struct {
	Next string `json:"next"`
}

type Token struct {
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int32  `json:"decimals"`
}

type Transaction struct {
	Ret []struct {
		ContractRet string `json:"contractRet"`
		Fee         int    `json:"fee"`
	} `json:"ret"`
	Signature        []string `json:"signature"`
	TxID             string   `json:"txID"`
	NetUsage         int      `json:"net_usage"`
	RawDataHex       string   `json:"raw_data_hex"`
	NetFee           int      `json:"net_fee"`
	EnergyUsage      int      `json:"energy_usage"`
	BlockNumber      int      `json:"blockNumber"`
	BlockTimestamp   int64    `json:"block_timestamp"`
	EnergyFee        int      `json:"energy_fee"`
	EnergyUsageTotal int      `json:"energy_usage_total"`
	RawData          struct {
		Contract []struct {
			Parameter struct {
				Value struct {
					Amount       int    `json:"amount"`
					OwnerAddress string `json:"owner_address"`
					ToAddress    string `json:"to_address"`
				} `json:"value"`
				TypeUrl string `json:"type_url"`
			} `json:"parameter"`
			Type string `json:"type"`
		} `json:"contract"`
		RefBlockBytes string `json:"ref_block_bytes"`
		RefBlockHash  string `json:"ref_block_hash"`
		Expiration    int64  `json:"expiration"`
		Timestamp     int64  `json:"timestamp"`
	} `json:"raw_data"`
	InternalTransactions []interface{} `json:"internal_transactions"`
}
type TransactionType string
