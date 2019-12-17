package 账单项目

import "fmt"

func menu() {

	key := ""    // 用来介绍用户输入的指令
	loop := true //
	flag := false
	// 设定账本金额
	balacne := 100000.00
	// 每次收支的金额
	money := 00.00
	// 每次收支的说明
	note := ""
	details := "收支\t账本金额\t收支金额\t 说	明"

	for {
		fmt.Println("-----------------家庭收入支出菜单------------------------")
		fmt.Println("					1.收入明细")
		fmt.Println("					2.登记收入")
		fmt.Println("					3.登记支出")
		fmt.Println("					4.退出软件")
		fmt.Println("请输入（1-4）：")
		fmt.Scanln(&key)
		switch key {
		case "1":
			if !flag {
				fmt.Println("没有交易计入")
				break
			}
			fmt.Println(details)
		case "2":
			fmt.Println("本次收入:")
			fmt.Scanln(&money)
			balacne += money
			fmt.Printf("本次收入说明本次收入:%v", note)
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balacne, money, note)
			flag = true
		case "3":
			fmt.Println("登记支出:")
			fmt.Scanln(&money)
			if money > balacne {
				fmt.Println("您的余额不足支出")
				break

			}
			balacne -= money
			fmt.Printf("本次支出说明本次收入:%v", note)
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balacne, money, note)
			flag = true
		case "4":
			loop = false
		case "你好":
			fmt.Println("Watch??")
		default:
			fmt.Println("请输入正确选择")
		}
		if !loop {
			break
		}
	}
	fmt.Println("退出")
}

type FamilyAccount struct {
	key     string
	loop    bool
	money   float64
	balance float64
	flag    bool
	note    string
	details string
}
// 工程模式调度
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		money:   00.00,
		balance: 100000.00,
		flag:    false,
		note:    "",
		details: "收支\t账本金额\t收支金额\t 说	明",
	}

}



func (this *FamilyAccount) MainMenu() {
	// 菜单
	for {
		fmt.Println("-----------------家庭收入支出菜单------------------------")
		fmt.Println("					1.收入明细")
		fmt.Println("					2.登记收入")
		fmt.Println("					3.登记支出")
		fmt.Println("					4.退出软件")
		fmt.Println("请输入（1-4）：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确选择")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("退出")

}




func (this *FamilyAccount)showDetails()  {
	// 账本显示
	if !this.flag {
		fmt.Println("没有交易计入")
		return
	}
	fmt.Println(this.details)

}


func (this *FamilyAccount)  income()  {
	// 收入
	fmt.Println("本次收入:")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Printf("本次收入说明本次收入:%v", this.note)
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true

}

func (this *FamilyAccount) pay() {
	// 支出方法
	fmt.Println("登记支出:")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("您的余额不足支出")
		return

	}
	this.balance -= this.money
	fmt.Printf("本次支出说明本次收入:%v", this.note)
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true

}

func (this *FamilyAccount) exit() {
	this.loop = false

}

func main() {

	//menu()
	b :=NewFamilyAccount()
	b.MainMenu()
}
