package conf

type Database struct {
	Host     string
	Username string
	Password string
	Port     string
	DbName   string
	Charset  string
}

type Redis struct {
	Host     string
	Port     string
	Password string
}

type ClickHouse struct {
	Host     string
	Username string
	Password string
	Port     string
	DbName   string
}

var (
	DbConf         = Database{}
	RedisConf      = Redis{}
	ClickHouseConf = ClickHouse{}
)
