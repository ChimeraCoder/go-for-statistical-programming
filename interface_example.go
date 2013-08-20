type Account interface {
    Deposit(int) error
}

type CheckingAccount struct {
    Balance int 
        superSecretId int64
}

func (destination CheckingAccount) Deposit(amount int) error {
    destination.Balance += amount
    return nil
}
