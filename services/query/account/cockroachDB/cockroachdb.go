package cockroachdb

import (
	"context"
	"html"
	"oneclick/config"
	"oneclick/entity"
	"strings"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


type CockroachDB struct {
	l  *zap.SugaredLogger
	db *gorm.DB
}

func NewCockroachDB(db *gorm.DB) (*CockroachDB, error) {
	return &CockroachDB{
		l:  zap.S(),
		db: db,
	}, nil
}


func (c *CockroachDB) SaveAccount(data entity.Accounts) error {
	ctx, cancel := context.WithTimeout(context.Background(), config.DefaultCockroachDbTimeout)
	defer cancel()
	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return saveAccount(tx, data)
		},
	); err != nil {
		c.l.Errorw("error save account ", "error", err)
		return err
	}
	return nil
}

func saveAccount(db *gorm.DB, data entity.Accounts) error {
	hashAccount, err := beforeSaveAccount(data)
	if err != nil {
		return err
	}
	return db.Clauses(clause.OnConflict{DoNothing: true}).FirstOrCreate(&hashAccount, entity.Accounts{Name: data.Name}).Error
}

func beforeSaveAccount(a entity.Accounts) (entity.Accounts, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)

	if err != nil {
		return entity.Accounts{}, err
	}

	a.Password = string(hashedPassword)
	a.Name = html.EscapeString(strings.TrimSpace(a.Name))
	return a, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// func loginCheck(db *gorm.DB, user model.AuthLog) (string, error) {
// 	var result entity.Accounts
// 	ctx, cancel := context.WithTimeout(context.Background(), config.DefaultCockroachDbTimeout)
// 	defer cancel()
// 	if err := db.WithContext(ctx).Where("name", user.Name).Find(&result).Error; err != nil {
// 		return "", err
// 	}
// 	if err := VerifyPassword(user.Password, result.Password); err != nil {
// 		return "", err
// 	}
// 	token, err := middleware.GenerateToken(result.ID)
// 	if err != nil {
// 		return "", err
// 	}

// 	return token, nil
// }

// func (c *CockroachDB) Login(user model.AuthLog) (string, error) {
// 	token, err := loginCheck(c.db, user)
// 	if err != nil {
// 		return "", err
// 	}
// 	return token, nil
// }
