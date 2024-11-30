package repo

var (
	// premium package
	RepoGetListPremiumPackage = `
	SELECT "id", "package_uid", "title", "description", "price", "is_active", "created_at", "updated_at",
	"created_by", "updated_by" FROM premium_package WHERE is_active IS TRUE 
	%s %s %s;`
	
	// premium package user
	RepoGetPremiumPackageUserByAccountID = `
	SELECT "id", "premium_package_id", "account_id", "purchase_date" FROM premium_package_user WHERE account_id = $1;`
)
