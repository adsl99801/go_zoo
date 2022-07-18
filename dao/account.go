package dao

import (
	"context"
	"time"

	"github.com/adsl99801/zoo/model"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// GetAllAccount is a function to get a slice of record(s) from account table in the pf database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllAccount(ctx context.Context, page, pagesize int64, order string) (results []*model.Account, totalRows int, err error) {

	resultOrm := DB.Model(&model.Account{})
	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(offset).Limit(pagesize)
	} else {
		resultOrm = resultOrm.Limit(pagesize)
	}

	if order != "" {
		resultOrm = resultOrm.Order(order)
	}

	if err = resultOrm.Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return results, totalRows, nil
}

// GetAccount is a function to get a single record from the account table in the pf database
// error - ErrNotFound, db Find error
func GetAccount(ctx context.Context, argID int32) (record *model.Account, err error) {
	record = &model.Account{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddAccount is a function to add a single record to account table in the pf database
// error - ErrInsertFailed, db save call failed
func AddAccount(ctx context.Context, record *model.Account) (result *model.Account, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateAccount is a function to update a single record from account table in the pf database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateAccount(ctx context.Context, argID int32, updated *model.Account) (result *model.Account, RowsAffected int64, err error) {

	result = &model.Account{}
	db := DB.First(result, argID)
	if err = db.Error; err != nil {
		return nil, -1, ErrNotFound
	}

	if err = Copy(result, updated); err != nil {
		return nil, -1, ErrUpdateFailed
	}

	db = db.Save(result)
	if err = db.Error; err != nil {
		return nil, -1, ErrUpdateFailed
	}

	return result, db.RowsAffected, nil
}

// DeleteAccount is a function to delete a single record from account table in the pf database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteAccount(ctx context.Context, argID int32) (rowsAffected int64, err error) {

	record := &model.Account{}
	db := DB.First(record, argID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
