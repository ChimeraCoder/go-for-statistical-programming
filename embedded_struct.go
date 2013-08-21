type Person struct {
    MainAccount CheckingAccount
}

type Bank struct {
    Accounts []CheckingAccount
    SavingsAccounts []struct{
        Balance int
        InterestRate float64
    }
}
