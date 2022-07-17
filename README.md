# IMDB-API

[![tests](https://github.com/Scrip7/imdb-api/actions/workflows/tests.yml/badge.svg)](https://github.com/Scrip7/imdb-api/actions/workflows/tests.yml)

A cross-platform [Go](https://go.dev) microservice application to scrap the [IMDB](https://imdb.com) website.

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

> ✅ = Done
>
> 🚧 = Under development
>
> ⚠️ = Failed to implement _(PRs are more than welcome)_
>
> ❌ = Disagreed to implement

| URL                                              | Description              | Status |
| ------------------------------------------------ | ------------------------ | :----: |
| `https://imdb.com/title/:id`                     | Title Single             |   🚧   |
| `https://imdb.com/chart/boxoffice`               | Top Box Office (US)      |   🚧   |
| `https://imdb.com/chart/moviemeter`              | Most Popular Movies      |   🚧   |
| `https://imdb.com/chart/top`                     | Top 250 Movies           |   🚧   |
| `https://imdb.com/chart/top-english-movies`      | Top Rated English Movies |   🚧   |
| `https://imdb.com/chart/toptv`                   | Top Rated TV Shows       |   🚧   |
| `https://imdb.com/chart/bottom`                  | Lowest Rated Movies      |   🚧   |
| `https://imdb.com/india/top-rated-indian-movies` | Top Rated Indian Movies  |   🚧   |

Is a route missing? Feel free to [open a new issue](https://github.com/Scrip7/imdb-api/issues) to let us know!

## Docker usage

Placeholder.

## Contributing

Pull Requests are always welcome.

## Disclaimer

> The "[IMDB-API](https://github.com/Scrip7/imdb-api)" repository is not affiliated, authorized, maintained, or endorsed by the [IMDB](https://en.wikipedia.org/wiki/IMDb) or any of its affiliates or subsidiaries (including [Amazon](<https://en.wikipedia.org/wiki/Amazon_(company)>)).
>
> This unofficial independent project is published as it is. Therefore, use it at your own risk and respect the copyrights of their content.

## License

See the `LICENSE` file for more information.
