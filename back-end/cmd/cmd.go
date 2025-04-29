package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"projeto/back-end/db"

	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shopspring/decimal"
)

func main() {

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("front-end/css"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("front-end/scripts"))))

	http.HandleFunc("/", handler)
	println("** Start Process.. **")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	user, password, host, porta, base := db.ConexaoMysql()

	dsn := user + ":" + password + "@tcp(" + host + ":" + porta + ")/" + base

	db, _ := sql.Open("mysql", dsn)

	defer db.Close()

	var valorComissao decimal.Decimal
	var nome string
	err := db.QueryRow("SELECT nome, valor_comissao FROM comissaocliente WHERE nome = ?", "Arthur Henrique").Scan(&nome, &valorComissao)

	if err != nil {
		if err == sql.ErrNoRows {
			valorComissao = decimal.Zero
		} else {
			log.Fatal(err)
		}
	}

	templatePath := `C:\ProgramaçãoGO\Projeto_Mysql_Web_GO_PY\front-end\template\index.htm`

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatal("Erro ao carregar o template:", err)
	}

	despesas := valorComissao.Div(decimal.NewFromFloat(2))

	valorComissaoFormatado := strings.Replace(valorComissao.StringFixed(2), ".", ",", 1)
	despesasFormatado := strings.Replace(despesas.StringFixed(2), ".", ",", 1)

	data := struct {
		Nome          string
		ValorComissao string
		MetaDespesas  string
	}{
		Nome:          nome,
		ValorComissao: valorComissaoFormatado,
		MetaDespesas:  despesasFormatado,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
