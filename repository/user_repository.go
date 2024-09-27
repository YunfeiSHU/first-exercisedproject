package repository

import (
	"gin-jwt-gorm/MySQL"
	"gin-jwt-gorm/domain/model"
	"log"
)

type userRepository struct {
	database MySQL.MysqlDataBase
	//因为mysqldatabase用的是匿名变量，故有DB的方法，也和自己实现的方法
}

//Mysql作为抽象层，我实现了CRUD的具体操作
//但为了方便使用，需要在repository创建userRepository，实现model.UserRepository;将repository和model层通过接口与类的方式的联系起来
//故利用mysql的CRUD实现，对应接口中方法
//--即mysql层与repository是通过结构体的嵌套，以使用CRUD

func NewUserRepository(db MySQL.MysqlDataBase) model.UserRepository {
	return &userRepository{
		database: db,
	}
}

func (u *userRepository) Create(user *model.User) error {
	user, err := u.database.CreateUser(user)
	if err != nil {
		return err
	}
	log.Println(user)
	return nil
}

func (u *userRepository) GetByName(name string) (*[]model.User, error) {
	users, err := u.database.FindByName(name)
	if err != nil {
		return nil, err
	}
	log.Println(users)
	return users, nil
}

func (u *userRepository) GetByEmail(email string) (*model.User, error) {
	user, err := u.database.FindOneByEmail(email)
	if err != nil {
		return nil, err
	}
	log.Println(user)
	return user, nil
}

func (u *userRepository) GetByID(id int) (*model.User, error) {
	user, err := u.database.FindOneByID(id)
	if err != nil {
		return nil, err
	}
	log.Println(user)
	return user, nil
}
