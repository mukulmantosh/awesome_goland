# Demo Walkthrough

### Open Shell in Kubernetes

![demo](./k8s_shell.gif)

Select the pod you want from the _Kubernetes_ view in the _Services_ tool window, then click on the _Run Shell_ button. A new tab will open and a terminal will be attached to the running pod.

By default, the IDE runs _/bin/bash_ as a shell. To run a different one, click **Show Settings** or open **Settings/Preferences | Build, Execution, Deployment | Kubernetes** and specify the shell that your pods use.
