package db

import (
	"cred-api/internal/services"
)

func CreatePerson(person *services.Person) error {
	return DB.Create(person).Error
}

func GetPeople(page, perPage int, name, surname, gender, nationality string, ageMin, ageMax int) ([]services.Person, error) {
	var people []services.Person
	query := DB.Offset((page - 1) * perPage).Limit(perPage)

	if name != "" {
		query = query.Where("name = ?", name)
	}

	if surname != "" {
		query = query.Where("surname = ?", surname)
	}

	if gender != "" {
		query = query.Where("gender = ?", gender)
	}

	if nationality != "" {
		query = query.Where("nationality = ?", nationality)
	}

	if ageMin > 0 {
		query = query.Where("age >= ?", ageMin)
	}

	if ageMax > 0 {
		query = query.Where("age <= ?", ageMax)
	}

	if err := query.Find(&people).Error; err != nil {
		return nil, err
	}

	return people, nil
}

func UpdatePerson(id string, updatedPerson *services.Person) error {
	return DB.Model(&services.Person{}).Where("id = ?", id).Updates(updatedPerson).Error
}

func DeletePerson(id string) error {
	return DB.Delete(&services.Person{}, id).Error
}
