# Xedule Proxy

This repo contains the source code for the Xedule auth server and the proxy back-end. The purpose of this repo is to simplify calls to Xedule and to exctract domain logic to a central place.

The backend checks the auth server periodically over http for new cookies. The auth provider checks Xedule periodically.

Puppeteer is a lot easier in Node compared to Go that's why its split up as a different codebase.