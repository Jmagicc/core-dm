package config

import (
	"gitee.com/chunanyong/zorm"
)

//dbDao 代表一个数据库,如果有多个数据库,就对应声明多个DBDao
var DBDao *zorm.DBDao

//// ctx默认应该有 web层传入,例如gin的c.Request.Context().这里只是模拟
//var ctx = context.Background()

// InitZOrmDao	01.初始化DBDao
func InitZOrmDao(DSN_STR string) (err error) {

	//自定义zorm日志输出
	zorm.LogCallDepth = 4 //日志调用的层级

	////记录异常日志的函数
	//zorm.FuncLogError = func(err error) {
	//	//log.Error(err)
	//	logs.Println(err)
	//}
	//
	////记录panic日志,默认使用defaultLogError实现
	//zorm.FuncLogPanic = func(err error) {
	//	logs.Error(err)
	//}
	//
	////打印sql的函数
	//zorm.FuncPrintSQL = func(sqlstr string, args []interface{}) {
	//	logs.Debug(sqlstr, args)
	//}

	//dbDaoConfig 数据库的配置.这里只是模拟,生产应该是读取配置配置文件,构造DataSourceConfig
	dbDaoConfig := zorm.DataSourceConfig{
		//DSN 数据库的连接字符串
		DSN: DSN_STR,
		//数据库驱动名称:mysql,postgres,oci8,sqlserver,sqlite3,clickhouse,dm,kingbase,aci 和DBType对应,处理数据库有多个驱动
		DriverName: "dm",
		//数据库类型(方言判断依据):mysql,postgresql,oracle,mssql,sqlite,clickhouse,dm,kingbase,shentong 和 DriverName 对应,处理数据库有多个驱动
		DBType: "dm",
		//MaxOpenConns 数据库最大连接数 默认50
		MaxOpenConns: 50,
		//MaxIdleConns 数据库最大空闲连接数 默认50
		MaxIdleConns: 50,
		//ConnMaxLifetimeSecond 连接存活秒时间. 默认600(10分钟)后连接被销毁重建.避免数据库主动断开连接,造成死连接.MySQL默认wait_timeout 28800秒(8小时)
		ConnMaxLifetimeSecond: 600,
		//PrintSQL 打印SQL.会使用FuncPrintSQL记录SQL
		//PrintSQL: false,
		//DefaultTxOptions 事务隔离级别的默认配置,默认为nil
		//DefaultTxOptions: nil,
		//如果是使用seata-golang分布式事务,建议使用默认配置
		//DefaultTxOptions: &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: false},

		//FuncSeataGlobalTransaction seata-golang分布式的适配函数,返回ISeataGlobalTransaction接口的实现
		//FuncSeataGlobalTransaction : MyFuncSeataGlobalTransaction,
	}

	// 根据dbDaoConfig创建dbDao, 一个数据库只执行一次,第一个执行的数据库为 defaultDao,后续zorm.xxx方法,默认使用的就是defaultDao
	DBDao, err = zorm.NewDBDao(&dbDaoConfig)

	return
}
