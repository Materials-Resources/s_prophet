package model

import (
	"github.com/uptrace/bun"
	"time"
)

type PricePage struct {
	bun.BaseModel              `bun:"table:price_page"`
	PricePageUid               int32     `bun:"price_page_uid,type:int,pk"`
	PricePageTypeCd            int16     `bun:"price_page_type_cd,type:smallint"`
	MfgClassId                 string    `bun:"mfg_class_id,type:varchar(255),nullzero"`
	SupplierId                 float64   `bun:"supplier_id,type:decimal(19,0),nullzero"`
	ProductGroupId             string    `bun:"product_group_id,type:varchar(8),nullzero"`
	DiscountGroupId            string    `bun:"discount_group_id,type:varchar(8),nullzero"`
	MajorGroupId               string    `bun:"major_group_id,type:varchar(8),nullzero"`
	Description                string    `bun:"description,type:varchar(255),nullzero"`
	PricingMethodCd            int16     `bun:"pricing_method_cd,type:smallint"`
	SourcePriceCd              int16     `bun:"source_price_cd,type:smallint,nullzero"`
	Price                      float64   `bun:"price,type:decimal(19,9),nullzero"`
	EffectiveDate              time.Time `bun:"effective_date,type:datetime"`
	ExpirationDate             time.Time `bun:"expiration_date,type:datetime"`
	ContractNumber             string    `bun:"contract_number,type:varchar(40),nullzero"`
	CalculationMethodCd        int16     `bun:"calculation_method_cd,type:smallint"`
	CalculationValue1          float64   `bun:"calculation_value1,type:decimal(19,6)"`
	CalculationValue2          float64   `bun:"calculation_value2,type:decimal(19,6)"`
	CalculationValue3          float64   `bun:"calculation_value3,type:decimal(19,6)"`
	CalculationValue4          float64   `bun:"calculation_value4,type:decimal(19,6)"`
	CalculationValue5          float64   `bun:"calculation_value5,type:decimal(19,6)"`
	CalculationValue6          float64   `bun:"calculation_value6,type:decimal(19,6)"`
	CalculationValue7          float64   `bun:"calculation_value7,type:decimal(19,6)"`
	CalculationValue8          float64   `bun:"calculation_value8,type:decimal(19,6)"`
	CalculationValue9          float64   `bun:"calculation_value9,type:decimal(19,6)"`
	CalculationValue10         float64   `bun:"calculation_value10,type:decimal(19,6)"`
	CalculationValue11         float64   `bun:"calculation_value11,type:decimal(19,6)"`
	CalculationValue12         float64   `bun:"calculation_value12,type:decimal(19,6)"`
	CalculationValue13         float64   `bun:"calculation_value13,type:decimal(19,6)"`
	CalculationValue14         float64   `bun:"calculation_value14,type:decimal(19,6)"`
	CalculationValue15         float64   `bun:"calculation_value15,type:decimal(19,6)"`
	Break1                     float64   `bun:"break1,type:decimal(19,9)"`
	Break2                     float64   `bun:"break2,type:decimal(19,9)"`
	Break3                     float64   `bun:"break3,type:decimal(19,9)"`
	Break4                     float64   `bun:"break4,type:decimal(19,9)"`
	Break5                     float64   `bun:"break5,type:decimal(19,9)"`
	Break6                     float64   `bun:"break6,type:decimal(19,9)"`
	Break7                     float64   `bun:"break7,type:decimal(19,9)"`
	Break8                     float64   `bun:"break8,type:decimal(19,9)"`
	Break9                     float64   `bun:"break9,type:decimal(19,9)"`
	Break10                    float64   `bun:"break10,type:decimal(19,9)"`
	Break11                    float64   `bun:"break11,type:decimal(19,9)"`
	Break12                    float64   `bun:"break12,type:decimal(19,9)"`
	Break13                    float64   `bun:"break13,type:decimal(19,9)"`
	Break14                    float64   `bun:"break14,type:decimal(19,9)"`
	OtherCost1                 float64   `bun:"other_cost1,type:decimal(19,9),nullzero"`
	OtherCost2                 float64   `bun:"other_cost2,type:decimal(19,9),nullzero"`
	OtherCost3                 float64   `bun:"other_cost3,type:decimal(19,9),nullzero"`
	OtherCost4                 float64   `bun:"other_cost4,type:decimal(19,9),nullzero"`
	OtherCost5                 float64   `bun:"other_cost5,type:decimal(19,9),nullzero"`
	OtherCost6                 float64   `bun:"other_cost6,type:decimal(19,9),nullzero"`
	OtherCost7                 float64   `bun:"other_cost7,type:decimal(19,9),nullzero"`
	OtherCost8                 float64   `bun:"other_cost8,type:decimal(19,9),nullzero"`
	OtherCost9                 float64   `bun:"other_cost9,type:decimal(19,9),nullzero"`
	OtherCost10                float64   `bun:"other_cost10,type:decimal(19,9),nullzero"`
	OtherCost11                float64   `bun:"other_cost11,type:decimal(19,9),nullzero"`
	OtherCost12                float64   `bun:"other_cost12,type:decimal(19,9),nullzero"`
	OtherCost13                float64   `bun:"other_cost13,type:decimal(19,9),nullzero"`
	OtherCost14                float64   `bun:"other_cost14,type:decimal(19,9),nullzero"`
	OtherCost15                float64   `bun:"other_cost15,type:decimal(19,9),nullzero"`
	Uom1                       string    `bun:"uom1,type:varchar(8),nullzero"`
	Uom2                       string    `bun:"uom2,type:varchar(8),nullzero"`
	Uom3                       string    `bun:"uom3,type:varchar(8),nullzero"`
	Uom4                       string    `bun:"uom4,type:varchar(8),nullzero"`
	Uom5                       string    `bun:"uom5,type:varchar(8),nullzero"`
	Uom6                       string    `bun:"uom6,type:varchar(8),nullzero"`
	Uom7                       string    `bun:"uom7,type:varchar(8),nullzero"`
	Uom8                       string    `bun:"uom8,type:varchar(8),nullzero"`
	Uom9                       string    `bun:"uom9,type:varchar(8),nullzero"`
	Uom10                      string    `bun:"uom10,type:varchar(8),nullzero"`
	Uom11                      string    `bun:"uom11,type:varchar(8),nullzero"`
	Uom12                      string    `bun:"uom12,type:varchar(8),nullzero"`
	Uom13                      string    `bun:"uom13,type:varchar(8),nullzero"`
	Uom14                      string    `bun:"uom14,type:varchar(8),nullzero"`
	TotalingMethodCd           int16     `bun:"totaling_method_cd,type:smallint"`
	TotalingBasisCd            int16     `bun:"totaling_basis_cd,type:smallint"`
	RowStatusFlag              int16     `bun:"row_status_flag,type:smallint"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`
	DateCreated                time.Time `bun:"date_created,type:datetime"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30)"`
	OtherCostTypeCd            int16     `bun:"other_cost_type_cd,type:smallint,nullzero"`
	OtherCostValue             float64   `bun:"other_cost_value,type:decimal(19,9),nullzero"`
	OtherCostSourceCd          int16     `bun:"other_cost_source_cd,type:smallint,nullzero"`
	CostCalculationMethodCd    int16     `bun:"cost_calculation_method_cd,type:smallint,nullzero"`
	CostCalculationValue       float64   `bun:"cost_calculation_value,type:decimal(19,4),nullzero"`
	CommissionCostTypeCd       int16     `bun:"commission_cost_type_cd,type:smallint,nullzero"`
	CommissionCostValue        float64   `bun:"commission_cost_value,type:decimal(19,9),nullzero"`
	CommissionCostSourceCd     int16     `bun:"commission_cost_source_cd,type:smallint,nullzero"`
	CommissionCostCalcMethodCd int16     `bun:"commission_cost_calc_method_cd,type:smallint,nullzero"`
	CommissionCostCalcValue    float64   `bun:"commission_cost_calc_value,type:decimal(19,4),nullzero"`
	PricePageId                string    `bun:"price_page_id,type:varchar(20),nullzero"`
	CustomerId                 float64   `bun:"customer_id,type:decimal(19,0),nullzero"`
	CustomerPartNo             string    `bun:"customer_part_no,type:varchar(40),nullzero"`
	CompanyId                  string    `bun:"company_id,type:varchar(8),nullzero"`
	InvMastUid                 int32     `bun:"inv_mast_uid,type:int,nullzero"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	CommCostCalcValue1         float64   `bun:"comm_cost_calc_value1,type:decimal(19,9),default:(0)"`
	CommCostCalcValue2         float64   `bun:"comm_cost_calc_value2,type:decimal(19,9),default:(0)"`
	CommCostCalcValue3         float64   `bun:"comm_cost_calc_value3,type:decimal(19,9),default:(0)"`
	CommCostCalcValue4         float64   `bun:"comm_cost_calc_value4,type:decimal(19,9),default:(0)"`
	CommCostCalcValue5         float64   `bun:"comm_cost_calc_value5,type:decimal(19,9),default:(0)"`
	CommCostCalcValue6         float64   `bun:"comm_cost_calc_value6,type:decimal(19,9),default:(0)"`
	CommCostCalcValue7         float64   `bun:"comm_cost_calc_value7,type:decimal(19,9),default:(0)"`
	CommCostCalcValue8         float64   `bun:"comm_cost_calc_value8,type:decimal(19,9),default:(0)"`
	CommCostCalcValue9         float64   `bun:"comm_cost_calc_value9,type:decimal(19,9),default:(0)"`
	CommCostCalcValue10        float64   `bun:"comm_cost_calc_value10,type:decimal(19,9),default:(0)"`
	CommCostCalcValue11        float64   `bun:"comm_cost_calc_value11,type:decimal(19,9),default:(0)"`
	CommCostCalcValue12        float64   `bun:"comm_cost_calc_value12,type:decimal(19,9),default:(0)"`
	CommCostCalcValue13        float64   `bun:"comm_cost_calc_value13,type:decimal(19,9),default:(0)"`
	CommCostCalcValue14        float64   `bun:"comm_cost_calc_value14,type:decimal(19,9),default:(0)"`
	CommCostCalcValue15        float64   `bun:"comm_cost_calc_value15,type:decimal(19,9),default:(0)"`
	CalculatorType             string    `bun:"calculator_type,type:char,default:('B')"`
	CurrencyId                 float64   `bun:"currency_id,type:decimal(19,0),nullzero"`
	ValuesCurrencyId           float64   `bun:"values_currency_id,type:decimal(19,0),nullzero"`
	ApplyPpToMroCd             int16     `bun:"apply_pp_to_mro_cd,type:smallint,default:((1104))"`
	DateLastSentOn832          time.Time `bun:"date_last_sent_on_832,type:datetime,nullzero"`
	DatePageDeleted            time.Time `bun:"date_page_deleted,type:datetime,nullzero"`
	PriceFamilyUid             int32     `bun:"price_family_uid,type:int,nullzero"`
	PeerPricePageUid           int32     `bun:"peer_price_page_uid,type:int,nullzero"`
	ContractLineUid            int32     `bun:"contract_line_uid,type:int,nullzero"`
	OnContractFlag             string    `bun:"on_contract_flag,type:char,default:('N')"`
	StrategicPriceAppliesToCd  int32     `bun:"strategic_price_applies_to_cd,type:int,default:((2636))"`
	RoundToNextDollar          int16     `bun:"round_to_next_dollar,type:tinyint,nullzero"`
	ApplyFreightFactor         string    `bun:"apply_freight_factor,type:char,default:('N')"`
	FreightFactorSourceCd      int32     `bun:"freight_factor_source_cd,type:int,nullzero"`
	NoChargeFlag               string    `bun:"no_charge_flag,type:char,nullzero"`
	NonStockItemsOnlyFlag      string    `bun:"non_stock_items_only_flag,type:char,default:('N')"`
	ApplyPpToSopCd             int16     `bun:"apply_pp_to_sop_cd,type:smallint,default:((1103))"`
	PriceOverride              string    `bun:"price_override,type:char,default:('N')"`
	RebateMargin               float64   `bun:"rebate_margin,type:decimal(19,9),nullzero"`
	UpperMarginVariance        float64   `bun:"upper_margin_variance,type:decimal(19,9),nullzero"`
	LowerMarginVariance        float64   `bun:"lower_margin_variance,type:decimal(19,9),nullzero"`
	ExcludeOrderLevelDiscFlag  string    `bun:"exclude_order_level_disc_flag,type:char,nullzero"`
	RolledItemPricingTypeCd    int32     `bun:"rolled_item_pricing_type_cd,type:int,nullzero"`
}
