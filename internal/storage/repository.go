package storage

import (
	"fmt"

	"github.com/zxc7563598/key-heat/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type gormRepo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &gormRepo{
		db: db,
	}
}

type Repository interface {
	GetStatsByDate(date string) (map[string]int, error)
	GetStatsByDateRange(startDate, endDate string) (map[string]int, error)
	GetKeyHistory(keyName, startDate, endDate string) (map[string]int, error)
	UpsertKeyCount(date, keyName string, delta int) error
}

// GetStatsByDate 获取指定日期的所有按键统计
// 返回 map[string]int，key 为按键名，value 为点击次数
func (r *gormRepo) GetStatsByDate(date string) (map[string]int, error) {
	var stats []model.KeyEvent
	err := r.db.Where("date = ?", date).Find(&stats).Error
	if err != nil {
		return nil, fmt.Errorf("查询日期 %s 的统计失败: %w", date, err)
	}
	result := make(map[string]int, len(stats))
	for _, stat := range stats {
		result[stat.KeyName] = stat.Count
	}
	return result, nil
}

// GetStatsByDateRange 获取日期范围内的所有按键统计
// startDate 和 endDate 格式: "2006-01-02"
func (r *gormRepo) GetStatsByDateRange(startDate, endDate string) (map[string]int, error) {
	var stats []model.KeyEvent
	err := r.db.Where("date BETWEEN ? AND ?", startDate, endDate).
		Find(&stats).Error
	if err != nil {
		return nil, fmt.Errorf("查询日期范围 %s ~ %s 的统计失败: %w", startDate, endDate, err)
	}
	// 合并统计：同一按键在日期范围内的总次数
	result := make(map[string]int)
	for _, stat := range stats {
		result[stat.KeyName] += stat.Count
	}
	return result, nil
}

// GetKeyHistory 获取单个按键在日期范围内的历史数据
// 返回 map[string]int，key 为日期，value 为当日次数
func (r *gormRepo) GetKeyHistory(keyName, startDate, endDate string) (map[string]int, error) {
	var stats []model.KeyEvent
	err := r.db.Where("key_name = ? AND date BETWEEN ? AND ?", keyName, startDate, endDate).
		Order("date ASC").
		Find(&stats).Error
	if err != nil {
		return nil, fmt.Errorf("查询按键 %s 的历史失败: %w", keyName, err)
	}
	result := make(map[string]int, len(stats))
	for _, stat := range stats {
		result[stat.Date] = stat.Count
	}
	return result, nil
}

// UpsertKeyCount 更新或插入统计
func (r *gormRepo) UpsertKeyCount(date, keyName string, delta int) error {
	stat := model.KeyEvent{
		Date:    date,
		KeyName: keyName,
		Count:   delta,
	}
	// Clauses 实现 SQLite 的 INSERT ... ON CONFLICT DO UPDATE
	return r.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "date"}, {Name: "key_name"}},
		DoUpdates: clause.Assignments(map[string]any{
			"count": gorm.Expr("count + ?", delta),
		}),
	}).Create(&stat).Error
}
