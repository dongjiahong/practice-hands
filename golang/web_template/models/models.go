package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type BaseDao struct {
	Conn *gorm.DB
}

type Config struct {
	Type        string `toml:"Type"`
	User        string `toml:"User"`
	Password    string `toml:"Password"`
	Host        string `toml:"Host"`
	Name        string `toml:"Name"`
	TablePrefix string `toml:"TablePrefix"`
}

type Model struct {
	ID         int    `gorm:"primary_key"`
	CreatedOn  string `json:"-"`
	ModifiedOn string `json:"-"`
	DeletedOn  int    `json:"-"`
}

func Init(conf *Config) {
	var err error
	db, err = gorm.Open(conf.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Name))
	if err != nil {
		log.Fatalf("models.Setup err: %v\n", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.TablePrefix + defaultTableName
	}

	db.LogMode(true)                                                                           // 打印sql语句
	db.SingularTable(true)                                                                     // 全局禁用表名复数，User的表名为user，而非users
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback) // gorm自动更新created_on字段
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback) // gorm自动更新modefied_on字段
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)                              // gorm自动更新deleted_on字段
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(150)

}

// updateTimeStampForCreateCallback will set `CreateOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Format("2006.01.02 15:04:05")
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Format("2006.01.02 15:04:05"))
	}
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf("UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

func GetTransaction() *gorm.DB {
	return db.Begin()
}

func NewBaseDao(conn *gorm.DB) BaseDao {
	if conn == nil {
		return BaseDao{
			Conn: db,
		}
	} else {
		return BaseDao{
			Conn: conn,
		}
	}
}
