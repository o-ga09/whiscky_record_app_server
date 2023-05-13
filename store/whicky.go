package store

import (
	"context"
	"errors"
	"fmt"
	"log"
	"main/entity"

	"github.com/go-sql-driver/mysql"
)

func (r *Repository) RegisterWhicky(ctx context.Context,db Execer,record *entity.WhickyRecord) error {
	record.DrankAt = r.Clocker.Now()

	sql := `INSERT INTO whicky_record (user_id,whisky_name,drankAt,taste,smell,evaluate, imageUrl) VALUES (?,?,?,?,?,?,?)`
	_, err := db.ExecContext(ctx,sql,record.UserID,record.Name,record.DrankAt,record.Taste,record.Smell,record.Evaluate, record.ImageURL)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err,&mysqlErr) && mysqlErr.Number == ErrCodeMYSQLDuplicateEntry {
			return fmt.Errorf("cannnot create sama name user: %w", ErrAlreadyEntry)
		}
		return err
	}
	return nil
}

func (r *Repository) GetWhickyRecord(ctx context.Context,db Queryer,uid string) (*[]entity.WhickyRecord, error) {
	records := make([]entity.WhickyRecord, 0)
	sql := `SELECT user_id,whisky_name,drankAt,taste,smell,evaluate,imageUrl FROM whicky_record WHERE user_id = ?`
	rows,err := db.QueryxContext(ctx,sql,uid)
	if err != nil {
		return nil,err
	}

    for rows.Next() {

        var record entity.WhickyRecord

        err := rows.Scan(
            &record.UserID,
            &record.Name,
            &record.DrankAt,
			&record.Taste,
			&record.Smell,
			&record.Evaluate,
            &record.ImageURL,
        )

        if err != nil {
            log.Fatal(err)
        }
        records = append(records, record)
    }
	return &records,nil
}