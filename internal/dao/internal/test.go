// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TestDao is the data access object for table test.
type TestDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns TestColumns // columns contains all the column names of Table for convenient usage.
}

// TestColumns defines and stores column names for table test.
type TestColumns struct {
	Id    string //
	Test  string //
	Test2 string //
}

// testColumns holds the columns for table test.
var testColumns = TestColumns{
	Id:    "id",
	Test:  "test",
	Test2: "test2",
}

// NewTestDao creates and returns a new DAO object for table data access.
func NewTestDao() *TestDao {
	return &TestDao{
		group:   "default",
		table:   "test",
		columns: testColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TestDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TestDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TestDao) Columns() TestColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TestDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TestDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TestDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
