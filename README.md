# cleango

[![release](https://github.com/thisdougb/cleango/actions/workflows/release.yaml/badge.svg)](https://github.com/thisdougb/cleango/actions/workflows/release.yaml)

#### Goal

A re-usable GoLang template for a web/api server, that saves me time and ensures I start projects with a good structure.

#### Strategy

- a runnable app as a template, so we start from a known-good
- include features most likely to be used, to avoid masses of boiler-plate code
- focus on making development easier and simpler

Use the button above "Use this template".

#### Get Started

When you template this repo it will contain 'thisdougb/cleango' in the pkg paths.
Here's how to reset those paths, using sed on Mac OS (at least), after you've cloned your new repo.

Substitute your GitHub name for _mygithubname_, and your repo name for _myproject_:
```
$ git clone git@github.com:thisdougb/myproject.git
$ cd myproject
$ find . \( -type d -name .git -prune \) -o -type f -print0 | xargs -0 sed -i '' -e 's/thisdougb/mygithubname/g'
$ find . \( -type d -name .git -prune \) -o -type f -print0 | xargs -0 sed -i '' -e 's/cleango/myproject/g'
```
Then run (assumes a local Redis instance):
```
$ go run -tags dev main.go
2022/07/21 11:33:12 server.go:27: Datastore connecting, host: 'localhost:6379', username:
2022/07/21 11:33:12 server.go:34: Datastore connected.
2022/07/21 11:33:12 server.go:46: webserver.Start(): listening on port 8080
```
You can test everything works:
```
$ curl -X POST http://localhost:8080/thing/enable/ -H "Content-Type: application/json" -d '{"thing_id": 1}'                                                                             
OK

$ redis-cli
127.0.0.1:6379> keys *
1) "app:thing:1:status"
127.0.0.1:6379> get "app:thing:1:status"
"1"
127.0.0.1:6379>
```

#### Logging

In [log.go](https://github.com/thisdougb/cleango/blob/refactor_logging/api/log.go#L7) we have a simple init() which sets the formatting for log statements.
It is easier and quicker to troubleshooting problems when you know where the log statements are from.

Ensuring filenames are descriptive, rather than main.go, helps here:
```
2022/07/21 11:33:12 enablething.go:27: error, ostrich 43723 has 8 legs.
```

#### Build Tags

I use build tags.
All test and mock files are _dev_, so excluded in the final build.

This also makes switching templating easy between environments.
```
$ go run -tags dev api/server.go
2022/07/21 11:33:12 server.go:46: webserver.Start(): listening on port 8080
```

#### Passing Datastore Reference

So [here](https://github.com/thisdougb/cleango/blob/main/api/handlers/env.go) I use an Env struct to reference Service pointers.
This allows seamless passing of the datastore connection (or mock) to the handlers.

I mashed up the Clean Architecture style with [this](https://www.alexedwards.net/blog/organising-database-access) blog post.
That's where my Env struct came from.

#### Use Case

Usecase is a core Clean Architecture idea, and a little vague.
I think of it as, 'an action that happens, like making a coffee.'
There's often multiple steps to produce an outcome.

An http handler depends on a usecase, but the usecase knows nothing about the http handler.

#### File Numbering

I use file name numbering for files that are part of the templating pattern.
This is purely to make scanning dirs and finding what you expect to always be there much quicker.

For example:
```
$ ls -l pkg/usecase/enablething
total 56
-rw-r--r--  1 thisdougb  staff  209 21 Apr 20:13 1_interface.go
-rw-r--r--  1 thisdougb  staff  181 21 Apr 20:14 2_service.go
-rw-r--r--  1 thisdougb  staff  179 21 Apr 19:43 3_mock.go
-rw-r--r--  1 thisdougb  staff   77 21 Apr 19:43 4_mock_reader.go
-rw-r--r--  1 thisdougb  staff  378 21 Apr 20:17 5_mock_writer.go
-rw-r--r--  1 thisdougb  staff  433 21 Apr 20:57 enablething.go
-rw-r--r--  1 thisdougb  staff  899 21 Apr 20:15 enablething_test.go
```
