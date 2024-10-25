package model

type ExportMode string

const (
	FULL ExportMode = "FULL"
	IMG  ExportMode = "IMG"
	TXT  ExportMode = "TXT"
)

type Exporter struct {
	ID       uint       `gorm:"primaryKey;autoIncrement;index"`
	Title    string     `gorm:"size:255"`
	Mode     ExportMode `gorm:"type:enum('FULL', 'IMG', 'TXT')"`
	NameRule string     `gorm:"size:255"`
}

func (Exporter) TableName() string {
	return "exporter"
}
