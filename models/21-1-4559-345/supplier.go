package model

import (
	"github.com/uptrace/bun"
	"time"
)

type Supplier struct {
	bun.BaseModel                   `bun:"table:supplier"`
	SupplierId                      float64   `bun:"supplier_id,type:decimal(19,0),pk"`
	BuyerId                         string    `bun:"buyer_id,type:varchar(16),nullzero"`
	DaysEarly                       float64   `bun:"days_early,type:decimal(19,0),default:(0)"`
	DaysLate                        float64   `bun:"days_late,type:decimal(19,0),default:(0)"`
	DeleteFlag                      string    `bun:"delete_flag,type:char"`
	DateCreated                     time.Time `bun:"date_created,type:datetime"`
	DateLastModified                time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy                string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	ControlValue                    string    `bun:"control_value,type:varchar(16),nullzero"`
	TargetValue                     float64   `bun:"target_value,type:decimal(19,9),nullzero"`
	ReviewCycle                     float64   `bun:"review_cycle,type:decimal(3,0),nullzero"`
	DateOfLastReview                time.Time `bun:"date_of_last_review,type:datetime,nullzero"`
	DefaultCarrier                  float64   `bun:"default_carrier,type:decimal(19,0),nullzero"`
	DefaultShippingInstructions     string    `bun:"default_shipping_instructions,type:varchar(255),nullzero"`
	AverageLeadTime                 float64   `bun:"average_lead_time,type:decimal(3,0),nullzero"`
	LeadTimeSource                  string    `bun:"lead_time_source,type:varchar(8),nullzero"`
	LeadTimeSafetyFactor            float64   `bun:"lead_time_safety_factor,type:decimal(19,4),nullzero"`
	SafetyStockDays                 float64   `bun:"safety_stock_days,type:decimal(19,4),nullzero"`
	Fct                             string    `bun:"fct,type:varchar(10),nullzero"`
	CurrencyId                      float64   `bun:"currency_id,type:decimal(19,0),nullzero"`
	SupplierName                    string    `bun:"supplier_name,type:varchar(255),nullzero"`
	TransitDays                     int32     `bun:"transit_days,type:int,default:(1)"`
	RmaRequired                     string    `bun:"rma_required,type:char,default:('N')"`
	CreatePoFromOe                  int32     `bun:"create_po_from_oe,type:int"`
	TpcxRationalized                string    `bun:"tpcx_rationalized,type:char,default:('N')"`
	FreightControlValue             string    `bun:"freight_control_value,type:varchar(16),nullzero"`
	FreightTargetValue              float64   `bun:"freight_target_value,type:decimal(19,9),default:(0)"`
	CreatedBy                       string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	CreateVendorRfqCd               int32     `bun:"create_vendor_rfq_cd,type:int,default:(986)"`
	VendorRfqDays                   int32     `bun:"vendor_rfq_days,type:int,default:(30)"`
	ShipDays                        int32     `bun:"ship_days,type:int,nullzero"`
	DeaCode                         string    `bun:"dea_code,type:char,nullzero"`
	DeaNumber                       string    `bun:"dea_number,type:char(255),nullzero"`
	DddReport                       string    `bun:"ddd_report,type:char,nullzero"`
	LegacyId                        string    `bun:"legacy_id,type:varchar(255),nullzero"`
	ParkerDivisionCd                string    `bun:"parker_division_cd,type:varchar(10),nullzero"`
	ParkerSupplierFlag              string    `bun:"parker_supplier_flag,type:char,nullzero"`
	CombineStockNsSpecialFlag       string    `bun:"combine_stock_ns_special_flag,type:char,default:('N')"`
	SafetyStockType                 int32     `bun:"safety_stock_type,type:int,nullzero"`
	ServiceLevelMeasure             int32     `bun:"service_level_measure,type:int,nullzero"`
	ServiceLevelPctGoal             float64   `bun:"service_level_pct_goal,type:decimal(19,2),nullzero"`
	UseContainersFlag               string    `bun:"use_containers_flag,type:char,default:('N')"`
	ContainerVolume                 float64   `bun:"container_volume,type:decimal(19,9),default:((0))"`
	ContainerWeight                 float64   `bun:"container_weight,type:decimal(19,9),default:((0))"`
	PoApprovalThresholdFlag         string    `bun:"po_approval_threshold_flag,type:char,nullzero"`
	SupplierSystemCd                int32     `bun:"supplier_system_cd,type:int,nullzero"`
	PrintPoFlag                     string    `bun:"print_po_flag,type:char,nullzero"`
	FaxPoFlag                       string    `bun:"fax_po_flag,type:char,nullzero"`
	EmailPoFlag                     string    `bun:"email_po_flag,type:char,nullzero"`
	EdiPoFlag                       string    `bun:"edi_po_flag,type:char,nullzero"`
	RestrictPedigreeItem            string    `bun:"restrict_pedigree_item,type:char,default:('Y')"`
	PallSupplierNumber              string    `bun:"pall_supplier_number,type:varchar(20),nullzero"`
	RestrictOtfItemFlag             string    `bun:"restrict_otf_item_flag,type:char,default:('N')"`
	SerialTrackingOption            string    `bun:"serial_tracking_option,type:char,default:('D')"`
	TermsDiscountExcludedFlag       string    `bun:"terms_discount_excluded_flag,type:char,default:('N')"`
	WeeklyTruckAllowedFlag          string    `bun:"weekly_truck_allowed_flag,type:char,nullzero"`
	RatelinxLocationId              int32     `bun:"ratelinx_location_id,type:int,nullzero"`
	DefaultPriceExpirationDate      time.Time `bun:"default_price_expiration_date,type:datetime,nullzero"`
	FreightFactor                   float64   `bun:"freight_factor,type:decimal(19,9),nullzero"`
	MinFreightFactor                float64   `bun:"min_freight_factor,type:decimal(19,9),nullzero"`
	DealerWarrantyClaimsCd          int32     `bun:"dealer_warranty_claims_cd,type:int,nullzero"`
	ParkerHannafinSupplierFlag      string    `bun:"parker_hannafin_supplier_flag,type:char,default:('N')"`
	ParkerAccountNo                 string    `bun:"parker_account_no,type:varchar(255),nullzero"`
	BrandNameCd                     int16     `bun:"brand_name_cd,type:smallint,nullzero"`
	RcdSupplierFlag                 string    `bun:"rcd_supplier_flag,type:char,nullzero"`
	SurchargeFlag                   string    `bun:"surcharge_flag,type:char,default:('N')"`
	PoDefaultMethod                 int32     `bun:"po_default_method,type:int,nullzero"`
	UseQtyRdyForContBuild           string    `bun:"use_qty_rdy_for_cont_build,type:char,default:('N')"`
	FidelitoneSupplierId            string    `bun:"fidelitone_supplier_id,type:varchar(255),nullzero"`
	DefaultServicebenchSupplierFlag string    `bun:"default_servicebench_supplier_flag,type:char,default:('N')"`
	DefaultPoPackingBasis           string    `bun:"default_po_packing_basis,type:varchar(255),nullzero"`
	DfltSupplierRedemptionMethod    int32     `bun:"dflt_supplier_redemption_method,type:int,nullzero"`
	SupplierRedemptionEmail         string    `bun:"supplier_redemption_email,type:varchar(255),nullzero"`
	SupplierRedemptionFaxNumber     string    `bun:"supplier_redemption_fax_number,type:varchar(20),nullzero"`
	ExportOrderInfoFlag             string    `bun:"export_order_info_flag,type:char,nullzero"`
	MinimumProcessPoCost            float64   `bun:"minimum_process_po_cost,type:decimal(19,9),nullzero"`
	BulkDiscountPercentage          float64   `bun:"bulk_discount_percentage,type:decimal(19,9),nullzero"`
	BulkContractNumberList          string    `bun:"bulk_contract_number_list,type:varchar(255),nullzero"`
	DoNotAutoCreateReturnsFlag      string    `bun:"do_not_auto_create_returns_flag,type:char,default:('N')"`
	CarrierShipMethodUid            int32     `bun:"carrier_ship_method_uid,type:int,nullzero"`
	DoNotCreateUnapprovedPoFlag     string    `bun:"do_not_create_unapproved_po_flag,type:char,nullzero"`
	DeviationLookbackPds            int32     `bun:"deviation_lookback_pds,type:int,nullzero"`
	DeviationMultOnePd              float64   `bun:"deviation_mult_one_pd,type:decimal(19,9),nullzero"`
	DeviationMultTwoPd              float64   `bun:"deviation_mult_two_pd,type:decimal(19,9),nullzero"`
	DeviationMultThreePd            float64   `bun:"deviation_mult_three_pd,type:decimal(19,9),nullzero"`
	MinSafetyStockDays              float64   `bun:"min_safety_stock_days,type:decimal(19,9),nullzero"`
	MaxSafetyStockDays              float64   `bun:"max_safety_stock_days,type:decimal(19,9),nullzero"`
	CritItemDeviationMult           float64   `bun:"crit_item_deviation_mult,type:decimal(19,9),nullzero"`
	CritMinSafetyStockDays          float64   `bun:"crit_min_safety_stock_days,type:decimal(19,9),nullzero"`
	CritMaxSafetyStockDays          float64   `bun:"crit_max_safety_stock_days,type:decimal(19,9),nullzero"`
	ExclSendLinkedInfoToRfnavFlag   string    `bun:"excl_send_linked_info_to_rfnav_flag,type:char,default:('N')"`
}
