# xen

Steps to execute -

1 . go get -v
2 . go run main.go

```
curl --location --request POST 'localhost:5000/cash_disbursements/initiate' \
--header 'authorization: secrettoken' \
--header 'Content-Type: application/json' \
--data-raw '
{
	"disbursement_code": "12345",
	"beneficiary_id_name": "Saurav",
	"location": {
		"branch_id": "branch_manila_412",
		"branch_name": "City Branch",
		"address": "One Global Place, Level 7-1",
		"city": "Taguig City",
		"phone_number": "639051269473",
		"operator_name": "John Doe",
		"metadata": {
			"extra_field_1": 20.0,
			"extra_field_2": "2020-05-22T16:04:35.057Z"
		}
	}
}
'
```

```
curl --location --request POST 'localhost:5000/cash_disbursements/commit' \
--header 'authorization: secrettoken' \
--header 'Content-Type: application/json' \
--data-raw '
{
	"ref_number": "dis-5594002ad0a04cd9b3f7",
	"amount": 2000
	
}
'
```
