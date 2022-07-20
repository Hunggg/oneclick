package cockroachdb

import (
	"oneclick/entity"
	"reflect"
	"testing"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func TestCockroachDB_SaveCategory(t *testing.T) {
	type fields struct {
		l  *zap.SugaredLogger
		db *gorm.DB
	}
	type args struct {
		data entity.Categories
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CockroachDB{
				l:  tt.fields.l,
				db: tt.fields.db,
			}
			if err := c.SaveCategory(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("CockroachDB.SaveCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCockroachDB_SaveBatchCategory(t *testing.T) {
	type fields struct {
		l  *zap.SugaredLogger
		db *gorm.DB
	}
	type args struct {
		listCategory []entity.Categories
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CockroachDB{
				l:  tt.fields.l,
				db: tt.fields.db,
			}
			if err := c.SaveBatchCategory(tt.args.listCategory); (err != nil) != tt.wantErr {
				t.Errorf("CockroachDB.SaveBatchCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCockroachDB_GetCategoryById(t *testing.T) {
	type fields struct {
		l  *zap.SugaredLogger
		db *gorm.DB
	}
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Categories
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CockroachDB{
				l:  tt.fields.l,
				db: tt.fields.db,
			}
			got, err := c.GetCategoryById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CockroachDB.GetCategoryById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CockroachDB.GetCategoryById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCockroachDB_GetListCategory(t *testing.T) {
	type fields struct {
		l  *zap.SugaredLogger
		db *gorm.DB
	}
	type args struct {
		ofset int
		limit int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Categories
		want1   uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CockroachDB{
				l:  tt.fields.l,
				db: tt.fields.db,
			}
			got, got1, err := c.GetListCategory(tt.args.ofset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("CockroachDB.GetListCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CockroachDB.GetListCategory() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CockroachDB.GetListCategory() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCockroachDB_DeleteCategoryById(t *testing.T) {
	type fields struct {
		l  *zap.SugaredLogger
		db *gorm.DB
	}
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CockroachDB{
				l:  tt.fields.l,
				db: tt.fields.db,
			}
			if err := c.DeleteCategoryById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("CockroachDB.DeleteCategoryById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCockroachDB_UpdateCategory(t *testing.T) {
	type fields struct {
		l  *zap.SugaredLogger
		db *gorm.DB
	}
	type args struct {
		data entity.Categories
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CockroachDB{
				l:  tt.fields.l,
				db: tt.fields.db,
			}
			if err := c.UpdateCategory(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("CockroachDB.UpdateCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
