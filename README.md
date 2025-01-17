# ChartMuseumUI
[![HitCount](http://hits.dwyl.io/idobry/chartmuseumui.svg)](http://hits.dwyl.io/idobry/chartmuseumui) [![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/dwyl/esta/issues)

<img src="./logo.png" width="300">

# ChartMuseumUI

ChartMuseumUI is a simple web app that (currently) provides GUI for your charts so you and your team can view and share the technologies your are using to any one at any time (in near future more capabilities will be added).
ChartMuseumUI was written in Go (Golang) with the help of Beego Framework.

<img src="./combine-gif.gif" width="2000">

## Getting Started

These instructions will get you with your very own private charts repository. You can run this on your localmachine, on cloud and basically on every machine that have docker and docker-compose installed.

### Usage

ChartMuseumUI using [ChartMuseum](https://github.com/helm/chartmuseum) as a backend so the best way would be to use docker-compose. 

For example, the following docker-compose file is defining ChartMuseum with Amazon S3 as a storage and exposing ChartMuseumUI on port 80 
```
version: '2.0'

services:
   ui:
     image: idobry/chartmuseumui:latest
     environment:
      CHART_MUSESUM_URL: "http://chartmuseum:8080"
     ports:
      - 80:8080
   chartmuseum:
     image: chartmuseum/chartmuseum:latest
     volumes:
       - ~/.aws:/root/.aws:ro
     restart: always
     environment:
      PORT: 8080
      DEBUG: 1
      STORAGE: "amazon"
      STORAGE_AMAZON_BUCKET: "chartmuseum-bucket"
      STORAGE_AMAZON_PREFIX: ""
      STORAGE_AMAZON_REGION: "eu-west-1"
     ports:
      - 8080:8080
```

Copy this file and run

```
docker-compose up 
```
Easy, right? now, we can add our private repository to our Helm client:
```
# choose any name you like
$ helm repo add chartmuseum <chartmuseum-url>
$ helm repo update
# to view our repos list
$ helm repo list
NAME        URL
stable      https://kubernetes-charts.storage.googleapis.com
incubator   http://storage.googleapis.com/kubernetes-charts-incubator
chartmuseum http://localhost:8080
```

Let's upload a chart into our private repository:

```
$ cd /chart/path
# create a chart package - this will create a .tgz file
$ helm package .
# copy packge name and run
$ curl -L --data-binary "@<packge-name>" <chartmuseum-url>/api/charts
```
   
In the browser, navigate to localhost and view your charts


## Built With

* [beego](https://beego.me/) - The web framework used
* [go](https://golang.org/) - Programing language
* [docker](https://www.docker.com/) - Packaged with docker


## Project Roadmap
* Add login screen
* Add more chartmuseum capabilitis:
   - Upload a chart
     - Support multiple
   - Delete a chart
     - Ask before deleting
     - Delete all versions button
     - Back to 'home' after delete all

## Contributing

Code contributions are very welcome. If you are interested in helping make chartmuseumui great then feel free!


## Authors

* **Ido Braunstain** - *Initial work*

See also the list of [contributors](https://github.com/idobry/contributors) who participated in this project.



