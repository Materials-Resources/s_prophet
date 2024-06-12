package model

import (
	"github.com/uptrace/bun"
	"time"
)

type JobPriceLine struct {
	bun.BaseModel               `bun:"table:job_price_line"`
	JobPriceLineUid             int32     `bun:"job_price_line_uid,type:int,pk"`
	JobPriceHdrUid              int32     `bun:"job_price_hdr_uid,type:int"`
	InvMastUid                  int32     `bun:"inv_mast_uid,type:int,nullzero"`
	Uom                         string    `bun:"uom,type:varchar(8),nullzero"`
	UnitSize                    float64   `bun:"unit_size,type:decimal(19,4),nullzero"`
	PricingMethod               int16     `bun:"pricing_method,type:smallint,nullzero"`
	SourcePrice                 int16     `bun:"source_price,type:smallint,nullzero"`
	Price                       float64   `bun:"price,type:decimal(19,9),nullzero"`
	QtyOrdered                  float64   `bun:"qty_ordered,type:decimal(19,9),nullzero"`
	QtyMaximum                  float64   `bun:"qty_maximum,type:decimal(19,9),nullzero"`
	OtherCostTypeCd             int16     `bun:"other_cost_type_cd,type:smallint,nullzero"`
	OtherCostValue              float64   `bun:"other_cost_value,type:decimal(19,9),nullzero"`
	OtherCostSourceCd           int16     `bun:"other_cost_source_cd,type:smallint,nullzero"`
	OtherCostCalcMethodCd       int16     `bun:"other_cost_calc_method_cd,type:smallint,nullzero"`
	OtherCostCalcValue          float64   `bun:"other_cost_calc_value,type:decimal(19,9),nullzero"`
	CommissionCostTypeCd        int16     `bun:"commission_cost_type_cd,type:smallint,nullzero"`
	CommissionCostValue         float64   `bun:"commission_cost_value,type:decimal(19,9),nullzero"`
	CommissionCostSourceCd      int16     `bun:"commission_cost_source_cd,type:smallint,nullzero"`
	CommissionCostCalcMethodCd  int16     `bun:"commission_cost_calc_method_cd,type:smallint,nullzero"`
	CommissionCostCalcValue     float64   `bun:"commission_cost_calc_value,type:decimal(19,9),nullzero"`
	RowStatusFlag               int32     `bun:"row_status_flag,type:int"`
	DateLastModified            time.Time `bun:"date_last_modified,type:datetime"`
	DateCreated                 time.Time `bun:"date_created,type:datetime"`
	LastMaintainedBy            string    `bun:"last_maintained_by,type:varchar(30)"`
	Multiplier                  float64   `bun:"multiplier,type:decimal(19,9),nullzero"`
	LineNo                      int32     `bun:"line_no,type:int"`
	SourceLocationId            float64   `bun:"source_location_id,type:decimal(19,0),nullzero"`
	CustomerPartNo              string    `bun:"customer_part_no,type:varchar(40),nullzero"`
	ExpirationDate              time.Time `bun:"expiration_date,type:datetime,nullzero"`
	VendorAuthNo                string    `bun:"vendor_auth_no,type:varchar(255),nullzero"`
	PoCost                      float64   `bun:"po_cost,type:decimal(19,9),nullzero"`
	DiscountGroupId             string    `bun:"discount_group_id,type:varchar(8),nullzero"`
	CustPoNo                    string    `bun:"cust_po_no,type:varchar(255),nullzero"`
	ItemCategoryUid             int32     `bun:"item_category_uid,type:int,nullzero"`
	SubCategoryUid              int32     `bun:"sub_category_uid,type:int,nullzero"`
	ProductGroupId              string    `bun:"product_group_id,type:varchar(8),nullzero"`
	CurrencyId                  float64   `bun:"currency_id,type:decimal(19,0),nullzero"`
	TerminalId                  float64   `bun:"terminal_id,type:decimal(19,0),nullzero"`
	ContractLineCostPageUid     int32     `bun:"contract_line_cost_page_uid,type:int,nullzero"`
	ContractLinePricePageUid    int32     `bun:"contract_line_price_page_uid,type:int,nullzero"`
	BudgetCd                    string    `bun:"budget_cd,type:varchar(255),nullzero"`
	ItemRevisionUid             int32     `bun:"item_revision_uid,type:int,nullzero"`
	PickFee                     float64   `bun:"pick_fee,type:decimal(19,9),nullzero"`
	LineStartDate               time.Time `bun:"line_start_date,type:datetime,nullzero"`
	InitialCommitmentAmount     float64   `bun:"initial_commitment_amount,type:decimal(19,9),default:((0))"`
	CommitmentAmount            float64   `bun:"commitment_amount,type:decimal(19,9),default:((0))"`
	TotalCommitmentAmount       float64   `bun:"total_commitment_amount,type:decimal(19,9),default:((0))"`
	CadPurchaseCost             float64   `bun:"cad_purchase_cost,type:decimal(19,9),nullzero"`
	StartingVirtualInventoryQty float64   `bun:"starting_virtual_inventory_qty,type:decimal(19,9),nullzero"`
	SnapshotCost                float64   `bun:"snapshot_cost,type:decimal(19,9),nullzero"`
	MinimumMccCode              string    `bun:"minimum_mcc_code,type:varchar(255),nullzero"`
	CommissionClassId           string    `bun:"commission_class_id,type:varchar(8),nullzero"`
	CommissionOverridePercent   float64   `bun:"commission_override_percent,type:decimal(19,9),nullzero"`
	AllDiscountGroupsFlag       string    `bun:"all_discount_groups_flag,type:char,default:('N')"`
	PocostingMethod             int16     `bun:"pocosting_method,type:smallint,nullzero"`
	PocostingSourcePoCostCd     int16     `bun:"pocosting_source_po_cost_cd,type:smallint,nullzero"`
	PocostingMultiplier         float64   `bun:"pocosting_multiplier,type:decimal(19,9),nullzero"`
	PocostingPoCost             float64   `bun:"pocosting_po_cost,type:decimal(19,9),nullzero"`
	PocostingPoContractNumber   string    `bun:"pocosting_po_contract_number,type:varchar(40),nullzero"`
	DefaultDisposition          string    `bun:"default_disposition,type:char,nullzero"`
}
