package Oracle

import (
	_ "github.com/mattn/go-oci8"
	"mega/atendimento/DashBoardSync/Types"

	"database/sql"
	"fmt"
)

func InitOracle(dbman *Types.DbManager) {

	conn, err := sql.Open("oci8", dbman.UserDb+"/"+dbman.PassDb+"@"+dbman.UrlDb)
	dbman.Conn = conn
	logError("InitOracle", err)

	err = conn.Ping()
	logError("InitOracle", err)

	/*res, err := dbman.Conn.Exec("ALTER DATABASE CHARACTER SET AL32UTF8;")
	logError("InitOracle", err)
	nm, err := res.RowsAffected()
	logError("InitOracle", err)*/

	//fmt.Println("Afetado: " + strconv.FormatInt(nm, 10))

}

func CloseConn(dbman *Types.DbManager) {

	err := dbman.Conn.Close()
	logError("CloseConn", err)

}

func SQLPrepare(sqlStatement string, dbman *Types.DbManager) {

	/*tx, err := dbman.Conn.Begin()
	dbman.TxConn = tx
	logError("SQLPrepare", err)*/

	stmt, err := dbman.Conn.Prepare(sqlStatement)
	dbman.Stmt = stmt
	logError("sqlPrepare", err)

}

func logError(fnc string, err error) {

	if err != nil {
		fmt.Println("Erro func[Oracle->" + fnc + "]: " + err.Error())
		panic(err)
	}

}
