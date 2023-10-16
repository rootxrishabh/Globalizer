# Globalizer
Globalizer is a kubernetes controller which to globalize your application to the world.

## Crux
This controller automatically creates service and ingress resource when you deploy your application using deployment resource, and removes the service and ingress as soon as you delete your deployment. The controller leverages client-go, api and apimachinery libraries in order to implement the given functionality.

## Steps
`make deploy` will take care of everything as long as you have a k8s cluster running : )

## Installation logic
1) Nginx-load-balancer is deployed as a component to create ingress
2) Namespace is created to allocate upcoming deployment.
3) Service account, cluster role and cluster role binding are created to grant permissions to the custom controller.
4) Deployment is created and the controller is deployed.

### Controller architecture
![image](https://github.com/rootxrishabh/Globalizer/assets/73490663/acbbcafe-f883-4ef1-b49c-d1ce47c80965)
