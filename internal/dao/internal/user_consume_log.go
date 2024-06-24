// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserConsumeLogDao is the data access object for table user_consume_log.
type UserConsumeLogDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns UserConsumeLogColumns // columns contains all the column names of Table for convenient usage.
}

// UserConsumeLogColumns defines and stores column names for table user_consume_log.
type UserConsumeLogColumns struct {
	Id          string // 主键
	UserId      string // 消费用户ID
	Amount      string // 消费使用的金额;单位：屋币
	ProductType string // 消费商品类型;0-小说VIP章节
	ProductId   string // 消费的的商品ID;例如：章节ID
	ProducName  string // 消费的的商品名;例如：章节名
	ProducValue string // 消费的的商品值;例如：1
	CreateTime  string // 创建时间
	UpdateTime  string // 更新时间
}

// userConsumeLogColumns holds the columns for table user_consume_log.
var userConsumeLogColumns = UserConsumeLogColumns{
	Id:          "id",
	UserId:      "user_id",
	Amount:      "amount",
	ProductType: "product_type",
	ProductId:   "product_id",
	ProducName:  "produc_name",
	ProducValue: "produc_value",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// NewUserConsumeLogDao creates and returns a new DAO object for table data access.
func NewUserConsumeLogDao() *UserConsumeLogDao {
	return &UserConsumeLogDao{
		group:   "default",
		table:   "user_consume_log",
		columns: userConsumeLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserConsumeLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserConsumeLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserConsumeLogDao) Columns() UserConsumeLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserConsumeLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserConsumeLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserConsumeLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
