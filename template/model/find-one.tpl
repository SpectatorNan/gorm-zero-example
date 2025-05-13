
func (m *default{{.upperStartCamelObject}}Model) FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error) {
	 formatDB := func(conn *gorm.DB) *gorm.DB {
	    return  conn.Model(&{{.upperStartCamelObject}}{}).Where("{{.originalPrimaryKey}} = ?", {{.lowerStartCamelPrimaryKey}})
	 }
	 var resp {{.upperStartCamelObject}}
	{{if .withCache}}{{.cacheKey}}
	err := m.QueryCtx(ctx, &resp, {{.cacheKeyVariable}}, func(conn *gorm.DB) error {
    		return formatDB(conn).First(&resp).Error
    	})
	 {{else}}
	err := m.conn.ExecCtx(ctx, func(conn *gorm.DB) error {
    		return formatDB(conn).Take(&resp).Error
    	})
	 {{end}}
	 if err != nil {
     		return nil, err
     	}
     	return &resp, nil
}

func (m *default{{.upperStartCamelObject}}Model) FindPageList(ctx context.Context, page *pagex.ListReq, orderBy pagex.OrderBy,
	orderKeys map[string]string, whereClause func(db *gorm.DB) *gorm.DB) ([]{{.upperStartCamelObject}}, int64, error) {
	{{if .withCache}}formatDB := func(conn *gorm.DB) (*gorm.DB, *gorm.DB) {
    		db := conn.Model(&{{.upperStartCamelObject}}{})
    		if whereClause != nil {
    			db = whereClause(db)
    		}
    		return db, nil
    	}
    	res, total, err := pagex.FindPageList[{{.upperStartCamelObject}}](ctx, m, page, orderBy, orderKeys, formatDB)
    	return res, total, err{{else}}conn := m.conn
                                      	formatDB := func() (*gorm.DB, *gorm.DB) {
                                      		db := conn.Model(&{{.upperStartCamelObject}}{})
                                      		if whereClause != nil {
                                      			db = whereClause(db)
                                      		}
                                      		return db, nil
                                      	}

                                      	res, total, err := pagex.FindPageListWithCount[{{.upperStartCamelObject}}](ctx, page, orderBy, orderKeys, formatDB)
                                      	return res, total, err{{end}}
}