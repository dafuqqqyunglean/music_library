package models

import "errors"

type Song struct {
	ID     int    `json:"id" db:"id"`
	Group  string `json:"group" binding:"required" db:"group_name"`
	Song   string `json:"song" binding:"required" db:"song"`
	Genre  string `json:"genre,omitempty" db:"genre"`
	Date   string `json:"date,omitempty" db:"date"`
	Lyrics string `json:"lyrics,omitempty" db:"lyrics"`
	Link   string `json:"link,omitempty" db:"link"`
}

type UpdateSongInput struct {
	Group  *string `json:"group"`
	Song   *string `json:"song"`
	Genre  *string `json:"genre,omitempty"`
	Date   *string `json:"date,omitempty"`
	Lyrics *string `json:"lyrics,omitempty"`
	Link   *string `json:"link,omitempty"`
}

func (i UpdateSongInput) Validate() error {
	if i.Group == nil && i.Song == nil && i.Genre == nil && i.Date == nil && i.Lyrics == nil && i.Link == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
