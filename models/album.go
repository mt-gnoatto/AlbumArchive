package models

type Album struct {
	Name   string `json:"name,omitempty"`
	Genre  string `json:"genre,omitempty"`
	Score  uint64 `json:"score,omitempty"`
	Liked  bool   `json:"liked,omitempty"`
	Played bool   `json:"played,omitempty"`
}
