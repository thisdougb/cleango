## Config

#### Goal

A simple way to _apply_ config settings to a Go app, when run locally, and in Docker/K8s, etc.

#### Strategy

- Use shell env vars to apply settings, this gives greatest flexibility and simplicity (locally easy, K8s secrets easy, Dockerfile easy, etc).
- Return reasonable defaults, when a setting is un-configured.
- Simple and readable usage in code, at the point of use.

#### Usage
Update the defaultValues map with a setting you want to use, giving a reasonable default value.

In code we can retrieve a config setting like this:
```
import (
  "github.com/thisdougb/cleango/config"
)

func main() {

  var cfg *config.Config // dynamic config settings

  print(cfg.ValueAsStr("REDIS_HOST"))  // string at the point of use
  print(cfg.ValueAsInt("MAX_THREADS")) // int at the point of use
  print(cfg.ValueAsBool("DEBUG_ON"))   // bool at the point of use
}
```

Note that we do not use the _envVarPrefix_ in code, which keeps the app code portable. Handling the type casting is also done for us in config, making the in-code method calls simpler and more readable.

An interesting side note is that settings are read dynamically, so we can change them in the environment without restarting the app.

#### Shell Environment Variables

To provide scope, when running multiple apps, the env vars should be prefixed with _envVarPrefix_.

```
$ env | grep MYAPP_
MYAPP_USERNAME=dev
MYAPP_PASSWORD=r2d28C3P0
```

#### Kubernetes Secrets

In the K8s manifest, create an env var from a secret like this:
```
env:
  - name: MYAPP_USERNAME
    valueFrom:
      secretKeyRef:
        name: myapp_credentials
        key: username
  - name: MYAPP_PASSWORD
    valueFrom:
      secretKeyRef:
        name: myapp_credentials
        key: password
```
And create the secret itself:
```
$ kubectl create secret generic myapp_credentials --from-literal=username=d3xg45sdf35 --from-literal=password=asdneqw234asck
```

#### Template Paths

Literal paths can be problematic between a dev env and production.

```
GetTemplatePath(fileName string) string
```

[This](https://github.com/thisdougb/cleango/blob/main/config/dev_config.go#L9) method is included via Go build tags (dev or !dev), and gives a simple way to set a file path between environments.  I use this to handle gohtml template files.

In your IDE/editor you need to ensure that your save/test actions are called with _-tags dev_.
