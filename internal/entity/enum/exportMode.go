package enum

type ExportMode string

const (
    FULL ExportMode = "FULL"
    IMG  ExportMode = "IMG"
    TXT  ExportMode = "TXT"
    _    ExportMode = ""
)
