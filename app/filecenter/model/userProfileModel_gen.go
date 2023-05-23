// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type (
	userProfileModel interface {
		Insert(ctx context.Context, data *UserProfile) (int64, error)                                // 返回的是受影响行数，如需获取自增id，请通过data参数获取
		TxInsert(ctx context.Context, t *gorm.DB, data *UserProfile, picture []byte) (string, error) // 用于事务新增数据，由上层(如rpc层)调用Transaction去实现，返回受影响的行数

		FindOne(ctx context.Context, id int64) (*UserProfile, error)                                                 // 通过主键id查找数据
		Finds(ctx context.Context, queries []Query, orders []Order) ([]*UserProfile, error)                          // 按条件查询，不分页
		FindsByPage(ctx context.Context, queries []Query, page *Page, orders []Order) ([]*UserProfile, int64, error) // 分页查询
		FindCount(ctx context.Context, queries []Query) (int64, error)

		Update(ctx context.Context, id int64, data *UserProfile) (int64, error)                                    // 通过主键id更新指定字段，零值不可更新，返回受影响行数
		TxUpdate(ctx context.Context, tx *gorm.DB, id int64, data *UserProfile) (int64, error)                     // 用于事务更新数据，由上层(如rpc层)调用Transaction去实现，零值不可更新，返回受影响行数
		UpdateOneMapById(ctx context.Context, id int64, data map[string]interface{}) (int64, error)                // 通过主键id更新字段，map参数为数据库需要更新的字段，返回受影响行数
		TxUpdateOneMapById(ctx context.Context, tx *gorm.DB, id int64, data map[string]interface{}) (int64, error) // 用于事务更新字段，由上层(如rpc层)调用Transaction去实现，返回受影响的行数

		Delete(ctx context.Context, id int64) (int64, error)                    // 通过主键id删除数据，返回受影响行数
		TxDelete(ctx context.Context, tx *gorm.DB, id int64) (int64, error)     // 用于事务删除数据，由上层(如rpc层)调用Transaction去实现，返回受影响行数
		Deletes(ctx context.Context, ids []int64) (int64, error)                // 通过主键id批量删除数据，返回受影响行数
		TxDeletes(ctx context.Context, tx *gorm.DB, ids []int64) (int64, error) // 用于事务批量删除数据，由上层(如rpc层)调用Transaction去实现，返回受影响行数

	}

	defaultUserProfileModel struct {
		DB *gorm.DB
	}

	UserProfile struct {
		Id int64 `db:"id"`

		UserId int64 `db:"user_id"`

		Md5 string `db:"md5"`
	}
)

func newUserProfileModel(conn *gorm.DB) *defaultUserProfileModel {
	return &defaultUserProfileModel{
		DB: conn,
	}
}

// Insert 返回的是受影响行数，如需获取自增id，请通过data参数获取
func (d *defaultUserProfileModel) Insert(ctx context.Context, data *UserProfile) (int64, error) {
	logx.WithContext(ctx).Infof("insert data:%+v", data)

	if data == nil {
		logx.WithContext(ctx).Errorf("insert error:%+v", "param invalid")
		return 0, InputParamInvalid
	}

	result := d.DB.Debug().WithContext(ctx).Create(data)
	if result.Error != nil {
		logx.WithContext(ctx).Errorf("insert error:%+v", result.Error)
		return 0, WriteDataFailed
	}

	return result.RowsAffected, nil
}

// TxInsert // 用于事务新增数据，由上层(如rpc层)调用Transaction去实现，返回受影响的行数
// tx：上层传递时请加上context

func (d *defaultUserProfileModel) TxInsert(ctx context.Context, t *gorm.DB, data *UserProfile, picture []byte) (string, error) {
	logx.WithContext(ctx).Infof("txInsert data:%+v", data)

	if data == nil {
		logx.WithContext(ctx).Errorf("txInsert error:%+v", "param invalid")
		return "", InputParamInvalid
	}

	tx := t.Begin() // 开启事务

	result := tx.Debug().WithContext(ctx).Create(data)
	if result.Error != nil {
		logx.WithContext(ctx).Errorf("txInsert error:%+v", result.Error)
		return "", WriteDataFailed
	}

	err := tx.Debug().WithContext(ctx).Where("`md5` = ?", data.Md5).First(&ProfilePicture{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logx.WithContext(ctx).Errorf("findOne error:%+v", err)
		return "", ReadDataFailed
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		result = tx.Debug().WithContext(ctx).Create(&ProfilePicture{
			Md5:     data.Md5,
			Picture: picture,
		})

		if result.Error != nil {
			logx.WithContext(ctx).Errorf("txInsert error:%+v", result.Error)
			return "", WriteDataFailed
		}
	}

	tx.Commit()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	return "xxxxx", nil
}

// Update 通过主键id更新指定字段，零值不可更新，返回受影响行数
func (d *defaultUserProfileModel) Update(ctx context.Context, id int64, data *UserProfile) (int64, error) {
	logx.WithContext(ctx).Infof("update data:%+v", data)

	if id <= 0 {
		logx.WithContext(ctx).Errorf("update error:%+v", "param invalid")
		return 0, InputParamInvalid
	}

	result := d.DB.Debug().WithContext(ctx).Where("`id` = ?", id).Updates(data)
	if result.Error != nil {
		logx.WithContext(ctx).Errorf("update error:%+v", result.Error)
		return 0, WriteDataFailed
	}

	return result.RowsAffected, nil
}

// TxUpdate 用于事务更新数据，由上层(如rpc层)调用Transaction去实现，零值不可更新，返回受影响行数
func (d *defaultUserProfileModel) TxUpdate(ctx context.Context, tx *gorm.DB, id int64, data *UserProfile) (int64, error) {
	logx.WithContext(ctx).Infof("txUpdate data:%+v", data)

	if id <= 0 {
		logx.WithContext(ctx).Errorf("txUpdate error:%+v", "param invalid")
		return 0, InputParamInvalid
	}

	result := tx.Debug().WithContext(ctx).Where("`id` = ?", id).Updates(data)
	if result.Error != nil {
		logx.WithContext(ctx).Errorf("update error:%+v", result.Error)
		return 0, WriteDataFailed
	}

	return result.RowsAffected, nil
}

// UpdateOneMapById 通过主键id更新字段，map参数为数据库需要更新的字段，返回受影响行数
func (d *defaultUserProfileModel) UpdateOneMapById(ctx context.Context, id int64, data map[string]interface{}) (int64, error) {
	logx.WithContext(ctx).Infof("updateOneMapById data:%+v", data)

	if id <= 0 {
		logx.WithContext(ctx).Errorf("updateOneMapById error:%+v", "param invalid")
		return 0, InputParamInvalid
	}

	result := d.DB.Debug().WithContext(ctx).Model(&UserProfile{}).Where("`id` = ?", id).Updates(data)
	if result.Error != nil {
		logx.WithContext(ctx).Errorf("updateOneMapById error:%+v", result.Error)
		return 0, WriteDataFailed
	}

	return result.RowsAffected, nil
}

// TxUpdateOneMapById 用于事务更新字段，由上层(如rpc层)调用Transaction去实现，返回受影响的行数
// tx：上层传递时请加上context
func (d *defaultUserProfileModel) TxUpdateOneMapById(ctx context.Context, tx *gorm.DB, id int64, data map[string]interface{}) (int64, error) {
	logx.WithContext(ctx).Infof("txUpdateOneMapById data:%+v", data)

	if id <= 0 {
		logx.WithContext(ctx).Errorf("txUpdateOneMapById error:%+v", "param invalid")
		return 0, InputParamInvalid
	}

	result := tx.Debug().WithContext(ctx).Model(&UserProfile{}).Where("`id` = ?", id).Updates(data)
	if result.Error != nil {
		logx.WithContext(ctx).Errorf("txUpdateOneMapById error:%+v", result.Error)
		return 0, WriteDataFailed
	}

	return result.RowsAffected, nil
}

// Delete 通过主键id删除数据，返回受影响行数
func (d *defaultUserProfileModel) Delete(ctx context.Context, id int64) (int64, error) {
	logx.WithContext(ctx).Infof("delete data:%+v", id)

	if id <= 0 {
		logx.WithContext(ctx).Errorf("delete error:%+v", "param invalid")
		return 0, InputParamInvalid
	}

	result := d.DB.Debug().WithContext(ctx).Where("`id` = ?", id).Delete(&UserProfile{})
	if result.Error != nil {
		logx.WithContext(ctx).Errorf("delete error:%+v", result.Error)
		return 0, WriteDataFailed
	}

	return result.RowsAffected, nil
}

// TxDelete 用于事务删除数据，由上层(如rpc层)调用Transaction去实现，返回受影响行数
func (d *defaultUserProfileModel) TxDelete(ctx context.Context, tx *gorm.DB, id int64) (int64, error) {
	logx.WithContext(ctx).Infof("txDelete data:%+v", id)

	if id <= 0 {
		logx.WithContext(ctx).Errorf("txDelete error:%+v", "param invalid")
		return 0, InputParamInvalid
	}

	result := tx.Debug().WithContext(ctx).Where("`id` = ?", id).Delete(&UserProfile{})
	if result.Error != nil {
		logx.WithContext(ctx).Errorf("txDelete error:%+v", result.Error)
		return 0, WriteDataFailed
	}

	return result.RowsAffected, nil
}

// Deletes 通过主键id批量删除数据，返回受影响行数
func (d *defaultUserProfileModel) Deletes(ctx context.Context, ids []int64) (int64, error) {
	logx.WithContext(ctx).Infof("deletes data:%+v", ids)

	if len(ids) == 0 {
		logx.WithContext(ctx).Errorf("deletes error:%+v", "param invalid")
		return 0, InputParamInvalid
	}

	result := d.DB.Debug().WithContext(ctx).Where("`id` in ?", ids).Delete(&UserProfile{})
	if result.Error != nil {
		logx.WithContext(ctx).Errorf("deletes error:%+v", result.Error)
		return 0, WriteDataFailed
	}

	return result.RowsAffected, nil
}

// TxDeletes 用于事务批量删除数据，由上层(如rpc层)调用Transaction去实现，返回受影响行数
func (d *defaultUserProfileModel) TxDeletes(ctx context.Context, tx *gorm.DB, ids []int64) (int64, error) {
	logx.WithContext(ctx).Infof("txDeletes data:%+v", ids)

	if len(ids) == 0 {
		logx.WithContext(ctx).Errorf("txDeletes error:%+v", "param invalid")
		return 0, InputParamInvalid
	}

	result := tx.Debug().WithContext(ctx).Where("`id` in ?", ids).Delete(&UserProfile{})
	if result.Error != nil {
		logx.WithContext(ctx).Errorf("txDeletes error:%+v", result.Error)
		return 0, WriteDataFailed
	}

	return result.RowsAffected, nil
}

// FindOne 通过主键id查找数据
func (d *defaultUserProfileModel) FindOne(ctx context.Context, id int64) (*UserProfile, error) {
	logx.WithContext(ctx).Infof("findOne data:%+v", id)

	if id <= 0 {
		logx.WithContext(ctx).Errorf("findOne error:%+v", "param invalid")
		return nil, InputParamInvalid
	}

	var result UserProfile
	err := d.DB.Debug().WithContext(ctx).Where("`id` = ?", id).First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		logx.WithContext(ctx).Errorf("findOne error:%+v", err)
		return nil, ReadDataFailed
	}

	return &result, nil
}

// Finds 按条件查询，不分页
// queries: 查询条件集合;
// orders: 排序字符串
func (d *defaultUserProfileModel) Finds(ctx context.Context, queries []Query, orders []Order) ([]*UserProfile, error) {
	logx.WithContext(ctx).Infof("finds queries: %+v, orders: %+v", queries, orders)

	var result []*UserProfile
	db := d.DB.Debug().WithContext(ctx)
	for _, query := range queries {
		db = db.Where(fmt.Sprintf("`%s` %s ?", query.Field, query.Condition), query.Value)
	}

	for _, order := range orders {
		if order.ASC {
			db = db.Order(fmt.Sprintf("%s asc", order.Field))
		} else {
			db = db.Order(fmt.Sprintf("%s desc", order.Field))
		}
	}

	err := db.Find(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		logx.WithContext(ctx).Errorf("finds error:%+v", err)
		return nil, ReadDataFailed
	}

	return result, nil
}

// FindsByPage 分页查询
// queries: 查询条件集合;
// page: 分页条件
// orders: 排序字符串
func (d *defaultUserProfileModel) FindsByPage(ctx context.Context, queries []Query, page *Page, orders []Order) ([]*UserProfile, int64, error) {
	logx.WithContext(ctx).Infof("findsByPage queries: %+v, page: %+v, orders: %+v", queries, page, orders)

	var result []*UserProfile
	var count int64
	db := d.DB.Debug().WithContext(ctx)
	for _, query := range queries {
		db = db.Where(fmt.Sprintf("`%s` %s ?", query.Field, query.Condition), query.Value)
	}

	err := db.Model(&UserProfile{}).Count(&count).Error
	if err != nil {
		logx.WithContext(ctx).Errorf("findsByPage error:%+v", err)
		return nil, 0, ReadDataFailed
	}

	for _, order := range orders {
		if order.ASC {
			db = db.Order(fmt.Sprintf("%s asc", order.Field))
		} else {
			db = db.Order(fmt.Sprintf("%s desc", order.Field))
		}
	}

	if page != nil {
		if page.PageIndex < 0 {
			page.PageIndex = 0
		}
		if page.PageSize <= 0 {
			page.PageSize = 20
		}
		if page.PageSize >= 100 {
			page.PageSize = 100
		}

		db = db.Offset(page.PageIndex * page.PageSize).Limit(page.PageSize)
	}

	err = db.Find(result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, ErrNotFound
		}
		logx.WithContext(ctx).Errorf("findsByPage error:%+v", err)
		return nil, 0, ReadDataFailed
	}

	return result, count, nil
}

// FindCount 按条件查询记录条数
// queries: 查询条件集合;
func (d *defaultUserProfileModel) FindCount(ctx context.Context, queries []Query) (int64, error) {
	logx.WithContext(ctx).Infof("findCount queries: %+v", queries)

	db := d.DB.Debug().WithContext(ctx)
	for _, query := range queries {
		db = db.Where(fmt.Sprintf("`%s` %s ?", query.Field, query.Condition), query.Value)
	}

	var count int64
	err := db.Model(&UserProfile{}).Count(&count).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, ErrNotFound
		}
		logx.WithContext(ctx).Errorf("findCount error:%+v", err)
		return 0, ReadDataFailed
	}

	return count, nil
}
