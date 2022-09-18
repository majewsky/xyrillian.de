# My standard service architecture

I have been working in the IT industry for over 10 years, first in a pure developer position, then ever so slowly
transitioning into an architect role. Over this time, experience with all the systems I have worked on has compounded
into a recurring design template that I apply to most of my services. In this post, I will describe this architecture,
working my way up from data model to component model to runtime environment, and provide justification for the principal
design decisions. Finally, I will acknowledge and discuss some limitations of this design.

The design presented here is mostly agnostic towards:

- *which programming language is used*: My services are implemented in [Go][golang], but the architecture should also
  be applicable to most other languages. (As a caveat, if you are using an opinionated framework like Ruby on Rails, it
  is usually a bad idea to go against your framework's opinions.)
- *the [microservice][wp-microservice]/[monolith][wp-monolith] distinction*: Most of my services happen to be
  microservices because of how [my team][sapcc]'s general stack looks like. But you can also build a monolith from this
  template. In fact, some of its characteristics are influenced by a monolith that I helped maintain for some years.

## Data model

I have some very tiny services where I can get away with storing all data in [object storage][wp-objectstorage], so no
database deployment is needed. Since we have an [OpenStack][openstack] installation anyway, [OpenStack Swift][os-swift]
is the obvious choice.

Even if I have a traditional database, large blobs still get offloaded into object storage, and the database will only
hold a reference to the object. This keeps the database footprint nice and small, which speeds up both regular accesses
as well as maintenance operations like backups and vacuums.

When I do need to store structured data, my database of choice is [PostgreSQL][postgres]. I do not have a very strong
reason to prefer PostgreSQL over other RDBMS like [MariaDB][mariadb], but there is a lot to like about it: It has decent
tooling. It has a phenomenal documentation. And most importantly, it is just rock solid: I trust it to never lose any
data and to always correctly enforce the constraints that I put in the schema. I do have some notes on PostgreSQL,
though, and they will be discussed in the "Limitations" section below.

Finally, and this one should be a no-brainer, but it does feel worth mentioning: Whenever my services write into
PostgreSQL, all writes belonging to one change are wrapped in a transaction. Besides helping to avoid accidental
inconsistencies, this also simplifies my deployment model: Most components can just be restarted whenever I want,
without any grace periods or controlled-shutdown procedures. Any transactions that were in flight will just be dropped
by the database, and when the component comes back up, it will just pick up where it left off. It is true that this
causes some duplicate work since the new process needs to reproduce the progress of the killed process, but this is
entirely outweighed by the reduced implementation complexity.

## Component model

There are up to three components that are run and scaled separately: APIs, an observer, and optionally workers.

### API component

As the name suggests, this component answers all requests coming in from outside. For my own services, the API is always
[REST-like][wp-rest], so we are looking at an HTTP API handling GET/POST/PUT/DELETE requests with JSON payloads. But in
principle, the type of protocol and payload does not matter. Care is taken to ensure that concurrent requests can be
parallelized (with the database and object storage acting as synchronization points), so the API component can easily be
scaled horizontally to fit the actual request load.

Unlike the other components, the API process needs some extra logic for controlled shutdown: If a shutdown would just
terminate the process right away, any unanswered API requests would have their respective connections closed, resulting
in unexpected errors on the client side. To avoid such errors, we use a two-step graceful shutdown procedure:

1. When SIGINT or SIGTERM is sent to the API process, the signal handler will stall for 10 seconds before actually
   initiating the shutdown. This gives the loadbalancer in front of the API ample time to drop the process from its
   backend pool, so no new requests will be sent to it.

2. After this first grace period, the API process will cease listening on its socket, and then wait up to 30 seconds for
   all remaining request handlers to complete. (Note that this step usually takes much less time since we will already
   have not received new requests for several seconds.)

### Observer component

- self-healing job loops: reconciliation of outside facts with our internal truth
- objects have DB columns like `next_validation_at` and `validation_error_message`, earlier approach `validated_at` has drawbacks in query construction and schedule flexibility
- not horizontally scaled, simplifies implementation; but it is a choke point, therefore:

### Worker component

- horizontally scalable
- relies on `FOR UPDATE SKIP LOCKED`; apparently supported in MariaDB with InnoDB since 2021 <https://mariadb.com/kb/en/changes-improvements-in-mariadb-106/>

## Runtime environment

- Docker (Alpine)
  - small images replicate faster
- Kubernetes
  - API and workers as Deployment with rolling-update
  - observer as Deployment with recreate
  - nginx-ingress in front of API for LB and TLS termination
- logging: container logs -> ELK
  - usually rather silent except for errors
  - request logs, incl. error messages on 5xx's, and sometimes paper trail for background jobs to aid in investigations
- monitoring: Prometheus
  - Prometheus metrics at application level: operation counters (both success and failure)
  - Prometheus metrics at DB level (via postgres-exporter): job loop backlog, number of objects with job error messages

## Limitations

- generally speaking: an architecture suited for "average everything", no consideration for extreme loads
  - current load example: Keppel Frankfurt = 100000 req/hour
- Postgres on Kubernetes with really dumb setup
  - "if it looks stupid and works, it ain't stupid"
  - problematic in case of high SLA or very high load
- 5xx results during DB restart
  - managed Postgres not an option because no cloud and strict data placement policies
- Postgres connection limit
  - run a tight ship on connection budget
  - pgbouncer is an option

[golang]: https://golang.org
[wp-microservice]: https://en.wikipedia.org/wiki/Microservices
[wp-monolith]: https://en.wikipedia.org/wiki/Monolithic_system
[sapcc]: https://github.com/sapcc
[openstack]: https://openstack.org
[wp-objectstorage]: https://en.wikipedia.org/wiki/Object_storage
[os-swift]: https://wiki.openstack.org/wiki/Swift
[postgres]: https://www.postgresql.org/
[mariadb]: https://mariadb.com/
[wp-rest]: https://en.wikipedia.org/wiki/Representational_state_transfer
