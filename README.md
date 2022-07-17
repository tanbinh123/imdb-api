# IMDb-API

[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/Scrip7/imdb-api)
[![tests](https://github.com/Scrip7/imdb-api/actions/workflows/tests.yml/badge.svg)](https://github.com/Scrip7/imdb-api/actions/workflows/tests.yml)
[![CodeFactor](https://www.codefactor.io/repository/github/scrip7/imdb-api/badge)](https://www.codefactor.io/repository/github/scrip7/imdb-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/Scrip7/imdb-api)](https://goreportcard.com/report/github.com/Scrip7/imdb-api)
[![License](https://img.shields.io/github/license/Scrip7/imdb-api?color=orange)](https://github.com/Scrip7/imdb-api/blob/main/LICENSE)
[![Contribute](https://img.shields.io/badge/PRs-welcome-blue.svg?color=d9ecde)](https://github.com/Scrip7/imdb-api/pulls)

<a href="https://github.com/Scrip7/imdb-api">
  <img align="right" src="https://ia.media-imdb.com/images/M/MV5BMTk3ODA4Mjc0NF5BMl5BcG5nXkFtZTgwNDc1MzQ2OTE@._V1_.png" width="110" />
  <img align="right" src="https://go.dev/images/go-logo-blue.svg" width="140" />
</a>

A cross-platform [Go](https://go.dev) microservice application to scrap the [IMDb](https://imdb.com) website.

It's supposed to run as an internal process in your backend architecture beside the other applications you're trying to develop.

> **Warning**
>
> Keep in mind this app does **not** have a rate-limiting middleware or authentication/authorization guards to protect itself against online spam threads.
>
> The public end-users should not have direct access to its interface unless you know what you're doing.

## Features

- [ ] Proxy support
- [ ] Swagger documentation
- [ ] Docker Hub image
- [ ] In-memory response caching
- [ ] Full resolution poster images
- [ ] [Slug](https://en.wikipedia.org/wiki/Slug) generation
- [ ] Extract videos (trailers)
- [ ] Casts list

## Routes coverage

In any stage of development, Pull Requests are more than welcome. 😊

> ✅ Done
>
> 🚧 Under development
>
> ❌ Failed to implement _(PRs are more than welcome)_

| URL                                                 | Title                                    | Status |
| --------------------------------------------------- | ---------------------------------------- | :----: |
| `https://imdb.com/title/:id`                        | Title Single                             |   🚧   |
| `https://imdb.com/search/title`                     | Advanced Search                          |   🚧   |
| `https://imdb.com/search/name`                      | Advanced Name Search                     |   🚧   |
| `https://imdb.com/search/common`                    | Collaborations Search                    |   🚧   |
| `https://imdb.com/news/top`                         | Top News                                 |   🚧   |
| `https://imdb.com/news/movie`                       | Movie News                               |   🚧   |
| `https://imdb.com/news/tv`                          | TV News                                  |   🚧   |
| `https://imdb.com/news/celebrity`                   | Celebrity News                           |   🚧   |
| `https://imdb.com/news/indie`                       | Indie News                               |   🚧   |
| `https://imdb.com/calendar`                         | Upcoming Releases                        |   🚧   |
| `https://imdb.com/chart/boxoffice`                  | Top Box Office (US)                      |   🚧   |
| `https://imdb.com/chart/moviemeter`                 | Most Popular Movies                      |   🚧   |
| `https://imdb.com/chart/top`                        | Top 250 Movies                           |   🚧   |
| `https://imdb.com/chart/top-english-movies`         | Top Rated English Movies                 |   🚧   |
| `https://imdb.com/chart/toptv`                      | Top Rated TV Shows                       |   🚧   |
| `https://imdb.com/chart/bottom`                     | Lowest Rated Movies                      |   🚧   |
| `https://imdb.com/india/upcoming`                   | Most Anticipated New Indian Movies/Shows |   🚧   |
| `https://imdb.com/india/top-rated-indian-movies`    | Top Rated Indian Movies                  |   🚧   |
| `https://imdb.com/india/top-rated-malayalam-movies` | Top Rated Malayalam Movies               |   🚧   |
| `https://imdb.com/india/top-rated-tamil-movies`     | Top Rated Tamil Movies                   |   🚧   |
| `https://imdb.com/india/top-rated-telugu-movies`    | Top Rated Telugu Movies                  |   🚧   |
| `https://imdb.com/india/tamil`                      | Trending Tamil Movies                    |   🚧   |
| `https://imdb.com/india/telugu`                     | Trending Telugu Movies                   |   🚧   |
| `https://imdb.com/event/all`                        | All Events Alphabetically                |   🚧   |

Is a route missing? Feel free to [open a new issue](https://github.com/Scrip7/imdb-api/issues) to let us know!

> **Note**
>
> To prevent duplication, please search to see if an issue exists for your request.

## Docker usage

Placeholder.

## Contributing

Pull Requests are always welcome.

## Disclaimer

> The "[IMDb-API](https://github.com/Scrip7/imdb-api)" repository is not affiliated, authorized, maintained, or endorsed by the [IMDb](https://en.wikipedia.org/wiki/IMDb) or any of its affiliates or subsidiaries (including [Amazon](<https://en.wikipedia.org/wiki/Amazon_(company)>)).
>
> This unofficial independent project is published as it is.
> 
> Therefore, use it at your own risk and respect the copyrights of their content.

## License

See the [`LICENSE`](https://github.com/Scrip7/imdb-api/blob/main/LICENSE) file for more information.
