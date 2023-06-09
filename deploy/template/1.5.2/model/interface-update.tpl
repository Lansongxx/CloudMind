Update(ctx context.Context, id int64, data *{{.upperStartCamelObject}}) (int64, error)                      // 通过主键id更新指定字段，零值不可更新，返回受影响行数
TxUpdate(ctx context.Context, tx *gorm.DB, id int64, data *{{.upperStartCamelObject}}) (int64, error)       // 用于事务更新数据，由上层(如rpc层)调用Transaction去实现，零值不可更新，返回受影响行数
UpdateOneMapById(ctx context.Context, id int64, data map[string]interface{}) (int64, error)                 // 通过主键id更新字段，map参数为数据库需要更新的字段，返回受影响行数
TxUpdateOneMapById(ctx context.Context, tx *gorm.DB, id int64, data map[string]interface{}) (int64, error)  // 用于事务更新字段，由上层(如rpc层)调用Transaction去实现，返回受影响的行数
