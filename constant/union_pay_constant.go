package constant

const (
	UNION_PAY_SUZHOU_BANK = 1
)

var (
	UNION_PAY_MAP = map[int]string{
		UNION_PAY_SUZHOU_BANK: "苏州银行聚合支付中心",
	}
	UNION_PAY_LIST = []map[string]interface{}{
		{
			"value": UNION_PAY_SUZHOU_BANK,
			"label": "苏州银行聚合支付中心",
		},
	}
)
