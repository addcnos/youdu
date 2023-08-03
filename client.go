package youdu

type Config struct {
	Addr   string
	Buin   int
	AppId  string
	AesKey string
}

type Client struct {
	config    *Config
	encryptor *Encryptor
	token     *token
}

func NewClient(config *Config) *Client {
	return &Client{
		config:    config,
		encryptor: NewEncryptorWithConfig(config),
	}
}
