package config

import "galaveg/utils/path"

type Config struct {
	App struct {
		Key               string   `mapstructure:"KEY"`
		PreviousKeys      []string `mapstructure:"PREVIOUS_KEYS"`
		Url               string   `mapstructure:"URL"`
		Debug             bool     `mapstructure:"DEBUG"`
		Timezone          string   `mapstructure:"TIMEZONE"`
		LogLevel          string   `mapstructure:"LOG_LEVEL"` // Example: panic, fatal, error, warn, info, debug, trace
		Folder            string   `mapstructure:"FOLDER"`
		Locale            string   `mapstructure:"LOCALE"`
		LocaleCookieKey   string   `mapstructure:"LOCALE_COOKIE_KEY"`
		DarkModeCookieKey string   `mapstructure:"DARK_MODE_COOKIE_KEY"`
	} `mapstructure:"APP"`
	Db struct {
		Prefix          string `mapstructure:"PREFIX"` // Example: glg_
		Host            string `mapstructure:"HOST"`
		Port            int    `mapstructure:"PORT"`
		Database        string `mapstructure:"DATABASE"`
		Username        string `mapstructure:"USERNAME"`
		Password        string `mapstructure:"PASSWORD"`
		Tls             string `mapstructure:"TLS"`     // Example: true, false, skip-verify, preferred, <name>
		Socket          string `mapstructure:"SOCKET"`  // Example: /var/run/mysqld/mysqld.sock
		Charset         string `mapstructure:"CHARSET"` // Example: utf8mb4
		MaxOpenConns    int    `mapstructure:"MAX_OPEN_CONNS"`
		MaxIdleConns    int    `mapstructure:"MAX_IDLE_CONNS"`
		ConnMaxLifetime int64  `mapstructure:"CONN_MAX_LIFETIME"`  // Example: int64(time.Hour)
		ConnMaxIdleTime int64  `mapstructure:"CONN_MAX_IDLE_TIME"` // Example: int64(5*time.Minute))
	} `mapstructure:"Db"`
	Http struct {
		Host         string   `mapstructure:"HOST"`
		Port         int      `mapstructure:"PORT"`
		Schema       string   `mapstructure:"SCHEMA"`        // Example: https, http
		AllowedHosts []string `mapstructure:"ALLOWED_HOSTS"` // Example: localhost,0.0.0.0,example.com
	} `mapstructure:"HTTP"`
	Mail struct {
		Host        string `mapstructure:"HOST"`
		Port        int    `mapstructure:"PORT"`
		Encryption  string `mapstructure:"ENCRYPTION"`
		Username    string `mapstructure:"USERNAME"`
		Password    string `mapstructure:"PASSWORD"`
		FromAddress string `mapstructure:"FROM_ADDRESS"`
		FromName    string `mapstructure:"FROM_NAME"`
	} `mapstructure:"MAIL"`
	Session struct {
		StoreUserKey            string `mapstructure:"STORE_USER_KEY"`
		CookieKey               string `mapstructure:"COOKIE_KEY"`
		RedisHost               string `mapstructure:"REDIS_HOST"`
		RedisPort               int    `mapstructure:"REDIS_PORT"`
		RedisUsername           string `mapstructure:"REDIS_USERNAME"`
		RedisPassword           string `mapstructure:"REDIS_PASSWORD"`
		RedisMaxIdleConnections int    `mapstructure:"REDIS_MAX_IDLE_CONNECTIONS"`
	} `mapstructure:"SESSION"`
}

func beforeReturn(c *Config) {
	c.App.Folder = strDefault(c.App.Folder, path.FindModuleRoot(path.Cwd()))
	c.App.PreviousKeys = strSliceFilter(c.App.PreviousKeys)
	c.Http.AllowedHosts = strSliceFilter(c.Http.AllowedHosts)
}
