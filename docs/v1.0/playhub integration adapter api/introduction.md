---
title: Introduction
excerpt: Playhub Integration Adapter Introduction.
category: 65bc99c69969910010b1426c
order: -1
hidden: false
slug: playhub-integration-adapter-intro
---

All requests from our platform to your should be processed in idempotent way. 
This means that if we send the same request multiple times, the result should be the same as if we sent the request only once. 
This is important to avoid any issues with your system in case of network issues or other problems.

we use the following IPs to access your endpoints:
staging environment: 34.76.77.205
production environment: 34.140.75.241