package arg

type Department struct {
	// 部门id
	Id int64 `json:"id"`
	//部门名称
	Name string `json:"name"`
	// 父部门id，根部门为1
	Parentid int64 `json:"parentid"`
	// 当前部门在父部门下的所有子部门中的排序值
	Order int64 `json:"order"`
	// 是否同步创建一个关联此部门的企业群，true表示是，false表示不是
	CreateDeptGroup bool `json:"createDeptGroup"`
	// 当部门群已经创建后，是否有新人加入部门会自动加入该群，true表示是，false表示不是
	AutoAddUser bool `json:"autoAddUser"`
	// 是否隐藏部门，true表示隐藏，false表示显示
	DeptHiding bool `json:"deptHiding"`
	// 可以查看指定隐藏部门的其他部门列表，如果部门隐藏，则此值生效，取值为其他的部门id组成的的字符串，使用“\|”符号进行分割
	DeptPermits string `json:"deptPermits"`
	// 可以查看指定隐藏部门的其他人员列表，如果部门隐藏，则此值生效，取值为其他的人员userid组成的的字符串，使用“\|”符号进行分割
	UserPermits string `json:"userPermits"`
	// 是否本部门的员工仅可见员工自己，为true时，本部门员工默认只能看到员工自己
	OuterDept bool `json:"outerDept"`
	// 本部门的员工仅可见员工自己为true时，可以配置额外可见部门，值为部门id组成的的字符串，使用“\|”符号进行分割
	OuterPermitDepts string `json:"outerPermitDepts"`
	// 本部门的员工仅可见员工自己为true时，可以配置额外可见人员，值为userid组成的的字符串，使用“\|”符号进行分割
	OuterPermitUsers string `json:"outerPermitUsers"`
	// 企业群群主
	OrgDeptOwner string `json:"orgDeptOwner"`
	// 部门的主管列表，取值为由主管的userid组成的字符串，不同的userid使用“\|”符号进行分割
	DeptManagerUseridList string `json:"deptManagerUseridList"`
	// 部门标识字段，开发者可用该字段来唯一标识一个部门，并与钉钉外部通讯录里的部门做映射
	SourceIdentifier string `json:"sourceIdentifier"`
	// 部门群是否包含子部门
	GroupContainSubDept bool `json:"groupContainSubDept"`
	// 部门自定义字段
	Ext string `json:"ext"`
}

type SimpleDepartment struct {
	// 部门id
	Id int64 `json:"id"`
	//部门名称
	Name string `json:"name"`
	// 父部门id，根部门为1
	Parentid int64 `json:"parentid"`
	// 是否同步创建一个关联此部门的企业群，true表示是，false表示不是
	CreateDeptGroup bool `json:"createDeptGroup"`
	// 当部门群已经创建后，是否有新人加入部门会自动加入该群，true表示是，false表示不是
	AutoAddUser bool `json:"autoAddUser"`
	// 部门自定义字段
	Ext string `json:"ext"`
}

type UserInfo struct {
	// 员工姓名
	Name string `json:"name"`
	// 员工在当前企业内的唯一标识，也称staffId
	Userid string `json:"userid"`
	// 是否是管理员，true：是，false：不是
	IsSys bool `json:"is_sys"`
	// 级别，1：主管理员，2：子管理员，100：老板，0：其他（如普通员工）
	SysLevel int32 `json:"sys_level"`
}

type (
	User struct {
		Unionid         string  `json:"unionid"`
		Userid          string  `json:"userid"`
		Name            string  `json:"name"`
		Tel             string  `json:"tel"`
		WorkPlace       string  `json:"workPlace"`
		Remark          string  `json:"remark"`
		Mobile          string  `json:"mobile"`
		Email           string  `json:"email"`
		OrgEmail        string  `json:"orgEmail"`
		Active          bool    `json:"active"`
		OrderInDepts    string  `json:"orderInDepts"`
		IsAdmin         bool    `json:"isAdmin"`
		IsBoss          bool    `json:"isBoss"`
		IsLeaderInDepts string  `json:"isLeaderInDepts"`
		IsHide          bool    `json:"isHide"`
		Department      []int64 `json:"department"`
		Position        string  `json:"position"`
		Avatar          string  `json:"avatar"`
		HiredDate       int64   `json:"hiredDate"`
		Jobnumber       string  `json:"jobnumber"`
		IsSenior        bool    `json:"isSenior"`
		StateCode       string  `json:"stateCode"`
		Roles           []*Role `json:"roles"`
	}
	IsLeaderInDepts struct {
		// 部门ID
		Key int32
		// 是否为主管
		Value bool
	}
	Role struct {
		// 角色id
		Id int64
		// 角色名称
		Name string
		// 角色组名称
		GroupName string
	}
)

type Userid struct {
	// 联系类型，0表示企业内部员工，1表示企业外部联系人
	ContactType int32 `json:"contactType"`
	// 员工id
	Userid string `json:"userid"`
}

type AdminList struct {
	// 管理员角色，1:主管理员，2:子管理员
	SysLevel int32 `json:"sys_level"`
	// 员工id
	Userid string `json:"userid"`
}

type UserList struct {
	// 成员列表
	Userlist []*User `json:"userlist"`
	// 在分页查询时返回，true代表还有下一页更多数据
	HasMore bool `json:"hasMore"`
}

type SimpleUser struct {
	// 成员名称
	Name string `json:"name"`
	// 员工id
	Userid string `json:"userid"`
}

type SimpleUserList struct {
	// 成员列表
	Userlist []*SimpleUser `json:"userlist"`
	// 在分页查询时返回，true代表还有下一页更多数据
	HasMore bool `json:"hasMore"`
}

type (
	// 授权方管理员信息
	AuthUserInfo struct {
		UserId string `json:"userId"`
	}
	// 授权方企业信息
	AuthCorpInfo struct {
		// 授权方企业id
		Corpid string `json:"corpid"`
		// 企业认证等级，0：未认证，1：高级认证，2：中级认证，3：初级认证
		AuthLevel int `json:"auth_level"`
		// 渠道码
		AuthChannel string `json:"auth_channel"`
		// 企业所属行业
		Industry     string `json:"industry"`
		FullCorpName string `json:"full_corp_name"`
		// 授权方企业名称
		CorpName string `json:"corp_name"`
		// 企业邀请链接
		InviteUrl string `json:"invite_url"`
		// 渠道类型,为了避免渠道码重复，可与渠道码共同确认渠道（可能为空。非空时当前只有满天星类型，值为STAR_ACTIVITY）
		AuthChannelType string `json:"auth_channel_type"`
		// 邀请码，只有自己邀请的企业才会返回邀请码，可用该邀请码统计不同渠道的拉新，否则值为空字符串
		InviteCode string `json:"invite_code"`
		// 企业是否认证
		IsAuthenticated bool `json:"is_authenticated"`
		// 序列号
		LicenseCode string `json:"license_code"`
		// 企业logo
		CorpLogoUrl string `json:"corp_logo_url"`
		// 授权企业所在省份
		CorpProvince string `json:"corp_province"`
		// 授权企业所在城市
		CorpCity string `json:"corp_city"`
	}
	// 授权的应用信息
	Agent struct {
		// 授权方应用id
		Agentid int64 `json:"agentid"`
		// 授权方应用名字
		AgentName string `json:"agent_name"`
		// 授权方应用头像
		LogoUrl string `json:"logo_url"`
		// 应用id
		Appid int64 `json:"appid"`
		// 对此微应用有管理权限的管理员userid
		AdminList []string `json:"admin_list"`
	}
	// 授权信息
	AuthInfo struct {
		Agent []*Agent `json:"agent"`
	}
	ChannelAgent struct {
		// 授权方应用名字
		AgentName string `json:"agent_name"`
		// 授权方应用id
		Agentid int64 `json:"agentid"`
		// 应用id
		Appid int64 `json:"appid"`
		// 授权方应用头像
		LogoUrl string `json:"logo_url"`
	}
	AuthMarketInfo struct {
	}
	// 授权的服务窗应用信息列表
	ChannelAuthInfo struct {
		ChannelAgent []*ChannelAgent `json:"channelAgent"`
	}
)
