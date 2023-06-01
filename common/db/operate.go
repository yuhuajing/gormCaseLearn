package db

import (
	"main/common/table"

	"github.com/jinzhu/gorm"
)

// 插入
func Insert(db *gorm.DB) {
	db.Create(&table.Product{Code: "L1212", Price: 1000})
}

// 查
func Read(db *gorm.DB, product *table.Product) {
	db.First(product, 1) // 查询id为1的product
	db.First(product, "code = ?", "L1212")
	db.Find(product)
}

// 更新 - 更新product的price为2000
func Modify(db *gorm.DB, product *table.Product) {
	db.Model(product).Update("Price", 2000)
}

// 删除指定条件
func DeleteRow(db *gorm.DB, str string, id uint) {
	db.Unscoped().Where(str, id).Delete(&table.Product{})
}

// 删除 - 删除表
func Delete(db *gorm.DB, product *table.Product) {
	db.DropTable(table.Product{})
}
