package models

import (
	"errors"
	"net"
	"time"
)

type HistoryElement struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	MovieID   *int      `json:"movie_id"`
	EpisodeID *int      `json:"episode_id"`
	WatchDate time.Time `json:"watch_date"`
	Progress  int       `json:"progress"`
}

func (db *DB) GetAllUserHistoryElements(userId int) ([]*HistoryElement, error) {
	query := `
		SELECT id, user_id, movie_id, episode_id, watch_date, progress
		FROM watch_history
		WHERE user_id = $1`
	rows, err := db.Conn.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	elements := []*HistoryElement{}
	for rows.Next() {
		elem := &HistoryElement{}
		err := rows.Scan(&elem.ID, &elem.UserID, &elem.MovieID, &elem.EpisodeID, &elem.WatchDate, &elem.Progress)
		if err != nil {
			continue
		}
		elements = append(elements, elem)
	}
	return elements, nil
}

func (db *DB) GetHistoryElementByID(id int) (*HistoryElement, error) {
	query := `
		SELECT id, user_id, movie_id, episode_id, watch_date, progress
		FROM watch_history
		WHERE id = $1`
	elem := &HistoryElement{}
	err := db.Conn.QueryRow(query, id).Scan(&elem.ID, &elem.UserID, &elem.MovieID, &elem.EpisodeID, &elem.WatchDate, &elem.Progress)
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		return nil, err
	}
	if err != nil {
		return nil, nil
	}
	return elem, nil
}

func (db *DB) AddHistoryElement(historyElement *HistoryElement) (*HistoryElement, error) {
	query := `
		INSERT INTO watch_history (user_id, movie_id, episode_id, watch_date, progress)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`
	err := db.Conn.QueryRow(
		query,
		historyElement.UserID,
		historyElement.MovieID,
		historyElement.EpisodeID,
		historyElement.WatchDate,
		historyElement.Progress,
	).Scan(&historyElement.ID)

	if err != nil {
		return nil, err
	}
	return historyElement, nil
}

func (db *DB) DeleteHistoryElement(id int) (*HistoryElement, error) {
	element, err := db.GetHistoryElementByID(id)

	if err != nil {
		return nil, err
	}

	if element == nil {
		return nil, nil
	}

	_, err = db.Conn.Exec("DELETE FROM watch_history WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return element, nil
}

func (db *DB) ClearUserHistoryElements(userId int) ([]*HistoryElement, error) {
	elements, err := db.GetAllUserHistoryElements(userId)

	if err != nil {
		return nil, err
	}

	if len(elements) == 0 {
		return nil, nil
	}

	_, err = db.Conn.Exec("DELETE FROM watch_history WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	return elements, nil
}

func (db *DB) UpdateHistoryElement(id int, historyElementUpdate HistoryElement) (*HistoryElement, error) {
	_, err := db.Conn.Exec(`
		UPDATE watch_history
		SET user_id = $1, movie_id = $2, episode_id = $3, watch_date = $4, progress = $5
		WHERE id = $6`,
		historyElementUpdate.UserID, historyElementUpdate.MovieID, historyElementUpdate.EpisodeID, historyElementUpdate.WatchDate, historyElementUpdate.Progress, id,
	)
	if err != nil {
		return nil, err
	}
	elem, _ := db.GetHistoryElementByID(id)
	return elem, nil
}
