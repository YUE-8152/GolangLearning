package user

import (
	"GolangLearning/model/db"
	"GolangLearning/model/domain/po"
	"errors"
)

func Create(user po.User) (err error) {
	//user = po.User{
	//	Name:     "王五",
	//	Age:      20,
	//	Password: "111111",
	//}
	if err = db.DB.Create(&user).Error; err != nil {
		err = errors.New("插入数据失败！")
	}
	return err
}

func QueryById(id int) (interface{}, error) {
	var user po.User
	var err error
	if err = db.DB.Where("id=?", id).First(&user).Error; err != nil {
		err = errors.New("根据id获取user信息失败！")
	}
	return user, err
}

func Update(user po.User) (err error) {
	var u po.User
	u.Id = user.Id
	u.Name = user.Name
	u.Age = user.Age
	u.Password = user.Password
	if err = db.DB.Model(&user).Where("id=?", u.Id).Updates(u).Error; err != nil {
		return errors.New("更新失败！")
	}
	return err
}

func Delete(id int) (err error) {
	if err = db.DB.Where("id = ?", id).Delete(&po.User{}).Error; err != nil {
		return errors.New("删除失败！")
	}
	return err
}

func QueryUser() (interface{}, error) {
	var user []po.User
	var err error
	if err = db.DB.Find(&user).Error; err != nil {
		err = errors.New("获取user列表失败！")
	}
	return user, err
}
