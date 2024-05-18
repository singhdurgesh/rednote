package configs

type Postgres struct {
	Host        string
	Port        int
	Database    string
	Username    string
	Password    string
	Sslmode     string
	Sslrootcert string
	Sslcert     string
	Sslkey      string
}
