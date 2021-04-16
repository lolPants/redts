# 🌠 rEDTS
> Remote [EDTS](https://bitbucket.org/Esvandiary/edts) API and CLI

## ❓ What is rEDTS
rEDTS, or Remote EDTS, is a system for running an EDTS instance on a remote server. This has the advantage of only requiring one database for many users, and outsourcing CPU time to an external system.

## 📡 Server
The server runs EDTS in the background with a simple Node.js HTTP API wrapper. It is deployed using the provided [Docker Images](https://github.com/lolPants/redts/packages/733073).

## 📻 Client
The client CLI forwards your CLI args to the remote server and proxies the response back to your terminal, as if you were running EDTS locally. You can find binaries as [GitHub Actions Artifacts](https://github.com/lolPants/redts/actions/workflows/cli.yml).
