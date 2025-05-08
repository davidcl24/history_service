package models

type HistoryElement struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	MovieID   int `json:"movie_id"`
	EpisodeID int `json:"episode_id"`
	WatchTime int `json:"watch_time"`
	Progress  int `json:"progress"`
}

func GetAllUserHistoryElements(userId int) []*HistoryElement {
	return nil
}

func GetHistoryElementByID(id int) *HistoryElement {
	return nil
}

func AddHistoryElement(historyElement *HistoryElement) *HistoryElement {
	return nil
}

func DeleteHistoryElement(id int) *HistoryElement {
	return nil
}

func ClearUserHistoryElements(userId int) []*HistoryElement {
	return nil
}

func UpdateHistoryElement(id int, historyElementUpdate HistoryElement) *HistoryElement {
	return nil
}
