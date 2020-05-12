package setting

import (
	"bytes"
	"fmt"
	"github.com/FengWuTech/commons/consul"
	"gopkg.in/yaml.v2"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	PageSize        int    `yaml:"PageSize"`
	PrefixUrl       string `yaml:"PrefixUrl"`
	RuntimeRootPath string `yaml:"RuntimeRootPath"`

	ImageSavePath  string   `yaml:"ImageSavePath"`
	ImageMaxSize   int      `yaml:"ImageMaxSize"`
	ImageAllowExts []string `yaml:"ImageAllowExts"`

	ExportSavePath string `yaml:"ExportSavePath"`
	QrCodeSavePath string `yaml:"QrCodeSavePath"`
	FontSavePath   string `yaml:"FontSavePath"`

	LogSavePath string `yaml:"LogSavePath"`
	LogSaveName string `yaml:"LogSaveName"`
	LogFileExt  string `yaml:"LogFileExt"`
	LogLevel    string `yaml:"LogLevel"`
	TimeFormat  string `yaml:"TimeFormat"`

	DivisionPrecision int `yaml:"DivisionPrecision"`

	CdnUrl     string `yaml:"CdnUrl"`
	QBoxBucket string `yaml:"QBoxBucket"`
	QBoxAccess string `yaml:"QBoxAccess"`
	QBoxSecret string `yaml:"QBoxSecret"`

	MinipStaffName                    string `yaml:"MinipStaffName"`
	MinipStaffQrCodeURL               string `yaml:"MinipStaffQrCodeURL"`
	MinipStaffAppID                   string `yaml:"MinipStaffAppID"`
	MinipStaffSecret                  string `yaml:"MinipStaffSecret"`
	MinipStaffOriID                   string `yaml:"MinipStaffOriID"`
	MinipStaffToken                   string `yaml:"MinipStaffToken"`
	MinipStaffEncodedAESKey           string `yaml:"MinipStaffEncodedAESKey"`
	MinipStaffTemplateMsgManageNotice string `yaml:"MinipStaffTemplateMsgManageNotice"`
	MinipStaffTemplateMsgOrderTimeout string `yaml:"MinipStaffTemplateMsgOrderTimeout"`
	MinipStaffTemplateMsgOrderRemind  string `yaml:"MinipStaffTemplateMsgOrderRemind"`
	MinipStaffTemplateMsgOrderFinish  string `yaml:"MinipStaffTemplateMsgOrderFinish"`
	MinipStaffTemplateMsgOrderNew     string `yaml:"MinipStaffTemplateMsgOrderNew"`

	MinipServiceProviderMchID  string `yaml:"MinipServiceProviderMchID"`
	MinipServiceProviderAppID  string `yaml:"MinipServiceProviderAppID"`
	MinipServiceProviderApiKey string `yaml:"MinipServiceProviderApiKey"`

	NoticeList string `yaml:"NoticeList"`
}

var AppSetting = &App{}

type Remote struct {
	MinipUserBaseUrl  string `yaml:"MinipUserBaseUrl"`
	MinipStaffBaseUrl string `yaml:"MinipStaffBaseUrl"`
	PayCenterBaseUrl  string `yaml:"PayCenterBaseUrl"`
	EsServiceUrl      string `yaml:"EsServiceUrl"`
	SentryDSN         string `yaml:"SentryDSN"`
}

var RemoteSetting = &Remote{}

type Server struct {
	RunMode      string        `yaml:"RunMode"`
	HttpPort     int           `yaml:"HttpPort"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
}

var ServerSetting = &Server{}

type Database struct {
	Type        string `yaml:"Type"`
	User        string `yaml:"User"`
	Password    string `yaml:"Password"`
	Host        string `yaml:"Host"`
	Name        string `yaml:"Name"`
	PmsName     string `yaml:"PmsName"`
	TablePrefix string `yaml:"TablePrefix"`
}

var DatabaseSetting = &Database{}

type DatabaseFlow struct {
	Type        string `yaml:"Type"`
	User        string `yaml:"User"`
	Password    string `yaml:"Password"`
	Host        string `yaml:"Host"`
	Name        string `yaml:"Name"`
	TablePrefix string `yaml:"TablePrefix"`
}

var DatabaseFlowSetting = &DatabaseFlow{}

type Redis struct {
	Host        string        `yaml:"Host"`
	Password    string        `yaml:"Password"`
	MaxIdle     int           `yaml:"MaxIdle"`
	MaxActive   int           `yaml:"MaxActive"`
	IdleTimeout time.Duration `yaml:"IdleTimeout"`
	MachineryDB int           `yaml:"MachineryDB"`
}

var RedisSetting = &Redis{}

type Sso struct {
	Host string `yaml:"Host"`
}

var SsoSetting = &Sso{}

type Cas struct {
	BaseUrl string `yaml:"BaseUrl"`
	AppID   string `yaml:"AppID"`
	ApiKey  string `yaml:"ApiKey"`
}

var CasSetting = &Cas{}

type PayCenter struct {
	BaseUrl string `yaml:"BaseUrl"`
	AppID   string `yaml:"AppID"`
	ApiKey  string `yaml:"ApiKey"`
}

var PayCenterSetting = &PayCenter{}

type Urm struct {
	URL       string `yaml:"URL"`
	AppID     string `yaml:"AppID"`
	AppSecret string `yaml:"AppSecret"`
}

var UrmSetting = &Urm{}

type Aliyun struct {
	RegionID     string `yaml:"RegionID"`
	AccessID     string `yaml:"AccessID"`
	AccessSecret string `yaml:"AccessSecret"`
	SliderAppKey string `yaml:"SliderAppKey"`
}

var AliyunSetting = &Aliyun{}

type TencentCloud struct {
	// 帐户层面
	AppId     string `yaml:"AppId"`
	SecretId  string `yaml:"SecretId"`
	SecretKey string `yaml:"SecretKey"`

	// 验证码
	CaptchaAppId        uint64 `yaml:"CaptchaAppId"`
	CaptchaAppSecretKey string `yaml:"CaptchaAppSecretKey"`
	// 对象存储cos
	CosBucket string `yaml:"CosBucket"`
	CosRegion string `yaml:"CosRegion"`
	//短信
	SmsAppID  string `yaml:"SmsAppID"`
	SmsAppKey string `yaml:"SmsAppKey"`
	SmsSignID string `yaml:"SmsSignID"`
}

var TencentCloudSetting = &TencentCloud{}

type Machinery struct {
	ResultBackend   string `yaml:"ResultBackend"`
	Broker          string `yaml:"Broker"`
	ResultsExpireIn int    `yaml:"ResultsExpireIn"`
}

var MachinerySetting = &Machinery{}

var cfg *ini.File

func Setup(env *string) {
	var (
		err      error
		cfgFiles []string
	)
	if *env == "dev" {
		cfgFiles = []string{"conf/dev.ini"}
	} else if *env == "test" {
		cfgFiles = []string{"conf/test.ini"}
	} else if *env == "prod" {
		cfgFiles = []string{"conf/prod.ini"}
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

// Setup initialize the configuration instance
func SetupFromRegistry(rc string, plt string) {
	consul.Setup(rc)

	kv, _, err := consul.Client().KV().Get(plt, nil)
	if err != nil {
		panic("读取配置文件失败")
	}

	var ret map[string]interface{}
	err = yaml.NewDecoder(bytes.NewReader(kv.Value)).Decode(&ret)
	if err != nil || ret == nil {
		panic("配置解析失败")
	}

	fun := func(key string, out interface{}) {
		if value, ok := ret[key]; ok {
			outData, err := yaml.Marshal(value)
			if err != nil {
				panic(fmt.Sprintf("配置Marshal异常: %v", err))
			}
			err = yaml.Unmarshal(outData, out)
			if err != nil {
				panic(fmt.Sprintf("配置Unmarshal异常: %v", err))
			}
		}
	}
	fun("app", AppSetting)
	fun("server", ServerSetting)
	fun("database", DatabaseSetting)
	fun("redis", RedisSetting)
	fun("sso", SsoSetting)
	fun("database_flow", DatabaseFlowSetting)
	fun("remote", RemoteSetting)
	fun("cas", CasSetting)
	fun("urm", UrmSetting)
	fun("aliyun", AliyunSetting)
	fun("tencent_cloud", TencentCloudSetting)
	fun("pay_center", PayCenterSetting)
	fun("machinery", MachinerySetting)

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
