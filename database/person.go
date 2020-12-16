package database

import "log"

type PersonHandler struct {
	DataBaseHandler *DBHandler
}

type Person struct {
	ID     uint
	Name   string
	Family string
	Age    uint
}

func NewPersonHandler(db *DBHandler) (*PersonHandler, error) {
	return &PersonHandler{DataBaseHandler: db}, nil
}

func (p *PersonHandler) GeTAllPeople() ([]*Person, error) {
	rows, err := p.DataBaseHandler.DB.Query("select * from `person`")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	people := []*Person{}
	for rows.Next() {
		p := Person{}
		if err = rows.Scan(&p.ID, &p.Name, &p.Family, &p.Age); err != nil {
			log.Println(err)
			continue
		}
		people = append(people, &p)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return people, nil
}

func (p *PersonHandler) GetPersonById(id int) (*Person, error) {
	row := p.DataBaseHandler.DB.QueryRow("select * from `person` where id = ?", id)
	if err := row.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	person := &Person{}
	if err := row.Scan(&person.ID, &person.Name, &person.Family, &person.Age); err != nil {
		log.Println(err)
		return nil, err
	}
	return person, nil
}

func (p *PersonHandler) GetPersonByName(name string) (*Person, error) {
	row := p.DataBaseHandler.DB.QueryRow("select * from `person` where name =?", name)
	if err := row.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	person := &Person{}
	if err := row.Scan(&person.ID, &person.Name, &person.Family, &person.Age); err != nil {
		log.Println(err)
		return nil, err
	}

	return person, nil
}

func (p *PersonHandler) AddPerson(person *Person) error {
	_, err := p.DataBaseHandler.DB.Exec("insert into `person` (`name`,`family`,`age`)values (?,?,?)", person.Name, person.Family, person.Age)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (p *PersonHandler) UpdatePerson(person *Person) error {
	_, err := p.DataBaseHandler.DB.Exec("update `person` set `name` = ?,`family`=?,`age`=? where (`id` = ? )", person.Name, person.Family, person.Age, person.ID)
	return err
}

func (p *PersonHandler) DeletePerson(person *Person) error {
	_, err := p.DataBaseHandler.DB.Exec("DELETE FROM `person` where (`id` = ?)", person.ID)
	return err
}
