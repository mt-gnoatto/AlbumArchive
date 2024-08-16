package models

type Album struct {
	Name   string `json:"name,omitempty"`
	Artist string `json:"artist,omitempty"`
	Genre  string `json:"genre,omitempty"`
	Score  uint64 `json:"score,omitempty"`
	Liked  bool   `json:"liked,omitempty"`
	Played bool   `json:"played,omitempty"`
}
