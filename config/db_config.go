package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

var db *sql.DB

// MySQLConfig 包含 MySQL 连接的配置选项
type MySQLConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

// NewMySQLConfig 创建一个 MySQLConfig 对象，并设置默认值
func NewMySQLConfig() *MySQLConfig {
	return &MySQLConfig{
		User:     "root",   // 默认用户名
		Password: "123456", // 默认密码
		Host:     "127.0.0.1",
		Port:     "3306",
		Database: "database",
	}
}

// WithUser 设置 MySQLConfig 的用户名
func (cfg *MySQLConfig) WithUser(user string) *MySQLConfig {
	cfg.User = user
	return cfg
}

// WithPassword 设置 MySQLConfig 的密码
func (cfg *MySQLConfig) WithPassword(password string) *MySQLConfig {
	cfg.Password = password
	return cfg
}

// WithHost 设置 MySQLConfig 的主机地址
func (cfg *MySQLConfig) WithHost(host string) *MySQLConfig {
	cfg.Host = host
	return cfg
}

// WithPort 设置 MySQLConfig 的端口号
func (cfg *MySQLConfig) WithPort(port string) *MySQLConfig {
	cfg.Port = port
	return cfg
}

// WithDatabase 设置 MySQLConfig 的数据库名
func (cfg *MySQLConfig) WithDatabase(database string) *MySQLConfig {
	cfg.Database = database
	return cfg
}

// InitDB 函数手动控制初始化顺序
func InitDB() {
	// 使用 options 模式初始化 MySQL 连接
	cfg := NewMySQLConfig().WithUser("root").WithPassword("123456").WithDatabase("short_url")
	// 构建数据源名称
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		SugarLogger.Error(" 无法连接到数据库:", zap.Error(err))
	}

	// 尝试连接数据库
	err = db.Ping()
	if err != nil {
		SugarLogger.Error(" 数据库连接失败:", zap.Error(err))
	}

	SugarLogger.Info(" 数据库连接成功！")
}

// GetDB 返回数据库实例
func GetDB() *sql.DB {
	return db
}
