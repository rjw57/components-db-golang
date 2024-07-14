// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/rjw57/components-db-golang/backend/model"
)

func newCabinet(db *gorm.DB, opts ...gen.DOOption) cabinet {
	_cabinet := cabinet{}

	_cabinet.cabinetDo.UseDB(db, opts...)
	_cabinet.cabinetDo.UseModel(&model.Cabinet{})

	tableName := _cabinet.cabinetDo.TableName()
	_cabinet.ALL = field.NewAsterisk(tableName)
	_cabinet.ID = field.NewInt64(tableName, "id")
	_cabinet.UUID = field.NewString(tableName, "uuid")
	_cabinet.Name = field.NewString(tableName, "name")
	_cabinet.CreatedAt = field.NewTime(tableName, "created_at")
	_cabinet.UpdatedAt = field.NewTime(tableName, "updated_at")

	_cabinet.fillFieldMap()

	return _cabinet
}

type cabinet struct {
	cabinetDo cabinetDo

	ALL       field.Asterisk
	ID        field.Int64
	UUID      field.String
	Name      field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (c cabinet) Table(newTableName string) *cabinet {
	c.cabinetDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cabinet) As(alias string) *cabinet {
	c.cabinetDo.DO = *(c.cabinetDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cabinet) updateTableName(table string) *cabinet {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.UUID = field.NewString(table, "uuid")
	c.Name = field.NewString(table, "name")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")

	c.fillFieldMap()

	return c
}

func (c *cabinet) WithContext(ctx context.Context) ICabinetDo { return c.cabinetDo.WithContext(ctx) }

func (c cabinet) TableName() string { return c.cabinetDo.TableName() }

func (c cabinet) Alias() string { return c.cabinetDo.Alias() }

func (c cabinet) Columns(cols ...field.Expr) gen.Columns { return c.cabinetDo.Columns(cols...) }

func (c *cabinet) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cabinet) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 5)
	c.fieldMap["id"] = c.ID
	c.fieldMap["uuid"] = c.UUID
	c.fieldMap["name"] = c.Name
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
}

func (c cabinet) clone(db *gorm.DB) cabinet {
	c.cabinetDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cabinet) replaceDB(db *gorm.DB) cabinet {
	c.cabinetDo.ReplaceDB(db)
	return c
}

type cabinetDo struct{ gen.DO }

type ICabinetDo interface {
	gen.SubQuery
	Debug() ICabinetDo
	WithContext(ctx context.Context) ICabinetDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICabinetDo
	WriteDB() ICabinetDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICabinetDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICabinetDo
	Not(conds ...gen.Condition) ICabinetDo
	Or(conds ...gen.Condition) ICabinetDo
	Select(conds ...field.Expr) ICabinetDo
	Where(conds ...gen.Condition) ICabinetDo
	Order(conds ...field.Expr) ICabinetDo
	Distinct(cols ...field.Expr) ICabinetDo
	Omit(cols ...field.Expr) ICabinetDo
	Join(table schema.Tabler, on ...field.Expr) ICabinetDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICabinetDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICabinetDo
	Group(cols ...field.Expr) ICabinetDo
	Having(conds ...gen.Condition) ICabinetDo
	Limit(limit int) ICabinetDo
	Offset(offset int) ICabinetDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICabinetDo
	Unscoped() ICabinetDo
	Create(values ...*model.Cabinet) error
	CreateInBatches(values []*model.Cabinet, batchSize int) error
	Save(values ...*model.Cabinet) error
	First() (*model.Cabinet, error)
	Take() (*model.Cabinet, error)
	Last() (*model.Cabinet, error)
	Find() ([]*model.Cabinet, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Cabinet, err error)
	FindInBatches(result *[]*model.Cabinet, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Cabinet) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICabinetDo
	Assign(attrs ...field.AssignExpr) ICabinetDo
	Joins(fields ...field.RelationField) ICabinetDo
	Preload(fields ...field.RelationField) ICabinetDo
	FirstOrInit() (*model.Cabinet, error)
	FirstOrCreate() (*model.Cabinet, error)
	FindByPage(offset int, limit int) (result []*model.Cabinet, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICabinetDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c cabinetDo) Debug() ICabinetDo {
	return c.withDO(c.DO.Debug())
}

func (c cabinetDo) WithContext(ctx context.Context) ICabinetDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cabinetDo) ReadDB() ICabinetDo {
	return c.Clauses(dbresolver.Read)
}

func (c cabinetDo) WriteDB() ICabinetDo {
	return c.Clauses(dbresolver.Write)
}

func (c cabinetDo) Session(config *gorm.Session) ICabinetDo {
	return c.withDO(c.DO.Session(config))
}

func (c cabinetDo) Clauses(conds ...clause.Expression) ICabinetDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cabinetDo) Returning(value interface{}, columns ...string) ICabinetDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cabinetDo) Not(conds ...gen.Condition) ICabinetDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cabinetDo) Or(conds ...gen.Condition) ICabinetDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cabinetDo) Select(conds ...field.Expr) ICabinetDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cabinetDo) Where(conds ...gen.Condition) ICabinetDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cabinetDo) Order(conds ...field.Expr) ICabinetDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cabinetDo) Distinct(cols ...field.Expr) ICabinetDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cabinetDo) Omit(cols ...field.Expr) ICabinetDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cabinetDo) Join(table schema.Tabler, on ...field.Expr) ICabinetDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cabinetDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICabinetDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cabinetDo) RightJoin(table schema.Tabler, on ...field.Expr) ICabinetDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cabinetDo) Group(cols ...field.Expr) ICabinetDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cabinetDo) Having(conds ...gen.Condition) ICabinetDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cabinetDo) Limit(limit int) ICabinetDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cabinetDo) Offset(offset int) ICabinetDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cabinetDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICabinetDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cabinetDo) Unscoped() ICabinetDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cabinetDo) Create(values ...*model.Cabinet) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cabinetDo) CreateInBatches(values []*model.Cabinet, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cabinetDo) Save(values ...*model.Cabinet) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cabinetDo) First() (*model.Cabinet, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Cabinet), nil
	}
}

func (c cabinetDo) Take() (*model.Cabinet, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Cabinet), nil
	}
}

func (c cabinetDo) Last() (*model.Cabinet, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Cabinet), nil
	}
}

func (c cabinetDo) Find() ([]*model.Cabinet, error) {
	result, err := c.DO.Find()
	return result.([]*model.Cabinet), err
}

func (c cabinetDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Cabinet, err error) {
	buf := make([]*model.Cabinet, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cabinetDo) FindInBatches(result *[]*model.Cabinet, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cabinetDo) Attrs(attrs ...field.AssignExpr) ICabinetDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cabinetDo) Assign(attrs ...field.AssignExpr) ICabinetDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cabinetDo) Joins(fields ...field.RelationField) ICabinetDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cabinetDo) Preload(fields ...field.RelationField) ICabinetDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cabinetDo) FirstOrInit() (*model.Cabinet, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Cabinet), nil
	}
}

func (c cabinetDo) FirstOrCreate() (*model.Cabinet, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Cabinet), nil
	}
}

func (c cabinetDo) FindByPage(offset int, limit int) (result []*model.Cabinet, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c cabinetDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cabinetDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cabinetDo) Delete(models ...*model.Cabinet) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cabinetDo) withDO(do gen.Dao) *cabinetDo {
	c.DO = *do.(*gen.DO)
	return c
}
