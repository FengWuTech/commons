package setting

import (
	"bytes"
	"fmt"
	"github.com/FengWuTech/commons/consul"
	"github.com/FengWuTech/commons/util/fileutil"
	"github.com/hashicorp/consul/api/watch"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
	"syscall"
	"time"

	"github.com/go-ini/ini"
)

type Config struct {
	Name     string `yaml:"name"`
	Registry struct {
		Consul struct {
			URL   string `yaml:"url"`
			Token string `yaml:"token"`
		}
	} `yaml:"registry"`
	Data struct {
		Dir string `yaml:"dir"`
	} `yaml:"data"`
}

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

	LogSavePath    string `yaml:"LogSavePath"`
	LogSaveName    string `yaml:"LogSaveName"`
	LogFileExt     string `yaml:"LogFileExt"`
	LogLevel       string `yaml:"LogLevel"`
	TimeFormat     string `yaml:"TimeFormat"`
	LogPushToRedis bool   `yaml:"LogPushToRedis"`

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

	//MinipServiceProviderMchID  string `yaml:"MinipServiceProviderMchID"`
	//MinipServiceProviderAppID  string `yaml:"MinipServiceProviderAppID"`
	//MinipServiceProviderApiKey string `yaml:"MinipServiceProviderApiKey"`

	NoticeList string `yaml:"NoticeList"`

	DispatchOrderMaxRemindTimes int `yaml:"DispatchOrderMaxRemindTimes"`
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

type Email struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
}

var EmailSetting = &Email{}

var MachinerySetting = &Machinery{}

var cfg *ini.File

func Setup(env *string) {
	var (
		iniFile  string
		yamlFile string
	)

	switch *env {
	case "dev":
		iniFile = "conf/dev.ini"
		yamlFile = "conf/dev.yaml"
	case "test":
		iniFile = "conf/test.ini"
		yamlFile = "conf/test.yaml"
	case "prod":
		iniFile = "conf/prod.ini"
		yamlFile = "conf/prod.yaml"
	default:
		panic("invalid env")
	}

	//试图从yaml加载配置文件
	cnt, err := ioutil.ReadFile(yamlFile)
	if err != nil || len(cnt) == 0 {
		fmt.Printf("yaml文件不存在，现从ini开始加载\n")
		setupFromIni(iniFile)
	} else {
		fmt.Printf("yaml文件存在，开始从yaml加载配置\n")
		setupFromRegistry(yamlFile)
	}
}

//从ini加载配置文件
func setupFromIni(cfgFile string) {
	cfgHandler, err := ini.Load(cfgFile)
	if err != nil {
		fmt.Printf("setting.Setup, fail to parse '%s': %v\n", cfgFile, err)
	} else {
		fmt.Printf("setting.Setup, parse '%s' success\n", cfgFile)
	}
	cfg = cfgHandler

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
	mapTo("email", EmailSetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		fmt.Printf("Cfg.MapTo %s err: %v\n", section, err)
	}
}

// Setup initialize the configuration instance
func setupFromRegistry(yamlFile string) {
	cnt, err := ioutil.ReadFile(yamlFile)
	if err != nil || len(cnt) == 0 {
		panic(fmt.Sprintf("配置文件加载失败: %v\n", err))
	}

	var config Config
	yaml.Unmarshal(cnt, &config)

	if strings.Index(config.Data.Dir, "~/") == 0 {
		cur, err := user.Current()
		if err != nil {
			panic(fmt.Sprintf("无法确定用户目录: %v", err))
		}
		config.Data.Dir = strings.ReplaceAll(config.Data.Dir, "~", cur.HomeDir)
	}

	if !fileutil.Exists(config.Data.Dir) {
		os.Mkdir(config.Data.Dir, 0777)
	}

	consul.Setup(config.Registry.Consul.URL, config.Registry.Consul.Token)

	var ret map[string]interface{}
	cfgFile := fmt.Sprintf("%s/%s.conf", config.Data.Dir, config.Name)
	kv, _, err := consul.Client().KV().Get(config.Name, nil)
	if err != nil {
		cnt, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			panic(fmt.Sprintf("读取本地文件失败 %v", err))
		}
		err = yaml.NewDecoder(bytes.NewReader(cnt)).Decode(&ret)
		if err != nil || ret == nil {
			panic(fmt.Sprintf("解析本地配置失败: %v\n", err))
		}
		fmt.Printf("配置中心无法连接，使用本地配置执行 %v\n", err)
	} else {
		err = yaml.NewDecoder(bytes.NewReader(kv.Value)).Decode(&ret)
		if err != nil || ret == nil {
			panic(fmt.Sprintf("解析注册中心配置失败: %v\n", err))
		}
		fmt.Printf("使用注册中心配置执行\n")
		ioutil.WriteFile(cfgFile, kv.Value, 0777)
	}

	fun := func(key string, out interface{}) {
		if value, ok := ret[key]; ok {
			outData, err := yaml.Marshal(value)
			if err != nil {
				panic(fmt.Sprintf("配置Marshal异常: %v\n", err))
			}
			err = yaml.Unmarshal(outData, out)
			if err != nil {
				panic(fmt.Sprintf("配置Unmarshal异常: %v\n", err))
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
	fun("email", EmailSetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second

	go WatchConfigChange(config)
}

func WatchConfigChange(config Config) {
	recvNum := 0
	for {
		plan, err := watch.Parse(map[string]interface{}{
			"type": "key",
			"key":  config.Name,
		})
		if err != nil || plan == nil {
			fmt.Printf("consul无法监听任务，将在10分钟后重试\n")
			return
		} else {
			plan.Handler = func(u uint64, raw interface{}) {
				if recvNum > 0 {
					fmt.Printf("程序启动后接收到配置更新消息，开始重启\n")
					syscall.Kill(syscall.Getpid(), syscall.SIGUSR2)
				} else {
					fmt.Printf("程序启动接受到配置推送，无需重启\n")
				}
				recvNum++
			}
			plan.Token = config.Registry.Consul.Token
			fmt.Printf("开始监听配置修改......\n")
			err = plan.Run(config.Registry.Consul.URL)
			fmt.Printf("监听配置修改失败，将在10分钟后重试: %v\n", err)
		}
		time.Sleep(time.Minute * 10)
	}
}
