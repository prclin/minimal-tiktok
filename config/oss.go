package config

type OSS struct {
	EndPoint        string
	AccessKeyId     string
	AccessKeySecret string
}

var DefaultOSS = &OSS{EndPoint: "", AccessKeyId: "", AccessKeySecret: ""}
