package sessiontypes

type AmazonConfigurationInput struct {
	RoleArn *string
	Region  *string
}

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

type SessionIdEncryptionOut struct {
	AssociatedData []byte
	CipherText     *string
	IV             []byte
}
