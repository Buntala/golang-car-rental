package driver

import (
	"car-rental/utilities/db"
	"fmt"
	"log"
)
func DBGetDriverAll() []DriverVal{
	conn := db.DbConnect()
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM driver ORDER BY DriverId DESC`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result []DriverVal
	for rows.Next() {
		//USE BIND LATER
		var driver DriverVal
		if err := rows.Scan(&driver.DriverId,&driver.Name,&driver.Nik,&driver.PhoneNumber,&driver.DailyCost); 
		err != nil {
				log.Fatal(err)
		}
		result = append(result,driver)
	}
	return result
}

func DBGetDriverOne(id string) []DriverVal{
	conn := db.DbConnect()
	defer conn.Close()
	query := fmt.Sprintf(`SELECT * FROM driver WHERE DriverId = %v DESC`, id)
	rows, err := conn.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result []DriverVal
	for rows.Next() {
		//USE BIND LATER
		var driver DriverVal
		if err := rows.Scan(&driver.DriverId,&driver.Name,&driver.Nik,&driver.PhoneNumber,&driver.DailyCost); 
		err != nil {
				log.Fatal(err)
		}
		result = append(result,driver)
	}
	return result
}

func DBInsertDriver(body Post_Rules) []DriverVal{
	conn := db.DbConnect()
	defer conn.Close()
	query := fmt.Sprintf(`INSERT INTO driver (name,nik,PhoneNumber,daily_cost) VALUES ('%v','%v','%v',%v) RETURNING *`, body.Name,body.Nik,body.PhoneNumber,body.DailyCost)
	rows, err := conn.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result []DriverVal
	for rows.Next() {
		//USE BIND LATER
		var driver DriverVal
		if err := rows.Scan(&driver.DriverId,&driver.Name,&driver.Nik,&driver.PhoneNumber,&driver.DailyCost); 
		err != nil {
				log.Fatal(err)
		}
		result = append(result,driver)
	}
	return result
}
func DBUpdateDriver(body Patch_Rules) []DriverVal{
	conn := db.DbConnect()
	defer conn.Close()
	query := fmt.Sprintf(`UPDATE driver SET name='%v',nik='%v',PhoneNumber='%v',daily_cost='%v' WHERE DriverId = %v RETURNING *`, body.Name,body.Nik,body.PhoneNumber,body.DailyCost, body.DriverId)
	rows, err := conn.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result []DriverVal
	for rows.Next() {
		//USE BIND LATER
		var driver DriverVal
		if err := rows.Scan(&driver.DriverId,&driver.Name,&driver.Nik,&driver.PhoneNumber,&driver.DailyCost); 
		err != nil {
				log.Fatal(err)
		}
		result = append(result,driver)
	}
	return result
}

func DBDeleteDriver(body Delete_Rules) []DriverVal{
	conn := db.DbConnect()
	defer conn.Close()
	query := fmt.Sprintf(`DELETE FROM driver WHERE DriverId = %v RETURNING *`, body.DriverId)
	fmt.Println(query)
	rows, err := conn.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result []DriverVal
	for rows.Next() {
		//USE BIND LATER
		var driver DriverVal
		if err := rows.Scan(&driver.DriverId,&driver.Name,&driver.Nik,&driver.PhoneNumber,&driver.DailyCost); 
		err != nil {
				log.Fatal(err)
		}
		result = append(result,driver)
	}
	return result
}