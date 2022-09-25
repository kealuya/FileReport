package handler

const (
	Select_Current_Header = `SELECT
								p.*,
								COUNT( DISTINCT d.doc_id ) AS doc_nums,
								count( DISTINCT f.update_user_id ) AS peo_nums,
								COUNT( f.update_date ) AS versions 
							FROM
								"Project" p
								LEFT JOIN Doc d ON d.pro_id = p.pro_id
								LEFT JOIN File f ON f.doc_id = d.doc_id 
							GROUP BY
								p.pro_id`
	Select_Product_Header = `SELECT
							f.doc_id,
							f.update_user_id,
							MAX( update_date ) AS update_date,
							u.user_name,
							p.pro_name,
							p.pro_id,
							u.phone_number 
						FROM
							File f
							LEFT JOIN Doc d ON d.doc_id = f.doc_id
							LEFT JOIN User u ON u.phone_number = f.update_user_id
							LEFT JOIN Project p ON p.pro_id = d.pro_id 
						GROUP BY
							p.pro_id`
	Select_doc_File = `SELECT
						doc.doc_id,
						doc.doc_type,
						file.version_show,
						doc.doc_name,
						file.file_name,
						doc.owner_id,
						doc.pro_id,
						doc.create_date,
						doc.is_discard,
						doc.is_owner_edit,
						doc.is_release,
						user.user_name AS owner,
						file.update_content,
						MAX( file.update_date ) AS update_date,
						update_user.user_name AS update_user 
					FROM
						doc
						INNER JOIN File file  ON file.doc_id = doc.doc_id
						INNER JOIN User user ON user.phone_number = doc.owner_id 
						INNER JOIN User update_user ON update_user.phone_number = file.update_user_id 
					GROUP BY
						doc.doc_id`
	Select_Last_Month_Update = `SELECT
							file.update_date,
							doc.doc_type,
							file.version_show,
							doc.doc_name,
							file.file_name,
							doc.owner_id,
							user.user_name AS owner,
							doc.doc_id,
							doc.pro_id,
							doc.is_discard,
							doc.is_owner_edit,
							doc.is_release,
							file.update_content,
						update_user.user_name AS update_user
						FROM
							File file
							INNER JOIN Doc doc ON doc.doc_id = file.doc_id
							INNER JOIN User user ON user.phone_number = doc.owner_id 
							INNER JOIN User update_user ON update_user.phone_number = file.update_user_id 
							WHERE file.update_date BETWEEN date('now', "-1 month") AND date('now',"+1 day")`
	Select_User = `SELECT * FROM user_info WHERE userid = ? `

	Update_User = `Update user_info SET user_role = ?,password = ? , modifier = ? , modify_time = ? WHERE userid = ? `

	Select_ProductInfo_By_Name = `SELECT product_name,last_update_time,last_creater FROM product_info  WHERE product_name = ?`

	Update_Product_info = `UPDATE product_info SET last_creater = ?,last_update_time=? WHERE product_name=?`

	Update_File = `UPDATE file_info SET major_version=?,minor_version=? , modifier = ? , modify_time = ? WHERE file_name = ? AND product_name = ?`

	Select_Latest_Trend = `SELECT *  FROM file_record  `

	Update_File_Info = `Update file_info SET major_version=?,minor_version=? , modifier = ? , modify_time = ? WHERE file_name = ? and product_name = ?`

	Update_Publish_File   = `Update file_info SET abolish_flag = ?,publish_flag = ? , modifier = ? , modify_time = ? WHERE file_name = ? and product_name = ?`
	Update_Authority_File = `Update file_info SET edit_flag = ?,publish_flag = ? , modifier = ? , modify_time = ? WHERE file_name = ? and product_name = ?`
	Update_Abolish_File   = `Update file_info SET abolish_flag = ?,publish_flag = ? , modifier = ? , modify_time = ? WHERE file_name = ? and product_name = ?`

	Select_Recent_Update = `SELECT file_name,major_version,minor_version,product_name,abolish_flag,publish_flag FROM file_info 
							WHERE modify_time BETWEEN date('now', "-1 month") AND date('now',"+1 day")`
	Select_File_Info = `SELECT file_name,major_version,minor_version,product_name,abolish_flag,publish_flag from file_info WHERE file_name = ? AND product_name = ? `
)
