# tock-csv-to-users

Tool for building a tock authentication configmap (tock helm chart) from a basic csv file

## Usage

```shell
$ csv-to-tock-users -f ./samples/users.csv -o auth-confimap.yaml
```

csv file sample :

```csv
email,password,org,roles
admin@app.com, password,app,
demo1@app.com,demo1,app,nlpUser|botUser|admin|technicalAdmin
demo2@app.com,arnauddemo2,demo,
```

Target Comfigmap :

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: admin-web-auth-cfg
  labels:
    app.kubernetes.io/name: admin-web-auth
    app.kubernetes.io/component: admin-web
data:
  tock_users:  "admin@app.com,demo1@app.com,demo2@app.com" # Identifiants (séparés par des virgules). Valeur par defaut `admin@app.com`
  tock_passwords: "password,demo1,demo2" # Mots de passe (séparés par des virgules). Valleur par defaut `password`` 
  tock_organizations: "app,app,demo" # Organisations (séparées par des virgules). Valleur par defaut `app``
  tock_roles: ",nlpUser|botUser|admin|technicalAdmin,admin," #  Rôles séparés par des | (puis par des virgules). Valeur par defaut vide , Dans cet exemple, Alice a le rôle botUser, alors que Bob a tous les rôles.
```

```
     ┌─────┐
   ┌─┤ CSV ├─────────────────────────────────────────────────────┐
   │ └─────┘                                                     │
   │email,password,org,roles                                     │
   │admin@app.com, password,app,                                 │
   │demo1@app.com,demo1,app,nlpUser|botUser|admin|technicalAdmin │
   │demo2@app.com,demo2,demo,                                    │
   │                                                             │
   └──────────────────────────┳──────────────────────────────────┘
                              ┃
                              ┃
                      csv-to-tock-users
                              ┃
     ┌────────────┐           ┃
   ┌─┤ ConfigMap  ├───────────▼───────────────────────────────────┐
   │ └────────────┘                                               │
   │apiVersion: v1                                                │
   │kind: ConfigMap                                               │
   │metadata:                                                     │
   │  name: admin-web-auth-cfg                                    │
   │  labels:                                                     │
   │    app.kubernetes.io/name: admin-web-auth                    │
   │    app.kubernetes.io/component: admin-web                    │
   │data:                                                         │
   │  tock_users: "admin@app.com,demo1@app.com,demo2@app.com"     │
   │  tock_passwords: " password,demo1,demo2"                     │
   │  tock_organizations: "app,app,demo"                          │
   │  tock_roles: ",nlpUser|botUser|admin|technicalAdmin,"        │
   │                                                              │
   └──────────────────────────────────────────────────────────────┘
```

## Tock Roles

Information about available roles can be found in Tock documentation:

https://doc.tock.ai/tock/master/admin/security.html#roles