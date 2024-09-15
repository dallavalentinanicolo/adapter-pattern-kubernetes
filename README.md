# README

## Multi-Channel Notification System in Go Using the Adapter Design Pattern

## Disclaimer

This project is a `education only` project created solely for the purpose of `learning Go` and demonstrating the `Adapter Design Pattern`.

The code and concepts presented here are for `educational` use only and are not intended for production environments or commercial applications. The implementation, including external libraries, is `simulated` to avoid incurring costs from external providers and simplified to support the learning goals.

## Overview

This project demonstrates the Adapter structural design pattern in Go by creating a multi-channel notification system. The system can send notifications through various channels: Email, SMS, and Push Notification.

The goal of the project is to show how the Adapter pattern can be used to integrate a new notification channel (Push Notifications) into an existing system that already supports Email and SMS, without changing the existing implementation.
Key Concepts

Notifier Interface: This is the common interface that all notification systems (Email, SMS, Push Notification) implement. It ensures that all notification methods have a SendNotification function.

Existing Systems (Email & SMS): The project initially supports Email and SMS notifications. These systems implement the Notifier interface directly.

External Push Notification System: Push Notifications are integrated through an Adapter because they come from a third-party library with a different interface. The Adapter allows the Push Notification system to be used as if it implemented the Notifier interface.

Adapter Pattern: The Adapter makes the Push Notification system compatible with the existing notification system by "adapting" the external interface to the common Notifier interface.

## How the Application Works

Fetching Pending Pods: The application connects to a Kubernetes cluster, checks for pending pods, and sends notifications if any are found.

Notification System: Notifications can be sent via:
Email: AWS or Simulated by printing to the console
SMS: Simulated by printing to the console.
Push Notifications: Integrated using the Adapter pattern, also simulated by printing to the console.

Prometheus Metrics: The application integrates with Prometheus to expose metrics about the number of pending pods.

HTTP Server: The application provides a simple HTTP server to view the current state of pending pods and check system health.

## Project Structure

main.go: The entry point of the application, which sets up the notification system, Prometheus metrics, and HTTP server.
clientk8s/: Manages Kubernetes client initialization.
prometheus/: Handles Prometheus metric exposure.
push/: Contains the Push Notification logic and the Adapter implementation.
resources/: Includes functionality for interacting with Kubernetes resources, like fetching pending pods.

## How to Run

Set up the environment: Make sure you have Go installed. Clone the repository and navigate to the project directory.

Run the application:

`go run *.go`

Access the server: Open your browser and go to `http://localhost:8080/` to see the pending pods, or `http://localhost:8080/check` for a health check.

Metrics: The Prometheus metrics are exposed at `http://localhost:8080/metrics`.

## Behavior

The application checks for pending Kubernetes pods every minute.
If there are any pending pods, a notification is sent via Email, SMS, and Push Notification (simulated by printing to the console).
If the number of pending pods changes (e.g., from some pending to none), a notification is sent.
Prometheus metrics are updated with the number of pending pods, allowing you to monitor them externally.