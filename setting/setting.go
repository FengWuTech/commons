package setting

import (
	"fmt"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	LogLevel    string
	TimeFormat  string

	DivisionPrecision int

	CdnUrl     string
	QBoxBucket string
	QBoxAccess string
	QBoxSecret string

	MinipStaffAppID         string
	MinipStaffSecret        string
	MinipStaffOriID         string
	MinipStaffToken         string
	MinipStaffEncodedAESKey string

	MinipServiceProviderMchID  string
	MinipServiceProviderAppID  string
	MinipServiceProviderApiKey string

	NoticeList string
}

var AppSetting = &App{}

type Remote struct {
	MinipUserBaseUrl  string
	MinipStaffBaseUrl string
	PayCenterBaseUrl  string
	EsServiceUrl      string
	SentryDSN         string
}

var RemoteSetting = &Remote{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	PmsName     string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type DatabaseFlow struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseFlowSetting = &DatabaseFlow{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
	MachineryDB int
}

var RedisSetting = &Redis{}

type Sso struct {
	Host string
}

var SsoSetting = &Sso{}

type Cas struct {
	BaseUrl string
	AppID   string
	ApiKey  string
}

var CasSetting = &Cas{}

type PayCenter struct {
	BaseUrl string
	AppID   string
	ApiKey  string
}

var PayCenterSetting = &PayCenter{}

type Urm struct {
	AppID     string
	AppSecret string
}

var UrmSetting = &Urm{}

type Aliyun struct {
	RegionID     string
	AccessID     string
	AccessSecret string
	SliderAppKey string
}

var AliyunSetting = &Aliyun{}

type TencentCloud struct {
	SecretId            string
	SecretKey           string
	CaptchaAppId        uint64
	CaptchaAppSecretKey string
}

var TencentCloudSetting = &TencentCloud{}

type Machinery struct {
	ResultBackend   string
	Broker          string
	ResultsExpireIn int
}

var MachinerySetting = &Machinery{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup(env *string) {
	var (
		err      error
		cfgFiles []string
	)
	if *env == "dev" {
		cfgFiles = []string{"conf/dev.ini"}
	} else if *env == "test" {
		cfgFiles = []string{"conf/test.ini"}
	} else {
		panic("invalid env")
	}
	for _, cfgFile := range cfgFiles {
		cfg, err = ini.Load(cfgFile)
		if err != nil {
			fmt.Printf("setting.Setup, fail to parse '%s': %v\n", cfgFile, err)
		} else {
			fmt.Printf("setting.Setup, parse '%s' success\n", cfgFile)
			break
		}
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)
	mapTo("sso", SsoSetting)
	mapTo("database_flow", DatabaseFlowSetting)
	mapTo("remote", RemoteSetting)
	mapTo("cas", CasSetting)
	mapTo("urm", UrmSetting)
	mapTo("aliyun", AliyunSetting)
	mapTo("tencent_cloud", TencentCloudSetting)
	mapTo("pay_center", PayCenterSetting)
	mapTo("machinery", MachinerySetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		fmt.Printf("Cfg.MapTo %s err: %v", section, err)
	}
}
