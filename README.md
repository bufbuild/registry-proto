# registry-proto

[![Build](https://github.com/bufbuild/registry-proto/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/bufbuild/registry-proto/actions/workflows/ci.yaml)
[![BSR](https://img.shields.io/badge/BSR-Module-0C65EC)](https://buf.build/bufbuild/registry)
![License](https://img.shields.io/github/license/bufbuild/registry-proto)
[![Slack](https://img.shields.io/badge/Slack-Buf-%23e01563)](https://buf.build/links/slack)

This is the public API for the [Buf Schema Registry](https://buf.build/product/bsr) (BSR),
and is designed to be suitable for use by any client.

## Getting started

- [Browse the module documentation](https://buf.build/bufbuild/registry) to learn how to use the API.
- Use a [Generated SDK](https://buf.build/bufbuild/registry/sdks) for your language,
  or [generated a client manually](https://buf.build/docs/generate/overview) using the `buf` CLI.

## Status: Stable

The [Buf CLI](https://buf.build/blog/buf-cli-next-generation) and
the [BSR web app](https://buf.build/blog/enhanced-buf-push-bsr-ui)
both depend on these APIs as of May 2024.

APIs that are v1 are stable and will not change in backward incompatible ways.
They are safe for any client to depend on in production.

APIs that are alpha or beta may change in backward incompatible ways, so they are unsafe to depend on in production.

## Feedback

Our goal is to make the BSR API a stable and useful primitive for customers to build on.

If you have a question, issue, feature request, or any other feedback about this API, please let us know!
You can:
- [File an issue](https://github.com/bufbuild/registry-proto/issues)
- [Ask a question in Slack](https://buf.build/links/slack)
- Email us at [feedback@buf.build](mailto:feedback@buf.build)

## Legal

Offered under the [Apache 2 license](https://github.com/bufbuild/registry-proto/blob/main/LICENSE)
