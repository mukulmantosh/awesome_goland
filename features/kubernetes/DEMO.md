# Demo Walkthrough


### Kubernetes 


- Clone the repo https://github.com/mukulmantosh/go_kubernetes.git 
- Switch to `demo` branch.


#### Instructions

1. Move inside `k8s`directory.
2. First deploy `ns.yaml` to create namespace
3. Deploy `db.yaml`
4. Make sure before going to `app.yaml`, the database is up and running. 
5. Now, deploy `app.yaml`
 

To check the status of pods and make sure they are in **the RUNNING** state, run the following command.

```shell
kubectl get pods -n go-k8s-demo
```

To delete, make sure you follow these steps:

- First delete the `app.yaml`
- Then delete the `db.yaml`
- Then delete the `ns.yaml`


#### Reference Video

- [Kickstarting Your Kubernetes Journey with Go and GoLand
](https://www.youtube.com/watch?v=GGy4paf6rm0)