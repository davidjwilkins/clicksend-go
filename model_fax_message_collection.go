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

// Array of FaxMessage items
type FaxMessageCollection struct {
	// Array of FaxMessage items
	Messages []FaxMessage `json:"messages"`
	// URL of file to send
	FileUrl string `json:"file_url"`
}
