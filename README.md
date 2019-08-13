# sparebank1-personligklient

Implementasjon av Sparebank 1 sitt "[PM Accounts - 1.4.0](https://developersparebank1.no/index.php?option=com_apiportal&view=apitester&usage=api&tab=tests&apiName=PM%20Accounts&apiId=21670ee3-3524-47bd-b450-690c9c23b0de&menuId=147&managerId=1&apiVersion=1.4.0)"-API

Klasser generert ved (go-swagger)[https://github.com/go-swagger/go-swagger] og oAuth2-flyt ved [Go oAuth2](https://github.com/golang/oauth2).


Denne personlige klienten eksporterer transaksjoner fra en valgt konto til en csv-fil i følgende format:
```
Date,Payee,Category,Memo,Outflow,Inflow
03/06/2019,KontoEierFornavn,,YOUNGS,103.000000,
03/06/2019,KontoEierFornavn,,Straksoverføring,120.000000,
03/06/2019,KontoEierFornavn,,Overførsel,,240.000000
```

## Getting started

1. Registrer deg for tilgang gjennom personlig klient hos Sparebank 1. Det kan du gjøre her: https://developersparebank1.no/personlig-klient
1. Kjør `go get github.com/johnksv/sparebank1-personligklient; cd $GOPATH/src/github.com/johnksv/sparebank1-personligklient`
1. `go run main.go`

Applikasjonen vil spørre det etter credentials du har fått ved registrering av den personlige klienten.

Du kan sette ofte brukte variabler som env-variabler.NB: Variablene blir tilgjengelig i historikken til shellet. Du kan også sette dem gjennom kjøring med feks intelliJ.
```
export ClientID=<Din clientID>
export ClientSecret=<Din ClientSecret>
export FinInst=<Din FinInst, feks: fid-nord-norge>
export RedirectURI=<Din RedirectURI>
#Har du allerede et accessToken kan du eksportere dette også, om ikke vil du få det når du kjører opp applikasjonen
#export AccessToken=<Ditt accessToken>
```

