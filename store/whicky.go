package store

import (
	"context"
	"errors"
	"fmt"
	"main/entity"

	"github.com/go-sql-driver/mysql"
)

func (r *Repository) RegisterWhicky(ctx context.Context,db Execer,record *entity.WhickyRecord) error {
	record.DrankAt = r.Clocker.Now()

	sql := `INSERT INTO whisky_record (user_id,whisky_name,drankAt,imageUrl) VALUES (?,?,?,?)`
	_, err := db.ExecContext(ctx,sql,record.UserID,record.Name,record.DrankAt,record.ImageURL)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err,&mysqlErr) && mysqlErr.Number == ErrCodeMYSQLDuplicateEntry {
			return fmt.Errorf("cannnot create sama name user: %w", ErrAlreadyEntry)
		}
		return err
	}
	return nil
}

func (r *Repository) GetWhickyRecord(ctx context.Context,db Queryer,uid entity.UserID) (*entity.WhickyRecord, error) {
	record := &entity.WhickyRecord{}
	sql := `SELECT user_id,whisky_name,drankAt,imageUrl FROM whicky_record WHERE user_id = ?`
	if err := db.GetContext(ctx,record,sql,uid); err != nil {
		return nil,err
	}
	return record,nil
}