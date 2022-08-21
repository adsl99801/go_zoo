package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


Table: weather
[ 0] city                                           VARCHAR(80)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 80      default: []
[ 1] temp_lo                                        INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 2] temp_hi                                        INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 3] prcp                                           FLOAT4               null: true   primary: false  isArray: false  auto: false  col: FLOAT4          len: -1      default: []
[ 4] date                                           DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []


JSON Sample
-------------------------------------
{    "city": "kApwDiAEDwOcuMDhmNILekkIs",    "temp_lo": 80,    "temp_hi": 2,    "prcp": 0.8808493878235294,    "date": "2043-12-04T15:36:04.730592613+08:00"}


Comments
-------------------------------------
[ 0] Warning table: weather does not have a primary key defined, setting col position 1 city as primary key
Warning table: weather primary key column city is nullable column, setting it as NOT NULL




*/

// Weather struct is a row record of the weather table in the pf database
type Weather struct {
	//[ 0] city                                           VARCHAR(80)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 80      default: []
	City string `gorm:"primary_key;column:city;type:VARCHAR;size:80;" json:"city"`
	//[ 1] temp_lo                                        INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	TempLo sql.NullInt64 `gorm:"column:temp_lo;type:INT4;" json:"temp_lo"`
	//[ 2] temp_hi                                        INT4                 null: true   primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	TempHi sql.NullInt64 `gorm:"column:temp_hi;type:INT4;" json:"temp_hi"`
	//[ 3] prcp                                           FLOAT4               null: true   primary: false  isArray: false  auto: false  col: FLOAT4          len: -1      default: []
	Prcp sql.NullFloat64 `gorm:"column:prcp;type:FLOAT4;" json:"prcp"`
	//[ 4] date                                           DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	Date time.Time `gorm:"column:date;type:DATE;" json:"date"`
}

var weatherTableInfo = &TableInfo{
	Name: "weather",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "city",
			Comment: ``,
			Notes: `Warning table: weather does not have a primary key defined, setting col position 1 city as primary key
Warning table: weather primary key column city is nullable column, setting it as NOT NULL
`,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(80)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       80,
			GoFieldName:        "City",
			GoFieldType:        "string",
			JSONFieldName:      "city",
			ProtobufFieldName:  "city",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "temp_lo",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "TempLo",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "temp_lo",
			ProtobufFieldName:  "temp_lo",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "temp_hi",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "TempHi",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "temp_hi",
			ProtobufFieldName:  "temp_hi",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "prcp",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "FLOAT4",
			DatabaseTypePretty: "FLOAT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "FLOAT4",
			ColumnLength:       -1,
			GoFieldName:        "Prcp",
			GoFieldType:        "sql.NullFloat64",
			JSONFieldName:      "prcp",
			ProtobufFieldName:  "prcp",
			ProtobufType:       "float",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "date",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "DATE",
			DatabaseTypePretty: "DATE",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "DATE",
			ColumnLength:       -1,
			GoFieldName:        "Date",
			GoFieldType:        "time.Time",
			JSONFieldName:      "date",
			ProtobufFieldName:  "date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (w *Weather) TableName() string {
	return "weather"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (w *Weather) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (w *Weather) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *Weather) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (w *Weather) TableInfo() *TableInfo {
	return weatherTableInfo
}
