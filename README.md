# Go-Simple-App
This serves as a lightweight demonstration application for debugging code issues.

To execute this on your Kubernetes cluster, please follow these steps:

1. Create a Kubernetes namespace called 'goapp' using the command: `kubectl create ns goapp`.

2. Compile service A and service B by running the following.
   - `cd http_go_demo_app/service_a`
   - `make`
   - `cd http_go_demo_app/service_b`
   - `make`


3. Navigate to your k8s folder and execute the following commands:
   - `kubectl apply -f service_a.yaml`
   - `kubectl apply -f service_n.yaml`


4. port forward service-a, run:
   - `kubectl port-forward svc/service-a 8080:80 -n goapp`

5. from terminal (your host) run:
   - `curl http://localhost:8080`

