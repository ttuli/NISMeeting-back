package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// UserEntity 用户实体结构体
type UserEntity struct {
	UserId            string    `json:"userId" db:"user_id"`
	Phone             string    `json:"phone" db:"phone"`
	NickName          string    `json:"nickName" db:"nick_name"`
	Sex               int       `json:"sex" db:"sex"`
	PersonalSignature string    `json:"personalSignature" db:"personal_signature"`
	CreateTime        time.Time `json:"createTime" db:"create_time"`
	AreaName          string    `json:"areaName" db:"area_name"`
	AreaCode          string    `json:"areaCode" db:"area_code"`
	Password          string    `json:"password" db:"password"`
}

type UserModel struct {
	Readconn  sqlx.SqlConn
	Writeconn sqlx.SqlConn
	tableName string
}

func NewUserModel(Readconn, Writeconn sqlx.SqlConn) *UserModel {
	return &UserModel{
		Readconn:  Readconn,
		Writeconn: Writeconn,
		tableName: "user",
	}
}

func (m *UserModel) Create(ctx context.Context, user *UserEntity) error {
	query := `
		INSERT INTO user (user_id, phone, nick_name, sex, personal_signature, area_name, area_code, password)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := m.Writeconn.ExecCtx(ctx, query,
		user.UserId,
		user.Phone,
		user.NickName,
		user.Sex,
		user.PersonalSignature,
		user.AreaName,
		user.AreaCode,
		user.Password,
	)

	if err != nil {
		return fmt.Errorf("创建用户失败: %v", err)
	}

	return nil
}

func (m *UserModel) GetByID(ctx context.Context, userID string) (*UserEntity, error) {
	query := `
		SELECT user_id, phone, nick_name, sex, personal_signature, create_time, area_name, area_code, password
		FROM user 
		WHERE user_id = ?
	`

	user := &UserEntity{}
	err := m.Readconn.QueryRowCtx(ctx, user, query, userID)

	if err != nil {
		return nil, fmt.Errorf("获取用户失败: %v", err)
	}

	return user, nil
}

func (m *UserModel) GetByPhone(ctx context.Context, phone string) (*UserEntity, error) {
	query := `
		SELECT user_id, phone, nick_name, sex, personal_signature, create_time, area_name, area_code, password
		FROM user 
		WHERE phone = ?
	`

	user := &UserEntity{}
	err := m.Readconn.QueryRowCtx(ctx, user, query, phone)

	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("获取用户失败: %v", err)
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return user, nil
}

func (m *UserModel) Update(ctx context.Context, user *UserEntity) error {
	query := `
		UPDATE user 
		SET phone = ?, nick_name = ?, sex = ?, personal_signature = ?, area_name = ?, area_code = ?
		WHERE user_id = ?
	`

	result, err := m.Writeconn.ExecCtx(ctx, query,
		user.Phone,
		user.NickName,
		user.Sex,
		user.PersonalSignature,
		user.AreaName,
		user.AreaCode,
		user.UserId,
	)

	if err != nil {
		return fmt.Errorf("更新用户失败: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("获取影响行数失败: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("用户不存在")
	}

	return nil
}

func (m *UserModel) Delete(ctx context.Context, userID string) error {
	query := `DELETE FROM user WHERE user_id = ?`

	result, err := m.Writeconn.ExecCtx(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("删除用户失败: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("获取影响行数失败: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("用户不存在")
	}

	return nil
}
