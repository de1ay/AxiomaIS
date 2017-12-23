package conf

const (
	// Server settings
	DOMAIN = "axiomais.ru"
	HTTP_PORT = ":80"
	HTTPS_PORT = ":443"
	PATH_STAIC_FILES = "./wwwroot/dist/static"
	PATH_INDEX = "./wwwroot/dist/index.html"
	MAIN_SITE_TLS = "./conf/certificates/main.pem"
	MAIN_SITE_TLS_KEY = "./conf/certificates/main_key.pem"
	WWW_SITE_TLS = "./conf/certificates/www.pem"
	WWW_SITE_TLS_KEY = "./conf/certificates/www_key.pem"
	API_SITE_TLS = "./conf/certificates/api.pem"
	API_SITE_TLS_KEY = "./conf/certificates/api_key.pem"
	// Database settings
	MYSQL_SERVER_ADDR = "127.0.0.1"
	MYSQL_LOGIN = "golang"
	MYSQL_PASSWORD = "e231d0799c926e09b6acd9655619746950339ea61b51153f715782b9a5e9f2a9"
	MYSQL_DB_NAME = "axioma"
	MYSQL_SERVER_PORT = ":3306"
	MYSQL_MAX_OPEN_CONNECTIONS = 151
	MYSQL_MAX_IDLE_CONNECTIONS = 0
	// Administrative settings
	LOGIN_MIN_LENGTH = 4
	PASSWORD_MIN_LENGTH = 6
)