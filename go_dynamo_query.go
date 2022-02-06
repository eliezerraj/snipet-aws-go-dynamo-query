package main

import (
	"fmt"
	"flag"
	"os"
	"context"
	"time"
	"math/rand"
	"strconv"

	"github.com/jaswdr/faker"
	"github.com/spf13/viper"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	entity "github.com/snipet-aws-go-dynamo-create-query/main/entity"
	service "github.com/snipet-aws-go-dynamo-create-query/main/service"
	repository "github.com/snipet-aws-go-dynamo-create-query/main/repository"
)

func main(){
	fmt.Println("Starting")

	option 	:= flag.String("option","","")
    pk  	:= flag.String("pk","","")
    sk  	:= flag.String("sk","","")
	table 	:= flag.String("table","","")
	order_id := flag.String("order_id","","")
	amount  := flag.String("amount","","")

	flag.Parse()
	fmt.Printf("table: %s option: %s pk: %v sk: %s \n", *option, *pk, *sk, *table)

	aws_region 		:= getEnvVar("AWS_REGION")
	aws_access_id 	:= getEnvVar("AWS_ACCESS_ID")
	aws_access_secret := getEnvVar("AWS_ACCESS_SECRET")

	//fmt.Printf("aws_region: %s aws_access_id: %s aws_access_secret: %v \n", aws_region, aws_access_id, aws_access_secret)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(aws_region),
		Credentials: credentials.NewStaticCredentials( aws_access_id , aws_access_secret , ""),},
	)
	if err != nil {
		fmt.Println("Erro Create aws Session: ",err.Error())
		os.Exit(1)
	}

	svc := dynamodb.New(sess)

	fmt.Println("svc : ",svc)

	repo := repository.NewRepository(svc, *table)
	service := service.NewService(repo)
	ctx := context.Background()

	switch *option {
		case "load_spleinir":
			rand.Seed(time.Now().UnixNano())

			randon_i := rand.Intn(53)
			f := faker.New()
			issuer := entity.Issuer{}
			issuer.Id = "id-" + strconv.Itoa(randon_i)
			issuer.Name = f.Person().Name()
			fmt.Println(issuer)

			bindata := entity.Bindata{}
			bindata.Issuer = issuer

			var table_name = "issuer"
			err := service.Save(ctx, &table_name ,issuer )
			if err != nil {
				fmt.Println("Erro Store: ",err.Error())
				os.Exit(1)
			}

			for i:=0; i < 1; i++{
				randon_p := rand.Intn(9000000)
				randon_u := rand.Intn(90000)
				person := entity.Person{}

				person.Id = "person_id-" + strconv.Itoa(randon_p)
				person.Issuer_id = "issuer_id-" + strconv.Itoa(randon_i)
				person.Unique_Id = "unique_id-" + strconv.Itoa(randon_u)
				person.Metada = "{ metadata... }"
				
				bindata.Person = person

				fmt.Println(person)

				table_name = "person"	
				err := service.Save(ctx, &table_name ,person )
				if err != nil {
					fmt.Println("Erro Store: ",err.Error())
					os.Exit(1)
				}
			}

			for i:=0; i < 1; i++{
				randon_prd := rand.Intn(30000)
				product := entity.Product{}

				product.Id = "product_id-" + strconv.Itoa(randon_prd)
				product.Issuer_id = "issuer_id-" + strconv.Itoa(randon_i)
				product.Status = "ACTIVE"
				product.Metada = "{ metadata... }"

				bindata.Product = product
				
				fmt.Println(product)
				table_name = "product"	
				err := service.Save(ctx, &table_name ,product )
				if err != nil {
					fmt.Println("Erro Store: ",err.Error())
					os.Exit(1)
				}
			}

			for i:=0; i < 1; i++{
				randon_acc := rand.Intn(70000)
				account := entity.Account{}

				account.Id = "account_id-" + strconv.Itoa(randon_acc)
				account.Issuer_id = "issuer_id-" + strconv.Itoa(randon_i)
				account.Status = "ACTIVE"
				account.Metada = "{ metadata... }"

				bindata.Account = account

				fmt.Println(account)
				table_name = "account"	
				err := service.Save(ctx, &table_name ,account )
				if err != nil {
					fmt.Println("Erro Store: ",err.Error())
					os.Exit(1)
				}
			}

			for i:=0; i < 1; i++{
				randon_bin := rand.Intn(1000000)

				bindata.Id = "car_id-" + strconv.Itoa(randon_bin)
				bindata.Issuer_id = "issuer_id-" + strconv.Itoa(randon_i)
				bindata.Status = "ACTIVE"

				fmt.Println(bindata)
			}

		case "load_invoice":
			rand.Seed(time.Now().UnixNano())

			randon_i := rand.Intn(100)
			randon_c := rand.Intn(20)
			randon_t := rand.Intn(2)
			
			for i:=0; i < 1; i++{
				f := faker.New()
				tenant := "tenant-"+ strconv.Itoa(randon_t)

				invoice := entity.Invoice{}
				invoice.Pk = "invoice-" + strconv.Itoa(randon_i)
				invoice.Sk = "customer-" + strconv.Itoa(randon_c)
				invoice.Name = "Ms. Customer Acme " + strconv.Itoa(randon_i)//f.Person().Name()
				invoice.Tenant = "tenant-"+ strconv.Itoa(randon_t)

				fmt.Println(invoice)
				err := service.AddInvoice(ctx, invoice)
				if err != nil {
					fmt.Println("Erro Store: ",err.Error())
					os.Exit(1)
				}

				var total_invoice float32 = 0.0
				qtd_invoice := 1
				for a:=0; a < 2; a++{
					randon_o := rand.Intn(200)
					invoice = entity.Invoice{}
					invoice.Pk = "invoice-" + strconv.Itoa(randon_i)
					invoice.Sk = "order-" + strconv.Itoa(randon_o)
					invoice.Tenant = tenant
					invoice.Sku  =  "sku-" + strconv.Itoa(rand.Intn(100))
					invoice.Name = f.Car().Model()
					invoice.Qtd  = rand.Intn(3) + 1
					invoice.Price = (rand.Float32()*10) + 1
					total_invoice = total_invoice + (invoice.Price * float32(invoice.Qtd))
		
					fmt.Println(invoice)
					err = service.AddInvoice(ctx, invoice)
					if err != nil {
						fmt.Println("Erro Store: ",err.Error())
						os.Exit(1)
					}
				}

				invoice = entity.Invoice{}
				invoice.Pk = "invoice-" + strconv.Itoa(randon_i)
				invoice.Sk = "invoice-" + strconv.Itoa(randon_i)
				invoice.Tenant = "tenant-"+ strconv.Itoa(randon_t)
				invoice.Amount = float32(qtd_invoice) * total_invoice
		
				fmt.Println(invoice)
				err = service.AddInvoice(ctx, invoice)
				if err != nil {
					fmt.Println("Erro Store: ",err.Error())
					os.Exit(1)
				}
			}

		case "query_invoice":
			result, err := service.QueryInvoice(ctx, *pk, *sk)
			if err != nil {
				fmt.Println("Erro QueryInvoice: ",err.Error())
				os.Exit(1)
			}
			fmt.Printf("result - > %s \n", result)
		case "query_invoice_gsi":
			result, err := service.QueryInvoiceGsi(ctx, *pk, *sk)
			if err != nil {
				fmt.Println("Erro QueryInvoiceGsi: ",err.Error())
				os.Exit(1)
			}
			fmt.Printf("result - > %s \n", result)
		case "update_transaction":
			value_amount, err := strconv.ParseFloat(*amount, 32)
			if err != nil {
				fmt.Println("Erro float conversion: ",err.Error())
				os.Exit(1)
			}
			result, err := service.UpdateInvoiceTransaction(ctx, *pk, *sk, *order_id, float32(value_amount))
			if err != nil {
				fmt.Println("Erro UpdateInvoiceTransaction: ",err.Error())
				os.Exit(1)
			}
			fmt.Printf("result - > %s \n", result)
		default:
			fmt.Printf("option: %s invalid \n", *option)  
		}
}

func getEnvVar(key string) string {
	fmt.Printf("Loading enviroment variable %s .. \n", key)
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file %s \n", err)
		os.Exit(1)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		fmt.Printf("Invalid type \n")
		os.Exit(1)
	}
	return value
}

