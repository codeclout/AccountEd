package membertypes

type CredentialsAWS struct {
	Value `json:"Value"`
}

type Value struct {
	AccessKeyID     string `json:"AccessKeyID"`
	SecretAccessKey string `json:"SecretAccessKey"`
	SessionToken    string `json:"SessionToken"`
	Source          string `json:"Source"`
	CanExpire       bool   `json:"CanExpire"`
	Expires         string `json:"Expires"`
}
