package confs


const (
	//DATABASEEXT db后缀
	DATABASEEXT = "database"
	//file后缀
	FILEEXT = "file"
	//解析key
	key = "confs"
)
var cfg map[string]interface{}
var GLOAB_URL string

//DB_EXT_TYPE 数据库后缀
var DB_EXT_TYPE = map[string]string{
	DATABASEEXT:"data.json",
	FILEEXT:"file.json",
}

//MYSQL_CONF mysql 配置结构体
type MYSQL_CONF struct {
	Host           string `json:"host"`
	Port           string `json:"port"`
	UserName       string `json:"username"`
	Password       string `json:"password"`
	DataBase       string `json:"database"`
	Charset        string `json:"charset"`
	ConnectTimeout int    `json:"connect_timeout"`
	ReadTimeout    int    `json:"read_timeout"`
	WriteTimeout   int    `json:"write_timeout"`
	MaxIdleConn    int    `json:"max_idle_conn"`
	MaxOpenConn    int    `json:"max_open_conn"`
}

//REDIS_CONF redis配置结构体
type REDIS_CONF struct {
	Host           string `json:"host"`
	Port           string `json:"port"`
	ConnectTimeout int    `json:"connect_timeout"`
	ReadTimeout    int    `json:"read_timeout"`
	WriteTimeout   int    `json:"write_timeout"`
	MaxIdle        int    `json:"max_idle"`
	IdleTimeout    int    `json:"idle_timeout"`
}