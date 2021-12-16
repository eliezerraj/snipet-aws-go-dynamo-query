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
	Tenant           string 	`json:"pk_gsi,omitmepty"`  
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