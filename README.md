# cleango

[![release](https://github.com/thisdougb/cleango/actions/workflows/release.yaml/badge.svg)](https://github.com/thisdougb/cleango/actions/workflows/release.yaml)

A template Go module, making it easy to start projects with consistent structure.
Use the button above "Use this template" to get your Go project off to a clean start.

* Leans on clean architecture
* Implements tests
* Build tags
* CI via GitHub Actions workflow

### Motivation
My aim with this repo is to give myself a template to copy, each time I start a new project in Go.

Underlying this I have, over the years, been living with various languages to try and find one that's fast to prototype whilst also being type safe.
I'm liking Go, and this template is an attempt to create a pattern for structuring Go projects loose-ishly based on Clean Architecture.

[Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) has lofty aims.
There aren't many resources out there that lay it out in Go.
I don't either, because Go doesn't map directly, or my projects don't fit neatly.

But I am interested in architecture that makes testing easy, so I persisted.
As a nice bonus I found it easy and fast to build APIs, because the architecture made it so.

So take this as a pragmatic approach to Clean Architecture using Go.

### Notes
Some notes on the code, or style, or implications.
I'm not a great Go coder, I'm still learning.
These notes are aimed at other learners.

Feedback is very much hoped for, to make this template and my Go-skills better.

#### Reset Paths
When you template this repo it will contain 'thisdougb/cleango' in the pkg paths.
Here's how to reset those paths, using sed on Mac OS (at least), after you've cloned your new repo.

Substitute your GitHub name for mygithubname, and your repo name for myproject:
```
$ git clone git@github.com:thisdougb/myproject.git
$ find myproject \( -type d -name .git -prune \) -o -type f -print0 | xargs -0 sed -i '' -e 's/thisdougb/mygithubname/g'
$ find myproject \( -type d -name .git -prune \) -o -type f -print0 | xargs -0 sed -i '' -e 's/cleango/myproject/g'
```

#### Build Tags
I love build tags.
This is a big win when using Go, because it's easy to switch config.

```
$ go run -tags dev api/server.go
2021/04/21 22:03:52 webserver.Start(): listening on port 8080
```

Here is an [example](https://github.com/thisdougb/cleango/blob/204df73075f69d8ff3fff555f1b739f40c060d3a/config/dev_config.go#L1) that says include this file when -tags is dev or test.
And in the GitHub action, [here](https://github.com/thisdougb/cleango/blob/204df73075f69d8ff3fff555f1b739f40c060d3a/.github/workflows/branches.yaml#L43) I run test with that tag.
Pulling in that particular file.

This is good because if you're test want to run quicker, you can set low limits for various things.

#### Passing Database Connection
This was a big challenge, as it is in most languages.
I found a solution which made a lot of sense, and fits right in with Clean Architecture.
It also makes testing and mocks really easy.

There are a lot of frameworks for http handling.
I dislike frameworks because they add dependencies, which often add more complexity that you need.
So [here](https://github.com/thisdougb/cleango/blob/main/api/handlers/env.go) I use an Env struct to reference Service pointers.
This allows seamless passing of the datastore connection (or mock) to the handlers.

I mashed up the Clean Architecture style with [this](https://www.alexedwards.net/blog/organising-database-access) blog post.
That's where my Env struct came from.

#### Use Case
Usecase is a core Clean Architecture idea, and a little vague.
I think of it as, 'an action that happens, like making a coffee.'
There's often multiple steps to produce an outcome.
An http handler depends on a usecase, but the usecase knows nothing about the http handler.

The usecase is also where I implement my datastore mocks.
Mock come after Interface.

#### Interface
The [interface](https://github.com/thisdougb/cleango/blob/main/pkg/usecase/enablething/1_interface.go) for a usecase is imported by the datastore, don't repeat yourself.
Interfaces in Go make code simpler and more robust, so learn interfaces.

[This](https://github.com/thisdougb/cleango/blob/2e28d75fb42b6559c34dab7fd86ac69aaacbeb8e/pkg/datastore/interface.go#L12) is where the datastore interface references the usecase interface.

#### Mocks
Mocks, I love mocks.
But mocking datastores and methods somehow always seems confusing.

To be clear, I'm mocking the datastore methods.
When my usecase makes a call to a datastore method, when testing it's my mocks that are injected.

[Here](https://github.com/thisdougb/cleango/blob/main/pkg/usecase/enablething/5_mock_writer.go) is the mock method, it mocks the call to the datastore method.
I have conditionals here, to simulate responses from the real method [here](https://github.com/thisdougb/cleango/blob/971877d70fe85886b42d81e1025da26a6b7978c4/pkg/datastore/redis/thing.go#L7).

And [here's](https://github.com/thisdougb/cleango/blob/971877d70fe85886b42d81e1025da26a6b7978c4/api/handlers/enablething_test.go#L24) another thing.
From a higher level method, I'm using data that triggers the mock to respond in a particular way for testing.
This might not be pure, but it's very handy and cuts down on a lot of mocking code.

I spent a lot of time figuring out mocking, in various languages.
Using Go, Clean Architecture, Env struct, etc, are where I'm currently at.
And it works well like this, without too much duplicate testing code.

Wider point.
Tests are not for 'production code', they are for all code.
Tests make your development quicker overall.
For years I thought tests were a hassle getting in the way of writing code.
If you're still thinking that, read [The Tortoise and the Hare](https://en.wikipedia.org/wiki/The_Tortoise_and_the_Hare).

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
-rw-r--r--  1 thisdougb  staff  433 21 Apr 20:57 main.go
-rw-r--r--  1 thisdougb  staff  899 21 Apr 20:15 main_test.go
```
