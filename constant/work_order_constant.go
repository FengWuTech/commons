package constant

const (
	PROCESS_REPAIR = "id_process_repair" //维修工单对应的BPMN ID

	PROCESS_REPAIR_NODE_ACCEPT = "id_accept"
	PROCESS_REPAIR_NODE_REPAIR = "id_repair"
	PROCESS_REPAIR_NODE_VERIFY = "id_verify"

	PROCESS_POOL_TYPE_REPAIR = 1 //维修工单

	//人相关
	APPEAL_TYPE_THING     = 1 //报事
	APPEAL_TYPE_COMPLAINT = 2 //投诉建议
	APPEAL_TYPE_PRAISE    = 3 //表扬赞赏
	APPEAL_TYPE_REPAIR    = 4 //报修
	//视觉(智能硬件)相关
	APPEAL_TYPE_VISION_FIRE_ESCAPE     = 5 //消防通道
	APPEAL_TYPE_VISION_DANGER_INVASION = 6 //危险区域
	APPEAL_TYPE_VISION_CAMERA_FAILURE  = 7 //设备故障
	APPEAL_TYPE_VISION_GARBAGE_LASTING = 8 //垃圾溢满

	//表扬投诉
	APPEAL_OBJECT_TYPE_CLEAN            = 1 //保洁
	APPEAL_OBJECT_TYPE_SECURITY         = 2 //保安
	APPEAL_OBJECT_TYPE_GREEN            = 3 //保绿
	APPEAL_OBJECT_TYPE_CUSTOMER_SERVICE = 4 //客服
	APPEAL_OBJECT_TYPE_REPAIR           = 5 //维修
	APPEAL_OBJECT_TYPE_OTHER            = 6 //其他

	//报事
	APPEAL_OBJECT_TYPE_THING_CLEAN            = 11 //保洁
	APPEAL_OBJECT_TYPE_THING_SECURITY         = 12 //保安
	APPEAL_OBJECT_TYPE_THING_GREEN            = 13 //保绿
	APPEAL_OBJECT_TYPE_THING_CUSTOMER_SERVICE = 14 //客服
	APPEAL_OBJECT_TYPE_THING_NEIGHBOR         = 15 //邻里
	APPEAL_OBJECT_TYPE_THING_OTHER            = 16 //其他

	//报修
	APPEAL_OBJECT_TYPE_REPAIR_LIGHT             = 21 //公共照明
	APPEAL_OBJECT_TYPE_REPAIR_WATER_ELECTRONICS = 22 //水电
	APPEAL_OBJECT_TYPE_REPAIR_UNIT_GATE         = 23 //单元门禁
	APPEAL_OBJECT_TYPE_REPAIR_ELEVATOR          = 24 //电梯
	APPEAL_OBJECT_TYPE_REPAIR_PARKING           = 25 //停车场
	APPEAL_OBJECT_TYPE_REPAIR_BOARD             = 26 //公告设备
	APPEAL_OBJECT_TYPE_REPAIR_OTHER             = 27 //其他

	APPEAL_ORDER_SOURCE_CUSTOMER_SERVICE = 1 //客服
	APPEAL_ORDER_SOURCE_WECHAT           = 2 //微信
	APPEAL_ORDER_SOURCE_INTEL_HARDWARE   = 3 //智慧硬件

	APPEAL_ORDER_STATUS_CREATED        = 1 //已创建
	APPEAL_ORDER_STATUS_ACCEPTED       = 2 //已接单
	APPEAL_ORDER_STATUS_DISPATCHED     = 3 //已分配
	APPEAL_ORDER_STATUS_REPAIRED       = 4 //维修完成
	APPEAL_ORDER_STATUS_RETURN_VISITED = 5 //回访完成
	APPEAL_ORDER_STATUS_FINISHED       = 6 //已结单

	DISPATCH_ORDER_STATUS_DISPATCHED = 1 //已派工
	DISPATCH_ORDER_STATUS_ACCEPTED   = 2 //已接单
	DISPATCH_ORDER_STATUS_REPAIRED   = 3 //已修复
	DISPATCH_ORDER_STATUS_VERIFIED   = 4 //已验证
	//DISPATCH_ORDER_STATUS_TRANSFER   = 5 //已转办
	//DISPATCH_ORDER_STATUS_ENTRUSTED  = 6 //已委托
	//DISPATCH_ORDER_STATUS_REJECTED   = 7 //已驳回
	//DISPATCH_ORDER_STATUS_FROZEN     = 8 //已冻结
	DISPATCH_ORDER_STATUS_CANCEL = 9 //已取消

	DISPATCH_ORDER_OPERATE_TYPE_APPOINTMENT = 1 //修改预约
	//DISPATCH_ORDER_OPERATE_TYPE_TRANSFER    = 2 //转派
	//DISPATCH_ORDER_OPERATE_TYPE_ENTRUST     = 3 //委托
	DISPATCH_ORDER_OPERATE_TYPE_FINISH = 4 //完工
	//DISPATCH_ORDER_OPERATE_TYPE_REJECT      = 5 //驳回
	//DISPATCH_ORDER_OPERATE_TYPE_FROZE       = 6 //冻结

	RETURN_VISIT_STATUS_CREATED             = 1 //已创建
	RETURN_VISIT_STATUS_FINISH_WITH_PROBLEM = 2 //回访结束，存在问题
	RETURN_VISIT_STATUS_FINISH_NO_PROBLEM   = 3 //回访结束，无问题

	RETURN_VISIT_RESULT_YES = 1 //回访问题回答yes
	RETURN_VISIT_RESULT_NO  = 0 //回访问题回答no

	DISPATCH_EMERGENCY_LEVEL_NORMAL = 1 //普通
	DISPATCH_EMERGENCY_LEVEL_URGENT = 2 //加急

	APPEAL_ORDER_EFFECT_EFFECTIVE    = 1 //有效的
	APPEAL_ORDER_EFFECT_COMMUNICABLE = 2 //可沟通的
	APPEAL_ORDER_EFFECT_INVALID      = 3 //无效的
)

func GetOrderFeConf(key string) []map[string]interface{} {
	return map[string][]map[string]interface{}{
		"APPEAL_TYPE": {
			{
				"value": APPEAL_TYPE_THING,
				"label": "报事问题",
			},
			{
				"value": APPEAL_TYPE_REPAIR,
				"label": "报修问题",
			},
			{
				"value": APPEAL_TYPE_COMPLAINT,
				"label": "投诉建议",
			},
			{
				"value": APPEAL_TYPE_PRAISE,
				"label": "表扬赞赏",
			},
			{
				"value": APPEAL_TYPE_VISION_FIRE_ESCAPE,
				"label": "消防通道报警",
			},
			{
				"value": APPEAL_TYPE_VISION_DANGER_INVASION,
				"label": "危险区域报警",
			},
			{
				"value": APPEAL_TYPE_VISION_CAMERA_FAILURE,
				"label": "摄像机故障",
			},
			{
				"value": APPEAL_TYPE_VISION_GARBAGE_LASTING,
				"label": "垃圾溢满报警",
			},
		},
		"APPEAL_TYPE_PMS": {
			{
				"value": APPEAL_TYPE_THING,
				"label": "报事问题",
			},
			{
				"value": APPEAL_TYPE_REPAIR,
				"label": "报修问题",
			},
			{
				"value": APPEAL_TYPE_COMPLAINT,
				"label": "投诉建议",
			},
			{
				"value": APPEAL_TYPE_PRAISE,
				"label": "表扬赞赏",
			},
		},
		"APPEAL_TYPE_VISION": {
			{
				"value": APPEAL_TYPE_VISION_FIRE_ESCAPE,
				"label": "消防通道报警",
			},
			{
				"value": APPEAL_TYPE_VISION_DANGER_INVASION,
				"label": "危险区域报警",
			},
			{
				"value": APPEAL_TYPE_VISION_CAMERA_FAILURE,
				"label": "摄像机故障",
			},
			{
				"value": APPEAL_TYPE_VISION_GARBAGE_LASTING,
				"label": "垃圾溢满报警",
			},
		},
		"APPEAL_OBJECT_TYPE": { //表扬投诉用
			{
				"value": APPEAL_OBJECT_TYPE_CLEAN,
				"label": "保洁",
			},
			{
				"value": APPEAL_OBJECT_TYPE_SECURITY,
				"label": "保安",
			},
			{
				"value": APPEAL_OBJECT_TYPE_GREEN,
				"label": "保绿",
			},
			{
				"value": APPEAL_OBJECT_TYPE_CUSTOMER_SERVICE,
				"label": "客服",
			},
			{
				"value": APPEAL_OBJECT_TYPE_REPAIR,
				"label": "维修",
			},
			{
				"value": APPEAL_OBJECT_TYPE_OTHER,
				"label": "其他",
			},
		},
		"APPEAL_OBJECT_TYPE_THING": { //报事用
			{
				"value": APPEAL_OBJECT_TYPE_THING_CLEAN,
				"label": "保洁",
			},
			{
				"value": APPEAL_OBJECT_TYPE_THING_SECURITY,
				"label": "保安",
			},
			{
				"value": APPEAL_OBJECT_TYPE_THING_GREEN,
				"label": "绿化",
			},
			{
				"value": APPEAL_OBJECT_TYPE_THING_CUSTOMER_SERVICE,
				"label": "客服",
			},
			{
				"value": APPEAL_OBJECT_TYPE_THING_NEIGHBOR,
				"label": "邻里",
			},
			{
				"value": APPEAL_OBJECT_TYPE_THING_OTHER,
				"label": "其他",
			},
		},
		"APPEAL_OBJECT_TYPE_REPAIR": { //报修用
			{
				"value": APPEAL_OBJECT_TYPE_REPAIR_LIGHT,
				"label": "公共照明",
			},
			{
				"value": APPEAL_OBJECT_TYPE_REPAIR_WATER_ELECTRONICS,
				"label": "水电",
			},
			{
				"value": APPEAL_OBJECT_TYPE_REPAIR_UNIT_GATE,
				"label": "单元门禁",
			},
			{
				"value": APPEAL_OBJECT_TYPE_REPAIR_ELEVATOR,
				"label": "电梯",
			},
			{
				"value": APPEAL_OBJECT_TYPE_REPAIR_PARKING,
				"label": "停车场",
			},
			{
				"value": APPEAL_OBJECT_TYPE_REPAIR_BOARD,
				"label": "公告设备",
			},
			{
				"value": APPEAL_OBJECT_TYPE_REPAIR_OTHER,
				"label": "其他",
			},
		},
		"APPEAL_ORDER_SOURCE": {
			{
				"value": APPEAL_ORDER_SOURCE_CUSTOMER_SERVICE,
				"label": "客服",
			},
			{
				"value": APPEAL_ORDER_SOURCE_WECHAT,
				"label": "微信",
			},
			{
				"value": APPEAL_ORDER_SOURCE_INTEL_HARDWARE,
				"label": "智慧硬件",
			},
		},
		"APPEAL_ORDER_STATUS": {
			{
				"value": APPEAL_ORDER_STATUS_CREATED,
				"label": "已创建",
			},
			{
				"value": APPEAL_ORDER_STATUS_ACCEPTED,
				"label": "已受理",
			},
			{
				"value": APPEAL_ORDER_STATUS_DISPATCHED,
				"label": "已派工",
			},
			{
				"value": APPEAL_ORDER_STATUS_REPAIRED,
				"label": "已处理",
			},
			{
				"value": APPEAL_ORDER_STATUS_RETURN_VISITED,
				"label": "已回访",
			},
			{
				"value": APPEAL_ORDER_STATUS_FINISHED,
				"label": "已结单",
			},
		},
		"APPEAL_ORDER_STATUS_C": {
			{
				"value": APPEAL_ORDER_STATUS_CREATED,
				"label": "已创建",
			},
			{
				"value": APPEAL_ORDER_STATUS_ACCEPTED,
				"label": "已受理",
			},
			{
				"value": APPEAL_ORDER_STATUS_DISPATCHED,
				"label": "已派工",
			},
			{
				"value": APPEAL_ORDER_STATUS_REPAIRED,
				"label": "已处理",
			},
			{
				"value": APPEAL_ORDER_STATUS_RETURN_VISITED,
				"label": "已处理",
			},
			{
				"value": APPEAL_ORDER_STATUS_FINISHED,
				"label": "已处理",
			},
		},
		"DISPATCH_ORDER_STATUS": {
			{
				"value": DISPATCH_ORDER_STATUS_DISPATCHED,
				"label": "已派工",
			},
			{
				"value": DISPATCH_ORDER_STATUS_ACCEPTED,
				"label": "已响应",
			},
			{
				"value": DISPATCH_ORDER_STATUS_REPAIRED,
				"label": "已处理",
			},
			{
				"value": DISPATCH_ORDER_STATUS_VERIFIED,
				"label": "已完成",
			},
			//{
			//	"value": DISPATCH_ORDER_STATUS_TRANSFER,
			//	"label": "已转单",
			//},
			//{
			//	"value": DISPATCH_ORDER_STATUS_ENTRUSTED,
			//	"label": "已委托",
			//},
			//{
			//	"value": DISPATCH_ORDER_STATUS_REJECTED,
			//	"label": "已驳回",
			//},
			//{
			//	"value": DISPATCH_ORDER_STATUS_FROZEN,
			//	"label": "已冻结",
			//},
			{
				"value": DISPATCH_ORDER_STATUS_CANCEL,
				"label": "已取消",
			},
		},
		"RETURN_VISIT_STATUS": {
			{
				"value": RETURN_VISIT_STATUS_CREATED,
				"label": "待回访",
			},
			{
				"value": RETURN_VISIT_STATUS_FINISH_WITH_PROBLEM,
				"label": "已回访，重修复",
			},
			{
				"value": RETURN_VISIT_STATUS_FINISH_NO_PROBLEM,
				"label": "已回访，无问题",
			},
		},
		"DISPATCH_ORDER_OPERATE_TYPE": {
			{
				"value": DISPATCH_ORDER_OPERATE_TYPE_APPOINTMENT,
				"label": "预约",
			},
			//{
			//	"value": DISPATCH_ORDER_OPERATE_TYPE_TRANSFER,
			//	"label": "转派",
			//},
			//{
			//	"value": DISPATCH_ORDER_OPERATE_TYPE_ENTRUST,
			//	"label": "委托",
			//},
			{
				"value": DISPATCH_ORDER_OPERATE_TYPE_FINISH,
				"label": "完工",
			},
			//{
			//	"value": DISPATCH_ORDER_OPERATE_TYPE_REJECT,
			//	"label": "驳回",
			//},
			//{
			//	"value": DISPATCH_ORDER_OPERATE_TYPE_FROZE,
			//	"label": "冻结",
			//},
		},
	}[key]
}

func ObjectTypeToName(type_ int) string {
	types := GetOrderFeConf("APPEAL_OBJECT_TYPE")
	for _, v := range types {
		value, _ := v["value"]
		if value == type_ {
			label, _ := v["label"]
			return label.(string)
		}
	}

	types = GetOrderFeConf("APPEAL_OBJECT_TYPE_THING")
	for _, v := range types {
		value, _ := v["value"]
		if value == type_ {
			label, _ := v["label"]
			return label.(string)
		}
	}

	types = GetOrderFeConf("APPEAL_OBJECT_TYPE_REPAIR")
	for _, v := range types {
		value, _ := v["value"]
		if value == type_ {
			label, _ := v["label"]
			return label.(string)
		}
	}

	return ""
}
