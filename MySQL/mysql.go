package MySQL

import (
	"gin-jwt-gorm/domain/model"
	"gorm.io/gorm"
)

//TODO 本段代码应该放在MySQL,作为抽象层
//TODO repository应该放与接口的方法

// MySQL是关系型数据量，直接DataBase作为接口，封装CRUD方法，最方便
type DataBase interface {
	//CRUD 接口

	//根据go对于指针操作的便利性，我们返回值传入指针：
	//如果是结构体只能传入空结构体，而不能传入nil
	//nil是值类型

	//根据id,email查询指定用户
	FindOneByID(id int) (*model.User, error)
	FindOneByEmail(email string) (*model.User, error)
	//根据name查询符合用户
	FindByName(name string) (*[]model.User, error)
	//创建用户
	CreateUser(user model.User) (*model.User, error)
	//修改用户:修改名称
	UpdateUserByEmail(email string) (*model.User, error)
	//删除用户
	DeleteByEmail(email string) (*model.User, error)

	//参数没有context，是因为对应MySQL,context用于控制操作执行的时限
}

type MysqlDataBase struct {
	*gorm.DB //匿名字段直接当*gorm.DB使用
}

/*原代码，采用非关系型数据库，
此处的结构体用于缓存---我通过MySQL,CRUD直接对users表处理，不考虑缓存
// 查询单个结构
type SingleUser struct {
	model.User
}

// 查询多个结构
type MultiUser struct {
	users []model.User
}*/

// TODO CRUD

func (mysqlDb *MysqlDataBase) FindOneByID(id int) (*model.User, error) {
	var user model.User
	//事务操作
	tx := mysqlDb.Begin()
	err := tx.Where("id = ?", id).First(&user).Error
	if err != nil {
		//先回滚，在return
		tx.Rollback()
		return nil, err
	}
	//执行成功，提交
	tx.Commit()
	return &user, nil
}
func (mysqlDb *MysqlDataBase) FindOneByEmail(email string) (*model.User, error) {
	var user model.User
	tx := mysqlDb.Begin()
	err := tx.Where("email = ?", email).First(&user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &user, nil
}

func (mysqlDb *MysqlDataBase) FindByName(name string) (*[]model.User, error) {
	var users []model.User
	tx := mysqlDb.Begin()
	err := tx.Where("name = ?", name).First(&users).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &users, nil
}

func (mysqlDb *MysqlDataBase) CreateUser(user *model.User) (*model.User, error) {
	tx := mysqlDb.Begin()
	err := tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return user, nil
}

func (mysqlDb *MysqlDataBase) UpdateUserByEmail(email string) (*model.User, error) {
	var user model.User
	tx := mysqlDb.Begin()
	err := tx.Model(&user).Where("email = ?", email).First(&user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &user, nil
}

func (mysqlDb *MysqlDataBase) DeleteByEmail(email string) (*model.User, error) {
	var user model.User
	tx := mysqlDb.Begin()
	err := tx.Model(&user).Where("email = ?", email).Delete(&user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &user, nil
}
