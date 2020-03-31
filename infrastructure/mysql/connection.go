package mysql

import (
	"corona_server/environment"
	"github.com/jinzhu/gorm"
	"golang.org/x/xerrors"
)

var (
	conn *gorm.DB
	err  error
)

var (
	// ErrNewConnection は DB 接続に関するエラー
	ErrNewConnection = xerrors.New("infrastructure-mysql-connection: failed to newConnection")
)

const errorCodeDuplicateEntry = "1062"
const errorCodeForeignKeyConstraint = "1452"

//Connection はDB 接続
func Connection() error {
	conn, err = newConnection()
	if err != nil {
		return xerrors.Errorf(ErrNewConnection.Error()+": %w", err)
	}
	return nil
}

func newConnection() (*gorm.DB, error) {
	if !dbStat(conn) {
		conf := makeConfig()
		if conn, err = gorm.Open("mysql", conf.build()); err != nil {
			return nil, xerrors.Errorf(ErrNewConnection.Error()+": %w", err)
		}

		if err = conn.DB().Ping(); err != nil {
			return nil, xerrors.Errorf(ErrNewConnection.Error()+": %w", err)
		}

		// For connection limits at appengine. https://cloud.google.com/sql/faq#sizeqps
		if conf.connMaxLifetime > 0 {
			conn.DB().SetConnMaxLifetime(conf.connMaxLifetime)
		}
		if conf.maxIdle > 0 {
			conn.DB().SetMaxIdleConns(conf.maxIdle)
		}
		if conf.maxOpen > 0 {
			conn.DB().SetMaxOpenConns(conf.maxOpen)
		}

		if environment.GetSharedEnvironments().GormLogMode {
			conn.LogMode(true)
		}
	}

	return conn, nil
}

// InTransaction は gorm のトランザクション処理を表す be/apiを参考
type InTransaction func(tx *gorm.DB) error

// DoInTransaction はトランザクションのヘルパー
// エラーが生じた場合ロークバック処理を行う
func DoInTransaction(fn InTransaction) error {
	conn, err := newConnection()
	if err != nil {
		return err
	}

	tx := conn.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	err = fn(tx)
	if err != nil {
		xerr := tx.Rollback().Error
		if xerr != nil {
			return xerr
		}
		return err
	}
	if err = tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

// DBコネクションが正常ならtrue
func dbStat(db *gorm.DB) bool {
	if db == nil {
		return false
	}
	if db.DB() == nil {
		return false
	}
	return db.DB().Ping() == nil
}
