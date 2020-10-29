# Docker

Common images used for my projects

## CI

These are images that are used for [CircleCI](https://circleci.com/). They have docker compose so we can run what is locally setup.

### Go

The golang image is based of `buster-node-browsers-legacy` image. We add ruby as we like testing with [nonnative](https://github.com/alexfalkowski/nonnative)

### Ruby

For ruby we can use the default images that are provided by [CircleCI](https://hub.docker.com/r/circleci/ruby/tags). To be consistent use the `stretch-node-browsers-legacy` images.

As an example

```sh
docker pull circleci/ruby:2.7-buster-node-browsers-legacy
```
