package mimecast

import (
	"net/http"
	"os"
)

var MimeCastGlobalConfig *MimeCastConfig
var MimeCastHttpClient *http.Client

func init() {
	MimeCastHttpClient = &http.Client{}
}

func InitMimeCastConfig() {
	MimeCastGlobalConfig = &MimeCastConfig{
		ApplicationId:  os.Getenv("MIMECAST_APPLICATION_ID"),
		ApplicationKey: os.Getenv("MIMECAST_APPLICATION_KEY"),
		AccessKey:      os.Getenv("MIMECAST_ACCESS_KEY"),
		SecretKey:      os.Getenv("MIMECAST_SECRET_KEY"),
	}
}

func SetMimeCastConfig(c MimeCastConfig) {
	MimeCastGlobalConfig = &c
}

type ApiEndpoint interface {
	Url() string
	RequestData() M
	RequestMeta() M
}

type MimeCastConfig struct {
	ApplicationId  string
	ApplicationKey string
	AccessKey      string
	SecretKey      string
}

type ResponseMetadata struct {
	RateLimit          int
	RateLimitReset     int
	RateLimitRemaining int
	NextToken          string
}

type Response struct {
	Status      int
	IsLastToken bool
	Data        []M
	MetaData    ResponseMetadata
}
