# About

An URL redirection application.

# Quick testing

```
export REDIR_1='{"Host":"localhost.foilen-lab.com","Redirection":"https://google.com","Permanent":false,"AppendQuery":false}'
./create-local-release.sh && build/bin/redirections
```

# Local testing

```
./create-local-release.sh

docker run -ti \
  --rm \
  --env 'REDIR_1={"Host":"localhost.foilen-lab.com","Redirection":"https://google.com","Permanent":false,"AppendQuery":false}' \
  redirections:main-SNAPSHOT
```

# Start it with Docker

```
docker run -ti \
  --rm \
  --env 'REDIR_1={"Host":"localhost.foilen-lab.com","Redirection":"https://google.com","Permanent":false,"AppendQuery":false}' \
  foilen/redirections
```

# Configuration

It is configurable via environment variables. All those starting with "REDIR_" are picked up. The format of the value is a json object with some options:
- host: the host that will be redirecting
- redirection: the full url where to redirect to
- permanent (default: false): to send 301 or 302 to tell if a permanent or temporary redirection
- appendQuery (default: true): to append everything that follows the hostname in the redirection
  - When true: http://localhost.foilen-lab.com/the/path goes to https://google.com/the/path
  - When false: http://localhost.foilen-lab.com/the/path goes to https://google.com

The port it will open is 80 or what is in HTTP_PORT.
