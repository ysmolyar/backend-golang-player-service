package models

type Player struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Team      string `json:"team"`
	Position  string `json:"position"`
	BatAvg    float64 `json:"batAvg"`
	HomeRuns  int    `json:"homeRuns"`
	RBI       int    `json:"rbi"`
} 