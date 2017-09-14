This project helps getting Docker host IP from containers.    
This is often required in real life (Consul, Kafka, ...).  


# Run the container
```
docker run --privileged --net=host --restart=always hjdr4/docker-metadata [iface]
```

If `iface` is not passed, the daemon returns the first interface without 'lo' nor 'docker' in its name.

The container requires some privileges to create a dummy interface and bind port 80.  


# Get host IP from containers
```
curl http://169.254.169.254/latest/meta-data/local-ipv4
```


# Build the container yourself
```
go build -o docker/docker-metadata ./main && docker build -t docker-metadata docker && rm -f docker/docker-metadata
```


# Licence
MIT

