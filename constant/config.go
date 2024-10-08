package constant

const (
	OWNER_NAME  = "Firhansyah"
	OWNER_EMAIL = "firhansyah@visionet.co.id"

	SERVICE_ID      = "fs.api.template"
	SERVICE_NAME    = "metaforce be golang"
	SERVICE_VERSION = "0.36"

	CONFIG_PATH             = "config.json"
	CONFIG_PATH_DEVELOPMENT = "config-development.json"

	CONFIG_PATH_DB_MONGO = "database_mongo"

	CONFIG_PATH_DB_HOST_MONGO       = CONFIG_PATH_DB_MONGO + ".host"
	CONFIG_PATH_DB_PORT_MONGO       = CONFIG_PATH_DB_MONGO + ".port"
	CONFIG_PATH_DB_USERNAME_MONGO   = CONFIG_PATH_DB_MONGO + ".username"
	CONFIG_PATH_DB_PASSWORD_MONGO   = CONFIG_PATH_DB_MONGO + ".password"
	CONFIG_PATH_DB_DATABASE_MONGO   = CONFIG_PATH_DB_MONGO + ".database"
	CONFIG_PATH_DB_SSL_MODE_MONGO   = CONFIG_PATH_DB_MONGO + ".sslmode"
	CONFIG_PATH_DB_AUTHSOURCE_MONGO = CONFIG_PATH_DB_MONGO + ".auth_source"

	CONFIG_PATH_APP = "app"

	CONFIG_PATH_APP_DEBUG          = CONFIG_PATH_APP + ".debug"
	CONFIG_PATH_APP_HOST           = CONFIG_PATH_APP + ".host"
	CONFIG_PATH_APP_PORT           = CONFIG_PATH_APP + ".port"
	CONFIG_PATH_APP_USE_TLS        = CONFIG_PATH_APP + ".use_tls"
	CONFIG_PATH_APP_CERT_FILE_PATH = CONFIG_PATH_APP + ".cert_file_path"
	CONFIG_PATH_APP_KEY_FILE_PATH  = CONFIG_PATH_APP + ".key_file_path"
	CONFIG_PATH_APP_TIME_ZONE      = CONFIG_PATH_APP + ".timezone"

	CONFIG_PATH_CORS                    = CONFIG_PATH_APP + ".cors"
	CONFIG_PATH_ALLOW_ORIGIN            = CONFIG_PATH_CORS + ".allow_origin"
	CONFIG_PATH_ALLOW_METHOD            = CONFIG_PATH_CORS + ".allow_method"
	CONFIG_PATH_ALLOW_HEADER            = CONFIG_PATH_CORS + ".allow_header"
	CONFIG_PATH_EXPOSE_HEADER           = CONFIG_PATH_CORS + ".expose_header"
	CONFIG_PATH_ALLOW_CREDENTIAL        = CONFIG_PATH_CORS + ".allow_credential"
	CONFIG_PATH_ALLOW_WILDCARD          = CONFIG_PATH_CORS + ".allow_wildcard"
	CONFIG_PATH_ALLOW_BROWSER_EXTENSION = CONFIG_PATH_CORS + ".allow_browser_extension"
	CONFIG_PATH_ALLOW_WEB_SOCKET        = CONFIG_PATH_CORS + ".allow_web_socket"
	CONFIG_PATH_ALLOW_FILE              = CONFIG_PATH_CORS + ".allow_file"
	CONFIG_PATH_MAX_AGE                 = CONFIG_PATH_CORS + ".max_age"

	CONFIG_PATH_DOMAIN = "domain"
	CONFIG_PATH_DB     = "database"

	CONFIG_PATH_DB_HOST     = CONFIG_PATH_DB + ".host"
	CONFIG_PATH_DB_PORT     = CONFIG_PATH_DB + ".port"
	CONFIG_PATH_DB_USERNAME = CONFIG_PATH_DB + ".username"
	CONFIG_PATH_DB_PASSWORD = CONFIG_PATH_DB + ".password"
	CONFIG_PATH_DB_DATABASE = CONFIG_PATH_DB + ".database"
	CONFIG_PATH_DB_DIALECT  = CONFIG_PATH_DB + ".dialect"
	CONFIG_PATH_DB_SSL_MODE = CONFIG_PATH_DB + ".sslmode"

	CONFIG_PATH_REDIS = "redis"

	CONFIG_PATH_REDIS_HOST     = CONFIG_PATH_REDIS + ".host"
	CONFIG_PATH_REDIS_PASSWORD = CONFIG_PATH_REDIS + ".password"
	CONFIG_PATH_REDIS_DB       = CONFIG_PATH_REDIS + ".db"

	CONFIG_PATH_SMTP = "smtp"

	CONFIG_PATH_SMTP_HOST     = CONFIG_PATH_SMTP + ".host"
	CONFIG_PATH_SMTP_PORT     = CONFIG_PATH_SMTP + ".port"
	CONFIG_PATH_SMTP_EMAIL    = CONFIG_PATH_SMTP + ".email"
	CONFIG_PATH_SMTP_PASSWORD = CONFIG_PATH_SMTP + ".password"
	CONFIG_PATH_SMTP_NAME     = CONFIG_PATH_SMTP + ".name"

	CONFIG_PATH_ONE_SIGNAL = "one_signal"

	CONFIG_PATH_ONE_SIGNAL_APP_ID       = CONFIG_PATH_ONE_SIGNAL + ".app_id"
	CONFIG_PATH_ONE_SIGNAL_REST_API_KEY = CONFIG_PATH_ONE_SIGNAL + ".rest_api_key"

	CONFIG_PATH_NEWRELIC          = "new_relic"
	CONFIG_PATH_NEWRELIC_APP_NAME = CONFIG_PATH_NEWRELIC + ".app_name"
	CONFIG_PATH_NEWRELIC_LICENSE  = CONFIG_PATH_NEWRELIC + ".license"

	CONFIG_PATH_MINIO = "minio"

	CONFIG_PATH_MINIO_HOST           = CONFIG_PATH_MINIO + ".host"
	CONFIG_PATH_MINIO_LOCATION       = CONFIG_PATH_MINIO + ".location"
	CONFIG_PATH_MINIO_ACCESS_KEY     = CONFIG_PATH_MINIO + ".access_key"
	CONFIG_PATH_MINIO_SECRET_KEY     = CONFIG_PATH_MINIO + ".secret_key"
	CONFIG_PATH_MINIO_SSL            = CONFIG_PATH_MINIO + ".ssl"
	CONFIG_PATH_MINIO_REPLACE_DOMAIN = CONFIG_PATH_MINIO + ".replace_domain"

	DEFAULT_APP_HOST  = "localhost"
	DEFAULT_APP_PORT  = "8080"
	DEFAULT_APP_DEBUG = false

	DEFAULT_DB_HOST     = "localhost"
	DEFAULT_DB_PORT     = "3306"
	DEFAULT_DB_USERNAME = "root"
	DEFAULT_DB_PASSWORD = ""
	DEFAULT_DB_DATABASE = "template"
	DEFAULT_DB_DIALECT  = DIALECT_MYSQL
	DEFAULT_DB_SSL_MODE = "disable"
	DEFAULT_TIME_ZONE   = "Asia/Jakarta"
	DEFAULT_PATH_SQLITE = SERVICE_NAME + ".db"

	DEFAULT_REDIS_HOST     = "localhost:6379"
	DEFAULT_REDIS_PASSWORD = ""
	DEFAULT_REDIS_DB       = 0

	DEFAULT_SMTP_HOST     = "smtp.gmail.com"
	DEFAULT_SMTP_PORT     = 587
	DEFAULT_SMTP_EMAIL    = "sdh.notification@gmail.com"
	DEFAULT_SMTP_PASSWORD = "Visionet*1!"
	DEFAULT_SMTP_NAME     = "San Diego Hills Notification"

	DEFAULT_ONE_SIGNAL_APP_ID       = "c90ea719-f8cb-4eb9-b946-bb6516927b94"
	DEFAULT_ONE_SIGNAL_REST_API_KEY = "MGM3MzQwMmUtYTliOC00NmU3LWI2NGMtMTNlNmM0ZWFiNjU1"

	DEFAULT_NEW_RELIC_APP_NAME = "sdh"
	DEFAULT_NEW_RELIC_LICENSE  = "0adb99dfc961708a1218481e50a266b49f9bNRAL"

	DEFAULT_MINIO_HOST           = "localhost:9000"
	DEFAULT_MINIO_LOCATION       = "us-east-1"
	DEFAULT_MINIO_ACCESS_KEY     = "4JFFD40NAS1768JPWVVV"
	DEFAULT_MINIO_SECRET_KEY     = "JBaIB6bwL2CDWD1hoDT6t7q+zKVQv+TBH8Cx+jhR"
	DEFAULT_MINIO_SSL            = false
	DEFAULT_MINIO_REPLACE_DOMAIN = ""

	MAX_LIMIT_QUERY     = 100
	DEFAULT_LIMIT_QUERY = 10

	DIALECT_MYSQL      = "mysql"
	DIALECT_POSTGRESQL = "postgres"
	DIALECT_SQL_SERVER = "mssql"
	DIALECT_SQLITE     = "sqlite"

	SWAGGER_HOST      = "swagger_host"
	JWT_SIGNATURE_KEY = "jwt_signature_key"

	ITOP_REST          = "itop_rest"
	ITOP_REST_URL      = ITOP_REST + ".url"
	ITOP_REST_VERSION  = ITOP_REST + ".version"
	ITOP_REST_USERNAME = ITOP_REST + ".username"
	ITOP_REST_PASSWORD = ITOP_REST + ".password"
	ITOP_REST_TOKEN    = ITOP_REST + ".token"

	MAX_RANGE = "max_range"

	CONFIG_RANGE_TIME_SURVEY          = "RANGE_TIME_SURVEY"
	CONFIG_MOBILE_SOP_VERSION         = "MOBILE_SOP_VERSION"
	CONFIG_MOBILE_INSTRUCTION_VERSION = "MOBILE_INSTRUCTION_VERSION"
	CONFIG_MOBILE_VIDEO_URL           = "MOBILE_VIDEO_URL"
	CONFIG_MAX_RANGE_SURVEY           = "MAX_RANGE_SURVEY"
)
