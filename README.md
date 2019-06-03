# sparebank1-personligklient

Implementasjon av Sparebank 1 sitt [PM Accounts - 1.4.0](https://developersparebank1.no/index.php?option=com_apiportal&view=apitester&usage=api&tab=tests&apiName=PM%20Accounts&apiId=21670ee3-3524-47bd-b450-690c9c23b0de&menuId=147&managerId=1&apiVersion=1.4.0)-API

Klasser generert ved (go-swagger)[https://github.com/go-swagger/go-swagger] og oAuth2-flyt ved [Go oAuth2](https://github.com/golang/oauth2).
Eksporterer transaksjoner fra en valgt konto til en csv-fil i følgende format:
```
Date,Payee,Category,Memo,Outflow,Inflow
03/06/2019,KontoEierFornavn,,YOUNGS,103.000000,
03/06/2019,KontoEierFornavn,,Straksoverføring,120.000000,
03/06/2019,KontoEierFornavn,,Overførsel,,240.000000
```


TODO:
- Spesifisere kategori 
