package document

type Document struct {
	Name string
	Document_id int
	Author string
	Pages int
	Width int
	Height int
}

func (d Document)PrintPreview() []byte {
  var data []byte
  return data
}
func (d Document)PageSetup() []byte {
	var data []byte
	return data

}
func (d Document) Print() []byte {
	var data []byte
	return data
}