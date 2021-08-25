package postgresdsn

type DSN struct {
	Host     string
	user     string
	password string
	dbname   string
	port     string
	sslmode  string
	timezone string
}

func CreateDSN(host string, user string, password string, dbname string, port string, sslmode string, timezone string) DSN {
	dsn := DSN{
		Host:     host,
		user:     user,
		password: password,
		dbname:   dbname,
		port:     port,
		sslmode:  sslmode,
		timezone: timezone,
	}
	return dsn
}

func (d *DSN) ReturnDSNAsString() string {
	return "host=" + d.Host + " user=" + d.user + " password=" + d.password + " dbname=" + d.dbname + " port=" + d.port + " sslmode=" + d.sslmode + " timezone=" + d.timezone
}


