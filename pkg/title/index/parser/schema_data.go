package parser

// GetSchemaOrgData extracts the first JSON string, which is usually on top of the web page's source code
// this is manually generated based on "schema.org" schemas,
// and it's unlikely to change in the future.
// func GetSchemaOrgData(doc *goquery.Document) (*schemaOrgData, error) {
// 	scriptJSON := doc.Find("script[type=\"application/ld+json\"]").First()
// 	var data schemaOrgData
// 	if err := json.Unmarshal([]byte(scriptJSON.Text()), &data); err != nil {
// 		return nil, err
// 	}
// 	return &data, nil
// }
