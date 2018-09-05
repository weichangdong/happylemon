package conf

const (
	//seccess
	CodeOk = 0
	//aes error1
	CodeAesErr = 1
	//param error
	CodeParaErr      = 3
	Code404Err       = 4
	Code500Err       = 5
	CodeTelRegedErr  = 6  //该手机号已注册过
	CodeSmsCodeErr   = 7  //短信验证码非法或者已经过期,请重新获取
	CodeBeBlockedErr = 8  //登录或者请求接口时,用户被封禁
	CodePassWdErr    = 9  //密码输入错误
	CodeFrost1Err    = 10 //密码输入错误次数太多,处于冻结期(1小时)
	CodeFrost2Err    = 11 //密码输入错误次数太多,处于冻结期(24小时)

	CodeThirdRegedErr = 13 //绑定第三方的时候,第三方账号已经注册过了
	CodeNoTelErr      = 14 //解绑的时候,这个账号没有绑定手机号码
	CodeBeCanceldErr  = 15 //登陆提醒,这个账号已经被注销了

	CodeReachTopErr          = 17 //创建圈子,达到3个上限了
	CodeNameRepeatErr        = 18 //创建圈子,圈子名称重复
	CodeWorldNotAllowErr     = 19 //各种创建,文字含有敏感信息,屏蔽字
	CodeSmsLimitMinute       = 20 //发送短信限制： 每分钟只能发送 1 次
	CodeSmsLimitHour         = 21 //发送短信限制： 每小时只能发送 5 次
	CodeSmsSendErr           = 22 //发送短信验证码失败
	CodeWelcodeMsgShowed     = 23 //圈子的欢迎语关闭了或者已经展示过欢迎语了
	CodeNoNewInfo            = 24 //没有更多的数据了,数据是最新的
	CodeFeedsNotExists       = 25 //帖子已经不存在
	CodeFeedsCloseTalk       = 26 //这个帖子关闭评论了
	CodeTalkTooFast          = 27 //评论频率过快
	CodeSoNoMatchErr         = 28 //搜索没有匹配的结果
	CodeNoOldInfo            = 29 //没有更多的数据了,数据是最旧的
	CodeNoGroupNotice        = 30 //圈子没有设置公告
	CodeGroupStatBeBlocked   = 31 //圈子状态-圈子被封禁
	CodeGroupStatBeDissolved = 32 //圈子状态-圈子被解散
	CodeWebSocketRetError    = 33 //grpc websocket 数据返回错误
	CodeWebSocketConnError   = 34 //grpc websocket 链接错误
	CodePushConnError        = 35 //grpc push 链接错误
	CodeWordConnError        = 36 //grpc word 链接错误
	CodeTalkNotExists        = 37 //评论已经不存在了

	CodeNotExixts          = 40 //圈子不存在或已被删除
	CodeReOperate          = 41 //请勿重复操作
	CodeGroupWatingApply   = 42 //已发送申请，等待审核
	CodeGroupDeniedJoin    = 43 //圈子拒绝申请加入
	CodeNotPermission      = 44 //本次操作无权限
	CodeMemNotJoinGroup    = 45 //对方未加入圈子或已退出
	CodeMemJoinedGroup     = 46 //对方已加入该圈子
	CodeMemNotGroupAdmin   = 47 //对方不是该圈子管理员
	CodeMemNotGroupCreator = 48 //非创建者：只有创建者才有权申请解散该圈子
	CodeGroupNameIsSet     = 49 //圈子名称已存在
	CodeIsBlacked          = 50 //无权操作：被拉黑
	CodeSmsLimitDay        = 51 //发送短信限制： 每天只能发送 10 次
	CodeAdminMaxLimit      = 52 //超过最大管理员数量
	CodeMobileNotReg       = 53 //手机号未注册(未绑定)
	CodeOrderDealToBlock   = 54 //信息已处理：被拉黑
	CodeOrderDealToRefuse  = 55 //信息已处理：被拒绝
	CodeOrderDealToAgree   = 56 //信息已处理：已同意
	CodeNumOrderExceed     = 57 //设置信息数量以超过规定数量

	CodeMyGroupNotExixts = 71 //没有我的圈子或推荐圈子
	CodeBannerNotExixts  = 72 //后台没有配置banner
	CodeThemeNotExixts   = 73 //后台没有配置主题
	CodeHeadlineOff      = 74 //上次保存记录客户端上传有问题
	CodeUserNotExixts    = 75 //用户不存在
	CodeSaveHeadline     = 76 //用户保存headline失败
	CodeDataVersion      = 77 //数据版本相同，不下发新数据
	CodeVersionToHigh    = 78 //app版本不匹配，无可下发数据
	CodeNoMoreData       = 79 //没有更多数据
)
