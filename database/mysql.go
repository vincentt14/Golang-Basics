package database

var connection string

// akan langsung di eksekusi saat package database di panggil
func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}