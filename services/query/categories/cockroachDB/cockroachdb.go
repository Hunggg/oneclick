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
		c.l.Errorw("error save categories ", "error", err)
		return err
	}
	return nil
}

func saveCategory(db *gorm.DB, data entity.Categories) error {
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&data).Error
}

func (c *CockroachDB) SaveBatchCategory(listCategory []entity.Categories) error {
	ctx, cancel := context.WithTimeout(context.Background(), config.DefaultCockroachDbTimeout)
	defer cancel()
	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return saveBatchCategory(c.db, listCategory)
		},
	); err != nil {
		c.l.Errorw("error save categories ", "error", err)
		return err
	}
	return nil
}

func saveBatchCategory(db *gorm.DB, listCategories []entity.Categories)  error {
	return db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(listCategories, 100).Error
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

func (c *CockroachDB) GetListCategory(ofset int, limit int) ([]entity.Categories, uint64, error) {
	return(getListCategory(c.db, ofset, limit))
}

func getListCategory(db *gorm.DB, offset int, limit int) ([]entity.Categories, uint64, error) {
	var result []entity.Categories
	var count int64
	ctx, cancel := context.WithTimeout(context.Background(), config.DefaultCockroachDbTimeout)
	defer cancel()
	if err := db.WithContext(ctx).Offset(offset).Limit(limit).Find(&result).Error; err != nil {
		return []entity.Categories{}, 0, err
	}

	if err := db.WithContext(ctx).Table("categories").Count(&count).Error; err != nil {
		return []entity.Categories{}, 0, err
	}

	return result, uint64(count), nil
}

func (c *CockroachDB) DeleteCategoryById(id uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), config.DefaultCockroachDbTimeout)
	defer cancel()
	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return deleteCategoryById(tx, id)
		},
	); err != nil {
		c.l.Errorw("error save categories ", "error", err)
		return err
	}
	return nil
}

func deleteCategoryById(db *gorm.DB, id uint64) error {
	return db.Delete(&entity.Categories{}, id).Error
}

