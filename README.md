# README

## Multi-Channel Notification System in Go Using the Adapter Design Pattern

## Disclaimer

This project is a `education only` project created solely for the purpose of `learning Go` and demonstrating the `Adapter Design Pattern` with Kubernetes.

The code and concepts presented here are for `educational` use only and are not intended for production environments or commercial applications. The implementation, including external libraries, is `simulated` to avoid incurring costs from external providers and simplified to support the learning goals.

## Overview

This project demonstrates the Adapter structural design pattern in Go by creating a multi-channel notification system. The system can send notifications through various channels: Email, SMS, and Push Notification via Telegram.

The goal of the project is to show how the Adapter pattern can be used to integrate a new notification channel (Push Notifications) into an existing system that already supports Email and SMS, without changing the existing implementation.

Adapter Pattern: The Adapter makes the Push Notification system compatible with the existing notification system by "adapting" the external interface to the common Notifier interface.

## How the Application Works

Fetching Pending Pods: The application connects to a Kubernetes cluster, checks for pending pods, and sends notifications if any are found.

Notification System: Notifications can be sent via:
Email: External Smtp server
SMS: Simulated by printing to the console.
Push Notifications Telegram: Integrated using the Adapter pattern, also simulated by printing to the console.

Prometheus Metrics: The application integrates with Prometheus to expose metrics about the number of pending pods.

HTTP Server: The application provides a simple HTTP server to view the current state of pending pods and check system health.

## Project Structure

main.go: The entry point of the application, which sets up the notification system, Prometheus metrics, and HTTP server.
clientk8s/: Manages Kubernetes client initialization.
prometheus/: Handles Prometheus metric exposure.
push/: Contains the Push Notification logic and the Adapter implementation.
resources/: Includes functionality for interacting with Kubernetes resources, like fetching pending pods.
my-pending-pods-helm/ contains helm-chart to test pod-pending
testing-pending-pod/ contains the manifest to test pod-pending

## How to Run

To test I used `kind`. Then you need to install `kubectl` and `make`. Then to setup the environment you can run makefile `make apply-all`

Set up the environment: Make sure you have Go installed. Clone the repository and navigate to the project directory.

Run the application:

`go run *.go`

Access the server: Open your browser and go to `http://localhost:8080/` to see the pending pods, or `http://localhost:8080/check` for a health check.

Metrics: The Prometheus metrics are exposed at `http://localhost:8080/metrics`.

## Behavior

The application checks for pending Kubernetes pods every minute.
If there are any pending pods, a notification is sent via Email, SMS(simulated by printing to the console), and Push Notification.
If the number of pending pods changes (e.g., from some pending to none), a notification is sent.
Prometheus metrics are updated with the number of pending pods, allowing you to monitor them externally.