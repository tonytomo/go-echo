package models

import (
	"go-echo/db"
)

type Warga struct {
	NoPpdb    string      `json:"noPPBB"`
	Nama      string      `json:"nama"`
	Alamat    string      `json:"alamat"`
	Kabupaten string      `json:"kabupaten"`
	Kecamatan string      `json:"kecamatan"`
	Desa      string      `json:"desa"`
	Rt        string      `json:"rt"`
	Rw        string      `json:"rw"`
	DataPbb   interface{} `json:"data_pbb"`
}

type Pbb struct {
	Tahun int `json:"tahun"`
	Pajak int `json:"pajak"`
	Denda int `json:"denda"`
}

type Trash struct {
	Id      int `json:"id"`
	Id_data int `json:"idData"`
	NoPpdb  int `json:"noPpdb"`
}

func FetchDataPbb() (Response, error) {
	var objW Warga
	var objP Pbb
	var objT Trash
	var arrobjW []Warga
	var arrobjP []Pbb
	var res Response

	con := db.CreateCon()

	sqlStatement1 := "SELECT * FROM warga"

	rows1, err := con.Query(sqlStatement1)

	if err != nil {
		return res, err
	}
	defer rows1.Close()

	for rows1.Next() {
		err = rows1.Scan(&objT.Id, &objW.NoPpdb, &objW.Nama, &objW.Alamat, &objW.Kabupaten, &objW.Kecamatan, &objW.Desa, &objW.Rt, &objW.Rw)
		if err != nil {
			return res, err
		}

		sqlStatement2 := "SELECT * FROM data_pbb WHERE no_ppbb = ?"

		rows2, err := con.Query(sqlStatement2, objT.Id)

		if err != nil {
			return res, err
		}
		defer rows2.Close()

		for rows2.Next() {
			err = rows2.Scan(&objT.Id_data, &objT.NoPpdb, &objP.Tahun, &objP.Pajak, &objP.Denda)
			if err != nil {
				return res, err
			}

			arrobjP = append(arrobjP, objP)
		}
		objW.DataPbb = arrobjP
		arrobjP = nil

		arrobjW = append(arrobjW, objW)
	}

	res.Status = true
	res.ErrorCode = "-"
	res.ErrorDescription = "-"
	res.Data = arrobjW

	return res, nil
}

// func StorePegawai(nama string, alamat string, telepon string) (Response, error) {
// 	var res Response

// 	con := db.CreateCon()

// 	sqlStatement := "INSERT INTO pegawai (nama, alamat, telepon) VALUES ($1, $2, $3) RETURNING ID"

// 	var lastInsertedId int64
// 	err := con.QueryRow(sqlStatement, nama, alamat, telepon).Scan(&lastInsertedId)
// 	if err != nil {
// 		return res, err
// 	}

// 	res.Status = true
// 	res.ErrorCode = "-"
// 	res.ErrorDescription = "-"
// 	res.Data = map[string]int64{
// 		"last_inserted_id": lastInsertedId,
// 	}

// 	return res, nil
// }

// func UpdatePegawai(id int, nama string, alamat string, telepon string) (Response, error) {
// 	var res Response

// 	con := db.CreateCon()

// 	sqlStatement := "UPDATE pegawai SET nama = $1, alamat = $2, telepon = $3 WHERE id = $4"

// 	result, err := con.Exec(sqlStatement, nama, alamat, telepon, id)
// 	if err != nil {
// 		return res, err
// 	}

// 	rowAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return res, err
// 	}

// 	res.Status = true
// 	res.ErrorCode = "-"
// 	res.ErrorDescription = "-"
// 	res.Data = map[string]int64{
// 		"rows_affected": rowAffected,
// 	}

// 	return res, nil
// }

// func DeletePegawai(id int) (Response, error) {
// 	var res Response

// 	con := db.CreateCon()

// 	sqlStatement := "DELETE FROM pegawai WHERE id = $1"

// 	result, err := con.Exec(sqlStatement, id)
// 	if err != nil {
// 		return res, err
// 	}

// 	rowDeleted, err := result.RowsAffected()
// 	if err != nil {
// 		return res, err
// 	}

// 	res.Status = true
// 	res.ErrorCode = "-"
// 	res.ErrorDescription = "-"
// 	res.Data = map[string]int64{
// 		"rows_deleted": rowDeleted,
// 	}

// 	return res, nil
// }
