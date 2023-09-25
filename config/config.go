package config

type MongoCategory struct {
	Host       string `env:"MONGO_CATEGORY_HOST" envDefault:"localhost"`
	Port       string `env:"MONGO_CATEGORY_PORT" envDefault:"27017"`
	Username   string `env:"MONGO_CATEGORY_USERNAME" envDefault:""`
	Password   string `env:"MONGO_CATEGORY_PASSWORD" envDefault:""`
	Database   string `env:"MONGO_CATEGORY_DATABASE" envDefault:"account"`
	Collection string `env:"MONGO_CATEGORY_COLLECTION" envDefault:"accounts"`
	Query      string `env:"MONGO_CATEGORY_QUERY" envDefault:""`
}

type I18n struct {
	Fallback string   `env:"I18N_FALLBACK_LANGUAGE" envDefault:"en"`
	Dir      string   `env:"I18N_DIR" envDefault:"./src/locales"`
	Locales  []string `env:"I18N_LOCALES" envDefault:"en,tr"`
}

type Server struct {
	Host  string `env:"SERVER_HOST" envDefault:"localhost"`
	Port  int    `env:"SERVER_PORT" envDefault:"3000"`
	Group string `env:"SERVER_GROUP" envDefault:"account"`
}

type Redis struct {
	Host string `env:"REDIS_HOST"`
	Port string `env:"REDIS_PORT"`
	Pw   string `env:"REDIS_PASSWORD"`
	Db   int    `env:"REDIS_DB"`
}

type CacheRedis struct {
	Host string `env:"REDIS_CACHE_HOST"`
	Port string `env:"REDIS_CACHE_PORT"`
	Pw   string `env:"REDIS_CACHE_PASSWORD"`
	Db   int    `env:"REDIS_CACHE_DB"`
}

type HttpHeaders struct {
	AllowedOrigins   string `env:"CORS_ALLOWED_ORIGINS" envDefault:"*"`
	AllowedMethods   string `env:"CORS_ALLOWED_METHODS" envDefault:"GET,POST,PUT,DELETE,OPTIONS"`
	AllowedHeaders   string `env:"CORS_ALLOWED_HEADERS" envDefault:"*"`
	AllowCredentials bool   `env:"CORS_ALLOW_CREDENTIALS" envDefault:"true"`
	Domain           string `env:"HTTP_HEADER_DOMAIN" envDefault:"*"`
}

type TokenSrv struct {
	Expiration int    `env:"TOKEN_EXPIRATION" envDefault:"3600"`
	Project    string `env:"TOKEN_PROJECT" envDefault:"empty"`
}

type Session struct {
	Topic string `env:"SESSION_TOPIC"`
}

type Topics struct {
	Category CategoryEvents
	Post     PostEvents
}

type CategoryEvents struct {
	Created               string `env:"STREAMING_TOPIC_CATEGORY_CREATED"`
	Updated               string `env:"STREAMING_TOPIC_CATEGORY_UPDATED"`
	Enabled               string `env:"STREAMING_TOPIC_CATEGORY_ENABLED"`
	Disabled              string `env:"STREAMING_TOPIC_CATEGORY_DISABLED"`
	Deleted               string `env:"STREAMING_TOPIC_CATEGORY_DELETED"`
	OrderUpdated          string `env:"STREAMING_TOPIC_CATEGORY_ORDER_UPDATED"`
	PostValidationSuccess string `env:"STREAMING_TOPIC_CATEGORY_POST_VALIDATION_SUCCESS"`
	PostValidationFailed  string `env:"STREAMING_TOPIC_CATEGORY_POST_VALIDATION_FAILED"`
}

type PostEvents struct {
	Created string `env:"STREAMING_TOPIC_POST_CREATED"`
	Updated string `env:"STREAMING_TOPIC_POST_UPDATED"`
}

type Nats struct {
	Url     string   `env:"NATS_URL" envDefault:"nats://localhost:4222"`
	Streams []string `env:"NATS_STREAMS" envDefault:""`
}

type CDN struct {
	Url string `env:"CDN_URL" envDefault:"http://localhost:3000"`
}

type App struct {
	Protocol string `env:"PROTOCOL" envDefault:"http"`
	DB       struct {
		Category MongoCategory
	}
	Server      Server
	HttpHeaders HttpHeaders
	I18n        I18n
	Topics      Topics
	Session     Session
	Nats        Nats
	Redis       Redis
	TokenSrv    TokenSrv
	CacheRedis  CacheRedis
	CDN         CDN
}
