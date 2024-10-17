package config

type DBConfig struct {
	Host                string `json:"host"                    validate:"required"`
	Port                int    `json:"port"                    validate:"required"`
	Name                string `json:"name"                    validate:"required"`
	User                string `json:"user"                    validate:"required"`
	Pass                string `json:"pass"                    validate:"required"`
	MaxIdleTimeInMinute int    `json:"max_idle_time_in_minute" validate:"required"`
	EnableSSLMode       bool   `json:"enable_ssl_mode"`
}

type Mode string

const DebugMode = Mode("debug")
const ReleaseMode = Mode("release")

type MongoDBConfig struct {
	Host                string `json:"host" `
	Port                int    `json:"port" `
	Database            string `json:"database" `
	User                string `json:"user" `
	Password            string `json:"password" `
	MaxIdleTimeInMinute int    `json:"max_idle_time_in_minute" `
	EnableSSL           bool   `json:"enable_ssl" `
}
type Config struct {
	Mode              Mode          `json:"mode"                       validate:"required"`
	ServiceName       string        `json:"service_name"               validate:"required"`
	HttpPort          int           `json:"http_port"                  validate:"required"`
	JwtSecret         string        `json:"jwt_secret"                 validate:"required"`
	DB                DBConfig      `json:"db"                         validate:"required"`
	MDB               MongoDBConfig `json:"mongodb"                    validate:"required"`
	MigrationSource   string        `json:"migrations"                 validate:"required"`
	HealthCheckRoute  string        `json:"HEALTH_CHECK_ROUTE"`
	RmqQueuePrefix    string        `json:"rmq_queue_prefix"           validate:"required"`
	ApiKeyEnabled     bool          `json:"API_KEY_ENABLED"`
	ApiKey            string        `json:"API_KEY"`
	RabbitmqURL       string        `json:"rmq_url"                    validate:"required"`
	RmqReconnectDelay int           `json:"rmq_reconnect_delay"        validate:"required"`
	RmqRetryInterval  int           `json:"rmq_retry_interval"         validate:"required"`
}

var config *Config

func init() {
	config = &Config{}
}

func GetConfig() Config {
	return *config
}
