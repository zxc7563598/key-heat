package storage

import (
	"fmt"
	"os"
	"path/filepath"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	sqlite "github.com/Tryanks/gorm-sqlite"
)

// DB 封装 GORM 数据库连接
type DB struct {
	*gorm.DB
	path string
}

// Config 数据库配置
type Config struct {
	// 数据库文件路径（如果不提供，将使用默认路径）
	Path string
	// 是否启用外键约束（默认启用）
	EnableForeignKey bool
	// 日志级别（默认 Silent，可改为 Info 用于调试）
	LogLevel logger.LogLevel
	// 是否跳过默认事务以提高性能（默认启用）
	SkipDefaultTransaction bool
}

// NewDB 创建新的数据库连接
func NewDB(cfg *Config) (*DB, error) {
	if cfg == nil {
		cfg = DefaultConfig()
	}

	// 确定数据库路径
	dbPath := cfg.Path
	if dbPath == "" {
		var err error
		dbPath, err = getDefaultDBPath()
		if err != nil {
			return nil, fmt.Errorf("获取默认数据库路径失败: %w", err)
		}
	}

	// 确保数据库目录存在
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("创建数据库目录失败: %w", err)
	}

	// 构建 DSN（数据源名称）
	// 使用纯 Go SQLite 驱动，无需 CGO
	dsn := dbPath

	// 如果需要启用外键约束
	if cfg.EnableForeignKey {
		// 纯 Go 驱动通过 _pragma 参数设置外键
		dsn = fmt.Sprintf("%s?_pragma=foreign_keys(1)", dbPath)
	}

	// 打开数据库连接
	gormDB, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(cfg.LogLevel),
		SkipDefaultTransaction: cfg.SkipDefaultTransaction,
	})
	if err != nil {
		return nil, fmt.Errorf("打开数据库失败: %w", err)
	}

	// 获取底层的 sql.DB 以设置连接池（可选）
	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, fmt.Errorf("获取底层数据库连接失败: %w", err)
	}

	// 设置连接池参数（SQLite 通常只允许单写，这里保守设置）
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetMaxIdleConns(1)

	return &DB{
		DB:   gormDB,
		path: dbPath,
	}, nil
}

// Close 关闭数据库连接
func (db *DB) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// GetPath 获取当前数据库文件路径
func (db *DB) GetPath() string {
	return db.path
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		EnableForeignKey:       true,
		LogLevel:               logger.Silent,
		SkipDefaultTransaction: true,
	}
}

// getDefaultDBPath 获取默认的数据库路径
func getDefaultDBPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("获取用户配置目录失败: %w", err)
	}

	// 跨平台应用数据目录
	// Windows: %APPDATA%/KeyHeat/
	// macOS:   ~/Library/Application Support/KeyHeat/
	// Linux:   ~/.config/KeyHeat/
	appDir := filepath.Join(configDir, "KeyHeat")
	dbPath := filepath.Join(appDir, "keyheat.db")

	return dbPath, nil
}
