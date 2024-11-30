package manager

import (
	"github.com/dwiangraeni/dealls/infra"
	"github.com/dwiangraeni/dealls/interfaces"
	"github.com/dwiangraeni/dealls/repo"
	"sync"
)

type RepoManager interface {
	AccountRepoManager() interfaces.IAccountRepo
	UserSwipeLogRepoManager() interfaces.IUserSwipeLogRepo
	PremiumPackageRepoManager() interfaces.IPremiumPackageRepo
	TransactionRepoManager() interfaces.ITransactionRepo
}

type repoManager struct {
	infra infra.Infra
}

func NewRepoManager(infra infra.Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}

var (
	accountRepoOnce sync.Once
	accountRepo     interfaces.IAccountRepo
)

func (r *repoManager) AccountRepoManager() interfaces.IAccountRepo {
	accountRepoOnce.Do(func() {
		accountRepo = repo.NewAccountRepo(r.infra.SQLDB())
	})

	return accountRepo
}

var (
	userSwipeLogRepoOnce sync.Once
	userSwipeLogRepo     interfaces.IUserSwipeLogRepo
)

func (r *repoManager) UserSwipeLogRepoManager() interfaces.IUserSwipeLogRepo {
	userSwipeLogRepoOnce.Do(func() {
		userSwipeLogRepo = repo.NewUserSwipeLogRepo(r.infra.SQLDB())
	})

	return userSwipeLogRepo
}

var (
	premiumPackageRepoOnce sync.Once
	premiumPackageRepo     interfaces.IPremiumPackageRepo
)

func (r *repoManager) PremiumPackageRepoManager() interfaces.IPremiumPackageRepo {
	premiumPackageRepoOnce.Do(func() {
		premiumPackageRepo = repo.NewPremiumPackageRepo(r.infra.SQLDB())
	})

	return premiumPackageRepo
}

var (
	transactionRepoOnce sync.Once
	transactionRepo     interfaces.ITransactionRepo
)

func (r *repoManager) TransactionRepoManager() interfaces.ITransactionRepo {
	transactionRepoOnce.Do(func() {
		transactionRepo = repo.NewTransactionRepo(r.infra.SQLDB())
	})

	return transactionRepo
}
