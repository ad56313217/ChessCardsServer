// auto generate by tools, do not modify 
// 自动生成，请勿手动修改
//
//
//golang的包的描述，注释 就是那个package的名字
package CCG_Login

// 自动生成的导入
import (
	"database/sql"
	"strconv"
	"github.com/glog"
	_ "github.com/go-sql-driver/mysql"
)

// golang的对应db的结构的描述，就是数据库里面的多少列，对应的go的结构体
type CCG_LoginDb struct {
	//由于在mysql的users表中name没有设置为NOT NULL,
	//所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,
	//但sql.NullString则可以接收nil值
	Id		int		`db:"Id"`
	UsrName		string		`db:"UsrName"`
	PassWord		string		`db:"PassWord"`
	Nick		string		`db:"Nick"`
	Udid		string		`db:"Udid"`
	HeadUrl		string		`db:"HeadUrl"`
	RegIp		string		`db:"RegIp"`
	RegDateTime		string		`db:"RegDateTime"`
	RegDeviceType		int		`db:"RegDeviceType"`
	UsrTel		string		`db:"UsrTel"`
	UsrEmail		string		`db:"UsrEmail"`
	UsrICard		string		`db:"UsrICard"`
	Res1		string		`db:"Res1"`
	Res2		string		`db:"Res2"`
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
			glog.Errorln("IsHasDb row.Next err:" + _e1.Error() )
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
	glog.Errorln("CreateDb db is has, name=" + strDbName)
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
		glog.Errorln("DeleteDb sql:" + strSql + ", err:" + _e.Error() )
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
		glog.Errorln("UsdDb sql:" + strSql + ", err:" + _e.Error() )
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
		glog.Errorln("IsHasTable no has db:" + strDbName)
		return false
	}

	strSql := "select TABLE_NAME from INFORMATION_SCHEMA.TABLES where TABLE_SCHEMA = '"
	strSql += strDbName + "' and TABLE_NAME='" + strTblName + "'"

	_r := poDb.QueryRow( strSql )
	strQueryDbName := ""
	if _e := _r.Scan( &strQueryDbName ); _e != nil{
		glog.Errorln( "IsHasTable, r.Scan err:" + _e.Error() )
		return false
	} else {
		if strQueryDbName == strTblName{
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
	strSql += " Id int( 9 ) primary key, "
	strSql += " UsrName VARCHAR( 16 ) NULL, "
	strSql += " PassWord VARCHAR( 16 ) NULL, "
	strSql += " Nick VARCHAR( 16 ) NULL, "
	strSql += " Udid VARCHAR( 64 ) NULL, "
	strSql += " HeadUrl VARCHAR( 256 ) NULL, "
	strSql += " RegIp VARCHAR( 16 ) NULL, "
	strSql += " RegDateTime VARCHAR( 32 ) NULL, "
	strSql += " RegDeviceType int( 9 ) NULL, "
	strSql += " UsrTel VARCHAR( 32 ) NULL, "
	strSql += " UsrEmail VARCHAR( 32 ) NULL, "
	strSql += " UsrICard VARCHAR( 18 ) NULL, "
	strSql += " Res1 VARCHAR( 32 ) NULL, "
	strSql += " Res2 VARCHAR( 32 ) NULL ) engine=MyISAM charset=utf8"
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("CreateTable sql:" + strSql + ", err:" + _e.Error() )
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
	//check input parameter
	if len(strTblName) <= 0 {
		glog.Errorln("DeleteTable no strTblName")
		panic("no db strTblName")
	}
	if IsHasTable(strDbName, strTblName, poDb) == false {
		return true
	}
	if UsdDb(strDbName, poDb) == false {
		glog.Errorln("DeleteTable select database failed")
		panic("DeleteTable select database failed")
	}
	//run sql fro delete database
	strSql := "DROP TABLE " + strTblName
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("DeleteTable sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}
//增加一个数据结构，到对应的db的table中
//how usd: oIn := OTest{} InsertToTbl("test", "test2", DB, &oIn)
//下面，还有一个FastInsertToTbl，不检查db，不选择db，不检测参数
func InsertToTbl(strDbName string, strTblName string, poDb *sql.DB, poIn *CCG_LoginDb) bool{ 
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
	strSql += "( Id, UsrName, PassWord, Nick, Udid, HeadUrl, RegIp, RegDateTime, RegDeviceType, UsrTel, UsrEmail, UsrICard, Res1, Res2 )"
	strSql += " values( '" + strconv.Itoa( poIn.Id ) + "', '" + poIn.UsrName + "', '" + poIn.PassWord + "', '" + poIn.Nick + "', '" + poIn.Udid + "', '" + poIn.HeadUrl + "', '" + poIn.RegIp + "', '" + poIn.RegDateTime + "', '" + strconv.Itoa( poIn.RegDeviceType ) + "', '" + poIn.UsrTel + "', '" + poIn.UsrEmail + "', '" + poIn.UsrICard + "', '" + poIn.Res1 + "', '" + poIn.Res2 + "')"
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("InsertToTbl, sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//快速增加一个结构体，到对应db的表中的，行上
//不选择db，不检查db，不检查表 什么都不检查
//传入参数：表名，db操作指针，数据结构
func FInsToTbl(strTblName string, poDb *sql.DB, poIn *CCG_LoginDb) bool{ 
	strSql := "insert INTO " + strTblName
	strSql += "( Id, UsrName, PassWord, Nick, Udid, HeadUrl, RegIp, RegDateTime, RegDeviceType, UsrTel, UsrEmail, UsrICard, Res1, Res2 )"
	strSql += " values( '" + strconv.Itoa( poIn.Id ) + "', '" + poIn.UsrName + "', '" + poIn.PassWord + "', '" + poIn.Nick + "', '" + poIn.Udid + "', '" + poIn.HeadUrl + "', '" + poIn.RegIp + "', '" + poIn.RegDateTime + "', '" + strconv.Itoa( poIn.RegDeviceType ) + "', '" + poIn.UsrTel + "', '" + poIn.UsrEmail + "', '" + poIn.UsrICard + "', '" + poIn.Res1 + "', '" + poIn.Res2 + "')"

	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FInsToTbl, sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}
//删除，有主key，会有一个主key的删除
//删除一行，通过主key，删除总是一行，没有什么删除一行中的一个或几个，没有
//如果，有主key，就需要用主key写一个删除
func DeleteByPriKey( strDbName string, strTblName string, poDb *sql.DB, iId int) bool{ 
	//check
	if len(strDbName) <= 0 {
		glog.Errorln("DeleteByPriKey, no strDbName")
		return false
	}
	if len(strTblName) <= 0 {
		glog.Errorln("DeleteByPriKey, no strTblName")
		return false
	}
	if UsdDb(strDbName, poDb) == false {
		return false
	}

	strSql := "delete from " + strTblName + " where Id=" + "'" + strconv.Itoa( iId ) + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("DeleteByPriKey, sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//用主key快速删除，什么都不检查
func FDelByPriKey( strTblName string, poDb *sql.DB, iId int) bool{ 
	strSql := "delete from " + strTblName + " where Id=" + "'" + strconv.Itoa( iId ) + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FDelByPriKey, sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//以结构中的字段，作为删除row的key的函数接口
//快速函数 以UsrName作为删除key
func FDelByUsrName( strTblName string, poDb *sql.DB, strUsrName string ) bool {
	strSql := "delete from " + strTblName + " where UsrName=" + "'" + strUsrName + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FDelByUsrName, sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//以结构中的字段，作为删除row的key的函数接口
//快速函数 以PassWord作为删除key
func FDelByPassWord( strTblName string, poDb *sql.DB, strPassWord string ) bool {
	strSql := "delete from " + strTblName + " where PassWord=" + "'" + strPassWord + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FDelByPassWord, sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//以结构中的字段，作为删除row的key的函数接口
//快速函数 以Nick作为删除key
func FDelByNick( strTblName string, poDb *sql.DB, strNick string ) bool {
	strSql := "delete from " + strTblName + " where Nick=" + "'" + strNick + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FDelByNick, sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//以结构中的字段，作为删除row的key的函数接口
//快速函数 以Udid作为删除key
func FDelByUdid( strTblName string, poDb *sql.DB, strUdid string ) bool {
	strSql := "delete from " + strTblName + " where Udid=" + "'" + strUdid + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FDelByUdid, sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//以结构中的字段，作为删除row的key的函数接口
//快速函数 以HeadUrl作为删除key
func FDelByHeadUrl( strTblName string, poDb *sql.DB, strHeadUrl string ) bool {
	strSql := "delete from " + strTblName + " where HeadUrl=" + "'" + strHeadUrl + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FDelByHeadUrl, sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//以结构中的字段，作为删除row的key的函数接口
//快速函数 以RegIp作为删除key
func FDelByRegIp( strTblName string, poDb *sql.DB, strRegIp string ) bool {
	strSql := "delete from " + strTblName + " where RegIp=" + "'" + strRegIp + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FDelByRegIp, sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//以结构中的字段，作为删除row的key的函数接口
//快速函数 以RegDateTime作为删除key
func FDelByRegDateTime( strTblName string, poDb *sql.DB, strRegDateTime string ) bool {
	strSql := "delete from " + strTblName + " where RegDateTime=" + "'" + strRegDateTime + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FDelByRegDateTime, sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//以结构中的字段，作为删除row的key的函数接口
//快速函数 以RegDeviceType作为删除key
func FDelByRegDeviceType( strTblName string, poDb *sql.DB, iRegDeviceType int ) bool {
	strSql := "delete from " + strTblName + " where RegDeviceType=" + "'" + strconv.Itoa( iRegDeviceType ) + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FDelByRegDeviceType, sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//以结构中的字段，作为删除row的key的函数接口
//快速函数 以UsrTel作为删除key
func FDelByUsrTel( strTblName string, poDb *sql.DB, strUsrTel string ) bool {
	strSql := "delete from " + strTblName + " where UsrTel=" + "'" + strUsrTel + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FDelByUsrTel, sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//以结构中的字段，作为删除row的key的函数接口
//快速函数 以UsrEmail作为删除key
func FDelByUsrEmail( strTblName string, poDb *sql.DB, strUsrEmail string ) bool {
	strSql := "delete from " + strTblName + " where UsrEmail=" + "'" + strUsrEmail + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FDelByUsrEmail, sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//以结构中的字段，作为删除row的key的函数接口
//快速函数 以UsrICard作为删除key
func FDelByUsrICard( strTblName string, poDb *sql.DB, strUsrICard string ) bool {
	strSql := "delete from " + strTblName + " where UsrICard=" + "'" + strUsrICard + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FDelByUsrICard, sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//以结构中的字段，作为删除row的key的函数接口
//快速函数 以Res1作为删除key
func FDelByRes1( strTblName string, poDb *sql.DB, strRes1 string ) bool {
	strSql := "delete from " + strTblName + " where Res1=" + "'" + strRes1 + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FDelByRes1, sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//以结构中的字段，作为删除row的key的函数接口
//快速函数 以Res2作为删除key
func FDelByRes2( strTblName string, poDb *sql.DB, strRes2 string ) bool {
	strSql := "delete from " + strTblName + " where Res2=" + "'" + strRes2 + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FDelByRes2, sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//查询类接口，都是快速查询，因为查询用的比较多
//查询类接口，都是快速查询，因为查询用的比较多
//快速获取所有
func FGetAll(strTblName string, poDb *sql.DB) []CCG_LoginDb { 
	_rr := make([]CCG_LoginDb, 0)
	strSql := "select * from " + strTblName
	_r, _e := poDb.Query(strSql)
	defer func() {
		if _r != nil {
			_r.Close() //可以关闭掉未scan连接一直占用
		}
	}()
	if _e != nil {
		glog.Errorln("FGetAll, sql:" + strSql + ", err:" + _e.Error() )
		return _rr
	}
	for _r.Next() {
		_oIn := CCG_LoginDb{}
		//不scan会导致连接不释放
		_e = _r.Scan(&_oIn.Id, &_oIn.UsrName, &_oIn.PassWord, &_oIn.Nick, &_oIn.Udid, &_oIn.HeadUrl, &_oIn.RegIp, &_oIn.RegDateTime, &_oIn.RegDeviceType, &_oIn.UsrTel, &_oIn.UsrEmail, &_oIn.UsrICard, &_oIn.Res1, &_oIn.Res2 )
		if _e != nil {
			glog.Errorln("FGetAll, r.Next err:" + _e.Error() )
		} else {
			_rr = append(_rr, _oIn)
		}
	}
	return _rr
}

//快速获取 通过Id获取某一行
func FGetById( strTblName string, poDb *sql.DB, iId int ) ( CCG_LoginDb, bool ) { 
	_oIn := CCG_LoginDb{}
	strSql := "select * from " + strTblName + " where Id=" + "'" + strconv.Itoa( iId ) + "'" 
	_r := poDb.QueryRow(strSql)
	if _e := _r.Scan(&_oIn.Id, &_oIn.UsrName, &_oIn.PassWord, &_oIn.Nick, &_oIn.Udid, &_oIn.HeadUrl, &_oIn.RegIp, &_oIn.RegDateTime, &_oIn.RegDeviceType, &_oIn.UsrTel, &_oIn.UsrEmail, &_oIn.UsrICard, &_oIn.Res1, &_oIn.Res2 ); _e != nil{ 
		glog.Errorln("FGetById, r.Scan err:" + _e.Error() )
		return _oIn, false
	}
	return _oIn, true
}

//快速获取 通过UsrName获取某一行
func FGetByUsrName( strTblName string, poDb *sql.DB, strUsrName string ) ( CCG_LoginDb, bool ) { 
	_oIn := CCG_LoginDb{}
	strSql := "select * from " + strTblName + " where UsrName=" + "'" + strUsrName + "'" 
	_r := poDb.QueryRow(strSql)
	if _e := _r.Scan(&_oIn.Id, &_oIn.UsrName, &_oIn.PassWord, &_oIn.Nick, &_oIn.Udid, &_oIn.HeadUrl, &_oIn.RegIp, &_oIn.RegDateTime, &_oIn.RegDeviceType, &_oIn.UsrTel, &_oIn.UsrEmail, &_oIn.UsrICard, &_oIn.Res1, &_oIn.Res2 ); _e != nil{ 
		glog.Errorln("FGetByUsrName, r.Scan err:" + _e.Error() )
		return _oIn, false
	}
	return _oIn, true
}

//快速获取 通过PassWord获取某一行
func FGetByPassWord( strTblName string, poDb *sql.DB, strPassWord string ) ( CCG_LoginDb, bool ) { 
	_oIn := CCG_LoginDb{}
	strSql := "select * from " + strTblName + " where PassWord=" + "'" + strPassWord + "'" 
	_r := poDb.QueryRow(strSql)
	if _e := _r.Scan(&_oIn.Id, &_oIn.UsrName, &_oIn.PassWord, &_oIn.Nick, &_oIn.Udid, &_oIn.HeadUrl, &_oIn.RegIp, &_oIn.RegDateTime, &_oIn.RegDeviceType, &_oIn.UsrTel, &_oIn.UsrEmail, &_oIn.UsrICard, &_oIn.Res1, &_oIn.Res2 ); _e != nil{ 
		glog.Errorln("FGetByPassWord, r.Scan err:" + _e.Error() )
		return _oIn, false
	}
	return _oIn, true
}

//快速获取 通过Nick获取某一行
func FGetByNick( strTblName string, poDb *sql.DB, strNick string ) ( CCG_LoginDb, bool ) { 
	_oIn := CCG_LoginDb{}
	strSql := "select * from " + strTblName + " where Nick=" + "'" + strNick + "'" 
	_r := poDb.QueryRow(strSql)
	if _e := _r.Scan(&_oIn.Id, &_oIn.UsrName, &_oIn.PassWord, &_oIn.Nick, &_oIn.Udid, &_oIn.HeadUrl, &_oIn.RegIp, &_oIn.RegDateTime, &_oIn.RegDeviceType, &_oIn.UsrTel, &_oIn.UsrEmail, &_oIn.UsrICard, &_oIn.Res1, &_oIn.Res2 ); _e != nil{ 
		glog.Errorln("FGetByNick, r.Scan err:" + _e.Error() )
		return _oIn, false
	}
	return _oIn, true
}

//快速获取 通过Udid获取某一行
func FGetByUdid( strTblName string, poDb *sql.DB, strUdid string ) ( CCG_LoginDb, bool ) { 
	_oIn := CCG_LoginDb{}
	strSql := "select * from " + strTblName + " where Udid=" + "'" + strUdid + "'" 
	_r := poDb.QueryRow(strSql)
	if _e := _r.Scan(&_oIn.Id, &_oIn.UsrName, &_oIn.PassWord, &_oIn.Nick, &_oIn.Udid, &_oIn.HeadUrl, &_oIn.RegIp, &_oIn.RegDateTime, &_oIn.RegDeviceType, &_oIn.UsrTel, &_oIn.UsrEmail, &_oIn.UsrICard, &_oIn.Res1, &_oIn.Res2 ); _e != nil{ 
		glog.Errorln("FGetByUdid, r.Scan err:" + _e.Error() )
		return _oIn, false
	}
	return _oIn, true
}

//快速获取 通过HeadUrl获取某一行
func FGetByHeadUrl( strTblName string, poDb *sql.DB, strHeadUrl string ) ( CCG_LoginDb, bool ) { 
	_oIn := CCG_LoginDb{}
	strSql := "select * from " + strTblName + " where HeadUrl=" + "'" + strHeadUrl + "'" 
	_r := poDb.QueryRow(strSql)
	if _e := _r.Scan(&_oIn.Id, &_oIn.UsrName, &_oIn.PassWord, &_oIn.Nick, &_oIn.Udid, &_oIn.HeadUrl, &_oIn.RegIp, &_oIn.RegDateTime, &_oIn.RegDeviceType, &_oIn.UsrTel, &_oIn.UsrEmail, &_oIn.UsrICard, &_oIn.Res1, &_oIn.Res2 ); _e != nil{ 
		glog.Errorln("FGetByHeadUrl, r.Scan err:" + _e.Error() )
		return _oIn, false
	}
	return _oIn, true
}

//快速获取 通过RegIp获取某一行
func FGetByRegIp( strTblName string, poDb *sql.DB, strRegIp string ) ( CCG_LoginDb, bool ) { 
	_oIn := CCG_LoginDb{}
	strSql := "select * from " + strTblName + " where RegIp=" + "'" + strRegIp + "'" 
	_r := poDb.QueryRow(strSql)
	if _e := _r.Scan(&_oIn.Id, &_oIn.UsrName, &_oIn.PassWord, &_oIn.Nick, &_oIn.Udid, &_oIn.HeadUrl, &_oIn.RegIp, &_oIn.RegDateTime, &_oIn.RegDeviceType, &_oIn.UsrTel, &_oIn.UsrEmail, &_oIn.UsrICard, &_oIn.Res1, &_oIn.Res2 ); _e != nil{ 
		glog.Errorln("FGetByRegIp, r.Scan err:" + _e.Error() )
		return _oIn, false
	}
	return _oIn, true
}

//快速获取 通过RegDateTime获取某一行
func FGetByRegDateTime( strTblName string, poDb *sql.DB, strRegDateTime string ) ( CCG_LoginDb, bool ) { 
	_oIn := CCG_LoginDb{}
	strSql := "select * from " + strTblName + " where RegDateTime=" + "'" + strRegDateTime + "'" 
	_r := poDb.QueryRow(strSql)
	if _e := _r.Scan(&_oIn.Id, &_oIn.UsrName, &_oIn.PassWord, &_oIn.Nick, &_oIn.Udid, &_oIn.HeadUrl, &_oIn.RegIp, &_oIn.RegDateTime, &_oIn.RegDeviceType, &_oIn.UsrTel, &_oIn.UsrEmail, &_oIn.UsrICard, &_oIn.Res1, &_oIn.Res2 ); _e != nil{ 
		glog.Errorln("FGetByRegDateTime, r.Scan err:" + _e.Error() )
		return _oIn, false
	}
	return _oIn, true
}

//快速获取 通过RegDeviceType获取某一行
func FGetByRegDeviceType( strTblName string, poDb *sql.DB, iRegDeviceType int ) ( CCG_LoginDb, bool ) { 
	_oIn := CCG_LoginDb{}
	strSql := "select * from " + strTblName + " where RegDeviceType=" + "'" + strconv.Itoa( iRegDeviceType ) + "'" 
	_r := poDb.QueryRow(strSql)
	if _e := _r.Scan(&_oIn.Id, &_oIn.UsrName, &_oIn.PassWord, &_oIn.Nick, &_oIn.Udid, &_oIn.HeadUrl, &_oIn.RegIp, &_oIn.RegDateTime, &_oIn.RegDeviceType, &_oIn.UsrTel, &_oIn.UsrEmail, &_oIn.UsrICard, &_oIn.Res1, &_oIn.Res2 ); _e != nil{ 
		glog.Errorln("FGetByRegDeviceType, r.Scan err:" + _e.Error() )
		return _oIn, false
	}
	return _oIn, true
}

//快速获取 通过UsrTel获取某一行
func FGetByUsrTel( strTblName string, poDb *sql.DB, strUsrTel string ) ( CCG_LoginDb, bool ) { 
	_oIn := CCG_LoginDb{}
	strSql := "select * from " + strTblName + " where UsrTel=" + "'" + strUsrTel + "'" 
	_r := poDb.QueryRow(strSql)
	if _e := _r.Scan(&_oIn.Id, &_oIn.UsrName, &_oIn.PassWord, &_oIn.Nick, &_oIn.Udid, &_oIn.HeadUrl, &_oIn.RegIp, &_oIn.RegDateTime, &_oIn.RegDeviceType, &_oIn.UsrTel, &_oIn.UsrEmail, &_oIn.UsrICard, &_oIn.Res1, &_oIn.Res2 ); _e != nil{ 
		glog.Errorln("FGetByUsrTel, r.Scan err:" + _e.Error() )
		return _oIn, false
	}
	return _oIn, true
}

//快速获取 通过UsrEmail获取某一行
func FGetByUsrEmail( strTblName string, poDb *sql.DB, strUsrEmail string ) ( CCG_LoginDb, bool ) { 
	_oIn := CCG_LoginDb{}
	strSql := "select * from " + strTblName + " where UsrEmail=" + "'" + strUsrEmail + "'" 
	_r := poDb.QueryRow(strSql)
	if _e := _r.Scan(&_oIn.Id, &_oIn.UsrName, &_oIn.PassWord, &_oIn.Nick, &_oIn.Udid, &_oIn.HeadUrl, &_oIn.RegIp, &_oIn.RegDateTime, &_oIn.RegDeviceType, &_oIn.UsrTel, &_oIn.UsrEmail, &_oIn.UsrICard, &_oIn.Res1, &_oIn.Res2 ); _e != nil{ 
		glog.Errorln("FGetByUsrEmail, r.Scan err:" + _e.Error() )
		return _oIn, false
	}
	return _oIn, true
}

//快速获取 通过UsrICard获取某一行
func FGetByUsrICard( strTblName string, poDb *sql.DB, strUsrICard string ) ( CCG_LoginDb, bool ) { 
	_oIn := CCG_LoginDb{}
	strSql := "select * from " + strTblName + " where UsrICard=" + "'" + strUsrICard + "'" 
	_r := poDb.QueryRow(strSql)
	if _e := _r.Scan(&_oIn.Id, &_oIn.UsrName, &_oIn.PassWord, &_oIn.Nick, &_oIn.Udid, &_oIn.HeadUrl, &_oIn.RegIp, &_oIn.RegDateTime, &_oIn.RegDeviceType, &_oIn.UsrTel, &_oIn.UsrEmail, &_oIn.UsrICard, &_oIn.Res1, &_oIn.Res2 ); _e != nil{ 
		glog.Errorln("FGetByUsrICard, r.Scan err:" + _e.Error() )
		return _oIn, false
	}
	return _oIn, true
}

//快速获取 通过Res1获取某一行
func FGetByRes1( strTblName string, poDb *sql.DB, strRes1 string ) ( CCG_LoginDb, bool ) { 
	_oIn := CCG_LoginDb{}
	strSql := "select * from " + strTblName + " where Res1=" + "'" + strRes1 + "'" 
	_r := poDb.QueryRow(strSql)
	if _e := _r.Scan(&_oIn.Id, &_oIn.UsrName, &_oIn.PassWord, &_oIn.Nick, &_oIn.Udid, &_oIn.HeadUrl, &_oIn.RegIp, &_oIn.RegDateTime, &_oIn.RegDeviceType, &_oIn.UsrTel, &_oIn.UsrEmail, &_oIn.UsrICard, &_oIn.Res1, &_oIn.Res2 ); _e != nil{ 
		glog.Errorln("FGetByRes1, r.Scan err:" + _e.Error() )
		return _oIn, false
	}
	return _oIn, true
}

//快速获取 通过Res2获取某一行
func FGetByRes2( strTblName string, poDb *sql.DB, strRes2 string ) ( CCG_LoginDb, bool ) { 
	_oIn := CCG_LoginDb{}
	strSql := "select * from " + strTblName + " where Res2=" + "'" + strRes2 + "'" 
	_r := poDb.QueryRow(strSql)
	if _e := _r.Scan(&_oIn.Id, &_oIn.UsrName, &_oIn.PassWord, &_oIn.Nick, &_oIn.Udid, &_oIn.HeadUrl, &_oIn.RegIp, &_oIn.RegDateTime, &_oIn.RegDeviceType, &_oIn.UsrTel, &_oIn.UsrEmail, &_oIn.UsrICard, &_oIn.Res1, &_oIn.Res2 ); _e != nil{ 
		glog.Errorln("FGetByRes2, r.Scan err:" + _e.Error() )
		return _oIn, false
	}
	return _oIn, true
}

//更新，只有快速接口，更新某一行的所有，以及更新某一列
//更新，只有快速接口，更新某一行的所有，以及更新某一列
//快速更新,通过主key，整个行，不检查
func FUpRowByPriKey( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, iId int ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "UsrName='" + poUp.UsrName + "', "
	strSql += "PassWord='" + poUp.PassWord + "', "
	strSql += "Nick='" + poUp.Nick + "', "
	strSql += "Udid='" + poUp.Udid + "', "
	strSql += "HeadUrl='" + poUp.HeadUrl + "', "
	strSql += "RegIp='" + poUp.RegIp + "', "
	strSql += "RegDateTime='" + poUp.RegDateTime + "', "
	strSql += "RegDeviceType='" + strconv.Itoa( poUp.RegDeviceType ) + "', "
	strSql += "UsrTel='" + poUp.UsrTel + "', "
	strSql += "UsrEmail='" + poUp.UsrEmail + "', "
	strSql += "UsrICard='" + poUp.UsrICard + "', "
	strSql += "Res1='" + poUp.Res1 + "', "
	strSql += "Res2='" + poUp.Res2 + "' "
	strSql += "where Id= '" + strconv.Itoa( iId ) + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("FUpRowByPriKey sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}
//快速更新,通过UsrName，整个行，不检查
func FUpRowByUsrName( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strUsrName string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "Id= '" + strconv.Itoa( poUp.Id ) + "', " 
	strSql += "PassWord= '" + poUp.PassWord + "', " 
	strSql += "Nick= '" + poUp.Nick + "', " 
	strSql += "Udid= '" + poUp.Udid + "', " 
	strSql += "HeadUrl= '" + poUp.HeadUrl + "', " 
	strSql += "RegIp= '" + poUp.RegIp + "', " 
	strSql += "RegDateTime= '" + poUp.RegDateTime + "', " 
	strSql += "RegDeviceType= '" + strconv.Itoa( poUp.RegDeviceType ) + "', " 
	strSql += "UsrTel= '" + poUp.UsrTel + "', " 
	strSql += "UsrEmail= '" + poUp.UsrEmail + "', " 
	strSql += "UsrICard= '" + poUp.UsrICard + "', " 
	strSql += "Res1= '" + poUp.Res1 + "', " 
	strSql += "Res2= '" + poUp.Res2 + "'" 
	strSql += "where UsrName='" + strUsrName + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//快速更新,通过PassWord，整个行，不检查
func FUpRowByPassWord( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strPassWord string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "Id= '" + strconv.Itoa( poUp.Id ) + "', " 
	strSql += "UsrName= '" + poUp.UsrName + "', " 
	strSql += "Nick= '" + poUp.Nick + "', " 
	strSql += "Udid= '" + poUp.Udid + "', " 
	strSql += "HeadUrl= '" + poUp.HeadUrl + "', " 
	strSql += "RegIp= '" + poUp.RegIp + "', " 
	strSql += "RegDateTime= '" + poUp.RegDateTime + "', " 
	strSql += "RegDeviceType= '" + strconv.Itoa( poUp.RegDeviceType ) + "', " 
	strSql += "UsrTel= '" + poUp.UsrTel + "', " 
	strSql += "UsrEmail= '" + poUp.UsrEmail + "', " 
	strSql += "UsrICard= '" + poUp.UsrICard + "', " 
	strSql += "Res1= '" + poUp.Res1 + "', " 
	strSql += "Res2= '" + poUp.Res2 + "'" 
	strSql += "where PassWord='" + strPassWord + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//快速更新,通过Nick，整个行，不检查
func FUpRowByNick( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strNick string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "Id= '" + strconv.Itoa( poUp.Id ) + "', " 
	strSql += "UsrName= '" + poUp.UsrName + "', " 
	strSql += "PassWord= '" + poUp.PassWord + "', " 
	strSql += "Udid= '" + poUp.Udid + "', " 
	strSql += "HeadUrl= '" + poUp.HeadUrl + "', " 
	strSql += "RegIp= '" + poUp.RegIp + "', " 
	strSql += "RegDateTime= '" + poUp.RegDateTime + "', " 
	strSql += "RegDeviceType= '" + strconv.Itoa( poUp.RegDeviceType ) + "', " 
	strSql += "UsrTel= '" + poUp.UsrTel + "', " 
	strSql += "UsrEmail= '" + poUp.UsrEmail + "', " 
	strSql += "UsrICard= '" + poUp.UsrICard + "', " 
	strSql += "Res1= '" + poUp.Res1 + "', " 
	strSql += "Res2= '" + poUp.Res2 + "'" 
	strSql += "where Nick='" + strNick + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//快速更新,通过Udid，整个行，不检查
func FUpRowByUdid( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strUdid string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "Id= '" + strconv.Itoa( poUp.Id ) + "', " 
	strSql += "UsrName= '" + poUp.UsrName + "', " 
	strSql += "PassWord= '" + poUp.PassWord + "', " 
	strSql += "Nick= '" + poUp.Nick + "', " 
	strSql += "HeadUrl= '" + poUp.HeadUrl + "', " 
	strSql += "RegIp= '" + poUp.RegIp + "', " 
	strSql += "RegDateTime= '" + poUp.RegDateTime + "', " 
	strSql += "RegDeviceType= '" + strconv.Itoa( poUp.RegDeviceType ) + "', " 
	strSql += "UsrTel= '" + poUp.UsrTel + "', " 
	strSql += "UsrEmail= '" + poUp.UsrEmail + "', " 
	strSql += "UsrICard= '" + poUp.UsrICard + "', " 
	strSql += "Res1= '" + poUp.Res1 + "', " 
	strSql += "Res2= '" + poUp.Res2 + "'" 
	strSql += "where Udid='" + strUdid + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//快速更新,通过HeadUrl，整个行，不检查
func FUpRowByHeadUrl( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strHeadUrl string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "Id= '" + strconv.Itoa( poUp.Id ) + "', " 
	strSql += "UsrName= '" + poUp.UsrName + "', " 
	strSql += "PassWord= '" + poUp.PassWord + "', " 
	strSql += "Nick= '" + poUp.Nick + "', " 
	strSql += "Udid= '" + poUp.Udid + "', " 
	strSql += "RegIp= '" + poUp.RegIp + "', " 
	strSql += "RegDateTime= '" + poUp.RegDateTime + "', " 
	strSql += "RegDeviceType= '" + strconv.Itoa( poUp.RegDeviceType ) + "', " 
	strSql += "UsrTel= '" + poUp.UsrTel + "', " 
	strSql += "UsrEmail= '" + poUp.UsrEmail + "', " 
	strSql += "UsrICard= '" + poUp.UsrICard + "', " 
	strSql += "Res1= '" + poUp.Res1 + "', " 
	strSql += "Res2= '" + poUp.Res2 + "'" 
	strSql += "where HeadUrl='" + strHeadUrl + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//快速更新,通过RegIp，整个行，不检查
func FUpRowByRegIp( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strRegIp string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "Id= '" + strconv.Itoa( poUp.Id ) + "', " 
	strSql += "UsrName= '" + poUp.UsrName + "', " 
	strSql += "PassWord= '" + poUp.PassWord + "', " 
	strSql += "Nick= '" + poUp.Nick + "', " 
	strSql += "Udid= '" + poUp.Udid + "', " 
	strSql += "HeadUrl= '" + poUp.HeadUrl + "', " 
	strSql += "RegDateTime= '" + poUp.RegDateTime + "', " 
	strSql += "RegDeviceType= '" + strconv.Itoa( poUp.RegDeviceType ) + "', " 
	strSql += "UsrTel= '" + poUp.UsrTel + "', " 
	strSql += "UsrEmail= '" + poUp.UsrEmail + "', " 
	strSql += "UsrICard= '" + poUp.UsrICard + "', " 
	strSql += "Res1= '" + poUp.Res1 + "', " 
	strSql += "Res2= '" + poUp.Res2 + "'" 
	strSql += "where RegIp='" + strRegIp + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//快速更新,通过RegDateTime，整个行，不检查
func FUpRowByRegDateTime( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strRegDateTime string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "Id= '" + strconv.Itoa( poUp.Id ) + "', " 
	strSql += "UsrName= '" + poUp.UsrName + "', " 
	strSql += "PassWord= '" + poUp.PassWord + "', " 
	strSql += "Nick= '" + poUp.Nick + "', " 
	strSql += "Udid= '" + poUp.Udid + "', " 
	strSql += "HeadUrl= '" + poUp.HeadUrl + "', " 
	strSql += "RegIp= '" + poUp.RegIp + "', " 
	strSql += "RegDeviceType= '" + strconv.Itoa( poUp.RegDeviceType ) + "', " 
	strSql += "UsrTel= '" + poUp.UsrTel + "', " 
	strSql += "UsrEmail= '" + poUp.UsrEmail + "', " 
	strSql += "UsrICard= '" + poUp.UsrICard + "', " 
	strSql += "Res1= '" + poUp.Res1 + "', " 
	strSql += "Res2= '" + poUp.Res2 + "'" 
	strSql += "where RegDateTime='" + strRegDateTime + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//快速更新,通过RegDeviceType，整个行，不检查
func FUpRowByRegDeviceType( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, iRegDeviceType int ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "Id= '" + strconv.Itoa( poUp.Id ) + "', " 
	strSql += "UsrName= '" + poUp.UsrName + "', " 
	strSql += "PassWord= '" + poUp.PassWord + "', " 
	strSql += "Nick= '" + poUp.Nick + "', " 
	strSql += "Udid= '" + poUp.Udid + "', " 
	strSql += "HeadUrl= '" + poUp.HeadUrl + "', " 
	strSql += "RegIp= '" + poUp.RegIp + "', " 
	strSql += "RegDateTime= '" + poUp.RegDateTime + "', " 
	strSql += "UsrTel= '" + poUp.UsrTel + "', " 
	strSql += "UsrEmail= '" + poUp.UsrEmail + "', " 
	strSql += "UsrICard= '" + poUp.UsrICard + "', " 
	strSql += "Res1= '" + poUp.Res1 + "', " 
	strSql += "Res2= '" + poUp.Res2 + "'" 
	strSql += "where RegDeviceType='" + strconv.Itoa( iRegDeviceType ) + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//快速更新,通过UsrTel，整个行，不检查
func FUpRowByUsrTel( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strUsrTel string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "Id= '" + strconv.Itoa( poUp.Id ) + "', " 
	strSql += "UsrName= '" + poUp.UsrName + "', " 
	strSql += "PassWord= '" + poUp.PassWord + "', " 
	strSql += "Nick= '" + poUp.Nick + "', " 
	strSql += "Udid= '" + poUp.Udid + "', " 
	strSql += "HeadUrl= '" + poUp.HeadUrl + "', " 
	strSql += "RegIp= '" + poUp.RegIp + "', " 
	strSql += "RegDateTime= '" + poUp.RegDateTime + "', " 
	strSql += "RegDeviceType= '" + strconv.Itoa( poUp.RegDeviceType ) + "', " 
	strSql += "UsrEmail= '" + poUp.UsrEmail + "', " 
	strSql += "UsrICard= '" + poUp.UsrICard + "', " 
	strSql += "Res1= '" + poUp.Res1 + "', " 
	strSql += "Res2= '" + poUp.Res2 + "'" 
	strSql += "where UsrTel='" + strUsrTel + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//快速更新,通过UsrEmail，整个行，不检查
func FUpRowByUsrEmail( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strUsrEmail string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "Id= '" + strconv.Itoa( poUp.Id ) + "', " 
	strSql += "UsrName= '" + poUp.UsrName + "', " 
	strSql += "PassWord= '" + poUp.PassWord + "', " 
	strSql += "Nick= '" + poUp.Nick + "', " 
	strSql += "Udid= '" + poUp.Udid + "', " 
	strSql += "HeadUrl= '" + poUp.HeadUrl + "', " 
	strSql += "RegIp= '" + poUp.RegIp + "', " 
	strSql += "RegDateTime= '" + poUp.RegDateTime + "', " 
	strSql += "RegDeviceType= '" + strconv.Itoa( poUp.RegDeviceType ) + "', " 
	strSql += "UsrTel= '" + poUp.UsrTel + "', " 
	strSql += "UsrICard= '" + poUp.UsrICard + "', " 
	strSql += "Res1= '" + poUp.Res1 + "', " 
	strSql += "Res2= '" + poUp.Res2 + "'" 
	strSql += "where UsrEmail='" + strUsrEmail + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//快速更新,通过UsrICard，整个行，不检查
func FUpRowByUsrICard( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strUsrICard string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "Id= '" + strconv.Itoa( poUp.Id ) + "', " 
	strSql += "UsrName= '" + poUp.UsrName + "', " 
	strSql += "PassWord= '" + poUp.PassWord + "', " 
	strSql += "Nick= '" + poUp.Nick + "', " 
	strSql += "Udid= '" + poUp.Udid + "', " 
	strSql += "HeadUrl= '" + poUp.HeadUrl + "', " 
	strSql += "RegIp= '" + poUp.RegIp + "', " 
	strSql += "RegDateTime= '" + poUp.RegDateTime + "', " 
	strSql += "RegDeviceType= '" + strconv.Itoa( poUp.RegDeviceType ) + "', " 
	strSql += "UsrTel= '" + poUp.UsrTel + "', " 
	strSql += "UsrEmail= '" + poUp.UsrEmail + "', " 
	strSql += "Res1= '" + poUp.Res1 + "', " 
	strSql += "Res2= '" + poUp.Res2 + "'" 
	strSql += "where UsrICard='" + strUsrICard + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//快速更新,通过Res1，整个行，不检查
func FUpRowByRes1( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strRes1 string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "Id= '" + strconv.Itoa( poUp.Id ) + "', " 
	strSql += "UsrName= '" + poUp.UsrName + "', " 
	strSql += "PassWord= '" + poUp.PassWord + "', " 
	strSql += "Nick= '" + poUp.Nick + "', " 
	strSql += "Udid= '" + poUp.Udid + "', " 
	strSql += "HeadUrl= '" + poUp.HeadUrl + "', " 
	strSql += "RegIp= '" + poUp.RegIp + "', " 
	strSql += "RegDateTime= '" + poUp.RegDateTime + "', " 
	strSql += "RegDeviceType= '" + strconv.Itoa( poUp.RegDeviceType ) + "', " 
	strSql += "UsrTel= '" + poUp.UsrTel + "', " 
	strSql += "UsrEmail= '" + poUp.UsrEmail + "', " 
	strSql += "UsrICard= '" + poUp.UsrICard + "', " 
	strSql += "Res2= '" + poUp.Res2 + "'" 
	strSql += "where Res1='" + strRes1 + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//快速更新,通过Res2，整个行，不检查
func FUpRowByRes2( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strRes2 string ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "Id= '" + strconv.Itoa( poUp.Id ) + "', " 
	strSql += "UsrName= '" + poUp.UsrName + "', " 
	strSql += "PassWord= '" + poUp.PassWord + "', " 
	strSql += "Nick= '" + poUp.Nick + "', " 
	strSql += "Udid= '" + poUp.Udid + "', " 
	strSql += "HeadUrl= '" + poUp.HeadUrl + "', " 
	strSql += "RegIp= '" + poUp.RegIp + "', " 
	strSql += "RegDateTime= '" + poUp.RegDateTime + "', " 
	strSql += "RegDeviceType= '" + strconv.Itoa( poUp.RegDeviceType ) + "', " 
	strSql += "UsrTel= '" + poUp.UsrTel + "', " 
	strSql += "UsrEmail= '" + poUp.UsrEmail + "', " 
	strSql += "UsrICard= '" + poUp.UsrICard + "', " 
	strSql += "Res1= '" + poUp.Res1 + "'" 
	strSql += "where Res2='" + strRes2 + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}

//通过主key，快速更新，某一行中的某一列
//如果要更新多列，还需要增加接口
//快速更新,通过主key，单cell，不检查
func FUpUsrNameByPriKey( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strUsrName string, iId int ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "UsrName='" + poUp.UsrName + "'" 
	strSql += "where Id= '" + strconv.Itoa( iId ) + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}
//通过主key，快速更新，某一行中的某一列
//如果要更新多列，还需要增加接口
//快速更新,通过主key，单cell，不检查
func FUpPassWordByPriKey( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strPassWord string, iId int ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "PassWord='" + poUp.PassWord + "'" 
	strSql += "where Id= '" + strconv.Itoa( iId ) + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}
//通过主key，快速更新，某一行中的某一列
//如果要更新多列，还需要增加接口
//快速更新,通过主key，单cell，不检查
func FUpNickByPriKey( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strNick string, iId int ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "Nick='" + poUp.Nick + "'" 
	strSql += "where Id= '" + strconv.Itoa( iId ) + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}
//通过主key，快速更新，某一行中的某一列
//如果要更新多列，还需要增加接口
//快速更新,通过主key，单cell，不检查
func FUpUdidByPriKey( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strUdid string, iId int ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "Udid='" + poUp.Udid + "'" 
	strSql += "where Id= '" + strconv.Itoa( iId ) + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}
//通过主key，快速更新，某一行中的某一列
//如果要更新多列，还需要增加接口
//快速更新,通过主key，单cell，不检查
func FUpHeadUrlByPriKey( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strHeadUrl string, iId int ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "HeadUrl='" + poUp.HeadUrl + "'" 
	strSql += "where Id= '" + strconv.Itoa( iId ) + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}
//通过主key，快速更新，某一行中的某一列
//如果要更新多列，还需要增加接口
//快速更新,通过主key，单cell，不检查
func FUpRegIpByPriKey( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strRegIp string, iId int ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "RegIp='" + poUp.RegIp + "'" 
	strSql += "where Id= '" + strconv.Itoa( iId ) + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}
//通过主key，快速更新，某一行中的某一列
//如果要更新多列，还需要增加接口
//快速更新,通过主key，单cell，不检查
func FUpRegDateTimeByPriKey( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strRegDateTime string, iId int ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "RegDateTime='" + poUp.RegDateTime + "'" 
	strSql += "where Id= '" + strconv.Itoa( iId ) + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}
//通过主key，快速更新，某一行中的某一列
//如果要更新多列，还需要增加接口
//快速更新,通过主key，单cell，不检查
func FUpRegDeviceTypeByPriKey( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, iRegDeviceType int, iId int ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "RegDeviceType='" + strconv.Itoa( poUp.RegDeviceType ) + "'" 
	strSql += "where Id= '" + strconv.Itoa( iId ) + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}
//通过主key，快速更新，某一行中的某一列
//如果要更新多列，还需要增加接口
//快速更新,通过主key，单cell，不检查
func FUpUsrTelByPriKey( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strUsrTel string, iId int ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "UsrTel='" + poUp.UsrTel + "'" 
	strSql += "where Id= '" + strconv.Itoa( iId ) + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}
//通过主key，快速更新，某一行中的某一列
//如果要更新多列，还需要增加接口
//快速更新,通过主key，单cell，不检查
func FUpUsrEmailByPriKey( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strUsrEmail string, iId int ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "UsrEmail='" + poUp.UsrEmail + "'" 
	strSql += "where Id= '" + strconv.Itoa( iId ) + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}
//通过主key，快速更新，某一行中的某一列
//如果要更新多列，还需要增加接口
//快速更新,通过主key，单cell，不检查
func FUpUsrICardByPriKey( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strUsrICard string, iId int ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "UsrICard='" + poUp.UsrICard + "'" 
	strSql += "where Id= '" + strconv.Itoa( iId ) + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}
//通过主key，快速更新，某一行中的某一列
//如果要更新多列，还需要增加接口
//快速更新,通过主key，单cell，不检查
func FUpRes1ByPriKey( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strRes1 string, iId int ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "Res1='" + poUp.Res1 + "'" 
	strSql += "where Id= '" + strconv.Itoa( iId ) + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}
//通过主key，快速更新，某一行中的某一列
//如果要更新多列，还需要增加接口
//快速更新,通过主key，单cell，不检查
func FUpRes2ByPriKey( strTblName string, poDb *sql.DB, poUp *CCG_LoginDb, strRes2 string, iId int ) bool { 
	strSql := "update " + strTblName + " set "
	strSql += "Res2='" + poUp.Res2 + "'" 
	strSql += "where Id= '" + strconv.Itoa( iId ) + "'" 
	_, _e := poDb.Exec(strSql)
	if _e != nil {
		glog.Errorln("sql:" + strSql + ", err:" + _e.Error() )
		return false
	}
	return true
}
