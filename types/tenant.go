package types

type Tenants struct {
	ID           string `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	Email        string `db:"email" json:"email"`
	Phone        string `db:"phone" json:"phone"`
	CompanySize  string `db:"company_size" json:"company_size"`
	SupportURL   string `db:"support_url" json:"support_url"`
	SupportEmail string `db:"support_email" json:"support_email"`
	CreatedAt    string `db:"created_at" json:"created_at"`
}

type TenantSource struct {
	TenantId string `db:"tenant_id" json:"tenant_id"`
	SourceId string `db:"source_id" json:"source_id"`
}

type TenantDTO struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	CompanySize string `json:"company_size" binding:"required"`
}
