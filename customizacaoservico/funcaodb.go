package customizacaoservico

import (
	"bytes"
	"strings"

	_ "gopkg.in/goracle.v1"
	_ "gopkg.in/goracle.v1/oracle"
	"gopkg.in/gorp.v1"
)

func Debug_geraInsertCustom(owner string, nometabela string, colunas []string) string {
	return geraInsertCustom(owner, nometabela, colunas)
}

func geraInsertCustom(owner string, nometabela string, colunas []string) string {
	var insert bytes.Buffer

	insert.WriteString("insert into " + owner + "." + nometabela + "(")
	var col bytes.Buffer
	var par bytes.Buffer
	for _, v := range colunas {
		if col.Len() != 0 {
			col.WriteString(", ")
			par.WriteString(", ")
		}
		col.WriteString(v)
		par.WriteString(":" + v)
	}
	insert.Write(col.Bytes())
	insert.WriteString(") \nvalues (")
	insert.Write(par.Bytes())
	insert.WriteString(")")

	return insert.String()
}

// retorna colunaspai, colunasfilha, erro
func carregaCamposFkCustom(ownerpai string, nometabelapai string, owner string, nometabela string, trans *gorp.Transaction) ([]string, []string, error) {

	var fk [2][]string

	consulta := "select (select LISTAGG(A.COLUMN_NAME, ';') WITHIN group(order by A.POSITION) PK" +
		"          from ALL_CONS_COLUMNS A" +
		"         where A.CONSTRAINT_NAME = PK.CONSTRAINT_NAME" +
		"           and A.OWNER = PK.OWNER" +
		"         group by A.TABLE_NAME, A.CONSTRAINT_NAME) COLUNASPAI," +
		"       (select LISTAGG(A.COLUMN_NAME, ';') WITHIN group(order by A.POSITION) PK" +
		"         from ALL_CONS_COLUMNS A" +
		"        where A.CONSTRAINT_NAME = FKREF.CONSTRAINT_NAME" +
		"         and A.OWNER = FKREF.OWNER" +
		"       group by A.TABLE_NAME, A.CONSTRAINT_NAME) COLUNASFILHO" +
		"          from ALL_CONSTRAINTS FKREF" +
		"         inner join ALL_CONSTRAINTS PK" +
		"            on (PK.CONSTRAINT_NAME = FKREF.R_CONSTRAINT_NAME and PK.OWNER = FKREF.R_OWNER)" +
		"         where FKREF.CONSTRAINT_TYPE = 'R' " +
		"           and PK.CONSTRAINT_TYPE = 'P' " +
		"              and PK.OWNER = :OWNERPAI " +
		"              and PK.TABLE_NAME = :TABELAPAI " +
		"           and FKREF.OWNER = :OWNER" +
		"           and FKREF.TABLE_NAME = :TABELA" +
		"         order by FKREF.OWNER" +
		"                 ,FKREF.TABLE_NAME"

	smt, err := trans.Prepare(consulta)
	if err != nil {
		return fk[0], fk[1], err
	}

	rows, err := smt.Query(ownerpai, nometabelapai, owner, nometabela)
	if err != nil {
		return fk[0], fk[1], err
	}
	defer rows.Close()

	for rows.Next() {
		var vColuna1, vColuna2 string
		rows.Scan(&vColuna1, &vColuna2)
		fk[0] = strings.Split(vColuna1, ";")
		fk[1] = strings.Split(vColuna2, ";")
	}

	return fk[0], fk[1], nil
}

func carregaCamposPkCustom(owner string, nometabela string, trans *gorp.Transaction) ([]string, error) {

	var pk []string

	consulta := "select COLUMN_NAME from ALL_CONS_COLUMNS CONSCOL, ALL_CONSTRAINTS CONS" +
		" where  CONSCOL.OWNER      = CONS.OWNER" +
		" and    CONSCOL.TABLE_NAME = CONS.TABLE_NAME" +
		" and    CONSCOL.CONSTRAINT_NAME = CONS.CONSTRAINT_NAME" +
		" and    CONSTRAINT_TYPE = 'P'" +
		" and    CONS.OWNER      = :OWNER" +
		" and    CONS.TABLE_NAME = :TABELA" +
		" order by POSITION"

	smt, err := trans.Prepare(consulta)
	if err != nil {
		return pk, err
	}

	rows, err := smt.Query(owner, nometabela)
	if err != nil {
		return pk, err
	}
	defer rows.Close()

	for rows.Next() {
		var vColuna string
		rows.Scan(&vColuna)
		pk = append(pk, vColuna)
	}

	return pk, nil
}

func geraSelectCustom(owner string, nometabela string, colunaswhere []string) string {

	var sel bytes.Buffer

	sel.WriteString("select * from ")
	sel.WriteString(owner + "." + nometabela)
	if len(colunaswhere) > 0 {
		var where bytes.Buffer
		for _, v := range colunaswhere {
			if where.Len() != 0 {
				where.WriteString(" and ")
			} else {
				where.WriteString(" where ")
			}
			where.WriteString(v)
			where.WriteString(" = :")
			where.WriteString(v)
		}
		sel.Write(where.Bytes())
	}

	return sel.String()
}

func geraUpdateCustom(owner string, nometabela string, colunas []string, colunaswhere []string) string {

	var sel bytes.Buffer

	sel.WriteString("update ")
	sel.WriteString(owner + "." + nometabela)
	var colset bytes.Buffer
	for _, v := range colunas {
		if colset.Len() != 0 {
			colset.WriteString("\n ,")
		} else {
			colset.WriteString("\n set ")
		}
		colset.WriteString(v)
		colset.WriteString(" = :")
		colset.WriteString(v)
	}
	sel.Write(colset.Bytes())
	if len(colunaswhere) > 0 {
		var where bytes.Buffer
		for _, v := range colunaswhere {
			if where.Len() != 0 {
				where.WriteString("\n and ")
			} else {
				where.WriteString("\n where ")
			}
			where.WriteString(v)
			where.WriteString(" = :P")
			where.WriteString(v)
		}
		sel.Write(where.Bytes())
	}

	return sel.String()
}

func geraDeleteCustom(owner string, nometabela string, colunaswhere []string) string {

	var sel bytes.Buffer

	sel.WriteString("delete from ")
	sel.WriteString(owner + "." + nometabela)
	if len(colunaswhere) > 0 {
		var where bytes.Buffer
		for _, v := range colunaswhere {
			if where.Len() != 0 {
				where.WriteString("\n and ")
			} else {
				where.WriteString("\n where ")
			}
			where.WriteString(v)
			where.WriteString(" = :")
			where.WriteString(v)
		}
		sel.Write(where.Bytes())
	}

	return sel.String()
}
