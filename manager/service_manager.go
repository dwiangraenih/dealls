package manager

import (
	"github.com/dwiangraeni/dealls/infra"
	"github.com/dwiangraeni/dealls/interfaces"
	"github.com/dwiangraeni/dealls/middleware"
	"github.com/dwiangraeni/dealls/service"
	"sync"
)

type ServiceManager interface {
	AuthService() interfaces.IAuthService
	AccountService() interfaces.IAccountService
	AccountManager() middleware.AccountToken
	UserSwipeLogService() interfaces.IUserSwipeLogService
	PremiumPackageService() interfaces.IPremiumPackageService
}

type serviceManager struct {
	infra infra.Infra
	repo  RepoManager
}

func NewServiceManager(infra infra.Infra) ServiceManager {
	return &serviceManager{
		repo:  NewRepoManager(infra),
		infra: infra,
	}
}

var (
	authServiceOnce sync.Once
	authService     interfaces.IAuthService
)

func (s *serviceManager) AuthService() interfaces.IAuthService {
	authServiceOnce.Do(func() {
		key := s.infra.Config().Sub("rsa")
		authService = service.NewAuthService(
			s.repo.AccountRepoManager(),
			key.GetString("public_key"),
			key.GetString("private_key"))
	})
	return authService
}

var (
	accountServiceOnce sync.Once
	accountService     interfaces.IAccountService
)

func (s *serviceManager) AccountService() interfaces.IAccountService {
	accountServiceOnce.Do(func() {
		accountService = service.NewAccountService(s.repo.AccountRepoManager())
	})
	return accountService
}

var (
	accountManagerOnce sync.Once
	accountManager     middleware.AccountToken
)

func (s *serviceManager) AccountManager() middleware.AccountToken {
	accountManagerOnce.Do(func() {
		key := s.infra.Config().Sub("rsa")
		accountManager = middleware.NewAccountToken(key.GetString("public_key"))
	})

	return accountManager
}

var (
	userSwipeLogServiceOnce sync.Once
	userSwipeLogService     interfaces.IUserSwipeLogService
)

func (s *serviceManager) UserSwipeLogService() interfaces.IUserSwipeLogService {
	userSwipeLogServiceOnce.Do(func() {
		userSwipeLogService = service.NewUserSwipeLogService(s.repo.UserSwipeLogRepoManager(), s.repo.AccountRepoManager())
	})
	return userSwipeLogService
}

var (
	premiumPackageServiceOnce sync.Once
	premiumPackageService     interfaces.IPremiumPackageService
)

func (s *serviceManager) PremiumPackageService() interfaces.IPremiumPackageService {
	premiumPackageServiceOnce.Do(func() {
		premiumPackageService = service.NewPremiumPackageService(s.repo.AccountRepoManager(), s.repo.PremiumPackageRepoManager(), s.repo.TransactionRepoManager())
	})
	return premiumPackageService
}
