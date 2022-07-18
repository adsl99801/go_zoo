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

// GetAllWeather is a function to get a slice of record(s) from weather table in the pf database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllWeather(ctx context.Context, page, pagesize int64, order string) (results []*model.Weather, totalRows int, err error) {

	resultOrm := DB.Model(&model.Weather{})
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

// GetWeather is a function to get a single record from the weather table in the pf database
// error - ErrNotFound, db Find error
func GetWeather(ctx context.Context, argCity string) (record *model.Weather, err error) {
	record = &model.Weather{}
	if err = DB.First(record, argCity).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddWeather is a function to add a single record to weather table in the pf database
// error - ErrInsertFailed, db save call failed
func AddWeather(ctx context.Context, record *model.Weather) (result *model.Weather, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateWeather is a function to update a single record from weather table in the pf database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateWeather(ctx context.Context, argCity string, updated *model.Weather) (result *model.Weather, RowsAffected int64, err error) {

	result = &model.Weather{}
	db := DB.First(result, argCity)
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

// DeleteWeather is a function to delete a single record from weather table in the pf database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteWeather(ctx context.Context, argCity string) (rowsAffected int64, err error) {

	record := &model.Weather{}
	db := DB.First(record, argCity)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
