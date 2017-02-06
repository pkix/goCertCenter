package main

import (
	"fmt"
	//"io/ioutil"
	certcenter "../../goCertCenter"
)

// Set your Authorization Token
func init() {
	certcenter.Bearer = "YourValidToken.oauth2.certcenter.com"
}

func main() {
/*
	//////////////////////////////////////////////////////

	// Get my profile information
	res, _ := certcenter.Profile()
	fmt.Println(res)

	//////////////////////////////////////////////////////

	// Inquire limit informations
	res, _ := certcenter.Limit()
	fmt.Println(res)

	//////////////////////////////////////////////////////

	// Get all valid ProductCodes
	res, _ := certcenter.Products()
	fmt.Println(res)

	//////////////////////////////////////////////////////

	// Inquire details about a product
	res, _ := certcenter.ProductDetails("GeoTrust.QuickSSLPremium")
	fmt.Println(res)

	//////////////////////////////////////////////////////

	// Get a Quote
	res, _ := certcenter.Quote(&certcenter.QuoteRequest{
		ProductCode: "Symantec.SecureSiteEV",
		SubjectAltNameCount: 0,
		ValidityPeriod: 24,
		ServerCount: 1,
	})
	fmt.Println(res)

	//////////////////////////////////////////////////////

	// Validate a CSR
	csr, _ := ioutil.ReadFile("csr")
	res, _ := certcenter.ValidateCSR(&certcenter.ValidateCSRRequest{
		CSR: string(csr),
	})
	fmt.Println(res)

	//////////////////////////////////////////////////////

	// Get the CA's User Agreement
	res, _ := certcenter.UserAgreement("Symantec.SecureSite")
	fmt.Println(res)

	//////////////////////////////////////////////////////

	// Get valid email addresses for email based DV validation
	res, _ := certcenter.ApproverList(&certcenter.ApproverListRequest{
		CommonName: "www.example.com",
		ProductCode: "GeoTrust.QuickSSLPremium",
	})
	fmt.Println(res)

	//////////////////////////////////////////////////////

	// Order a certificate
	csr, _ := ioutil.ReadFile("csr")
	res, _ := certcenter.Order(&certcenter.OrderRequest{
			OrderParameters: certcenter.OrderParameters{
				ProductCode: "RapidSSL.RapidSSL",
				CSR: string(csr),
				ValidityPeriod: 24,
				ApproverEmail:"domains@certcenter.com",
			},
			AdminContact: certcenter.Contact{
				FirstName: "John",
				LastName: "Doe",
				Phone: "+1 212 999 999",
				Email: "john.doe@example.com",
			},
			TechContact: certcenter.Contact{
				FirstName: "John",
				LastName: "Doe",
				Phone: "+1 212 999 999",
				Email: "john.doe@example.com",
			},
		},
	)
	fmt.Println(res)

	//////////////////////////////////////////////////////

	// Re-set the approvers email address
	res, _ := certcenter.PutApproverEmail(&certcenter.PutApproverEmailRequest{
		CertCenterOrderID: 6277108,
		ApproverEmail: "domains@certcenter.com",
	})
	fmt.Println(res)

	//////////////////////////////////////////////////////

	// Re-send the approver email to the approvers address
	res, _ := certcenter.ResendApproverEmail(&certcenter.ResendApproverEmailRequest{
		CertCenterOrderID: 6277108,
	})
	fmt.Println(res)

	//////////////////////////////////////////////////////

	// Get filtered orders
	res, _ := certcenter.Orders(&certcenter.OrdersRequest{
		Status: "COMPLETE", // COMPLETE, PENDING, CANCELLED, REVOKED
		ProductType: "SSL", // SSL, CODESIGN, SMIME, TRUSTSEAL
		CommonName: "%",
		IncludeFulfillment: true,
		IncludeOrderParameters: true,
		IncludeBillingDetails: true,
		IncludeContacts: true,
		IncludeOrganizationInfos: true,
	})
	fmt.Println(res)

*/


	fmt.Println("")
	return
}
