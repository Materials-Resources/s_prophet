package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CrmRun struct {
	bun.BaseModel              `bun:"table:crm_run"`
	RunNumber                  int32     `bun:"run_number,type:int,pk"`
	CompanyId                  string    `bun:"company_id,type:varchar(8),pk"`
	CustomerId                 float64   `bun:"customer_id,type:decimal(19,0),pk"`
	CustomerName               string    `bun:"customer_name,type:varchar(255),nullzero"`
	Last30DaysSales            float64   `bun:"last_30_days_sales,type:decimal(19,2),nullzero"`
	Last60DaysSales            float64   `bun:"last_60_days_sales,type:decimal(19,2),nullzero"`
	Last90DaysSales            float64   `bun:"last_90_days_sales,type:decimal(19,2),nullzero"`
	Last365DaysSales           float64   `bun:"last_365_days_sales,type:decimal(19,2),nullzero"`
	Last365DaysCommCost        float64   `bun:"last_365_days_comm_cost,type:decimal(19,9),nullzero"`
	Last365DaysOtherCost       float64   `bun:"last_365_days_other_cost,type:decimal(19,9),nullzero"`
	Last365DaysFreightBilled   float64   `bun:"last_365_days_freight_billed,type:decimal(19,4),nullzero"`
	Last365DaysFreightUnbilled float64   `bun:"last_365_days_freight_unbilled,type:decimal(19,4),nullzero"`
	PreviousYearSales          float64   `bun:"previous_year_sales,type:decimal(19,2),nullzero"`
	DateAcctOpened             time.Time `bun:"date_acct_opened,type:datetime,nullzero"`
	LastHardTouchDate          time.Time `bun:"last_hard_touch_date,type:datetime,nullzero"`
	OpenOrderValue             float64   `bun:"open_order_value,type:decimal(19,9),nullzero"`
	OpenQuoteValue             float64   `bun:"open_quote_value,type:decimal(19,9),nullzero"`
	OpenOpportunityValue       float64   `bun:"open_opportunity_value,type:decimal(19,9),nullzero"`
	NumberOfOpportunities      int32     `bun:"number_of_opportunities,type:int,nullzero"`
	AverageDso                 float64   `bun:"average_dso,type:decimal(19,9),nullzero"`
	YtdPrctAmtWon              float64   `bun:"ytd_prct_amt_won,type:decimal(19,9),nullzero"`
	YtdPrctLinesWon            float64   `bun:"ytd_prct_lines_won,type:decimal(19,9),nullzero"`
	RmaValueLast365            float64   `bun:"rma_value_last_365,type:decimal(19,9),nullzero"`
	OrderValueLast365          float64   `bun:"order_value_last_365,type:decimal(19,9),nullzero"`
	Last365DaysCogs            float64   `bun:"last_365_days_cogs,type:decimal(19,9),nullzero"`
	DaysSinceLastOrder         int32     `bun:"days_since_last_order,type:int,nullzero"`
	DaysSinceFirstOrder        int32     `bun:"days_since_first_order,type:int,nullzero"`
	NumberOfOrders             int32     `bun:"number_of_orders,type:int,nullzero"`
	DaysSinceLastQuote         int32     `bun:"days_since_last_quote,type:int,nullzero"`
	DaysSinceFirstQuote        int32     `bun:"days_since_first_quote,type:int,nullzero"`
	NumberOfQuotes             int32     `bun:"number_of_quotes,type:int,nullzero"`
	SicCode                    float64   `bun:"sic_code,type:decimal(6,0),nullzero"`
	SicDescription             string    `bun:"sic_description,type:varchar(40),nullzero"`
	LeadSourceId               string    `bun:"lead_source_id,type:varchar(8),nullzero"`
	SourceDescription          string    `bun:"source_description,type:varchar(30),nullzero"`
	Class1id                   string    `bun:"class_1id,type:varchar(255),nullzero"`
	Class2id                   string    `bun:"class_2id,type:varchar(255),nullzero"`
	Class3id                   string    `bun:"class_3id,type:varchar(255),nullzero"`
	Class4id                   string    `bun:"class_4id,type:varchar(255),nullzero"`
	Class5id                   string    `bun:"class_5id,type:varchar(255),nullzero"`
	SalesrepId                 string    `bun:"salesrep_id,type:varchar(16),nullzero"`
	PhysAddress1               string    `bun:"phys_address1,type:varchar(50),nullzero"`
	PhysAddress2               string    `bun:"phys_address2,type:varchar(50),nullzero"`
	PhysCity                   string    `bun:"phys_city,type:varchar(50),nullzero"`
	PhysState                  string    `bun:"phys_state,type:varchar(50),nullzero"`
	PhysPostalCode             string    `bun:"phys_postal_code,type:varchar(10),nullzero"`
	PhysCountry                string    `bun:"phys_country,type:varchar(50),nullzero"`
	CentralPhoneNumber         string    `bun:"central_phone_number,type:varchar(20),nullzero"`
	NewCustomerFlag            string    `bun:"new_customer_flag,type:char,nullzero"`
	SalesrepName               string    `bun:"salesrep_name,type:varchar(255),nullzero"`
	InRewardsProgramFlag       string    `bun:"in_rewards_program_flag,type:char,nullzero"`
	NextTaskToComplete         string    `bun:"next_task_to_complete,type:varchar(16),nullzero"`
	LastCompletedTask          string    `bun:"last_completed_task,type:varchar(16),nullzero"`
	CustomerCategoryUid        int32     `bun:"customer_category_uid,type:int,nullzero"`
	WarehouseSizeCd            int32     `bun:"warehouse_size_cd,type:int,nullzero"`
	RetailSizeCd               int32     `bun:"retail_size_cd,type:int,nullzero"`
	Roa0To30                   float64   `bun:"roa_0_to_30,type:decimal(19,4),nullzero"`
	Roa31To60                  float64   `bun:"roa_31_to_60,type:decimal(19,4),nullzero"`
	Roa61To90                  float64   `bun:"roa_61_to_90,type:decimal(19,4),nullzero"`
	Gap0To30                   float64   `bun:"gap_0_to_30,type:decimal(19,4),nullzero"`
	Gap31To60                  float64   `bun:"gap_31_to_60,type:decimal(19,4),nullzero"`
	Gap61To90                  float64   `bun:"gap_61_to_90,type:decimal(19,4),nullzero"`
	MissedBuyFlag              string    `bun:"missed_buy_flag,type:char,nullzero"`
	CreditStatus               string    `bun:"credit_status,type:char(8),nullzero"`
	AverageOrderSize           float64   `bun:"average_order_size,type:decimal(19,4),nullzero"`
	AvgDaysBetweenOrders       int32     `bun:"avg_days_between_orders,type:int,nullzero"`
	DateLastInvoiced           time.Time `bun:"date_last_invoiced,type:datetime,nullzero"`
	TotalOrders                int32     `bun:"total_orders,type:int,nullzero"`
	TotalQuotes                int32     `bun:"total_quotes,type:int,nullzero"`
	PhysAddress3               string    `bun:"phys_address3,type:varchar(50),nullzero"`
	CostToCarryLateInvoices    float64   `bun:"cost_to_carry_late_invoices,type:decimal(19,2),default:((0))"`
}
