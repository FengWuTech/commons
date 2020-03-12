package constant

const (
	// 问卷状态
	TopicStatusInit  = 0 //未开始
	TopicStatusStart = 1 //进行中
	TopicStatusEnd   = 2 //已结束

	// 微信发送筛选
	SurveyWechatUserReceiverAll  = 0 //所有订阅用户
	SurveyWechatUserReceiverUser = 1 //业主
	// 微信发送业主筛选
	SurveyWechatUserVotedIgnore = 0 //所有业主不论投票
	SurveyWechatUserVotedNot    = 1 //所有未投票用户

	// 问卷类型
	SurveyTopicTemplateTrue  = 1 //模版类型
	SurveyTopicTemplateFalse = 0 //普通类型
)

const (
	_                           = iota
	SurveyQuestionTypeSingle    //单选
	SurveyQuestionTypeMultiple  //多选
	SurveyQuestionTypeUserInput //用户填空
)
