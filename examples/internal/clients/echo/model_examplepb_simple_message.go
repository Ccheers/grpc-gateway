/*
 * Echo Service
 *
 * Echo Service API consists of a single service which returns a message.
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package echo

// SimpleMessage represents a simple message sent to the Echo service.
type ExamplepbSimpleMessage struct {
	// Id represents the message identifier.
	Id         string                  `json:"id"`
	Num        string                  `json:"num"`
	LineNum    string                  `json:"lineNum"`
	Lang       string                  `json:"lang"`
	Status     *ExamplepbEmbedded      `json:"status"`
	En         string                  `json:"en"`
	No         *ExamplepbEmbedded      `json:"no"`
	ResourceId string                  `json:"resourceId"`
	NId        *ExamplepbNestedMessage `json:"nId"`
}
