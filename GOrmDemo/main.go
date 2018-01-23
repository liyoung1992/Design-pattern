package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlService struct {
	db *gorm.DB
	status bool
}
//打开mysql连接
func (my* MysqlService) open() (*gorm.DB,error){
	if my.db == nil {
		var err error
		my.db,err = gorm.Open("mysql","root:123456@(127.0.0.1:3306)/godb?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			fmt.Println("mysql connect error:",err)
			return nil,err
		}
	}
	return my.db,nil
}
func (my* MysqlService) close() (error){
	if my.db != nil {
		return my.db.Close()
	}
	return nil
}


func main() {
	var  service MysqlService
	//------------------------------连接、关闭数据库-----------------------------------------
	//连接数据库
	service.open()
	defer service.close()

	type User struct {
		gorm.Model
		Name string `gorm:"default:'galeone'"`
		Age int
	}
	//------------------------------创建表-----------------------------------------
	// 创建表
	if !service.db.HasTable(&User{}) {
		service.db.CreateTable(&User{})
	}
	//CRUD ：读写数据
	user := User{Name:"liyang",Age:18}
	user2 := User{Age:18}

	//主键为空返回true
	flag1 := service.db.NewRecord(user)
	service.db.Create(&user)
	service.db.Create(&user2)
	//创建user后返回false
	flag2 := service.db.NewRecord(&user)
	fmt.Println("flag1:",flag1)
	fmt.Println("flag2:",flag2)
	//----------------------------------查询-------------------------------------
	//查询
	var user_search1,user_search2,user_search3 User
	//查找第一条记录
	service.db.First(&user_search1)
	//查找最后一条记录
	service.db.Last(&user_search2)
	//where查询
	service.db.Where("name = ?", "liyang").First(&user_search3)
	fmt.Println("user_search1=",user_search1)
	fmt.Println("user_search2=",user_search2)
	fmt.Println("user_search3=",user_search3)

	var user_search4,user_search5,user_search6 User
	service.db.Where(&User{Name: "liyang", Age: 18}).First(&user_search4)
	service.db.Where("name in (?)", []string{"555", "liyang"}).Find(&user_search5)
	service.db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&user_search6)
	fmt.Println("user_search4=",user_search4)
	fmt.Println("user_search5=",user_search5)
	fmt.Println("user_search6=",user_search6)
	// 更多查询请参考官方文档，基本差不多
	//--------------------------更新---------------------------------------------
	var user_update1,user_update2,user_update3 User
	service.db.First(&user_update1)
	user_search1.Name = "update1"
	user_search1.Age = 999
	service.db.Save(&user_search1)

	service.db.Last(&user_update2)
	service.db.Model(&user_update2).Update("name","update2")

	service.db.Where("name = ?", "666").Delete(User{})
	service.db.Where(&User{Name: "liyang"}).First(&user_update3)
	service.db.Model(&user_update3).Updates(User{Name: "update3", Age: 18})
	//--------------------------删除---------------------------------------------
	//通过主键删除，否则会删除全部，当然这是软删除
	var user_del User
	user_del.ID = 19
	service.db.Delete(&user_del)
	service.db.Where("age = ?",999).Delete(&User{})
//	service.db.Where("name = ?", "666").Delete(User{})

}