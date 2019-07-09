// auto generate by tools, do not modify 
// 自动生成，请勿手动修改
//
//
//golang的包的描述，注释 就是那个package的名字
package GoTest

// 自动生成的导入
import (
	"database/sql"
	"strconv"
	"github.com/glog"
	_ "github.com/go-sql-driver/mysql"
)

// golang的对应db的结构的描述，就是数据库里面的多少列，对应的go的结构体
type TestDb struct {
	//由于在mysql的users表中name没有设置为NOT NULL,
	//所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,
	//但sql.NullString则可以接收nil值
	Code		string		`db:"code"`
	OldCode		int		`db:"oldcode"`
	Name		string		`db:"当前的名字"`
	OldName		string		`db:"曾经使用的名字"`
	DayLine		string		`db:"刷新日期"`
	Remarks		string		`db:"备注"`
}


//how usd:
/*
if GoTest.IsHasDb("test1", DB) == true {
	//GoTest.DeleteDb("test1", DB)
} else {
	GoTest.CreateDb("test1", DB)
}
if GoTest.IsHasTable("test1", "test2", DB) == true {
	//GoTest.DeleteTable("test1", "test2", DB)
} else {
	GoTest.CreateTable("test1", "test2", DB)
}
if GoTest.UsdDb("test1", DB) == false {
}
//对表，进行相关的处理
oIn := GoTest.TestDb{
	Code:    "123",
	OldCode: 123,
	Name:    "123",
	OldName: "123",
	DayLine: "123",
	Remarks: "456789",
}
_b := GoTest.FInsToTbl("test2", DB, &oIn)
if _b == false {
	//return
}
_rr := GoTest.FGetAll("test2", DB)
for i, v := range _rr {
	fmt.Println(i, v)
}
_r, _b := GoTest.FGetByCode("test2", DB, "123")
if _b == true {
	fmt.Println(_r)
	_r.OldCode = 456
	_r.Name = "456"
	_r.OldName = "456"
	_r.DayLine = "456"
	_r.Remarks = "123456"
	GoTest.FUpRowByPriKey("test2", DB, &_r, _r.Code)
}
GoTest.FDelByPriKey("test2", DB, "123")
*/

// 下面是相关的函数接口，首字母大写，表示是public，go暗扣
//

//首先，生成db，需要的相关接口，总是有的
//检查一个db是否存在，如果，传入的参数错误，会抛出异常，要注意!!!
func IsHasDb(strDbName string, poDb *sql.DB) bool {
	//check input parameter
	if len(strDbName) <= 0 {
		glog.Errorln("IsHasDb no dbname")
		panic("no db name")
	}
	if poDb == nil {
		glog.Errorln("IsHasDb no db pointer")
		panic("no db pointer")
	}

	//需要先选择mysql的数据库
	_, _e := poDb.Exec("USE information_schema")
	if _e != nil {
		glog.Errorln("sql:USE information_schema, err:" + _e.Error() )
		panic("USE information_schema failed")
	}

	//执行查询db是否存在的sql
	rows, _e1 := poDb.Query("select SCHEMA_NAME from SCHEMATA where SCHEMA_NAME = ? ", strDbName)
	defer func() {
		if rows != nil {
			rows.Close() //可以关闭掉未scan连接一直占用
		}
	}()
	if _e1 != nil {
		glog.Errorln("select SCHEMA_NAME from SCHEMATA where SCHEMA_NAME " + strDbName + "err:" + _e1.Error() )
		panic("select db by name failed")
	}
	for rows.Next() {
		strGetDbName := ""

		//不scan会导致连接不释放
		_e1 = rows.Scan(&strGetDbName)
		if _e1 != nil {
			glog.Errorln("IsHasDb row.Next failed, err:" + _e1.Error() )
		}
		if strGetDbName == strDbName {
			return true
		}
	}
	return false
}

//创建一个数据库 传入参数有问题，会抛出异常，注意！！！
//数据库为：默认创建
//如果要设定数据库的字符集，请手动设定
func CreateDb(strDbName string, poDb *sql.DB) bool {
	//check input parameter
	if len(strDbName) <= 0 {
		glog.Errorln("CreateDb no dbname")
		panic("no db name")
	}
	if poDb == nil {
		glog.Errorln("CreateDb no db pointer")
		panic("no db pointer")
	}

	//check is has, if has, return false
	if IsHasDb(strDbName, poDb) == true {
	glog.Errorln("CreateDb db is has, name = %s", strDbName)
		return false
	}

	//run sql fro create database
	strSql := "create database if not exists "
	strSql += strDbName
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("CreateDb sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//删除数据库 尽量别用
func DeleteDb(strDbName string, poDb *sql.DB) bool {
	//check input parameter
	if len(strDbName) <= 0 {
		glog.Errorln("DeleteDb no dbname")
		panic("no db name")
	}
	if poDb == nil {
		glog.Errorln("DeleteDb no db pointer")
		panic("no db pointer")
	}

	//check is has, if has, return false
	if IsHasDb(strDbName, poDb) == false {
		return true
	}

	//run sql fro delete database
	strSql := "drop database if exists " + strDbName
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("DeleteDb sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}

//选择一个数据库
func UsdDb(strDbName string, poDb *sql.DB) bool {
	//check input parameter
	if len(strDbName) <= 0 {
		glog.Errorln("UsdDb no dbname")
		panic("no db name")
	}
	if poDb == nil {
		glog.Errorln("UsdDb no db pointer")
	panic("no db pointer")
	}

	if IsHasDb(strDbName, poDb) == false {
		return false
	}
	strSql := "USE " + strDbName
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("UsdDb sql:" + strSql + "err:" + _e.Error() )
		panic(strSql)
	}
	return true
}
//还需要生成表的常规的接口
//表是否存在
//传入参数为：数据库的名字 表的名字 数据库的操作指针
//传入参数，名字的size为0，会异常，db操作指针为nil，会异常
//how usd: _b := IsHasTable( "数据库名字", "表名字", dbpointer )
func IsHasTable(strDbName string, strTblName string, poDb *sql.DB) bool {
	//check input parameter
	if len(strTblName) <= 0 {
		glog.Errorln("IsHasTable strTblName size 0")
		panic("no strTblName")
	}

	if IsHasDb(strDbName, poDb) == false {
		glog.Errorln("IsHasTable no db name : %s", strDbName)
		return false
	}

	strSql := "select TABLE_NAME from INFORMATION_SCHEMA.TABLES where TABLE_SCHEMA = '"
	strSql += strDbName + "' and TABLE_NAME='" + strTblName + "'"

	rows, _e1 := poDb.Query(strSql)
	defer func() {
		if rows != nil {
			rows.Close() //可以关闭掉未scan连接一直占用
		}
	}()
	if _e1 != nil {
		glog.Errorln("IsHasTable sql:" + strSql + "err:" + _e1.Error() )
		return false
	}
	for rows.Next() {
		strGetTblName := ""

		//不scan会导致连接不释放
		_e1 = rows.Scan(&strGetTblName)
		if _e1 != nil {
			glog.Errorln("IsHasTable row.Next failed, err:" + _e1.Error() )
		}
		if strGetTblName == strTblName {
			return true
		}
	}
	return false
}

//创建表，这个复杂
func CreateTable(strDbName string, strTblName string, poDb *sql.DB) bool {
	if IsHasDb(strDbName, poDb) == false {
		return false
	}
	if UsdDb(strDbName, poDb) == false {
		return false
	}

	strSql := "CREATE TABLE " + strTblName + " ("
	strSql += " code VARCHAR( 64 ) primary key, "
	strSql += " oldcode TEXT( 64 ) NULL, "
	strSql += " 当前的名字 TEXT( 64 ) NULL, "
	strSql += " 曾经使用的名字 TEXT( 64 ) NULL, "
	strSql += " 刷新日期 TEXT( 64 ) NULL, "
	strSql += " 备注 TEXT( 1024 ) NULL ) engine=MyISAM charset=utf8"
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("CreateTable sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}

//删除表
func DeleteTable(strDbName string, strTblName string, poDb *sql.DB) bool {
	//check is has, if has, return false
	if IsHasDb(strDbName, poDb) == false {
		return true
	}
	if UsdDb(strDbName, poDb) == false {
		glog.Errorln("DeleteTable select database failed")
		panic("DeleteTable select database failed")
	}

	//check input parameter
	if len(strTblName) <= 0 {
		glog.Errorln("DeleteTable no strTblName")
		panic("no db strTblName")
	}

	if IsHasTable(strDbName, strTblName, poDb) == false {
		return true
	}

	//run sql fro delete database
	strSql := "DROP TABLE " + strTblName
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("DeleteTable sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}
//增加一个数据结构，到对应的db的table中
//how usd: oIn := OTest{} InsertToTbl("test", "test2", DB, &oIn)
//下面，还有一个FastInsertToTbl，不检查db，不选择db，不检测参数
func InsertToTbl(strDbName string, strTblName string, poDb *sql.DB, poIn *TestDb) bool{ 
	//check
	if len(strDbName) <= 0 {
		glog.Errorln("InsertToTbl, no db name")
		return false
	}
	if len(strTblName) <= 0 {
		glog.Errorln("InsertToTbl, no tbl name")
		return false
	}
	if poDb == nil {
		glog.Errorln("InsertToTbl, db pointer is nil")
		return false
	}
	if poIn == nil {
		glog.Errorln("InsertToTbl, oIn pointer is nil")
		return false
	}
	//check is has table
	if IsHasTable(strDbName, strTblName, poDb) == false {
		return false
	}
	//usd db
	if UsdDb(strDbName, poDb) == false {
		return false
	}
	strSql := "insert INTO " + strTblName
	strSql += "( code, oldcode, 当前的名字, 曾经使用的名字, 刷新日期, 备注 )"
	strSql += " values( " + poIn.Code + ", " + strconv.Itoa( poIn.OldCode ) + ", " + poIn.Name + ", " + poIn.OldName + ", " + poIn.DayLine + ", " + poIn.Remarks + ")"
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("InsertToTbl, sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}

//快速增加一个结构体，到对应db的表中的，行上
//不选择db，不检查db，不检查表 什么都不检查
//传入参数：表名，db操作指针，数据结构
func FInsToTbl(strTblName string, poDb *sql.DB, poIn *TestDb) bool{ 
	strSql := "insert INTO " + strTblName
	strSql += "( code, oldcode, 当前的名字, 曾经使用的名字, 刷新日期, 备注 )"
	strSql += " values( " + poIn.Code + ", " + strconv.Itoa( poIn.OldCode ) + ", " + poIn.Name + ", " + poIn.OldName + ", " + poIn.DayLine + ", " + poIn.Remarks + ")"

	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FInsToTbl, sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}
//删除，有主key，会有一个主key的删除
//删除一行，通过主key，删除总是一行，没有什么删除一行中的一个或几个，没有
//如果，有主key，就需要用主key写一个删除
func DeleteByPriKey( strDbName string, strTblName string, poDb *sql.DB, strCode string) bool{ 
	//check
	if len(strDbName) <= 0 {
		glog.Errorln("DeleteByPriKey, no strDbName")
		return false
	}
	if len(strTblName) <= 0 {
		glog.Errorln("DeleteByPriKey, no strTblName")
		return false
	}
	if len(strCode) <= 0 {
		glog.Errorln("DeleteByPriKey, no PrimaryKey")
		return false
	}
	if UsdDb(strDbName, poDb) == false {
		return false
	}

	strSql := "delete from " + strTblName + " where code=" + strCode
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("DeleteFromTableByCurCode, sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}

//用主key快速删除，什么都不检查
func FDelByPriKey( strTblName string, poDb *sql.DB, strCode string) bool{ 
	strSql := "delete from " + strTblName + " where code=" + strCode
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("DeleteFromTableByCurCode, sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}

//以结构中的字段，作为删除row的key的函数接口
//快速函数 以OldCode作为删除key
func FDelByOldCode( strTblName string, poDb *sql.DB, iOldCode int ) bool {
	strSql := "delete from " + strTblName + " where oldcode=" + strconv.Itoa( iOldCode )
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FDelByOldCode, sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}

//以结构中的字段，作为删除row的key的函数接口
//快速函数 以Name作为删除key
func FDelByName( strTblName string, poDb *sql.DB, strName string ) bool {
	strSql := "delete from " + strTblName + " where 当前的名字=" + strName
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FDelByName, sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}

//以结构中的字段，作为删除row的key的函数接口
//快速函数 以OldName作为删除key
func FDelByOldName( strTblName string, poDb *sql.DB, strOldName string ) bool {
	strSql := "delete from " + strTblName + " where 曾经使用的名字=" + strOldName
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FDelByOldName, sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}

//以结构中的字段，作为删除row的key的函数接口
//快速函数 以DayLine作为删除key
func FDelByDayLine( strTblName string, poDb *sql.DB, strDayLine string ) bool {
	strSql := "delete from " + strTblName + " where 刷新日期=" + strDayLine
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FDelByDayLine, sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}

//以结构中的字段，作为删除row的key的函数接口
//快速函数 以Remarks作为删除key
func FDelByRemarks( strTblName string, poDb *sql.DB, strRemarks string ) bool {
	strSql := "delete from " + strTblName + " where 备注=" + strRemarks
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FDelByRemarks, sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}

//查询类接口，都是快速查询，因为查询用的比较多
//查询类接口，都是快速查询，因为查询用的比较多
//快速获取所有
func FGetAll(strTblName string, poDb *sql.DB) []TestDb { 
	_rr := make([]TestDb, 0)
	strSql := "select * from " + strTblName
	_r, _e := poDb.Query(strSql)
	defer func() {
		if _r != nil {
			_r.Close() //可以关闭掉未scan连接一直占用
		}
	}()
	if _e != nil {
		glog.Errorln("FGetAll, sql:" + strSql + "err:" + _e.Error() )
		return _rr
	}
	for _r.Next() {
		_oIn := TestDb{}
		//不scan会导致连接不释放
		_e = _r.Scan(&_oIn.Code, &_oIn.OldCode, &_oIn.Name, &_oIn.OldName, &_oIn.DayLine, &_oIn.Remarks )
		if _e != nil {
			glog.Errorln("FGetAll, r.Next, err:" + _e.Error() )
		} else {
			_rr = append(_rr, _oIn)
		}
	}
	return _rr
}

//快速获取 通过Code获取某一行
func FGetByCode( strTblName string, poDb *sql.DB, strCode string ) ( TestDb, bool ) { 
	_oIn := TestDb{}
	strSql := "select * from " + strTblName + " where code=" + strCode
	_r := poDb.QueryRow(strSql)
	if _e := _r.Scan(&_oIn.Code, &_oIn.OldCode, &_oIn.Name, &_oIn.OldName, &_oIn.DayLine, &_oIn.Remarks ); _e != nil{ 
		glog.Errorln("FGetByCode, r.Next, err:" + _e.Error() )
		return _oIn, false
	}
	return _oIn, true
}

//快速获取 通过OldCode获取某一行
func FGetByOldCode( strTblName string, poDb *sql.DB, iOldCode int ) ( TestDb, bool ) { 
	_oIn := TestDb{}
	strSql := "select * from " + strTblName + " where oldcode=" + strconv.Itoa( iOldCode )
	_r := poDb.QueryRow(strSql)
	if _e := _r.Scan(&_oIn.Code, &_oIn.OldCode, &_oIn.Name, &_oIn.OldName, &_oIn.DayLine, &_oIn.Remarks ); _e != nil{ 
		glog.Errorln("FGetByOldCode, r.Next, err:" + _e.Error() )
		return _oIn, false
	}
	return _oIn, true
}

//快速获取 通过Name获取某一行
func FGetByName( strTblName string, poDb *sql.DB, strName string ) ( TestDb, bool ) { 
	_oIn := TestDb{}
	strSql := "select * from " + strTblName + " where 当前的名字=" + strName
	_r := poDb.QueryRow(strSql)
	if _e := _r.Scan(&_oIn.Code, &_oIn.OldCode, &_oIn.Name, &_oIn.OldName, &_oIn.DayLine, &_oIn.Remarks ); _e != nil{ 
		glog.Errorln("FGetByName, r.Next, err:" + _e.Error() )
		return _oIn, false
	}
	return _oIn, true
}

//快速获取 通过OldName获取某一行
func FGetByOldName( strTblName string, poDb *sql.DB, strOldName string ) ( TestDb, bool ) { 
	_oIn := TestDb{}
	strSql := "select * from " + strTblName + " where 曾经使用的名字=" + strOldName
	_r := poDb.QueryRow(strSql)
	if _e := _r.Scan(&_oIn.Code, &_oIn.OldCode, &_oIn.Name, &_oIn.OldName, &_oIn.DayLine, &_oIn.Remarks ); _e != nil{ 
		glog.Errorln("FGetByOldName, r.Next, err:" + _e.Error() )
		return _oIn, false
	}
	return _oIn, true
}

//快速获取 通过DayLine获取某一行
func FGetByDayLine( strTblName string, poDb *sql.DB, strDayLine string ) ( TestDb, bool ) { 
	_oIn := TestDb{}
	strSql := "select * from " + strTblName + " where 刷新日期=" + strDayLine
	_r := poDb.QueryRow(strSql)
	if _e := _r.Scan(&_oIn.Code, &_oIn.OldCode, &_oIn.Name, &_oIn.OldName, &_oIn.DayLine, &_oIn.Remarks ); _e != nil{ 
		glog.Errorln("FGetByDayLine, r.Next, err:" + _e.Error() )
		return _oIn, false
	}
	return _oIn, true
}

//快速获取 通过Remarks获取某一行
func FGetByRemarks( strTblName string, poDb *sql.DB, strRemarks string ) ( TestDb, bool ) { 
	_oIn := TestDb{}
	strSql := "select * from " + strTblName + " where 备注=" + strRemarks
	_r := poDb.QueryRow(strSql)
	if _e := _r.Scan(&_oIn.Code, &_oIn.OldCode, &_oIn.Name, &_oIn.OldName, &_oIn.DayLine, &_oIn.Remarks ); _e != nil{ 
		glog.Errorln("FGetByRemarks, r.Next, err:" + _e.Error() )
		return _oIn, false
	}
	return _oIn, true
}

//更新，只有快速接口，更新某一行的所有，以及更新某一列
//更新，只有快速接口，更新某一行的所有，以及更新某一列
//快速更新,通过主key，整个行，不检查
func FUpRowByPriKey( strTblName string, poDb *sql.DB, poUp *TestDb, strCode string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "oldcode=" + strconv.Itoa( poUp.OldCode ) + ", "
	strSql += "当前的名字=" + poUp.Name + ", "
	strSql += "曾经使用的名字=" + poUp.OldName + ", "
	strSql += "刷新日期=" + poUp.DayLine + ", "
	strSql += "备注=" + poUp.Remarks + " "
	strSql += "where code=" + strCode
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FUpRowByPriKey sql" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}
//快速更新,通过OldCode，整个行，不检查
func FUpRowByOldCode( strTblName string, poDb *sql.DB, poUp *TestDb, iOldCode int ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "code=" + poUp.Code + ", "
	strSql += "当前的名字=" + poUp.Name + ", "
	strSql += "曾经使用的名字=" + poUp.OldName + ", "
	strSql += "刷新日期=" + poUp.DayLine + ", "
	strSql += "备注=" + poUp.Remarks + " "
	strSql += "where oldcode=" + strconv.Itoa( iOldCode )
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}

//快速更新,通过Name，整个行，不检查
func FUpRowByName( strTblName string, poDb *sql.DB, poUp *TestDb, strName string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "code=" + poUp.Code + ", "
	strSql += "oldcode=" + strconv.Itoa( poUp.OldCode ) + ", "
	strSql += "曾经使用的名字=" + poUp.OldName + ", "
	strSql += "刷新日期=" + poUp.DayLine + ", "
	strSql += "备注=" + poUp.Remarks + " "
	strSql += "where 当前的名字=" + strName
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}

//快速更新,通过OldName，整个行，不检查
func FUpRowByOldName( strTblName string, poDb *sql.DB, poUp *TestDb, strOldName string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "code=" + poUp.Code + ", "
	strSql += "oldcode=" + strconv.Itoa( poUp.OldCode ) + ", "
	strSql += "当前的名字=" + poUp.Name + ", "
	strSql += "刷新日期=" + poUp.DayLine + ", "
	strSql += "备注=" + poUp.Remarks + " "
	strSql += "where 曾经使用的名字=" + strOldName
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}

//快速更新,通过DayLine，整个行，不检查
func FUpRowByDayLine( strTblName string, poDb *sql.DB, poUp *TestDb, strDayLine string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "code=" + poUp.Code + ", "
	strSql += "oldcode=" + strconv.Itoa( poUp.OldCode ) + ", "
	strSql += "当前的名字=" + poUp.Name + ", "
	strSql += "曾经使用的名字=" + poUp.OldName + ", "
	strSql += "备注=" + poUp.Remarks + " "
	strSql += "where 刷新日期=" + strDayLine
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}

//快速更新,通过Remarks，整个行，不检查
func FUpRowByRemarks( strTblName string, poDb *sql.DB, poUp *TestDb, strRemarks string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "code=" + poUp.Code + ", "
	strSql += "oldcode=" + strconv.Itoa( poUp.OldCode ) + ", "
	strSql += "当前的名字=" + poUp.Name + ", "
	strSql += "曾经使用的名字=" + poUp.OldName + ", "
	strSql += "刷新日期=" + poUp.DayLine + " "
	strSql += "where 备注=" + strRemarks
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}

//通过主key，快速更新，某一行中的某一列
//如果要更新多列，还需要增加接口
//快速更新,通过主key，单cell，不检查
func FUpOldCodeByPriKey( strTblName string, poDb *sql.DB, poUp *TestDb, iOldCode int, strCode string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "oldcode=" + strconv.Itoa( poUp.OldCode )
	strSql += "where code=" + strCode
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}
//通过主key，快速更新，某一行中的某一列
//如果要更新多列，还需要增加接口
//快速更新,通过主key，单cell，不检查
func FUpNameByPriKey( strTblName string, poDb *sql.DB, poUp *TestDb, strName string, strCode string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "当前的名字=" + poUp.Name
	strSql += "where code=" + strCode
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}
//通过主key，快速更新，某一行中的某一列
//如果要更新多列，还需要增加接口
//快速更新,通过主key，单cell，不检查
func FUpOldNameByPriKey( strTblName string, poDb *sql.DB, poUp *TestDb, strOldName string, strCode string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "曾经使用的名字=" + poUp.OldName
	strSql += "where code=" + strCode
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}
//通过主key，快速更新，某一行中的某一列
//如果要更新多列，还需要增加接口
//快速更新,通过主key，单cell，不检查
func FUpDayLineByPriKey( strTblName string, poDb *sql.DB, poUp *TestDb, strDayLine string, strCode string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "刷新日期=" + poUp.DayLine
	strSql += "where code=" + strCode
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}
//通过主key，快速更新，某一行中的某一列
//如果要更新多列，还需要增加接口
//快速更新,通过主key，单cell，不检查
func FUpRemarksByPriKey( strTblName string, poDb *sql.DB, poUp *TestDb, strRemarks string, strCode string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "备注=" + poUp.Remarks
	strSql += "where code=" + strCode
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + "err:" + _e.Error() )
		return false
	}
	return true
}
