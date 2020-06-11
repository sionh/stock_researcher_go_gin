package entity

import (
	"github.com/jinzhu/gorm"
)

type Shikiho struct {
	gorm.Model
	Year                  int     `gorm:"type:INTEGER;index:indexShikihosOn_year"`
	Quarter               int     `gorm:"type:INTEGER;index:indexShikihosOn_quarter"`
	Code                  int     `gorm:"type:INTEGER;index:indexShikihosOnCode"`
	CompName              string  `gorm:"type:TEXT"`
	URL                   string  `gorm:"type:TEXT"`
	Feature               string  `gorm:"type:TEXT"`
	Topic                 string  `gorm:"type:TEXT"`
	Outlook               string  `gorm:"type:TEXT"`
	NextSales             int     `gorm:"type:INTEGER"`
	NextOperationgIncome  int     `gorm:"type:INTEGER"`
	NextOrdinaryIncome    int     `gorm:"type:INTEGER"`
	NextNetIncome         int     `gorm:"type:INTEGER"`
	NextNetIncomePerShare float32 `gorm:"type:FLOAT"`
	NextDividend          float32 `gorm:"type:FLOAT"`
	PrevSales             int     `gorm:"type:INTEGER"`
	PrevOperatingIncome   int     `gorm:"type:INTEGER"`
	PrevOrdinaryIncome    int     `gorm:"type:INTEGER"`
	PrevNetIncome         int     `gorm:"type:INTEGER"`
	PrevNetIncomePerShare float32 `gorm:"type:FLOAT"`
	PrevDividend          float32 `gorm:"type:FLOAT"`
	OperationgCf          int     `gorm:"type:INTEGER"`
	PrevOperatingCf       int     `gorm:"type:INTEGER"`
	InvestmentCf          int     `gorm:"type:INTEGER"`
	PrevInvestmentCf      int     `gorm:"type:INTEGER"`
	FinancingCf           int     `gorm:"type:INTEGER"`
	PrevFinancingCf       int     `gorm:"type:INTEGER"`
	Cash                  int     `gorm:"type:INTEGER"`
	PrevCash              int     `gorm:"type:INTEGER"`
	TotalAssets           int     `gorm:"type:INTEGER"`
	OwnedCapital          int     `gorm:"type:INTEGER"`
	OwnedCapitalRatio     string  `gorm:"type:TEXT"`
	Capital               int     `gorm:"type:INTEGER"`
	RetainedEarnings      int     `gorm:"type:INTEGER"`
	InterestBearingDebt   int     `gorm:"type:INTEGER"`
	Establishment         string  `gorm:"type:TEXT"`
	StockListing          string  `gorm:"type:TEXT"`
	Settlement            string  `gorm:"type:TEXT"`
	TypeOfBusiness        string  `gorm:"type:TEXT"`
	Supplier              string  `gorm:"type:TEXT"`
	NumOfEmployee         string  `gorm:"type:TEXT"`
	HeadOffice            string  `gorm:"type:TEXT"`
	CompareCompany        string  `gorm:"type:TEXT"`
	Certificate           string  `gorm:"type:TEXT"`
	SalesContact          string  `gorm:"type:TEXT"`
	Consolidated          string  `gorm:"type:TEXT"`
	ConsolidatedBusiness  string  `gorm:"type:TEXT"`
	Bank                  string  `gorm:"type:TEXT"`
}
