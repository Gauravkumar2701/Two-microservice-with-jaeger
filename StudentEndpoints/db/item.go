

package db

import (
	"StudentEndpoints/models"
	_ "database/sql"

)

func (db Database) GetAllStudent() (*models.ItemList, error) {
	list := &models.ItemList{}

	rows, err := db.Conn.Query("SELECT * FROM student ORDER BY ID DESC")
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Skills, &item.CreatedAt)
		if err != nil {
			return list, err
		}
		list.Student = append(list.Student, item)
	}
	return list, nil
}

func (db Database) AddItem(item *models.Item) error {
	var id int
	var createdAt string
	query := `INSERT INTO student(name, skills) VALUES ($1, $2) RETURNING id, created_at`
	err := db.Conn.QueryRow(query, item.Name, item.Skills).Scan(&id, &createdAt)
	if err != nil {
		return err
	}

	item.ID = id
	item.CreatedAt = createdAt
	return nil
}
