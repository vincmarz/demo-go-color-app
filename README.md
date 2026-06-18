# demo-go-color-app

demo-go-color-app e una semplice applicazione demo scritta in Go.

Applicazione espone un web server HTTP sulla porta 8080 e mostra una pagina HTML con colore di sfondo configurabile tramite variabile ambiente.

Il progetto e utilizzato nel laboratorio DevOps su OpenShift per validare un flusso applicativo basato su:

- repository Git applicativo;
- build container tramite OpenShift Pipelines, Tekton e Buildah;
- pubblicazione immagine su OpenShift Internal Registry;
- deploy dichiarativo tramite OpenShift GitOps e Argo CD;
- modifica applicativa visibile tramite cambio colore della pagina.

## Struttura repository

```text
demo-go-color-app/
|-- Containerfile
|-- README.md
|-- go.mod
`-- main.go
```

## Endpoint

| Endpoint | Descrizione |
|---|---|
| / | Pagina HTML demo. |
| /healthz | Endpoint di health check. Restituisce ok. |

## Variabili ambiente

| Variabile | Default | Descrizione |
|---|---|---|
| PAGE_COLOR | #1E90FF | Colore di sfondo della pagina HTML. |
| APP_VERSION | v1 | Versione applicativa mostrata nella pagina. |

## Esecuzione locale

```bash
go build -o app main.go
./app
curl http://localhost:8080/healthz
```

Output atteso:

```text
ok
```

## Build container

```bash
buildah bud -f Containerfile -t demo-go-color-app:latest .
```

## Repository GitOps correlato

Il deploy applicativo viene gestito da repository GitOps separato:

```text
https://github.com/vincmarz/demo-app-gitops.git
```

## Sicurezza

Non committare token, kubeconfig, dockerconfig, Secret Kubernetes, password, chiavi private o certificati privati.

Controllo consigliato prima del commit:

```bash
grep -RniE 'ghp_|github_pat_|BEGIN RSA|PRIVATE KEY|AKIA|ASIA|config.json' . --exclude-dir=.git || true
```

## Stato iniziale

```text
Repository: https://github.com/vincmarz/demo-go-color-app.git
Branch: main
Commit iniziale: 6406c84
Applicazione: demo-go-color-app
Linguaggio: Go
Porta applicativa: 8080
Health endpoint: /healthz
Container build: Containerfile
```
