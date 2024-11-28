package response

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	Username      string `json:"username"`
	AccountMaskID string `json:"account_mask_id"`
	AccountRole   string `json:"account_role"`
	Name          string `json:"name"`
	CreatedAt     string `json:"created_at"`
	CreatedBy     string `json:"created_by"`
	UpdatedAt     string `json:"updated_at"`
	UpdatedBy     string `json:"updated_by"`
}
