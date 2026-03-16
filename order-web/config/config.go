package config

type OrderSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type GoodsSrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type InventorySrvConfig struct {
	Name string `mapstructure:"name" json:"name"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host" json:"host"`
	Port      uint64 `mapstructure:"port" json:"port"`
	Namespace string `mapstructure:"namespace" json:"namespace"`
	Group     string `mapstructure:"group" json:"group"`
	Dataid    string `mapstructure:"dataid" json:"dataid"`
}

type AlipayConfig struct {
	AppID        string `mapstructure:"appid" json:"appid"`
	PrivateKey   string `mapstructure:"private_key" json:"private_key"`
	AliPublicKey string `mapstructure:"ali_public_key" json:"ali_public_key"`
	NotifyURL    string `mapstructure:"notify_url" json:"notify_url"`
	ReturnURL    string `mapstructure:"return_url" json:"return_url"`
}

type ServerConfig struct {
	Name             string             `mapstructure:"name" json:"name"`
	Host             string             `mapstructure:"host" json:"host"`
	Port             int                `mapstructure:"port" json:"port"`
	Tags             []string           `mapstructure:"tags" json:"tags"`
	OrderSrvConf     OrderSrvConfig     `mapstructure:"order_srv" json:"order_srv"`
	GoodsSrvConf     GoodsSrvConfig     `mapstructure:"goods_srv" json:"goods_srv"`
	InventorySrvConf InventorySrvConfig `mapstructure:"inventory_srv" json:"inventory_srv"`
	JWTInfo          JWTConfig          `mapstructure:"jwt" json:"jwt"`
	ConsulInfo       ConsulConfig       `mapstructure:"consul" json:"consul"`
	AlipayInfo       AlipayConfig       `mapstructure:"alipay" json:"alipay"`
}
