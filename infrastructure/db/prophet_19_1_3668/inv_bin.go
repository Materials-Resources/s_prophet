package prophet_19_1_3668

import (
	"time"

	"github.com/uptrace/bun"
)

type InvBin struct {
	bun.BaseModel    `bun:"table:inv_bin"`
	CompanyId        string    `bun:"company_id,type:varchar(8),pk"`
	LocationId       float64   `bun:"location_id,type:decimal(19,0),pk"`
	Bin              string    `bun:"bin,type:varchar(10),pk"`
	Quantity         float64   `bun:"quantity,type:decimal(19,9),default:(0)"`
	DateCreated      time.Time `bun:"date_created,type:datetime"`
	DateLastModified time.Time `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string    `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	InvMastUid       int32     `bun:"inv_mast_uid,type:int,pk"`
	QtyAllocated     float64   `bun:"qty_allocated,type:decimal(19,9),default:(0)"`
	InvBinUid        int32     `bun:"inv_bin_uid,type:int"`
	RowStatusFlag    int32     `bun:"row_status_flag,type:int,default:(1037)"`
}
