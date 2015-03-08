package main

import (
	"fmt"
	"math"
)

func main() {
	//GetThreeBitNum()
	//LadderOfProfitDistribution()
	CompleteSquareNumber()
}

/*
有 1、2、3、4 个数字，能组成多少个互不相同且无重复数字的三位数？都是多少？

程序分析：
可填在百位、十位、个位的数字都是 1、2、3、4。
组成所有的排列后再去 掉不满足条件的排列。
*/
func GetThreeBitNum() {
	/*
	   一下为三重循环
	*/
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {
				/*确保 i 、j 、k 三位互不相同*/
				if i != k && i != j && j != k {
					fmt.Println(i, j, k)
				}
			}
		}
	}

}

/*
阶梯利润分配
题目：企业发放的奖金根据利润提成。利润(I)低于或等于10万元时，奖金可提成10%；利润高于10万元，低于20万元，
低于10万元的部分按10%提成，高于10万元的部分，可提成7.5%；20万到40万之间时，高于20万元的部分，可提成5%；
40万到60万之间时高于40万元的部分，可提成3%；60万到100万之间时，高于60万元的部分，可提成1.5%，高于100万元时，
超过100万元的部分按1%提成，从键盘输入当月利润I，求应发放奖金总数？
1.程序分析：请利用数轴来分界，定位。注意定义时需把奖金定义成长整型。
2.程序源代码：
*/
func LadderOfProfitDistribution() {
	var I float32 = 0.0
	var bonus float32 = 0.0
	fmt.Println("请输入利润")
	fmt.Scanf("%f\n", &I)
	switch {
	case I > 1000000:
		bonus = (I - 1000000) * 0.01
		I = 1000000
		fallthrough ////必须是最后一个语句
	case I > 600000:
		bonus += (I - 600000) * 0.015
		I = 600000
		fallthrough
	case I > 400000:
		bonus += (I - 400000) * 0.03
		I = 400000
		fallthrough
	case I > 200000:
		bonus += (I - 200000) * 0.05
		I = 200000
		fallthrough
	case I > 100000:
		bonus += (I - 100000) * 0.075
		I = 100000
		fallthrough
	default:
		bonus += I * 0.1
	}
	fmt.Printf("提成总计：%f\n", bonus)
}

/*
题目：一个整数，它加上 100 后是一个完全平方数，再加上 268 又是一个完全平方数，请问
该数是多少？
1.程序分析：在10万以内判断，先将该数加上100后再开方，再将该数加上268后再开方，如果开方后的结果满足如下条件，即是结果。请看具体分析：
*/
func CompleteSquareNumber() {
	fmt.println("enter CompleteSquareNumber")
	i := 0
	for {
		x := int(math.Sqrt(float64(i + 100)))
		y := int(math.Sqrt(float64(i + 168)))
		if x+x == i+100 && y*y == i+168 {
			fmt.Println(i)
			break
		}
		fmt.Println(i)
		i++
	}
}
