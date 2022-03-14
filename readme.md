#Install instructions
```
 1. go mod vendor #gets all dependencies
 2. go install    #install to $GOBIN
 3. locate & add /go/bin to $PATH (ex: /root/go/bin)
```

#Usage
```
pluck -v common_config.yaml -t template.gotext | tail -n +2 > docker_compose.yml  #creates docker-compose file
```

_Note:-_
-
check ./example dir for samples
