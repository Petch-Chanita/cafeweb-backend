package dto

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
	CafeID   string `json:"cafe_id" binding:"required"`
}

type RequestData struct {
	ID     string `json:"user_id"` // รับ id ของผู้ใช้
	CafeID string `json:"cafe_id"` // รับ cafe_id ที่เกี่ยวข้อง
}
