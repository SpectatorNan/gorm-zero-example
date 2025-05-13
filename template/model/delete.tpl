
func (m *default{{.upperStartCamelObject}}Model) Delete(ctx context.Context, tx *gorm.DB, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) error {
	{{if .withCache}}data, err:=m.FindOne(ctx, {{.lowerStartCamelPrimaryKey}})
	if err!=nil{
        if errors.Is(err, ErrNotFound) {
                return nil
        }
		return err
	}
	 err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
        if tx != nil {
            db = tx
        }
        return db.Delete(&{{.upperStartCamelObject}}{}, {{.lowerStartCamelPrimaryKey}}).Error
	}, m.GetCacheKeys(data)...){{else}}err := m.conn.ExecCtx(ctx, func(conn *gorm.DB) error {
                                                                   		db := conn
                                                                   		if tx != nil {
                                                                   			db = tx
                                                                   		}
                                                                   		return db.Delete(&{{.upperStartCamelObject}}{}, {{.lowerStartCamelPrimaryKey}}).Error
                                                                   	}){{end}}
	return err
}

func (m *default{{.upperStartCamelObject}}Model) BatchDelete(ctx context.Context, tx *gorm.DB, datas []{{.upperStartCamelObject}}) error {
	{{if .withCache}}err := batchx.BatchExecCtx(ctx, m, datas, func(conn *gorm.DB) error {
db := conn
		for _, v := range datas {
			if err := db.Delete(&v).Error; err != nil {
				return err
			}
		}
		return nil
    	},tx){{else}}err := batchx.BatchNoCacheExecCtx[{{.upperStartCamelObject}}](ctx, m.conn, func(db *gorm.DB) error {
                                 		for _, data := range datas {
                                 			if err := db.Delete(&data).Error; err != nil {
                                 				return err
                                 			}
                                 		}
                                 		return nil
                                 	}, tx){{end}}
	return err
}

