# Order Microservice - Go
**Stack**: *Go*, *MySQL*, *Kafka*, *Docker*

## Context
This repository contains a small example of a microservice, 100% operational. When changing the status of the order to TRUE, a message is sent to Kafka, running on Docker, so that another microservice (the download one) creates the link to download a fictitious purchase.

This code is part of a tutorial available at [dev.to/markgerald](https://dev.to/markgerald)
