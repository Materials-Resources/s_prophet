package model

import (
	"github.com/uptrace/bun"
	"time"
)

type MassUpdateDefinitionDetail struct {
	bun.BaseModel                 `bun:"table:mass_update_definition_detail"`
	MassUpdateDefinitionDetailUid int32     `bun:"mass_update_definition_detail_uid,type:int,pk,identity"`
	MassUpdateDefinitionUid       int32     `bun:"mass_update_definition_uid,type:int"`
	ColumnName                    string    `bun:"column_name,type:varchar(255)"`
	IncludeFlag                   string    `bun:"include_flag,type:char"`
	RowStatusFlag                 int32     `bun:"row_status_flag,type:int"`
	DateCreated                   time.Time `bun:"date_created,type:datetime,default:(getdate())"`
	CreatedBy                     string    `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	DateLastModified              time.Time `bun:"date_last_modified,type:datetime,default:(getdate())"`
	LastMaintainedBy              string    `bun:"last_maintained_by,type:varchar(255),default:(suser_sname())"`
}
