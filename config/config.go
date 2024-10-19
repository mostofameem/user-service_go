package config

import "sync"

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

type GrpcRetryPolicyConfig struct {
	MaxAttempts          int      `json:"max_attempts"               validate:"required"`
	InitialBackoff       string   `json:"initial_backoff_in_seconds" validate:"required"`
	MaxBackoff           string   `json:"max_backoff_in_seconds"     validate:"required"`
	BackoffMultiplier    float64  `json:"backoff_multiplier"         validate:"required"`
	RetryableStatusCodes []string `json:"retryable_status_codes"     validate:"required"`
}
type GrpcUrlsConfig struct {
	User  string `json:"user"              validate:"required"`
	Posts string `json:"posts"        validate:"required"`
}
type Config struct {
	Mode                    Mode                  `json:"mode"                       validate:"required"`
	ServiceName             string                `json:"service_name"               validate:"required"`
	HttpPort                int                   `json:"http_port"                  validate:"required"`
	GrpcPort                int                   `json:"grpc_port"                        validate:"required"`
	JwtSecret               string                `json:"jwt_secret"                 validate:"required"`
	DB                      DBConfig              `json:"db"                         validate:"required"`
	MDB                     MongoDBConfig         `json:"mongodb"                    validate:"required"`
	MigrationSource         string                `json:"migrations"                 validate:"required"`
	HealthCheckRoute        string                `json:"HEALTH_CHECK_ROUTE"`
	RmqQueuePrefix          string                `json:"rmq_queue_prefix"           validate:"required"`
	ApiKeyEnabled           bool                  `json:"API_KEY_ENABLED"`
	ApiKey                  string                `json:"API_KEY"`
	RabbitmqURL             string                `json:"rmq_url"                    validate:"required"`
	RmqReconnectDelay       int                   `json:"rmq_reconnect_delay"        validate:"required"`
	RmqRetryInterval        int                   `json:"rmq_retry_interval"         validate:"required"`
	GrpcUrls                GrpcUrlsConfig        `json:"grpc_urls"                             validate:"required"`
	GrpcRetryPolicy         GrpcRetryPolicyConfig `json:"grpc_retry_policy"                             validate:"required"`
	RedisURL                string                `json:"redis_url"                  validate:"required"`
	RedisSearchPrefix       string                `json:"redis_search_prefix"        validate:"required"`
	BranchIdRetentionPeriod int                   `json:"branch_id_retention_period" validate:"required"`
}

var config *Config
var cnfOnce = sync.Once{}

func GetConfig() *Config {
	cnfOnce.Do(func() {
		loadConfig()
	})

	return config
}
