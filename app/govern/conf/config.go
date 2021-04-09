package conf

import (
	"fmt"
	"github.com/sfshf/sprout/pkg/util/json"
	"github.com/spf13/viper"
	"os"
	"time"
)

var (
	C config

	EnvConfigFile      = "CONFIG_FILE"
	EnvCasbinModelFile = "CASBIN_MODEL_FILE"
	EnvMongoUri        = "MONGO_URI"
	EnvMongoDatabase   = "MONGO_DATABASE"
	EnvHttpHost        = "HTTP_HOST"
	EnvCertFile        = "HTTP_CERT_FILE"
	EnvCertKeyFile     = "HTTP_CERT_KEY_FILE"
)

func Init() {
	fpath := os.Getenv(EnvConfigFile)
	if fpath == "" {
		fpath = "app/govern/conf/config.toml"
	}
	viper.Reset()
	viper.SetConfigFile(fpath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&C); err != nil {
		panic(err)
	}
	if c := os.Getenv(EnvCasbinModelFile); c != "" {
		C.Casbin.Model = c
	}
	if c := os.Getenv(EnvMongoUri); c != "" {
		C.MongoDB.ServerUri = c
	}
	if c := os.Getenv(EnvMongoDatabase); c != "" {
		C.MongoDB.Database = c
	}
	if c := os.Getenv(EnvHttpHost); c != "" {
		C.HTTP.Host = c
	}
	if c := os.Getenv(EnvCertFile); c != "" {
		C.HTTP.CertFile = c
	}
	if c := os.Getenv(EnvCertKeyFile); c != "" {
		C.HTTP.CertKeyFile = c
	}
	if C.PrintConfig {
		fmt.Printf("%s\n", json.MarshalIndent2String(C))
	}
}

type config struct {
	RunMode     string
	WWW         string
	Swagger     bool
	PrintConfig bool
	Global      global
	Root        *root
	MongoDB     mongoDB
	JWTAuth     jwtAuth
	Casbin      casbin
	PicCaptcha  picCaptcha
	HTTP        http
	Log         log
	CORS        cors
	Redis       redis
}

type global struct {
	AppName        string
	AppIcon        string
	DateFormat     string
	DatetimeFormat string
	TimeZone       string
}

type root struct {
	Account   string
	Password  string
	SessionId string
}

type mongoDB struct {
	ServerUri string
	Database  string
}

type jwtAuth struct {
	Enable     bool
	SigningKey string
	Expired    int64
}

type casbin struct {
	Enable           bool
	Debug            bool
	Model            string
	AutoLoad         bool
	AutoLoadInternal int
}

type picCaptcha struct {
	Enable      bool
	Length      int
	Width       int
	Height      int
	MaxSkew     float64
	DotCount    int
	Threshold   int
	Expiration  time.Duration
	RedisStore  bool
	RedisDB     int
	RedisPrefix string
}

type http struct {
	Host             string
	Port             int
	CertFile         string
	CertKeyFile      string
	ShutdownTimeout  int
	MaxContentLength int
	MaxLoggerLength  int
}

type log struct {
	SkipStderr      bool
	Log2Mongo       bool
	MongoCollection string
	MaxWorkers      int
	MaxBuffers      int
}

type cors struct {
	Enable           bool
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	MaxAge           int
}

type redis struct {
	Addr     string
	Passwrod string
}
