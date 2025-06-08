package system

func GetName() string {
	rows, err := db.Query("SELECT name FROM system")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var name string
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			panic(err)
		}
	}
	return name
}
