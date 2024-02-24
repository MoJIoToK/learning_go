package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sync"
)

type BankCommand string

const (
	GOROUTINE_DEPOSIT                = 10
	GOROUTINE_WITHDRAWAL             = 5
	ALPHA_FOR_DELAY                  = 10 - 5
	BETA_FOR_DELAY                   = 6
	RANDOM_DEPOSIT                   = 11
	RANDOM_WITHDRAWAL                = 6
	balanceCommand       BankCommand = "balance"
	depositCommand       BankCommand = "deposit"
	withdrawalCommand    BankCommand = "withdrawal"
	exitCommand          BankCommand = "exit"
)

// BankClient is interface that will be implemented by the structure Client.
type BankClient interface {
	// Deposit deposits given amount to clients account.
	Deposit(amount int)

	// Withdrawal withdraws given amount from clients account.
	// return error if clients balance less the withdrawal amount
	Withdrawal(amount int) error

	// Balance returns clients balance
	Balance() int
}

// Client is structure of client of the bank.
type Client struct {
	rwmx          sync.RWMutex
	clientDeposit int
}

// NewClient is constructor for struct - Client
func NewClient() *Client {
	return &Client{}
}

// Implemented method from the BankClient interface.
func (c *Client) Deposit(amount int) {
	c.rwmx.Lock()
	defer c.rwmx.Unlock()
	c.clientDeposit += amount
}

// Implemented method from the BankClient interface.
func (c *Client) Withdrawal(amount int) error {
	c.rwmx.Lock()
	defer c.rwmx.Unlock()
	if amount > c.clientDeposit {
		return errors.New("Операция отклонена! На счёте недостаточно средств!")
	}
	c.clientDeposit -= amount
	return nil
}

// Implemented method from the BankClient interface.
func (c *Client) Balance() int {
	c.rwmx.Lock()
	defer c.rwmx.Unlock()
	return c.clientDeposit
}

func main() {

	client := NewClient()
	//wgDeposit, wgWithdrawal := sync.WaitGroup{}, sync.WaitGroup{}
	//
	//for i := 0; i <= GOROUTINE_DEPOSIT; i++ {
	//	r := rand.Intn(ALPHA_FOR_DELAY) + BETA_FOR_DELAY
	//	time.Sleep(time.Duration(r) * 100 * time.Millisecond)
	//	wgDeposit.Add(1)
	//	go func() {
	//		random := rand.Intn(RANDOM_DEPOSIT)
	//		//balanceBeforeDepo := client.Balance()
	//		client.Deposit(random)
	//		//fmt.Printf("Goroutine Deposit - %d. Balance before deposit - %d. Deposit - %d. Balance After - %d.\n", i, balanceBeforeDepo, random, client.Balance())
	//		wgDeposit.Done()
	//	}()
	//}
	//
	//for i := 0; i <= GOROUTINE_WITHDRAWAL; i++ {
	//	r := rand.Intn(ALPHA_FOR_DELAY) + BETA_FOR_DELAY
	//	time.Sleep(time.Duration(r) * 100 * time.Millisecond)
	//	wgWithdrawal.Add(1)
	//	go func() {
	//		random := rand.Intn(RANDOM_WITHDRAWAL)
	//		//balanceBeforeWith := client.Balance()
	//		err := client.Withdrawal(random)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		//fmt.Printf("Goroutine Withdrawal - %d. Balance before Withdrawal - %d. Withdrawal - %d. Balance After - %d.\n", i, balanceBeforeWith, random, client.Balance())
	//		wgWithdrawal.Done()
	//	}()
	//}
	//
	//wgDeposit.Wait()
	//wgWithdrawal.Wait()

	menu(client)

}

func menu(client *Client) {
	fmt.Println("_____________________")
	fmt.Println("Добро пожаловать в мини приложение банка ФГ 'Бюро'!\n" +
		"В этом приложении Вы можете:\n " +
		"- узнать баланс Вашего счета, команда - balance\n" +
		"- пополнить баланс Вашего счета, команда - deposit\n" +
		"- снять денежные средства с Вашего счета, команда - withdrawal\n" +
		"- выйти из аккаунта, команда - exit")

	scanner := bufio.NewScanner(os.Stdin)
	var amount int
	for scanner.Scan() {
		command := scanner.Text()
		switch BankCommand(command) {
		case balanceCommand:
			fmt.Println("Ваш баланс - ", client.Balance())
		case depositCommand:
			fmt.Println("Введите сумму зачисления: ")
			amount, err := scan(amount)
			if err != nil {
				fmt.Println(err)
				break
			} else {
				client.Deposit(amount)
				fmt.Println("Денежные средства успешно зачислены на Ваш счёт!")
			}
			//fmt.Scanln(&amount)
			fmt.Println("Ваш баланс - ", client.Balance())
		case withdrawalCommand:
			fmt.Println("Введите сумму списания: ")
			amount, err := scan(amount)
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Scanln(&amount)
			client.Withdrawal(amount)
			fmt.Println("Денежные средства успешно сняты с Вашего счёта!")
			fmt.Println("Ваш баланс - ", client.Balance())
		case exitCommand:
			fmt.Println("Будем рады видеть Вас снова! До свидания!")
			os.Exit(1)
		default:
			fmt.Println("Неизвестная команда! Вы можете использовать команды: balance , deposit , withdrawal , exit")
			continue
		}

	}
}

func scan(amount int) (int, error) {

	_, err := fmt.Scan(&amount)
	if err != nil {
		return 0, errors.New("Введено некоректное число!")
	}
	return amount, nil
}

//В момент старта запускает 10 горутин, каждая из которых с промежутком от 0.5 секунд
//до 1 секунды зачисляет на счёт клиента случайную сумму от 1 до 10.

//Так же запускается 5 горутин, которые с промежутком 0.5 секунд до 1 секунды снимают с клиента случайную сумму
//от 1 до 5. Если снятие невозможно, в консоль выводится сообщение об ошибке, и приложение продолжает работу.

//Если пользователь введет в консоль слово balance — приложение выведет в консоль текущий баланс клиента.
//deposit — запрашивается сумма (целое число) — которая добавляется на счёт пользователя
//withdrawal — запрашивается сумма (целое число) — которая выводится со счёта пользователя.Если запрашиваемая сумма больше текущего баланса пользователя, то пользователю должно быть показано сообщение о том, что его баланс недостаточен (и, очевидно, операция не должна быть выполнена).
//Если пользователь введет слово exit — приложение завершит работу.
//При вводе любой другой команды приложение выведет сообщение "Unsupported command. You can use commands: balance, deposit, withdrawal, exit".
