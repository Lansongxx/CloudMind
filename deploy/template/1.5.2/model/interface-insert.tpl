Insert(ctx context.Context, data *{{.upperStartCamelObject}}) (int64, error) // 返回的是受影响行数，如需获取自增id，请通过data参数获取
TxInsert(ctx context.Context, tx *gorm.DB, data *{{.upperStartCamelObject}}) (int64, error)       // 用于事务新增数据，由上层(如rpc层)调用Transaction去实现，返回受影响的行数
