package util

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

// MySQLConfig 包含MySQL连接的配置选项
type MySQLConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

// NewMySQLConfig 创建一个MySQLConfig对象，并设置默认值
func NewMySQLConfig() *MySQLConfig {
	return &MySQLConfig{
		User:     "root",   // 默认用户名
		Password: "123456", // 默认密码
		Host:     "127.0.0.1",
		Port:     "3306",
		Database: "database",
	}
}

// WithUser 设置MySQLConfig的用户名
func (cfg *MySQLConfig) WithUser(user string) *MySQLConfig {
	cfg.User = user
	return cfg
}

// WithPassword 设置MySQLConfig的密码
func (cfg *MySQLConfig) WithPassword(password string) *MySQLConfig {
	cfg.Password = password
	return cfg
}

// WithHost 设置MySQLConfig的主机地址
func (cfg *MySQLConfig) WithHost(host string) *MySQLConfig {
	cfg.Host = host
	return cfg
}

// WithPort 设置MySQLConfig的端口号
func (cfg *MySQLConfig) WithPort(port string) *MySQLConfig {
	cfg.Port = port
	return cfg
}

// WithDatabase 设置MySQLConfig的数据库名
func (cfg *MySQLConfig) WithDatabase(database string) *MySQLConfig {
	cfg.Database = database
	return cfg
}

// init函数在包被导入时自动执行
func init() {
	// 使用options模式初始化MySQL连接
	cfg := NewMySQLConfig().WithUser("root").WithPassword("123456").WithDatabase("short_url")
	// 构建数据源名称
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("无法连接到数据库:", err)
	}

	// 尝试连接数据库
	err = db.Ping()
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	fmt.Println("数据库连接成功！")
}

// GetDB 返回数据库实例
func GetDB() *sql.DB {
	return db
}
