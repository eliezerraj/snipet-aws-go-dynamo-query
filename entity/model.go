package entity

type Order struct {
	Sku 			string 	 `json:"sku,omitempty"`
	Name			string 	 `json:"name,omitempty"`
	Qtd 			int		 `json:"qtd,omitempty"`
    Price 			float32  `json:"price,omitempty"`
	Tenant          string 	 `json:"tenant,omitmepty"`  
}

type Invoice struct {
	Pk      		 string     `json:"pk,omitempty"`
	Sk    		 	 string  	`json:"sk,omitempty"`
	Tenant           string 	`json:"tenant_gsi,omitmepty"`  
	Amount 			 float32    `json:"amount,omitempty"` 
	Sku 			 string 	`json:"sku,omitempty"`
	Name			 string 	`json:"name,omitempty"`
	Qtd 			 int		`json:"qtd,omitempty"`
    Price 			 float32    `json:"price,omitempty"`
}

type Customer struct {
	Id 		int `json:"id,omitempty"`
    Name 	string `json:"name,omitempty"`
	Email 	string `json:"email,omitempty"`
	Tenant  string 	`json:"tenant,omitmepty"`  
}

//-----------------------------------

type Person struct {
	Id 			string `json:"person_id,omitempty"`
	Issuer_id	string `json:"issuer_id,omitempty"`
	Unique_Id	string `json:"unique_customer_id,omitempty"`
	Status		string `json:"status,omitempty"`
    Metada 		string `json:"metadata,omitempty"`
	ExpirationDate	int	`json:"expiration_date,omitempty"`
}

type Account struct {
	Id 				string `json:"account_id,omitempty"`
	Issuer_id		string `json:"issuer_id,omitempty"`
    External_Id 	string `json:"external_id,omitempty"`
	Status			string `json:"status,omitempty"`
	Metada 			string `json:"metadata,omitempty"`
	ExpirationDate	int	`json:"expiration_date,omitempty"`
}

type Product struct {
	Id 				string `json:"product_id,omitempty"`
	Issuer_id		string `json:"issuer_id,omitempty"`
	Status			string `json:"status,omitempty"`
	Metada 			string `json:"metadata,omitempty"`
	ExpirationDate	int	`json:"expiration_date,omitempty"`
}

type Issuer struct {
	Id 				string `json:"issuer_id,omitempty"`
	Name			string `json:"name,omitempty"`
	Status			string `json:"status,omitempty"`
	Metada 			string `json:"metadata,omitempty"`
	ExpirationDate	int	`json:"expiration_date,omitempty"`
}

type Bindata struct {
	Id 					string 	`json:"card_id,omitempty"`
	Issuer_id			string 	`json:"issuer_id,omitempty"`
	Account_id			string 	`json:"account_id,omitempty"`
	Product_id			string 	`json:"product_id,omitempty"`
	Person_id			string 	`json:"person_id,omitempty"`
	Status				string 	`json:"status,omitempty"`
	Account  			Account `json:"account,omitmepty"`  
    Person 				Person 	`json:"person,omitempty"`
	Product 			Product `json:"product,omitempty"`
	Issuer  			Issuer 	`json:"issuer,omitmepty"`  
	ExpirationDate		int	`json:"expiration_date,omitempty"`
}