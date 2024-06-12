package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Vendor struct {
	bun.BaseModel              `bun:"table:vendor"`
	VendorId                   float64   `bun:"vendor_id,type:decimal(19,0),pk"`
	CompanyId                  string    `bun:"company_id,type:varchar(8),pk"`
	Default1099Type            float64   `bun:"default_1099_type,type:decimal(3,0)"`
	ApAccountNo                string    `bun:"ap_account_no,type:varchar(32),nullzero"`
	DefaultPurchAcctNo         string    `bun:"default_purch_acct_no,type:varchar(32),nullzero"`
	DefaultTermsId             string    `bun:"default_terms_id,type:varchar(2),nullzero"`
	Class1id                   string    `bun:"class_1id,type:varchar(255),nullzero"`
	Class2id                   string    `bun:"class_2id,type:varchar(255),nullzero"`
	Class3id                   string    `bun:"class_3id,type:varchar(255),nullzero"`
	Class4id                   string    `bun:"class_4id,type:varchar(255),nullzero"`
	Class5id                   string    `bun:"class_5id,type:varchar(255),nullzero"`
	DeleteFlag                 string    `bun:"delete_flag,type:char"`
	DateCreated                time.Time `bun:"date_created,type:datetime"`
	DateLastModified           time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy           string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	VendorIdString             string    `bun:"vendor_id_string,type:varchar(22)"`
	DefaultPayFreightTo        string    `bun:"default_pay_freight_to,type:char,nullzero"`
	CurrencyId                 float64   `bun:"currency_id,type:decimal(3,0)"`
	EdiOrPaper                 string    `bun:"edi_or_paper,type:char"`
	SecurityInfo               string    `bun:"security_info,type:varchar(10),nullzero"`
	InterchgReceiverId         string    `bun:"interchg_receiver_id,type:varchar(15),nullzero"`
	IntlSan                    string    `bun:"intl_san,type:varchar(15),nullzero"`
	FederalIdNumber            string    `bun:"federal_id_number,type:varchar(15),nullzero"`
	DefaultInvoiceDesc         string    `bun:"default_invoice_desc,type:varchar(30),nullzero"`
	PayableGroupId             string    `bun:"payable_group_id,type:varchar(8),nullzero"`
	TqmTracking                string    `bun:"tqm_tracking,type:char,nullzero"`
	AlwaysTakeTerms            string    `bun:"always_take_terms,type:char,default:('N')"`
	JobIdRequired              string    `bun:"job_id_required,type:char,nullzero"`
	WorkersCompExpirationDate  time.Time `bun:"workers_comp_expiration_date,type:datetime,nullzero"`
	GeneralLiabExpirationDate  time.Time `bun:"general_liab_expiration_date,type:datetime,nullzero"`
	VendorName                 string    `bun:"vendor_name,type:varchar(255),nullzero"`
	TradingPartnerName         string    `bun:"trading_partner_name,type:varchar(255),nullzero"`
	TrackRebates               string    `bun:"track_rebates,type:char,default:('N')"`
	RebateAccountNo            string    `bun:"rebate_account_no,type:varchar(32),nullzero"`
	RebateAllowanceAccountNo   string    `bun:"rebate_allowance_account_no,type:varchar(32),nullzero"`
	PurchaseAcctGroupHdrUid    int32     `bun:"purchase_acct_group_hdr_uid,type:int,nullzero"`
	CreatedBy                  string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	AllowRebatePayImpVariance  string    `bun:"allow_rebate_pay_imp_variance,type:char,default:('Y')"`
	RebatePayImpVarianceType   int32     `bun:"rebate_pay_imp_variance_type,type:int,default:(1612)"`
	LegacyId                   string    `bun:"legacy_id,type:varchar(255),nullzero"`
	WildcardAcctInVoucherEntry string    `bun:"wildcard_acct_in_voucher_entry,type:char,default:('N')"`
	AttorneyFeeFlag            string    `bun:"attorney_fee_flag,type:char,default:('N')"`
	Irs1099State1              string    `bun:"IRS_1099_state_1,type:varchar(255),nullzero"`
	Irs1099State1IdNo          string    `bun:"IRS_1099_state_1_id_no,type:varchar(255),nullzero"`
	Irs1099State2              string    `bun:"IRS_1099_state_2,type:varchar(255),nullzero"`
	Irs1099State2IdNo          string    `bun:"IRS_1099_state_2_id_no,type:varchar(255),nullzero"`
	TaxpayerIdNoType           int32     `bun:"taxpayer_id_no_type,type:int,default:(300)"`
	WarrantyReceivableAcct     string    `bun:"warranty_receivable_acct,type:varchar(32),nullzero"`
	WarrantyRevenueAcct        string    `bun:"warranty_revenue_acct,type:varchar(32),nullzero"`
	WarrantyAllowanceAcct      string    `bun:"warranty_allowance_acct,type:varchar(32),nullzero"`
	CommissionReceivableAcct   string    `bun:"commission_receivable_acct,type:varchar(32),nullzero"`
	CommissionRevenueAcct      string    `bun:"commission_revenue_acct,type:varchar(32),nullzero"`
	CommissionAllowanceAcct    string    `bun:"commission_allowance_acct,type:varchar(32),nullzero"`
	TrackCoresFlag             string    `bun:"track_cores_flag,type:char,nullzero"`
	BankCoresFlag              string    `bun:"bank_cores_flag,type:char,nullzero"`
	IsrTaxWithheldFlag         string    `bun:"isr_tax_withheld_flag,type:char,nullzero"`
	IsrTaxPct                  float64   `bun:"isr_tax_pct,type:decimal(19,9),nullzero"`
	IsrTaxAcct                 string    `bun:"isr_tax_acct,type:varchar(32),nullzero"`
	VendorOnHoldFlag           string    `bun:"vendor_on_hold_flag,type:char,nullzero"`
	InvoiceLineVarianceAmt     float64   `bun:"invoice_line_variance_amt,type:decimal(19,9),nullzero"`
	InvoiceLineVarianceType    string    `bun:"invoice_line_variance_type,type:char,nullzero"`
	MaxCustomerSalesDiscount   float64   `bun:"max_customer_sales_discount,type:decimal(19,9),nullzero"`
	VendorTypeCd               int32     `bun:"vendor_type_cd,type:int,default:((2981))"`
	SeparateChecksFlag         string    `bun:"separate_checks_flag,type:char,default:('N')"`
	CreditLimit                float64   `bun:"credit_limit,type:decimal(19,4),nullzero"`
	AsbMainAccountNo           string    `bun:"asb_main_account_no,type:varchar(255),nullzero"`
}
