package model

import (
	"github.com/uptrace/bun"
	"time"
)

type OePickTicket struct {
	bun.BaseModel            `bun:"table:oe_pick_ticket"`
	CompanyId                string    `bun:"company_id,type:varchar(8)"`
	PickTicketNo             float64   `bun:"pick_ticket_no,type:decimal(19,0),pk"`
	OrderNo                  string    `bun:"order_no,type:varchar(8)"`
	PrintDate                time.Time `bun:"print_date,type:datetime,nullzero"`
	CarrierId                float64   `bun:"carrier_id,type:decimal(19,0),nullzero"`
	TrackingNo               string    `bun:"tracking_no,type:varchar(40),nullzero"`
	Instructions             string    `bun:"instructions,type:varchar(255),nullzero"`
	FreightOut               float64   `bun:"freight_out,type:decimal(19,4),nullzero"`
	FreightIn                float64   `bun:"freight_in,type:decimal(19,4),nullzero"`
	ShipDate                 time.Time `bun:"ship_date,type:datetime,nullzero"`
	InvoiceNo                float64   `bun:"invoice_no,type:decimal(19,0),nullzero"`
	PickAndHold              string    `bun:"pick_and_hold,type:char"`
	DateCreated              time.Time `bun:"date_created,type:datetime"`
	DateLastModified         time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy         string    `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	PrintedFlag              string    `bun:"printed_flag,type:char,default:('Y')"`
	LocationId               float64   `bun:"location_id,type:decimal(19,0)"`
	DeleteFlag               string    `bun:"delete_flag,type:char"`
	DirectShipment           string    `bun:"direct_shipment,type:char,nullzero"`
	FrontCounterTax          string    `bun:"front_counter_tax,type:char,nullzero"`
	TotalTax                 float64   `bun:"total_tax,type:decimal(19,2),nullzero"`
	FrontCounter             string    `bun:"front_counter,type:char"`
	InvoiceIdWhenShipped     string    `bun:"invoice_id_when_shipped,type:varchar(10),nullzero"`
	Auxiliary                string    `bun:"auxiliary,type:char,default:('N')"`
	InvoiceBatchUid          int32     `bun:"invoice_batch_uid,type:int,default:(1)"`
	FreightCodeUid           int32     `bun:"freight_code_uid,type:int,nullzero"`
	DeliveryListStatus       int32     `bun:"delivery_list_status,type:int,default:(1024)"`
	RetrievedByWms           string    `bun:"retrieved_by_wms,type:char,default:('N')"`
	Picker                   string    `bun:"picker,type:varchar(255),nullzero"`
	LineItemSortSequence     int32     `bun:"line_item_sort_sequence,type:int,default:(1)"`
	ConfirmableRowStatusFlag int32     `bun:"confirmable_row_status_flag,type:int,default:('1980')"`
	WeightOverrideFlag       string    `bun:"weight_override_flag,type:char,nullzero"`
	OverrideValue            float64   `bun:"override_value,type:decimal(19,9),nullzero"`
	CreatedBy                string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	OePickTicketTypeCd       int32     `bun:"oe_pick_ticket_type_cd,type:int,nullzero"`
	OutgoingFreightCost      float64   `bun:"outgoing_freight_cost,type:decimal(19,9),nullzero"`
	QualityControl           string    `bun:"quality_control,type:varchar(255),nullzero"`
	PackingListPrintedFlag   string    `bun:"packing_list_printed_flag,type:char,nullzero"`
	DeaPickTicketFlag        string    `bun:"dea_pick_ticket_flag,type:char,nullzero"`
	Packer                   string    `bun:"packer,type:varchar(255),nullzero"`
	Checker                  string    `bun:"checker,type:varchar(255),nullzero"`
	ReviewShipmentFlag       string    `bun:"review_shipment_flag,type:char,default:('N')"`
	OriginalFreightOut       float64   `bun:"original_freight_out,type:decimal(19,9),nullzero"`
	ExportedToRouterFlag     string    `bun:"exported_to_router_flag,type:char,nullzero"`
	PriceConfirmedFlag       string    `bun:"price_confirmed_flag,type:char,nullzero"`
	ShipConfirmedFlag        string    `bun:"ship_confirmed_flag,type:char,nullzero"`
	OrigSortOrder            int32     `bun:"orig_sort_order,type:int,nullzero"`
	VicsBolNumber            string    `bun:"vics_bol_number,type:varchar(17),nullzero"`
	OutFreightCost           float64   `bun:"out_freight_cost,type:decimal(19,9),nullzero"`
	FreightOutEditedFlag     string    `bun:"freight_out_edited_flag,type:char,nullzero"`
	PrintPricesOnPackinglist string    `bun:"print_prices_on_packinglist,type:char,nullzero"`
	PregeneratedInvoiceNo    float64   `bun:"pregenerated_invoice_no,type:decimal(19,0),nullzero"`
	PrepaidFreightOut        float64   `bun:"prepaid_freight_out,type:decimal(19,9),nullzero"`
	PrintCanadianB3FormsFlag string    `bun:"print_canadian_b3_forms_flag,type:char,default:('N')"`
	ProNumber                string    `bun:"pro_number,type:varchar(255),nullzero"`
	PackingListPrintBy       string    `bun:"packing_list_print_by,type:varchar(30),nullzero"`
	PackingListPrintDate     time.Time `bun:"packing_list_print_date,type:datetime,nullzero"`
	SlabFlag                 string    `bun:"slab_flag,type:char,default:('N')"`
	ShippingAccount          string    `bun:"shipping_account,type:varchar(255),nullzero"`
	SidNo                    string    `bun:"sid_no,type:varchar(255),nullzero"`
	SentToAttDate            time.Time `bun:"sent_to_att_date,type:datetime,nullzero"`
	UserConfirmedPickFlag    string    `bun:"user_confirmed_pick_flag,type:char,nullzero"`
	Scan                     string    `bun:"scan,type:varchar(255),nullzero"`
	ActualFedexFreightOut    float64   `bun:"actual_fedex_freight_out,type:decimal(19,4),nullzero"`
	DiffFedexChargeFlag      string    `bun:"diff_fedex_charge_flag,type:char,default:('N')"`
	PickConfirmedFlag        string    `bun:"pick_confirmed_flag,type:char,nullzero"`
	SentToTrackaboutFlag     string    `bun:"sent_to_trackabout_flag,type:char,nullzero"`
	ShippingRouteUid         int32     `bun:"shipping_route_uid,type:int,nullzero"`
	BolNumber                string    `bun:"bol_number,type:varchar(255),nullzero"`
	DriverId                 string    `bun:"driver_id,type:varchar(16),nullzero"`
	SalesMarketGroupUid      int32     `bun:"sales_market_group_uid,type:int,nullzero"`
	ArnNumber                string    `bun:"arn_number,type:varchar(255),nullzero"`
	ScanPackUid              int32     `bun:"scan_pack_uid,type:int,nullzero"`
	ReferenceNo              string    `bun:"reference_no,type:varchar(255),nullzero"`
	PickCompleteFlag         string    `bun:"pick_complete_flag,type:char,nullzero"`
	RfnavStopNo              int32     `bun:"rfnav_stop_no,type:int,nullzero"`
	RfnavPickStatusCd        int32     `bun:"rfnav_pick_status_cd,type:int,nullzero"`
}
