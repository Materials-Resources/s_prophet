package model

import (
	"github.com/uptrace/bun"
	"time"
)

type CustDefaults struct {
	bun.BaseModel              `bun:"table:cust_defaults"`
	ArAccountNo                string    `bun:"ar_account_no,type:varchar(32)"`
	RevenueAccountNo           string    `bun:"revenue_account_no,type:varchar(32)"`
	CosAccountNo               string    `bun:"cos_account_no,type:varchar(32)"`
	AllowedAccountNo           string    `bun:"allowed_account_no,type:varchar(32)"`
	TermsAccountNo             string    `bun:"terms_account_no,type:varchar(32)"`
	FreightAccountNo           string    `bun:"freight_account_no,type:varchar(32)"`
	BrokerageAccountNo         string    `bun:"brokerage_account_no,type:varchar(32)"`
	PriceFileId                float64   `bun:"price_file_id,type:decimal(19,0),nullzero"`
	CompanyId                  string    `bun:"company_id,type:varchar(8),pk"`
	LocationId                 float64   `bun:"location_id,type:decimal(19,0),nullzero"`
	DeleteFlag                 string    `bun:"delete_flag,type:char"`
	DateCreated                time.Time `bun:"date_created,type:datetime"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	Fob                        string    `bun:"fob,type:varchar(20),nullzero"`
	DefaultDisposition         string    `bun:"default_disposition,type:char,nullzero"`
	DeferredRevenueAccountNo   string    `bun:"deferred_revenue_account_no,type:varchar(32),nullzero"`
	PickTicketType             string    `bun:"pick_ticket_type,type:varchar(2),nullzero"`
	FinanceChgRevenueAccountNo string    `bun:"finance_chg_revenue_account_no,type:varchar(32),nullzero"`
	InvoiceType                string    `bun:"invoice_type,type:varchar(2),nullzero"`
	AcceptableWaitTime         float64   `bun:"acceptable_wait_time,type:decimal(5,0)"`
	CreditLimit                float64   `bun:"credit_limit,type:decimal(22,4),default:(0)"`
	PendingPaymentAccountNo    string    `bun:"pending_payment_account_no,type:varchar(32),nullzero"`
	TermsId                    string    `bun:"terms_id,type:varchar(2),nullzero"`
	PricingMethodCd            int16     `bun:"pricing_method_cd,type:tinyint,nullzero"`
	SourcePriceCd              int32     `bun:"source_price_cd,type:int,nullzero"`
	Multiplier                 float64   `bun:"multiplier,type:decimal(19,9),nullzero"`
	PriceLibraryId             string    `bun:"price_library_id,type:varchar(20),nullzero"`
	SalesrepId                 string    `bun:"salesrep_id,type:varchar(16),nullzero"`
	DefaultBranch              string    `bun:"default_branch,type:varchar(8),nullzero"`
	PreferredLocationId        float64   `bun:"preferred_location_id,type:decimal(19,0),nullzero"`
	CustomerTypeUid            int32     `bun:"customer_type_uid,type:int,nullzero"`
	AllowAdvanceBilling        string    `bun:"allow_advance_billing,type:char,default:('N')"`
	AdvanceBillAccountNo       string    `bun:"advance_bill_account_no,type:varchar(32),nullzero"`
	DataIdentifierGroupUid     int32     `bun:"data_identifier_group_uid,type:int,nullzero"`
	CostCenterTrackingOption   int32     `bun:"cost_center_tracking_option,type:int,nullzero"`
	SignatureRequired          string    `bun:"signature_required,type:char,default:('N')"`
	PackingBasis               string    `bun:"packing_basis,type:varchar(16),nullzero"`
	InvoiceSurchargePct        float64   `bun:"invoice_surcharge_pct,type:decimal(19,9),default:(0)"`
	LaborRate                  float64   `bun:"labor_rate,type:decimal(19,9),nullzero"`
	TaxableFlag                string    `bun:"taxable_flag,type:char,default:('N')"`
	AllowLineItemFreightFlag   string    `bun:"allow_line_item_freight_flag,type:char,default:('N')"`
	CarrierId                  float64   `bun:"carrier_id,type:decimal(19,0),nullzero"`
	ShippingRouteUid           int32     `bun:"shipping_route_uid,type:int,nullzero"`
	ServiceTermsId             string    `bun:"service_terms_id,type:varchar(2),nullzero"`
	DownpaymentPercentage      float64   `bun:"downpayment_percentage,type:decimal(19,2),default:((0.00))"`
	ReqPymtUponReleaseOfItems  string    `bun:"req_pymt_upon_release_of_items,type:char,default:('N')"`
	IncludeDpSummaryOnInvoices string    `bun:"include_dp_summary_on_invoices,type:char,default:('N')"`
	JobNumberRequiredFlag      string    `bun:"job_number_required_flag,type:char,default:('N')"`
	UseLastMarginPricingFlag   string    `bun:"use_last_margin_pricing_flag,type:char,default:('N')"`
	InvoiceCompCostCdTier1     int32     `bun:"invoice_comp_cost_cd_tier1,type:int,default:((2110))"`
	DaysUntilQuoteExpires      int32     `bun:"days_until_quote_expires,type:int,nullzero"`
	InvoiceCompCostCdTier2     int32     `bun:"invoice_comp_cost_cd_tier2,type:int,default:((2111))"`
	InvoiceCompCostCdTier3     int32     `bun:"invoice_comp_cost_cd_tier3,type:int,default:((1195))"`
	CustomerTypeCd             int32     `bun:"customer_type_cd,type:int,nullzero"`
	ConsInvSummaryFilename     string    `bun:"cons_inv_summary_filename,type:varchar(255),nullzero"`
	ConsInvDetailFilename      string    `bun:"cons_inv_detail_filename,type:varchar(255),nullzero"`
	InvoiceFilename            string    `bun:"invoice_filename,type:varchar(255),nullzero"`
	StatementFilename          string    `bun:"statement_filename,type:varchar(255),nullzero"`
	UseVendorItemTermsFlag     string    `bun:"use_vendor_item_terms_flag,type:char,nullzero"`
	RmaRevenueAccountNo        string    `bun:"rma_revenue_account_no,type:varchar(32),nullzero"`
	DealerWrrtyClaimsAccountNo string    `bun:"dealer_wrrty_claims_account_no,type:varchar(32),nullzero"`
	ArBatchType                string    `bun:"ar_batch_type,type:char,default:('C')"`
	ConsBackordersFlag         string    `bun:"cons_backorders_flag,type:char,nullzero"`
	TaxGroupId                 string    `bun:"tax_group_id,type:varchar(10),nullzero"`
	CreditStatus               string    `bun:"credit_status,type:varchar(8),nullzero"`
	CurrencyId                 float64   `bun:"currency_id,type:decimal(19,0)"`
	ApplyConvenienceFeeFlag    string    `bun:"apply_convenience_fee_flag,type:char,nullzero"`
	ApplyFuelSurchargesCd      int32     `bun:"apply_fuel_surcharges_cd,type:int,nullzero"`
	AppliedFuelchargesToDsFlag string    `bun:"applied_fuelcharges_to_ds_flag,type:char,nullzero"`
	SuppressZeroDollarFlag     string    `bun:"suppress_zero_dollar_flag,type:char,nullzero"`
	Class5id                   string    `bun:"class_5id,type:varchar(255),nullzero"`
}
