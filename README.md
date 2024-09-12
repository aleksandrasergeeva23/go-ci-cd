# Go CI/CD Pipeline Automation Tool

This project is a simple, lightweight Continuous Integration/Continuous Deployment (CI/CD) pipeline tool built using Go. The tool runs custom-defined stages like build, test, and deployment in Docker containers and integrates with Kubernetes for final deployments.

## Features

- **Pipeline as Code**: Define build, test, and deployment stages in a YAML configuration file.
- **Docker Integration**: Each stage runs inside Docker containers.
- **Concurrent Builds**: Uses Goroutines for running pipeline stages concurrently.
- **Monitoring Dashboard**: (Optional) Real-time monitoring of pipeline stages via a web interface.
- **Error Handling & Retries**: Automatic retries for failed stages with detailed logs.
- **GitHub Actions Support**: Automatically trigger pipeline runs for every new commit.
- **Kubernetes Deployment**: The final deployment stage pushes the application to a Kubernetes cluster.