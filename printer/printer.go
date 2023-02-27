package printer

type Printer interface {
	PrintPreview() []byte
	PageSetup() []byte
	Print() []byte
}

func PagePreview(art Printer) bool {
	return true
}

func PageSetup(art Printer) bool {
	return true
}

func Print(art Printer) bool {
	
}
