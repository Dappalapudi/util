
# env package

All services should rely and respect environment variable ENV. No exceptions!

Example:

    ENV=dev ./accounts

Should start accounts service pointing to the databases setup in Development environment.

**env** package abstracts ENV environment variable so all services can use calls to **env** package to discover the environment they are running in.

Services during start up **must** check environment they are running using `env.Is()` call and then open database connections and make other steps required for
particular environment.

Let say there is *events* service that uses Redis for pub/sub. Start it with: 

    ENV=stg ./events

As long as you are on VPN connection to access Staging Redis this will start events service. That service will use Staging Redis instance to start listening on incoming events.

## Benefits

The big benefit of standard environment is that local execution is possible without *docker-compose* and other container tech.

Your instance of the service will point to the same databases that QA Analysts/Testers have access to. If they will have some test data setup with failing tests, you should be able to replicate that scenario
on your local instance of the service because databases will be the same.

