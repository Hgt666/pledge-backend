package task

import (
	"easy-swap/dal"
	// "easy-swap/dao"
	"easy-swap/model"
	"fmt"
	"math/big"
	"time"

	"github.com/robfig/cron"
)

// StartCalcJob 启动定时任务：每分钟跑一次（练手足够）
func StartCalcJob() {
	c := cron.New()
	// 每分钟执行一次
	err := c.AddFunc("* * * * *", func() {
		fmt.Println("【定时任务】开始计算利息、LTV、健康度...")
		CalcAllLoanInterest()
		// CalcAllDepositInterest()
		// CalcHealthAndLTV()
		fmt.Println("【定时任务】计算完成")
	})
	if err != nil {
		fmt.Println("定时任务启动失败", err)
		return
	}
	c.Start()
	fmt.Println("✅ 利息/健康度 定时任务已启动")
}

// ====================== 1. 计算所有借款利息 ======================
func CalcAllLoanInterest() {
	var loans []model.UserPledgeLoan
	dal.DB.Where("status = ?", 0).Find(&loans)

	yearSec := int64(365 * 24 * 3600)

	for _, loan := range loans {
		// 借款秒数
		duration := time.Now().Unix() - loan.CreatedAt.Unix()
		if duration <= 0 {
			continue
		}

		// 金额转 big.Int
		loanAmt, _ := new(big.Int).SetString(loan.LoanAmount, 10)
		apr := loan.InterestAmount // 这里直接用数据库的利息字段当作 APR，练手固定值 300 (3%)，实际项目中应该有专门的 APR 字段

		aprBig := big.NewInt(0)
		aprBig.SetString(apr,10)
		// 利息 = 金额 * APR * 秒数 / 年秒数
		interest := new(big.Int).Mul(loanAmt, aprBig)
		interest = new(big.Int).Mul(interest, big.NewInt(duration))
		interest = new(big.Int).Div(interest, big.NewInt(yearSec))

		// 更新利息
		dal.DB.Model(&loan).Update("interest_amount", interest.String())
	}
}

// ====================== 2. 计算所有存款利息 ======================
func CalcAllDepositInterest() {
	var deposits []model.UserDeposit
	dal.DB.Where("status = ?", 0).Find(&deposits)

	yearSec := int64(365 * 24 * 3600)

	for _, d := range deposits {
		duration := time.Now().Unix() - d.CreatedAt.Unix()
		if duration <= 0 {
			continue
		}

		amt, _ := new(big.Int).SetString(d.Amount, 10)
		apr := int64(300) // 3% 固定练手

		interest := new(big.Int).Mul(amt, big.NewInt(apr))
		interest = new(big.Int).Mul(interest, big.NewInt(duration))
		interest = new(big.Int).Div(interest, big.NewInt(yearSec*10000))

		dal.DB.Model(&d).Update("interest", interest.String())
	}
}

// ====================== 3. 计算 LTV + 健康度 ======================
func CalcHealthAndLTV() {
	var loans []model.UserPledgeLoan
	dal.DB.Where("status = ?", 0).Find(&loans)

	for _, loan := range loans {
		// 1. 获取质押资产价格
		var price model.PriceCache
		dal.DB.First(&price, "symbol = ?", "ETH")

		// 2. 计算 LTV
		collateralValue, _ := new(big.Int).SetString(loan.CollateralAmount, 10)
		loanValue, _ := new(big.Int).SetString(loan.LoanAmount, 10)

		ltv := float64(loanValue.Uint64())*price.Price / float64(collateralValue.Uint64())
		healthRate := 0.8 / ltv // 清算线 80%

		// 3. 更新
		dal.DB.Model(&loan).Updates(map[string]any{
			"health_rate": healthRate,
			"ltv":         ltv,
		})

		// 4. 健康度 <=1 标记可清算
		if healthRate <= 1 {
			dal.DB.Model(&loan).Update("status", 3)
		}
	}
}