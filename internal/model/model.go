package model

import (
	"blog_service/global"
	"blog_service/pkg/setting"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettings) (*gorm.DB, error) {
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&local=Local",
		databaseSetting.DBUsename,
		databaseSetting.DBPassword,
		fmt.Sprintf("%s:%s", databaseSetting.DBHost, databaseSetting.DBPort),
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}

	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallBack)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallBack)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}

func updateTimeStampForCreateCallBack(score *gorm.Scope) {
	if !score.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := score.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := score.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallBack(score *gorm.Scope) {
	if _, ok := score.Get("gorm:update_column"); !ok {
		_ = score.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallback(score *gorm.Scope) {
	if !score.HasError() {
		var extraOption string
		if str, ok := score.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deleteOnField, hasDeletedOnField := score.FieldByName("DeletedOn")
		isDelField, hasIsDelField := score.FieldByName("IsDel")
		if !score.Search.Unscoped && hasDeletedOnField && hasIsDelField {
			now := time.Now().Unix()
			score.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v, %v=%v%v%v",
				score.QuotedTableName(),
				score.Quote(deleteOnField.DBName),
				score.AddToVars(now),
				score.Quote(isDelField.DBName),
				score.AddToVars(1),
				addExtraSpaceIfExist(score.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			score.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				score.QuotedTableName(),
				addExtraSpaceIfExist(score.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			))
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}

	return ""
}
