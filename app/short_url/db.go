package short_url

import (
	"ShortUrl/common/db"
	"log"
)

// GetUrlInfo 获取链接信息
func GetUrlInfo() []*UrlInfo {
	conn := db.MySQLDatabase()
	sql := `SELECT 
    			id,
				ref_code,
				origin_url,
				page_view 
			FROM 
			    url_info`
	infos := make([]*UrlInfo, 0)
	err := conn.Select(&infos, sql)
	if err != nil {
		log.Println("查询出错:", err)
		return infos
	}
	return infos
}

// SaveUrlInfo 保存链接信息
func SaveUrlInfo(ui *UrlInfo) {
	conn := db.MySQLDatabase()
	rs, err := conn.Exec("INSERT INTO url_info(ref_code,origin_url) VALUE(?,?)", ui.RefCode, ui.OriginUrl)
	if err != nil {
		log.Println("插入新数据出错", err)
		return
	}
	id, err := rs.LastInsertId()
	log.Printf("插入新数据ID%d", id)
}

// UpdatePageViewInfo 更新链接访问量
func UpdatePageViewInfo(ui *UrlInfo) {
	conn := db.MySQLDatabase()
	rs, err := conn.Exec("UPDATE url_info SET page_view = ? WHERE ref_code = ?", ui.PageView, ui.RefCode)
	if err != nil {
		log.Println("插入新数据出错", err)
		return
	}
	id, err := rs.LastInsertId()
	log.Printf("插入新数据ID%d", id)
}

// DeleteUrlInfo 删除链接信息
func DeleteUrlInfo(ui *UrlInfo) {
	conn := db.MySQLDatabase()
	rs, err := conn.Exec("DELETE FROM url_info WHERE ref_code = ?", ui.RefCode)
	if err != nil {
		log.Println("插入新数据出错", err)
		return
	}
	id, err := rs.LastInsertId()
	log.Printf("插入新数据ID%d", id)
}

// PageViewRank 链接访问量排行榜
func PageViewRank() []*UrlInfo {
	conn := db.MySQLDatabase()
	sql := `SELECT 
    			id,
				ref_code,
				origin_url,
				page_view 
			FROM 
			    url_info
			ORDER BY page_view DESC	`
	infos := make([]*UrlInfo, 0)
	err := conn.Select(&infos, sql)
	if err != nil {
		log.Println("查询出错:", err)
		return infos
	}
	return infos
}
