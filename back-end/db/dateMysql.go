package db

import (
	"database/sql"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/shopspring/decimal"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	user, password, host, porta, base := ConexaoMysql()

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

	despesas := valorComissao.Div(decimal.NewFromFloat(7))

	resultado := valorComissao.Div(decimal.NewFromFloat(2))

	valorComissaoFormatado := strings.Replace(valorComissao.StringFixed(2), ".", ",", 1)
	despesasFormatado := strings.Replace(despesas.StringFixed(2), ".", ",", 1)
	quantiDespesa := strings.Replace(resultado.StringFixed(2), ".", ",", 1)

	data := struct {
		Nome           string
		ValorComissao  string
		MetaDespesas   string
		DespesaPessoal string
	}{
		Nome:          nome,
		ValorComissao: valorComissaoFormatado,
		MetaDespesas:  despesasFormatado,

		DespesaPessoal: quantiDespesa,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
