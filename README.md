# tok8s
A Continuous Delivery Tool For Google Container Registry

### Table of Contents
- [tok8s]
  * [High Level Design](#high-level-design)
  * [Setup](#setup)

## High Level Design

Google Container Registry (GCR) build trigger can be configured to trigger on git commits on a branch or new tags. The trigger can then kick off a multi-stage docker build after all test pass and push the newly built  docker image into GCR. Upon a successful docker image push into GCR, a message will be published to the grc topic of the project as long as the topic exists. tok8s will subscribe to the topic and get the docker image from the message, it will automatically deploy the docker image to kubernetes.

The design diagram can be found at [TODO:]


## Setup
* Check out the repo
* If dep isn't installed, install dep
```
brew install dep
brew upgrade dep
```
* Install dependencies
```
$ dep ensure
```


 