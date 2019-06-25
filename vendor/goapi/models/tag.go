package models

import (
	db "goapi/databases"
        "log"
)
type Tag struct {
	ID int `json:"tag_id"`
	Name string `json:"tag_name"`
}

func (t *Tag) GetTags() (tags []Tag, err error) {
	tags = make([]Tag, 0)
	rows, err := db.SqlDB.Query("SELECT tag_id, tag_name FROM tb_tag")
	defer rows.Close()

	if err != nil {
		return
	}
	
	for rows.Next() {
		var tag Tag
		rows.Scan(&tag.ID, &tag.Name)
		tags = append(tags, tag)
	}

        log.Println(tags)
	if err = rows.Err(); err != nil {
		return
	}
	return
}
