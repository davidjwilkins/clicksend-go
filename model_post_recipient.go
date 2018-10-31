/*
 * ClickSend v3 REST API
 *
 *  This is the official [ClickSend](https://clicksend.com) SDK.  *You'll need to create a free account to use the API. You can register [here](https://www.clicksend.com/signup).*  You can use our API documentation along with the SDK. Our API docs can be found [here](https://developers.clicksend.com). 
 *
 * API version: 3.1.0
 * Contact: support@clicksend.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// PostRecipient model
type PostRecipient struct {
	// Name of address
	AddressName string `json:"address_name"`
	// First line of address
	AddressLine1 string `json:"address_line_1"`
	// Second line of address
	AddressLine2 string `json:"address_line_2"`
	// City
	AddressCity string `json:"address_city"`
	// State
	AddressState string `json:"address_state"`
	// Postal code
	AddressPostalCode string `json:"address_postal_code"`
	// Country
	AddressCountry string `json:"address_country"`
	// ID of return address to use
	ReturnAddressId int32 `json:"return_address_id"`
	// When to send letter (0/null=now)
	Schedule int32 `json:"schedule,omitempty"`
}
