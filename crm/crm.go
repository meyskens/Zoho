package crm

import (
	zoho "github.com/schmorrison/Zoho"
	"math/rand"
	"time"
)

type CRMModule string

// Proper names for CRM modules
const (
	AccountsModule       CRMModule = "Accounts"
	CallsModule          CRMModule = "Calls"
	CampaignsModule      CRMModule = "Campaigns"
	CasesModule          CRMModule = "Cases"
	ContactsModule       CRMModule = "Contacts"
	CustomModule         CRMModule = "Custom"
	DealsModule          CRMModule = "Deals"
	EventsModule         CRMModule = "Events"
	InvoicesModule       CRMModule = "Invoices"
	LeadsModule          CRMModule = "Leads"
	PotentialsModule     CRMModule = "Potentials"
	PriceBooksModule     CRMModule = "PriceBooks"
	ProductsModule       CRMModule = "Products"
	PurchaseOrdersModule CRMModule = "PurchaseOrders"
	QuotesModule         CRMModule = "Quotes"
	SalesOrdersModule    CRMModule = "SalesOrders"
	SolutionsModule      CRMModule = "Solutions"
	TasksModule          CRMModule = "Tasks"
	VendorsModule        CRMModule = "Vendors"
)

// API is used for interacting with the Zoho CRM API
// the exposed methods are primarily access to CRM modules which provide access to CRM Methods
type API struct {
	*zoho.Zoho
	id string
}

// New returns a *crm.API with the provided zoho.Zoho as an embedded field
func New(z *zoho.Zoho) *API {
	id := func() string {
		var id []byte
		keyspace := "abcdefghijklmnopqrutuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		for i := 0; i < 25; i++ {
			source := rand.NewSource(time.Now().UnixNano())
			rnd := rand.New(source)
			id = append(id, keyspace[rnd.Intn(len(keyspace))])
		}
		return string(id)
	}()

	return &API{
		Zoho: z,
		id:   id,
	}
}
