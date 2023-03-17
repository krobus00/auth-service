package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Env :nodoc:
func Env() string {
	return viper.GetString("env")
}

// LogLevel :nodoc:
func LogLevel() string {
	return viper.GetString("log_level")
}

// GRPCPort :nodoc:
func GRPCport() string {
	return viper.GetString("ports.grpc")
}

// DatabaseDSN :nodoc:
func DatabaseDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		DatabaseUsername(),
		DatabasePassword(),
		DatabaseHost(),
		DatabaseName(),
		DatabaseSSLMode())
}

// DatabaseHost :nodoc:
func DatabaseHost() string {
	return viper.GetString("database.host")
}

// DatabaseName :nodoc:
func DatabaseName() string {
	return viper.GetString("database.database")
}

// DatabaseUsername :nodoc:
func DatabaseUsername() string {
	return viper.GetString("database.username")
}

// DatabasePassword :nodoc:
func DatabasePassword() string {
	return viper.GetString("database.password")
}

// DatabaseSSLMode :nodoc:
func DatabaseSSLMode() string {
	if viper.IsSet("database.sslmode") {
		return viper.GetString("database.sslmode")
	}
	return "disable"
}

// DatabasePingInterval :nodoc:
func DatabasePingInterval() time.Duration {
	if viper.GetInt("database.ping_interval") <= 0 {
		return DefaultDatabasePingInterval
	}
	return time.Duration(viper.GetInt("database.ping_interval")) * time.Millisecond
}

// DatabaseRetryAttempts :nodoc:
func DatabaseRetryAttempts() float64 {
	if viper.GetInt("database.retry_attempts") > 0 {
		return float64(viper.GetInt("database.retry_attempts"))
	}
	return DefaultDatabaseRetryAttempts
}

// DatabaseMaxIdleConns :nodoc:
func DatabaseMaxIdleConns() int {
	if viper.GetInt("database.max_idle_conns") <= 0 {
		return DefaultDatabaseMaxIdleConns
	}
	return viper.GetInt("database.max_idle_conns")
}

// DatabaseMaxOpenConns :nodoc:
func DatabaseMaxOpenConns() int {
	if viper.GetInt("database.max_open_conns") <= 0 {
		return DefaultDatabaseMaxOpenConns
	}
	return viper.GetInt("database.max_open_conns")
}

// DatabaseConnMaxLifetime :nodoc:
func DatabaseConnMaxLifetime() time.Duration {
	if !viper.IsSet("database.conn_max_lifetime") {
		return DefaultDatabaseConnMaxLifetime
	}
	return time.Duration(viper.GetInt("database.conn_max_lifetime")) * time.Millisecond
}

// DisableCaching :nodoc:
func DisableCaching() bool {
	return viper.GetBool("redis.disable_caching")
}

// RedisHost :nodoc:
func RedisCacheHost() string {
	return viper.GetString("redis.cache_host")
}

// RedisDialTimeout :nodoc:
func RedisDialTimeout() time.Duration {
	cfg := viper.GetString("redis.dial_timeout")
	return parseDuration(cfg, 5*time.Second)
}

// RedisWriteTimeout :nodoc:
func RedisWriteTimeout() time.Duration {
	cfg := viper.GetString("redis.write_timeout")
	return parseDuration(cfg, 2*time.Second)
}

// RedisReadTimeout :nodoc:
func RedisReadTimeout() time.Duration {
	cfg := viper.GetString("redis.read_timeout")
	return parseDuration(cfg, 2*time.Second)
}

// RedisCacheTTL :nodoc:
func RedisCacheTTL() time.Duration {
	cfg := viper.GetString("cache_ttl")
	return parseDuration(cfg, DefaultRedisCacheTTL)
}

// TokenSecret :nodoc:
func TokenSecret() string {
	return viper.GetString("jwt.secret_key")
}

// AccessTokenDuration :nodoc:
func AccessTokenDuration() time.Duration {
	cfg := viper.GetString("jwt.access_token_duration")
	return parseDuration(cfg, 15*time.Minute)
}

// RefreshTokenDuration :nodoc:
func RefreshTokenDuration() time.Duration {
	cfg := viper.GetString("jwt.refresh_token_duration")
	return parseDuration(cfg, 24*time.Hour)
}

// BcryptCost :nodoc:
func BcryptCost() int {
	if viper.GetInt("bcrypt.cost") > 4 && viper.GetInt("bcrypt.cost") < 31 {
		return viper.GetInt("redis.bcrypt.cost")
	}
	return 10
}

// BcryptSalt :nodoc:
func BcryptSalt() string {
	return viper.GetString("bcrypt.salt")
}

func parseDuration(in string, defaultDuration time.Duration) time.Duration {
	dur, err := time.ParseDuration(in)
	if err != nil {
		return defaultDuration
	}
	return dur
}

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("config not found")
		}
		return err
	}
	return nil
}
