# cleango

[![release](https://github.com/thisdougb/cleango/actions/workflows/release.yaml/badge.svg)](https://github.com/thisdougb/cleango/actions/workflows/release.yaml)

#### Goal

A re-usable GoLang template for a web/api server, that saves me time and ensures I start projects with a good structure.

#### Strategy

- a runnable app as a template, so we start from a known-good
- include features most likely to be used, to avoid masses of boiler-plate code
- an easy method to create releases
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
2023/03/14 08:40:01 main.go:32: Datastore connecting, host: 'localhost:6379', username: 
2023/03/14 08:40:01 main.go:39: Datastore connected.
2023/03/14 08:40:01 main.go:55: webserver.Start(): listening on port 8080
```
You can test everything works:

```
$ curl http://localhost:8080   
<!DOCTYPE html>
<html lang="en">
    <head></head>
    <body>
        <h4>Hello World!</h4>
    </body>
</html>
```

and 

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

#### Release Packages

 A GitHub action uses goreleaser to automatically build [release packages](https://github.com/thisdougb/cleango/releases).
 The release includes the README file, and the LICENSE file.

```
$ git checkout main   
Switched to branch 'main'
Your branch is up to date with 'origin/main'.

$ git tag -a v0.1.0 -m "initial release"

$ git push origin v0.1.0                                                                        
Enumerating objects: 1, done.
Counting objects: 100% (1/1), done.
Writing objects: 100% (1/1), 191 bytes | 191.00 KiB/s, done.
Total 1 (delta 0), reused 0 (delta 0), pack-reused 0
To github.com:thisdougb/cleango.git
 * [new tag]         v0.1.0 -> v0.1.0
 ```

#### Emedded Files

Using the _embed_ module we can include the static template files in the resulting Go binary.
This means a deployable Go app that uses html templates will work.

On startup the following is printed to help understanding:

```
2023/03/14 08:40:01 env.go:87: dir .
2023/03/14 08:40:01 env.go:87: dir templates
2023/03/14 08:40:01 env.go:90: file: templates/footer.gohtml
2023/03/14 08:40:01 env.go:90: file: templates/header.gohtml
2023/03/14 08:40:01 env.go:90: file: templates/index.gohtml
2023/03/14 08:40:01 env.go:80: embed FS file: templates/footer.gohtml
2023/03/14 08:40:01 env.go:80: embed FS file: templates/header.gohtml
2023/03/14 08:40:01 env.go:80: embed FS file: templates/index.gohtml
```

The _embed_ module can be hard to figure out from the docs.
The key thing is to know the FS contains file paths.

But we use the templates using the template name (string), which is coincidentally the file name.
The template name (string) is used in the define statement within the template file.
These two strings must match, otherwise a blank html page is served.

```
2023/03/14 08:40:02 handle_homepage.go:47: templates: ; defined templates are: "index.gohtml", "header.gohtml", "footer.gohtml"
```

Logging shows us the file and line that prints the above output, so it can be followed for understanding.

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
