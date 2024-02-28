package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
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
	c.rwmx.RLock()
	defer c.rwmx.RUnlock()
	return c.clientDeposit
}

func main() {

	fmt.Println("Подождите несколько секунд, пока загружаются данные!")
	client := NewClient()

	//creation of the gorutin group.
	wgDeposit, wgWithdrawal := sync.WaitGroup{}, sync.WaitGroup{}

	for i := 0; i <= GOROUTINE_DEPOSIT; i++ {

		//r is a variable in which a random number is written in the range from 0 to 10 for financial transactions.
		r := rand.Intn(ALPHA_FOR_DELAY) + BETA_FOR_DELAY

		//delay for goroutine.
		time.Sleep(time.Duration(r) * 100 * time.Millisecond)

		//setting the wgDeposit group counter to 1.
		wgDeposit.Add(1)
		go func() {
			random := rand.Intn(RANDOM_DEPOSIT)
			client.Deposit(random)

			//Indicates that the group element has completed its execution. The internal counter decreases its value.
			wgDeposit.Done()
		}()
	}

	for i := 0; i <= GOROUTINE_WITHDRAWAL; i++ {
		r := rand.Intn(ALPHA_FOR_DELAY) + BETA_FOR_DELAY
		time.Sleep(time.Duration(r) * 100 * time.Millisecond)
		wgWithdrawal.Add(1)
		go func() {
			random := rand.Intn(RANDOM_WITHDRAWAL)
			err := client.Withdrawal(random)
			if err != nil {
				fmt.Println(err)
			}
			wgWithdrawal.Done()
		}()
	}

	//awaiting completion of all goroutines from the group.
	wgDeposit.Wait()
	wgWithdrawal.Wait()

	menu(client)

}

// Menu is a function that displays the menu and also reads commands from the console.
// According to the entered command, one of the methods of the Client structure is called.
func menu(client *Client) {
	fmt.Println("_____________________")
	fmt.Println("Добро пожаловать в мини приложение банка ФГ 'Бюро'!\n" +
		"В этом приложении Вы можете:\n" +
		"- узнать баланс Вашего счета, команда - balance\n" +
		"- пополнить баланс Вашего счета, команда - deposit\n" +
		"- снять денежные средства с Вашего счета, команда - withdrawal\n" +
		"- выйти из аккаунта, команда - exit")

	scanner := bufio.NewScanner(os.Stdin)
	enterCommand()

	for scanner.Scan() {
		command := scanner.Text()
		switch BankCommand(command) {
		case balanceCommand:
			fmt.Println("Ваш баланс - ", client.Balance())
			enterCommand()
		case depositCommand:
			fmt.Print("Введите сумму зачисления: ")
			amount, err := scan(scanner)
			if err != nil {
				fmt.Println(err)
				enterCommand()
				continue
			} else {
				client.Deposit(amount)
				fmt.Println("Денежные средства успешно зачислены на Ваш счёт!")
				fmt.Println("Ваш баланс - ", client.Balance())
				enterCommand()
			}
		case withdrawalCommand:
			fmt.Print("Введите сумму списания: ")
			amount, err := scan(scanner)
			if err != nil {
				fmt.Println(err)
				enterCommand()
				continue
			}
			if err := client.Withdrawal(amount); err != nil {
				fmt.Println(err)
				enterCommand()
				continue
			}
			fmt.Println("Денежные средства успешно сняты с Вашего счёта!")
			fmt.Println("Ваш баланс - ", client.Balance())
			enterCommand()
		case exitCommand:
			fmt.Println("Будем рады видеть Вас снова! До свидания!")
			os.Exit(1)
		default:
			fmt.Println("Неизвестная команда! Вы можете использовать команды: balance , deposit , withdrawal , exit")
			enterCommand()
			continue
		}

	}
}

// Scan a function for reading the amount of deposit/debit. This function checks that a number has been entered by the user.
// If a non-number is entered, an error is returned.
func scan(scanner *bufio.Scanner) (int, error) {

	scanner.Scan()
	amount, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, errors.New("Введено некоректное число!")
	}
	return amount, nil
}

// EnterCommand is function displays a prompt to enter a command. You can write any text into it.
func enterCommand() {
	fmt.Print("\nВведите команду: ")
}
