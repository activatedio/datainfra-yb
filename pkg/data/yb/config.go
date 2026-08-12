package yb

// DefaultPort is the default YSQL port for a YugabyteDB cluster.
const DefaultPort = 5433

// Config represents the configuration for a YugabyteDB YSQL connection
type Config struct {

	// Hosts specifies the YSQL host addresses of the YugabyteDB cluster nodes.
	Hosts []string `json:"hosts"`

	// Port specifies the YSQL port. Defaults to 5433 when unset.
	Port int `json:"port"`

	// Username specifies the database user to connect as.
	Username string `json:"username"`

	// Password specifies the password for the database user.
	Password string `json:"password"`

	// Name specifies the name of the database to connect to.
	Name string `json:"name"`

	// EnableDefaultTransaction determines whether gorm wraps every write in a transaction.
	EnableDefaultTransaction bool `json:"enableDefaultTransaction"`

	// EnableSQLLogging determines whether SQL statement logging is enabled.
	EnableSQLLogging bool `json:"enableSqlLogging"`

	// MaxIdleConns optionally overrides the maximum number of idle connections in the pool.
	MaxIdleConns *int `json:"maxIdleConns"`
}
