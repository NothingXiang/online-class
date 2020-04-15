package req

// 检查分页参数
func CheckPage(page, limit int) bool {

	return page > 0 && limit > 0
}
