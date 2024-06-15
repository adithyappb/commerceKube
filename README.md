## CommerceKube

## Overview
CommerceKube is a containerized microservices-based application using Docker and Kubernetes. It includes services for user management, product catalog, and order processing.

## Features

- **Microservices Architecture**: Separate services for users, products, and orders.
- **Containerization**: Each service runs in its own Docker container.
- **Kubernetes Deployment**: Deployments, services, and ingress configurations for Kubernetes.
- **Helm Charts**: Manage Kubernetes resources using Helm.

## Requirements

- Docker
- Kubernetes
- Helm

## Setup

# Docker Compose

# Build and run the services:
   docker-compose up --build

# Create the namespace:
kubectl apply -f k8s/namespace.yaml

# Deploy the services:
1) kubectl apply -f k8s/user-service-deployment.yaml
2) kubectl apply -f k8s/product-service-deployment.yaml
3) kubectl apply -f k8s/order-service-deployment.yaml
4) kubectl apply -f k8s/ingress.yaml

# Helm
# Install the Helm chart:
helm install microservices helm/

# Usage
1) Access the services via the ingress controller. Update your hosts file to point microservices.local to your Kubernetes cluster IP.
2) User Service: http://microservices.local/users
3) Product Service: http://microservices.local/products
4) Order Service: http://microservices.local/orders
