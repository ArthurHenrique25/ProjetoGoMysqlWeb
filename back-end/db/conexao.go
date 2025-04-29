package db

func ConexaoMysql() (string, string, string, string, string) {
	user := "root"
	password := "Sadore@123"
	host := "127.0.0.1"
	porta := "3306"
	base := "Base_Comissao"

	return user, password, host, porta, base
}
