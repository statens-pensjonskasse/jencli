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
  --jpl-deploy-job job/INC/job/jpl-deploy/job/main
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
jencli deploy --image batchrapportering --branch master --use-branch-postfix --tag latest --swarm utv --env utv
```

Hvis du står på en branch og ønsker å deploye kan du kjøre

```shell
jencli deploy --swarm utv --env utv --use-branch-postfix
```

Dersom du bruker nytt taggeregime for feature-branches (`latest_<branch>`) kan du bruke `--use-branch-tag`-flagget,
e.g.

```shell
jencli deploy --swarm team --env tmmmed1 --stack teammmed1 --use-branch-tag
```

