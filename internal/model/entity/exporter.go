package entity

type ExportMode string
type Exporter struct {
    ID       uint       `gorm:"primaryKey;autoIncrement;index"`
    Title    string     `gorm:"size:255"`
    Mode     ExportMode `gorm:"type:enum('FULL', 'IMG', 'TXT')"`
    NameRule string     `gorm:"size:255; default:''"`
}

func (Exporter) TableName() string {
    return "exporter"
}
