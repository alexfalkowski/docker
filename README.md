# Docker

Common images used for my projects

These are images that are used for [CircleCI](https://circleci.com/). They have docker compose so we can run what is locally setup.

## Go

The golang image is based of `buster-node-browsers-legacy` image.

## Ruby

For ruby we can use the default images that are provided by [CircleCI](https://hub.docker.com/r/circleci/ruby/tags). To be consistent use the `buster-node-browsers-legacy` images.

As an example

```sh
docker pull circleci/ruby:2.7-buster-node-browsers-legacy
```

## Java

The java image is based of `buster-node-browsers-legacy` image.

## Scala

The scala image is based of `buster-node-browsers-legacy` image.

## DB

These images provide databases that we can use while testing.

### HBase

This image builds a [Apache HBase](https://hbase.apache.org/book.html) database.
