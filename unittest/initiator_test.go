package unittest

import (
	"github.com/dwiangraeni/dealls/interfaces"
	"github.com/dwiangraeni/dealls/service"
	"github.com/dwiangraeni/dealls/utils"
)

type MockAccountService struct {
	accountRepo interfaces.IAccountRepo
	hashCursor  utils.HashInterface
}

func MockNewAccountService(ms MockAccountService) interfaces.IAccountService {
	return service.NewAccountService(ms.accountRepo)
}

type MockAuthService struct {
	accountRepo interfaces.IAccountRepo
	publicKey   string
	privateKey  string
	utilsPass   utils.PasswordHasher
}

func MockNewAuthService(ms MockAuthService) interfaces.IAuthService {
	return service.NewAuthService(ms.accountRepo, ms.publicKey, ms.privateKey, ms.utilsPass)
}

type MockPremiumPackageService struct {
	accountRepo        interfaces.IAccountRepo
	premiumPackageRepo interfaces.IPremiumPackageRepo
	hashCursor         utils.HashInterface
	transactionRepo    interfaces.ITransactionRepo
}

func MockNewPremiumPackageService(ms MockPremiumPackageService) interfaces.IPremiumPackageService {
	return service.NewPremiumPackageService(ms.accountRepo, ms.premiumPackageRepo, ms.transactionRepo)
}

type MockUserSwipeLogService struct {
	userSwipeLogRepo   interfaces.IUserSwipeLogRepo
	accountRepo        interfaces.IAccountRepo
	premiumPackageRepo interfaces.IPremiumPackageRepo
	maxSwipeADay       int
}

func MockNewUserSwipeLogService(ms MockUserSwipeLogService) interfaces.IUserSwipeLogService {
	return service.NewUserSwipeLogService(ms.userSwipeLogRepo, ms.accountRepo, ms.premiumPackageRepo, ms.maxSwipeADay)
}
