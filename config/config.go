package config

func GetDatabaseConfig() map[string]string {
	return map[string]string{
		"user":     "pandora",
		"password": "pandora",
		"host":     "127.0.0.1",
		"port":     "3306",
		"dbname":   "facturador",
	}
}
