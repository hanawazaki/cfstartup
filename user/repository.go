package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// newbie notes
// 1
// aturan yang akan digunakan oleh function NewRepository
// (user User) adalah diambil dari struct User
// (User,error) artinya nilai balikan adalah User atau error

// 2
// method Save adalah method yang dimiliki oleh instance/object r dengan type struct repository
// nama method,param,dan nilai balikannya mengacu pada interface Repositoy
// r.db.Create(&user).Error berfungsi untuk melakukan penyimpanan ke db dengan pengecekan error
// yg disimpan di variable err, lalu dicek jika ada error maka return user,err
// jika tidak akan return user,nil dan dimasukan ke struct repository

// 3
// function newRepository digunakan untuk mem passing nilai db yang sudah diinsert
// agar db bisa diakses di main, maka param nya db dengan type *gorm.DB
// dengan nilai balikan *repository (dibuat pointer agar nilai nya update)
// di dalam NewRepository me-return nilai db dalam struct repository
// yang sudah memiliki nilai agar bisa diakses di main

// 4
// struct repository dibuat melalui function NewRepository
