package repository

import (
	"database/sql"
	"happy-end/internal/model"
)

type LogsRepository struct {
	DB *sql.DB
}

func NewLogsRepository(db *sql.DB) *LogsRepository {
	return &LogsRepository{DB: db}
}

func (r *LogsRepository) GetLogs() ([]model.Logs, error) {
	rows, err := r.DB.Query("SELECT id, char, dialog FROM logs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []model.Logs
	for rows.Next() {
		var p model.Logs
		if err := rows.Scan(&p.ID, &p.Char, &p.Dialog); err != nil {
			return nil, err
		}
		logs = append(logs, p)
	}
	return logs, nil
}

func (r *LogsRepository) GetLogById(id int) (*model.Logs, error) {
	row, err := r.DB.Query("SELECT id, char, dialog FROM logs WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var log model.Logs
	if err := row.Scan(&log.ID, &log.Char, &log.Dialog); err != nil {
		return nil, err
	}
	return &log, nil
}

func (r *LogsRepository) CreateLog(log *model.Logs) (*model.Logs, error) {
	row, err := r.DB.Query("INSERT INTO logs (char, dialog) VALUES ($1, $2) RETURNING id, char, dialog", log.Char, log.Dialog)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	if err := row.Scan(&log.ID, &log.Char, &log.Dialog); err != nil {
		return nil, err
	}
	return log, nil
}

func (r *LogsRepository) UpdateLog(id int, log *model.Logs) (*model.Logs, error) {
	row, err := r.DB.Query("UPDATE logs SET char = $1, dialog = $2 WHERE id = $3 RETURNING id, char, dialog", log.Char, log.Dialog, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	if err := row.Scan(&log.ID, &log.Char, &log.Dialog); err != nil {
		return nil, err
	}
	return log, nil
}

func (r *LogsRepository) DeleteLog(id int) error {
	row, err := r.DB.Query("DELETE FROM logs WHERE id = $1", id)
	if err != nil {
		return err
	}
	defer row.Close()

	return nil
}
