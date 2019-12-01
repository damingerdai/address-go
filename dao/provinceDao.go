package dao

import (
	"damingerdai/address/config"
	"damingerdai/address/database"
	"damingerdai/address/models"
)

var conf *config.DBConfig

func init() {
	env := config.New()
	a := config.NewDBConfig(env)
	conf = &a
}

func ListProvinces() []*models.Province {
	conn := database.GetConnection(conf)
	defer conn.Close()

	rows, err := conn.Query("SELECT _id, name, province_id FROM province")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	result := make([]*models.Province, 32)
	for rows.Next() {
		//将行数据保存到record字典
		var id, provinceId int
		var name string
		err = rows.Scan(&id, &name, &provinceId)
		if err != nil {
			panic(err.Error())
		}
		province := models.Province{
			Id:         id,
			Name:       name,
			ProvinceId: provinceId,
		}

		result = append(result, &province)

	}

	return result

}
