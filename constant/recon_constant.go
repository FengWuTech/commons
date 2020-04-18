package constant

const (
	RECON_STATUS_WAITED    = 1
	RECON_STATUS_GENERATED = 2
	RECON_STATUS_SUBMITTED = 3
	RECON_STATUS_CONFIRMED = 4
	RECON_STATUS_FILED     = 5
)

var RECON_STATUS_MAP = []map[string]interface{}{
	{
		"label": "生成中",
		"value": RECON_STATUS_WAITED,
	},
	{
		"label": "已生成",
		"value": RECON_STATUS_GENERATED,
	},
	{
		"label": "审核中",
		"value": RECON_STATUS_SUBMITTED,
	},
	{
		"label": "已确认",
		"value": RECON_STATUS_CONFIRMED,
	},
	{
		"label": "已归档",
		"value": RECON_STATUS_FILED,
	},
}
