# Quick Start


## Install

`athenapdf-service` (Weaver) is packaged, and distributed as a [Docker][docker] image.

Before starting, ensure your Docker environment is [set up][docker], and ready-to-use. **For OSX / Windows users**, ensure your [Docker Machine][docker-machine] is prepared, and the appropriate environment variables are established.

It can be installed using:

```bash
docker pull arachnysdocker/athenapdf-service
```

You can run the same command above to get updates.


## Deployment

### Running

Copy [`conf/sample.env`][sample] (from source), and modify it accordingly.

```bash
# Docker Engine
docker run -p 8080:8080 --rm arachnysdocker/athenapdf-service

# Docker Compose
# Uses `conf/sample.env` by default, change it in `docker-compose.yml`
docker-compose up
```

### Configuration

`athenapdf-service` expects the configuration variables to be set in the environment.

They can be passed to the container via `--env-file` (as seen above) or `--env` (refer to Docker's [documentation][docker-run] on `run` for more information).

See [`conf/sample.env`][sample] for a list of available configuration options.

#### Statsd

[Statsd][statsd] is used for capturing time-series metrics, and it can be used to build lovely dashboards to visualise them.

![Grafana](https://s3-eu-west-1.amazonaws.com/athena-pdf-public/examples/grafana.png)

Stats are sent to the following buckets:

Bucket | Type | Description
--- | --- | ---
`conversion_duration` | Timer | Time taken for a successful conversion
`success` | Counter | Incremented for every successful conversion
`conversion_timeout` | Counter | Incremented for every conversion work that timed out (the timeout can be increased through `WEAVER_WORKER_TIMEOUT`)
`s3_upload_error` | Counter | Incremented when a conversion has failed to be uploaded to S3
`conversion_error` | Counter | Incremented when a conversion error has occurred
`cloudconvert` | Counter | Incremented when converting with CloudConvert as a fallback
`conversion_failed` | Counter | Incremented when a conversion has failed

### Amazon Web Services

At [Arachnys][arachnys], it is deployed using Amazon's _new_ [EC2 Elastic Container Service (ECS)][ecs]. We run one container (task) per container instance (EC2 instance with Amazon's Docker agent), and we route requests across multiple instances through [Elastic Load Balancer][elb].

### Docker Cloud

To be added.

### Scaling

Weaver's Docker image is completely stateless, and it was built to be scalable, both horizontally, and vertically. It works out of the box without any user configuration.

To scale horizontally, simply deploy more instances of the microservice, and set up a load balancer or a stateful task manager to distribute the traffic across your instances.

If you are scaling vertically (better hardware), increase the number of concurrent workers, and the size of the work queue accordingly.


[statsd]: https://github.com/etsy/statsd
[docker]: https://www.docker.com/
[docker-machine]: https://docs.docker.com/mac/step_one/
[docker-run]: https://docs.docker.com/engine/reference/commandline/run/
[arachnys]: https://www.arachnys.com/?utm_campaign=athena&utm_medium=external%20website&utm_source=github&utm_content=readme
[ecs]: https://aws.amazon.com/ecs/
[elb]: https://aws.amazon.com/elasticloadbalancing/
[sample]: ../conf/sample.env
