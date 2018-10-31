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

// PostPostcard model
type PostPostcard struct {
	// Postcard file URLs
	FileUrls []string `json:"file_urls"`
	// Array of recipients
	Recipients []PostRecipient `json:"recipients"`
}
