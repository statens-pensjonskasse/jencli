# Jencli

Et enkelt program for å starte jobber på Jenkins.
P.t. snakker `jencli` kun med en eksperimentell manuell deployjob for JPL.

## Installasjon

Kjør `make install` for å installere `jencli` under `${GOPATH}/bin` eller `make install-linux` for å installere
under `/usr/local/bin` (krever sudo).

## Konfigurasjon

Skaff et API Token fra Jenkins under konfigurasjon på kontoen din.

```shell
jencli config set \
  --jenkins-url http://jenkins.spk.no \
  --user <brukernavn> \
  --token <APIToken> \
  --slack <@slackbrukernavn> \
  --jpl-deploy-job job/DEV/job/JPL-Deploy/job/main
```

Dette lager en `config.yaml`-fil under `$HOME/.config/jencli/` men konfigurasjonen.

## Deploy

Kjør

```shell
jencli deploy --help
```

for en liste over argumenter.

For å deploye f.eks. `INFRA/batchrapportering` kan man kjøre

```shell
jencli deploy \
  --image batchrapportering \
  --branch master \
  --use-branch-postfix \
  --tag latest \
  --swarm utv \
  --env utv
```

Hvis du står i et repo og ønsker å deploye brancher du jobber på kan du kjøre

```shell
jencli deploy --swarm utv --env utv --use-branch-postfix
```

da vil `branch` utledes av hvilken branch du er på og `image` uteledes av mappa du står i.


Dersom du bruker nytt taggeregime for feature-branches (`latest_<branch>`) kan du bruke `--use-branch-tag`-flagget,
e.g.

```shell
jencli deploy --swarm team --env teammmmed1 --stack tmmmed1 --use-branch-tag
```

Eksempel på deploy til team-miljø

```shell
jencli deploy \
  --image opptjening-pensjon-simulering-for-nav-ws \
  --branch feature/SPKPRODUKT-5482_fra_jetty_til_spring_boot \
  --use-branch-postfix \
  --tag latest \
  --swarm team \
  --env teamaustralis3 \
  --stack taus3 \
  --slack medlemsdata-cd
```

Eksempel på bruk av harbor
```shell
jencli deploy \
  --swarm test \
  --env kpt \
  --stack kpt  \
  --use-branch-tag \
  --use-harbor
```