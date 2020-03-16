package constant

const (
	SUZHOU_BANK = 1
)

var (
	BANK_MAP = map[int]string{
		SUZHOU_BANK: "苏州银行",
	}
	BANK_LIST = []map[string]interface{}{
		{
			"value": SUZHOU_BANK,
			"label": "苏州银行",
		},
	}
)
