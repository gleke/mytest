// 设备信息表
type DmDeviceInfo struct {
	ID             int64             `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	ProjectID      int64             `gorm:"column:project_id;type:bigint;default:0;NOT NULL"`            // 项目ID(雪花ID)
	AreaID         int64             `gorm:"column:area_id;type:bigint;default:0;NOT NULL"`               // 项目区域ID(雪花ID)
	ProductID      string            `gorm:"column:product_id;type:char(11);NOT NULL"`                    // 产品id
	DeviceName     string            `gorm:"column:device_name;type:varchar(100);NOT NULL"`               // 设备名称
	DeviceAlias    string            `gorm:"column:device_alias;type:varchar(100);NOT NULL"`              // 设备别名
	Position       stores.Point      `gorm:"column:position;type:point;NOT NULL"`                         // 设备的位置(默认百度坐标系BD09)
	Secret         string            `gorm:"column:secret;type:varchar(50);NOT NULL"`                     // 设备秘钥
	Cert           string            `gorm:"column:cert;type:varchar(512);NOT NULL"`                      // 设备证书
	Imei           string            `gorm:"column:imei;type:varchar(15);NOT NULL"`                       // IMEI号信息
	Mac            string            `gorm:"column:mac;type:varchar(17);NOT NULL"`                        // MAC号信息
	Version        string            `gorm:"column:version;type:varchar(64);NOT NULL"`                    // 固件版本
	HardInfo       string            `gorm:"column:hard_info;type:varchar(64);NOT NULL"`                  // 模组硬件型号
	SoftInfo       string            `gorm:"column:soft_info;type:varchar(64);NOT NULL"`                  // 模组软件版本
	MobileOperator int64             `gorm:"column:mobile_operator;type:smallint;default:1;NOT NULL"`     // 移动运营商:1)移动 2)联通 3)电信 4)广电
	Phone          sql.NullString    `gorm:"column:phone;type:varchar(20)"`                               // 手机号
	Iccid          sql.NullString    `gorm:"column:iccid;type:varchar(20)"`                               // SIM卡卡号
	Address        string            `gorm:"column:address;type:varchar(512);NOT NULL"`                   // 所在地址
	Tags           map[string]string `gorm:"column:tags;type:json;serializer:json;NOT NULL;default:'{}'"` // 设备标签
	IsOnline       int64             `gorm:"column:is_online;type:smallint;default:2;NOT NULL"`           // 是否在线,1是2否
	FirstLogin     sql.NullTime      `gorm:"column:first_login;type:TIMESTAMP WITH TIME ZONE"`            // 激活时间
	LastLogin      sql.NullTime      `gorm:"column:last_login;type:TIMESTAMP WITH TIME ZONE"`             // 最后上线时间
	LogLevel       int64             `gorm:"column:log_level;type:smallint;default:1;NOT NULL"`           // 日志级别:1)关闭 2)错误 3)告警 4)信息 5)调试
	stores.Time
	ProductInfo *DmProductInfo `gorm:"foreignKey:ProductID;references:ProductID"` // 添加外键

}

// 产品信息表
type DmProductInfo struct {
	ProductID    string            `gorm:"column:product_id;type:char(11);primary_key;NOT NULL"`        // 产品id
	ProductName  string            `gorm:"column:product_name;type:varchar(100);NOT NULL"`              // 产品名称
	ProductImg   string            `gorm:"column:product_img;type:varchar(200)"`                        // 产品图片
	ProductType  int64             `gorm:"column:product_type;type:smallint;default:1"`                 // 产品状态:1:开发中,2:审核中,3:已发布
	AuthMode     int64             `gorm:"column:auth_mode;type:smallint;default:1"`                    // 认证方式:1:账密认证,2:秘钥认证
	DeviceType   int64             `gorm:"column:device_type;type:smallint;default:1"`                  // 设备类型:1:设备,2:网关,3:子设备
	CategoryID   int64             `gorm:"column:category_id;type:integer;default:1"`                   // 产品品类
	NetType      int64             `gorm:"column:net_type;type:smallint;default:1"`                     // 通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN
	DataProto    int64             `gorm:"column:data_proto;type:smallint;default:1"`                   // 数据协议:1:自定义,2:数据模板
	AutoRegister int64             `gorm:"column:auto_register;type:smallint;default:1"`                // 动态注册:1:关闭,2:打开,3:打开并自动创建设备
	Secret       string            `gorm:"column:secret;type:varchar(50)"`                              // 动态注册产品秘钥
	Desc         string            `gorm:"column:description;type:varchar(200)"`                        // 描述
	DevStatus    string            `gorm:"column:dev_status;type:varchar(20);NOT NULL"`                 // 产品状态
	Tags         map[string]string `gorm:"column:tags;type:json;serializer:json;NOT NULL;default:'{}'"` // 产品标签

	stores.Time
	//Devices []*DmDeviceInfo    `gorm:"foreignKey:ProductID;references:ProductID"` // 添加外键
}
