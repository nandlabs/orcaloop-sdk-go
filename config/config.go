package config

type ActionSvcConfig struct {
	Name     string          `json:"name" yaml:"name"`
	Listener *ListenerConfig `json:"listener" yaml:"listener"`
}

type ListenerConfig struct {
	ListenHost     string `json:"listen_host" yaml:"listen_host" bson:"listen_host" mapstructure:"listen_host"`
	ListenPort     int16  `json:"listen_port" yaml:"listen_port" bson:"listen_port" mapstructure:"listen_port"`
	ReadTimeout    int64  `json:"read_timeout,omitempty" yaml:"read_timeout,omitempty" bson:"read_timeout,omitempty" mapstructure:"read_timeout,omitempty"`
	WriteTimeout   int64  `json:"write_timeout,omitempty" yaml:"write_timeout,omitempty" bson:"write_timeout,omitempty" mapstructure:"write_timeout,omitempty"`
	EnableTLS      bool   `json:"enable_tls" yaml:"enable_tls" bson:"enable_tls" mapstructure:"enable_tls"`
	PrivateKeyPath string `json:"private_key_path,omitempty" yaml:"private_key_path,omitempty" bson:"private_key_path,omitempty" mapstructure:"private_key,omitempty"`
	CertPath       string `json:"cert_path,omitempty" yaml:"cert_path,omitempty" bson:"cert_path,omitempty" mapstructure:"cert,omitempty"`
}
