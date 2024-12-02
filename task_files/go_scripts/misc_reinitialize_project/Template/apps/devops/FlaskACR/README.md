# Preview


in local folder


```sh
docker build -t [PROJECT_NAME]prodcontainerregistry.azurecr.io/flask-backend-preview:0.2.4 -f .\apps\devops\FlaskACR\Dockerfile.preview.alpine .


```

in az-cli-helper container

```sh

az acr login --name [PROJECT_NAME]prodcontainerregistry; az acr update -n [PROJECT_NAME]prodcontainerregistry --admin-enabled true; docker push [PROJECT_NAME]prodcontainerregistry.azurecr.io/flask-backend-preview:0.2.4

```

# prod



in local folder


```sh
docker build -t [PROJECT_NAME]prodcontainerregistry.azurecr.io/flask-backend-prod:0.2.4 -f .\apps\devops\FlaskACR\Dockerfile.prod.alpine .
```

in az-cli-helper container


```sh

az acr login --name [PROJECT_NAME]prodcontainerregistry; az acr update -n [PROJECT_NAME]prodcontainerregistry --admin-enabled true;docker push [PROJECT_NAME]prodcontainerregistry.azurecr.io/flask-backend-prod:0.2.4

# then



```



