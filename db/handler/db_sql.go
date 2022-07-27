package handler

const (
	Select_User = `SELECT * FROM user_info WHERE userid = ? `

	Update_User = `Update user_info SET user_role = ?,password = ? , modifier = ? , modify_time = ? WHERE userid = ? `

	Select_ProductInfo_By_Name = `SELECT product_name,last_update_time,last_creater FROM product_info  WHERE product_name = ?`

	Update_Product_info = `UPDATE product_info SET last_creater = ?,last_update_time=? WHERE product_name=?`

	Update_File = `UPDATE file_info SET major_version=?,minor_version=? , modifier = ? , modify_time = ? WHERE file_name = ? AND product_name = ?`

	Select_Current_Header = `SELECT product_name, COUNT( DISTINCT file_name ) AS file_number, COUNT( DISTINCT creater ) AS person_number, COUNT( * ) AS version_count
							FROM file_record  GROUP BY product_name`

	Select_Latest_Trend = `SELECT *  FROM file_record  `

	Update_File_Info = `Update file_info SET major_version=?,minor_version=? , modifier = ? , modify_time = ? WHERE file_name = ? and product_name = ?`

	Update_Publish_File = `Update file_info SET abolish_flag = ?,publish_flag = ? , modifier = ? , modify_time = ? WHERE file_name = ? and product_name = ?`

	Update_Abolish_File = `Update file_info SET abolish_flag = ?,publish_flag = ? , modifier = ? , modify_time = ? WHERE file_name = ? and product_name = ?`

	Select_Product_Header = `SELECT id,product_name,last_update_time,last_creater FROM product_info `

	Select_Recent_Update = `SELECT file_name,major_version,minor_version,product_name,abolish_flag,publish_flag FROM file_info 
							WHERE modify_time BETWEEN date('now', "-1 month") AND date('now',"+1 day")`
	Select_File_Info = `SELECT file_name,major_version,minor_version,product_name,abolish_flag,publish_flag from file_info WHERE file_name = ? AND product_name = ? `
)
