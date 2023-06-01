package main

import (
	"main/common/db"
	"main/common/table"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	datab, err := db.Connect()
	if err != nil {
		panic("连接数据库失败")
	}
	defer datab.Close()
	// 自动迁移模式
	datab.AutoMigrate(&table.Product{})
	// 插入
	db.Insert(datab)
	// 读取
	//var product Product
	//read(db,&product)
	//fmt.Println(product)
	//deleteRow(db, "id=?", 15)
	//delete(db, &Product{})
}
