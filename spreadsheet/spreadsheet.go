package spreadsheet

type Spreadsheet struct {
	Name string
	Spreadsheet_id int
	Author string
	Sheets int
	CellWidth int
	CellHeight int
}

func (s Spreadsheet)PrintPreview() []byte {
  var data []byte
  return data
}
func (s Spreadsheet)PageSetup() []byte {
	var data []byte
	return data

}
func (s Spreadsheet) Print() []byte {
	var data []byte
	return data
}