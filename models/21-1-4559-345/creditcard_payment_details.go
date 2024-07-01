package gen

import (
	"github.com/uptrace/bun"
	"time"
)

type CreditcardPaymentDetails struct {
	bun.BaseModel                   `bun:"table:creditcard_payment_details"`
	CreditcardPaymentDetailsUid     int32     `bun:"creditcard_payment_details_uid,type:int,pk"`                   // UID of this table
	PaymentNumber                   float64   `bun:"payment_number,type:decimal(19,0)"`                            // Payment number from ar_payment_details table
	FirstName                       string    `bun:"first_name,type:varchar(255),nullzero"`                        // Credit card holders first name
	LastName                        string    `bun:"last_name,type:varchar(255),nullzero"`                         // Credit card holders last name
	StreetAddress1                  string    `bun:"street_address1,type:varchar(255),nullzero"`                   // Credit card holders street address 1
	StreetAddress2                  string    `bun:"street_address2,type:varchar(255),nullzero"`                   // Credit card holders street address 2
	City                            string    `bun:"city,type:varchar(25),nullzero"`                               // Credit card holders city
	State                           string    `bun:"state,type:char(2),nullzero"`                                  // Credit card holders state
	ZipCode                         string    `bun:"zip_code,type:varchar(10),nullzero"`                           // Credit card holders zip code
	OriginalAuthAmount              float64   `bun:"original_auth_amount,type:decimal(19,4),nullzero"`             // Original authorization amount
	AuthAmount                      float64   `bun:"auth_amount,type:decimal(19,4),nullzero"`                      // Actual authorization amount
	TransSettled                    string    `bun:"trans_settled,type:char(1)"`                                   // A flag to indicate whether the credit card transaction has been settled
	Taker                           string    `bun:"taker,type:varchar(30)"`                                       // ID of the user who took the order.
	DateCreated                     time.Time `bun:"date_created,type:datetime,default:(getdate())"`               // Indicates the date/time this record was created.
	LastMaintainedBy                string    `bun:"last_maintained_by,type:varchar(30),default:(suser_sname())"`  // ID of the user who last maintained this record
	DateLastModified                time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`         // Indicates the date/time this record was last modified.
	BatchNumber                     string    `bun:"batch_number,type:varchar(20),nullzero"`                       // Batch number from Protobase
	RetrievalRefNumber              string    `bun:"retrieval_ref_number,type:varchar(20),nullzero"`               // A reference number from Protobase
	CustomerVerificationValue       string    `bun:"customer_verification_value,type:varchar(10),nullzero"`        // Customer verification value from credit card
	UserSpecifiedAmount             float64   `bun:"user_specified_amount,type:decimal(19,4),default:(0)"`         // Amount user put in payment amount field on remittance tab in order entry
	SwitchIssueNumber               string    `bun:"switch_issue_number,type:varchar(2),nullzero"`                 // Switch Issue Number
	RealexReferenceNumber           string    `bun:"realex_reference_number,type:varchar(255),nullzero"`           // Unique ID sent to Realex authorization
	RealexOrderId                   string    `bun:"realex_order_id,type:varchar(255),nullzero"`                   // Unique ID returned from Realex initial auth
	PaymentAccountId                string    `bun:"payment_account_id,type:varchar(255),nullzero"`                // Contains the unique identifier that identifies the account associated with the payment; this value is typically generated by an external source.
	AddressEditedFlag               string    `bun:"address_edited_flag,type:char(1),default:('N')"`               // A column to indicate when the cardholder address has been edited.  If it has been edited, the address will need to get sent along for credit card transactions to the payment processing host
	CommercialCardResponseCode      string    `bun:"commercial_card_response_code,type:varchar(30),nullzero"`      // Commercial Card Response Code returned by payment processing host
	CardBrand                       string    `bun:"card_brand,type:varchar(30),nullzero"`                         // Brand of card used for the transaction (Visa, MasterCard, etc.)
	MarketCodeValue                 string    `bun:"market_code_value,type:varchar(255),nullzero"`                 // Market Code value (integration provider-specific)
	SwipedTransactionFlag           string    `bun:"swiped_transaction_flag,type:char(1),default:('N')"`           // Indicates whether the card data was captured using a magnetic stripe reader device
	EcommerceTransactionFlag        string    `bun:"ecommerce_transaction_flag,type:char(1),default:('N')"`        // Indicates whether the transaction originated from an ECommerce environment
	TransactionCardholderLocationCd int32     `bun:"transaction_cardholder_location_cd,type:int,default:((2815))"` // Location of the cardholder at the time of the transaction (see code group 2097 -
	StreetAddress3                  string    `bun:"street_address3,type:varchar(50),nullzero"`                    // Address line 3
}