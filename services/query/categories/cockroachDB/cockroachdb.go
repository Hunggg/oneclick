package cockroachdb

import (
	"context"
	"oneclick/config"
	"oneclick/entity"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CockroachDB struct {
	l *zap.SugaredLogger
	db *gorm.DB
}

func NewCockroachDB(db *gorm.DB) (*CockroachDB, error) {
	return &CockroachDB{
		l: zap.S(),
		db: db,
	}, nil
}


func (c *CockroachDB) SaveCategory(data entity.Categories) error {
	ctx, cancel := context.WithTimeout(context.Background(), config.DefaultCockroachDbTimeout)
	defer cancel()
	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return saveCategory(tx, data)
		},
	); err != nil {
		c.l.Errorw("error save wearable on chain data", "error", err)
		return err
	}
	return nil
}

func saveCategory(db *gorm.DB, data entity.Categories) error {
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&data).Error
}

func (c *CockroachDB) GetCategoryById(id uint64) (entity.Categories, error) {
	var result entity.Categories
	ctx, cancel := context.WithTimeout(context.Background(), config.DefaultCockroachDbTimeout)
	defer cancel()
	if err := c.db.WithContext(ctx).Where("category_id", id).Find(&result).Error; err != nil {
		c.l.Errorw("error get information of category", "error", err, "category_id", id)
		return entity.Categories{}, err
	}
	return result, nil
}
