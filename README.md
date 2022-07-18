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

> **Warning**
>
> Keep in mind this app does **not** have rate-limiting or authentication features to protect itself against spam threads.
>
> The end-users should **not** have direct access to its interface unless you understand what you're doing.

## Features

- [ ] Proxy support
- [ ] Swagger documentation
- [ ] [Slug](https://en.wikipedia.org/wiki/Slug) generator
- [ ] In-memory caching
- [ ] CLI tool
- [ ] Docker Hub image

## Installation

```
Placeholder.
```

## Routes coverage

In any stage of development, Pull Requests are more than welcome. 😊

> ✅ Done
>
> 🚧 Under development
>
> 💭 The route is known, but haven't fully thought how to implement it, yet!
>
> ❌ Failed to implement

| URL                                         | Scope             | Module                            | Implementation | Tests |
| ------------------------------------------- | ----------------- | --------------------------------- | :------------: | :---: |
| `imdb.com/title/:id`                        | Title             | Index                             |       🚧       |   -   |
| `imdb.com/title/:id/mediaindex`             | Title             | Photo Gallery                     |       💭       |   -   |
| `imdb.com/title/:id/videogallery`           | Title             | Video Gallery                     |       💭       |   -   |
| `imdb.com/title/:id/news`                   | Title             | News                              |       💭       |   -   |
| `imdb.com/title/:id/fullcredits`            | Title → Details   | Full Cast & Crew                  |       💭       |   -   |
| `imdb.com/title/:id/releaseinfo`            | Title → Details   | Release Info (Dates)              |       💭       |   -   |
| `imdb.com/title/:id/externalsites`          | Title → Details   | External Sites                    |       💭       |   -   |
| `imdb.com/title/:id/companycredits`         | Title → Details   | Company Credits                   |       💭       |   -   |
| `imdb.com/title/:id/locations`              | Title → Details   | Filming & Production              |       💭       |   -   |
| `imdb.com/title/:id/technical`              | Title → Details   | Technical Specifications          |       💭       |   -   |
| `imdb.com/title/:id/taglines`               | Title → Storyline | Taglines                          |       💭       |   -   |
| `imdb.com/title/:id/plotsummary`            | Title → Storyline | Plot Summary                      |       💭       |   -   |
| `imdb.com/title/:id/keywords`               | Title → Storyline | Plot Keywords                     |       💭       |   -   |
| `imdb.com/title/:id/parentalguide`          | Title → Storyline | Parents Guide                     |       💭       |   -   |
| `imdb.com/title/:id/trivia`                 | Title → Inform    | Trivia                            |       💭       |   -   |
| `imdb.com/title/:id/trgoofsivia`            | Title → Inform    | Goofs                             |       💭       |   -   |
| `imdb.com/title/:id/crazycredits`           | Title → Inform    | Crazy Credits                     |       💭       |   -   |
| `imdb.com/title/:id/quotes`                 | Title → Inform    | Quotes                            |       💭       |   -   |
| `imdb.com/title/:id/alternateversions`      | Title → Inform    | Alternate Versions                |       💭       |   -   |
| `imdb.com/title/:id/movieconnections`       | Title → Inform    | Connections                       |       💭       |   -   |
| `imdb.com/title/:id/soundtrack`             | Title → Inform    | Soundtracks                       |       💭       |   -   |
| `imdb.com/title/:id/awards`                 | Title → Opinion   | Awards                            |       💭       |   -   |
| `imdb.com/title/:id/faq`                    | Title → Opinion   | FAQ                               |       💭       |   -   |
| `imdb.com/title/:id/reviews`                | Title → Opinion   | User Reviews                      |       💭       |   -   |
| `imdb.com/title/:id/ratings`                | Title → Opinion   | User Ratings                      |       💭       |   -   |
| `imdb.com/title/:id/externalreviews`        | Title → Opinion   | External Reviews                  |       💭       |   -   |
| `imdb.com/title/:id/episodes`               | Title → TV        | Episode List                      |       💭       |   -   |
| `imdb.com/title/:id/tvschedule`             | Title → TV        | Schedule                          |       💭       |   -   |
| `imdb.com/search/title`                     | Search            | Title                             |       💭       |   -   |
| `imdb.com/search/name`                      | Search            | Name                              |       💭       |   -   |
| `imdb.com/search/common`                    | Search            | Collaborations                    |       💭       |   -   |
| `imdb.com/chart/boxoffice`                  | Chart             | Box Office                        |       💭       |   -   |
| `imdb.com/chart/moviemeter`                 | Chart             | Most Popular Movies               |       💭       |   -   |
| `imdb.com/chart/top`                        | Chart             | Top 250 Movies                    |       💭       |   -   |
| `imdb.com/chart/top-english-movies`         | Chart             | Top Rated English Movies          |       💭       |   -   |
| `imdb.com/chart/toptv`                      | Chart             | Top Rated TV Shows                |       💭       |   -   |
| `imdb.com/chart/bottom`                     | Chart             | Lowest Rated Movies               |       💭       |   -   |
| `imdb.com/news/top`                         | News              | Top                               |       💭       |   -   |
| `imdb.com/news/movie`                       | News              | Movie                             |       💭       |   -   |
| `imdb.com/news/tv`                          | News              | TV                                |       💭       |   -   |
| `imdb.com/news/celebrity`                   | News              | Celebrity                         |       💭       |   -   |
| `imdb.com/news/indie`                       | News              | Indie                             |       💭       |   -   |
| `imdb.com/india/upcoming`                   | Indian            | Most Anticipated New Movies/Shows |       💭       |   -   |
| `imdb.com/india/top-rated-indian-movies`    | Indian            | Top Rated Movies                  |       💭       |   -   |
| `imdb.com/india/top-rated-malayalam-movies` | Indian            | Top Rated Malayalam Movies        |       💭       |   -   |
| `imdb.com/india/top-rated-tamil-movies`     | Indian            | Top Rated Tamil Movies            |       💭       |   -   |
| `imdb.com/india/top-rated-telugu-movies`    | Indian            | Top Rated Telugu Movies           |       💭       |   -   |
| `imdb.com/india/tamil`                      | Indian            | Trending Tamil Movies             |       💭       |   -   |
| `imdb.com/india/telugu`                     | Indian            | Trending Telugu Movies            |       💭       |   -   |
| `imdb.com/event/all`                        | Event             | All Events Alphabetically         |       💭       |   -   |
| `imdb.com/calendar`                         | -                 | Upcoming Releases                 |       💭       |   -   |

Is a route missing? Feel free to [open a new issue](https://github.com/Scrip7/imdb-api/issues) to let us know!

> **Note**
>
> To prevent duplication, please search to see if an issue exists for your request.

## Docker usage

Placeholder.

## Motivation

My teacher signed me on a project when I was a junior programmer.

I had to make a Social platform bot (like Telegram or Discord bots) where users could search for Movies and TV Series, look up their detail, create playlists, and share them with others.

After smashing my head onto the keyboard for almost two weeks and a half, I finally found a free IMDb API, which stopped working after a while.

Back in the day, I didn't know how to integrate my application to the IMDb website properly, so sadly, I had to abandon the project.

But now, when I look back at it after many years, I can imagine how hard it could be if someone tries to develop an application sightly or even heavily dependent on the IMDb website.

There are so many reverse-engineered, untrustworthy platforms out there that would take your money in exchange for providing an unstable API.

So I thought of developing a microservice that developers could use in the software they're trying to create.

I didn't profoundly plan how I would create such an application.

But because I was motivated, I created its repository and will slowly work on building small chunks of it, starting from the documentation and overall expectations.

Side note:

I came from having six years of TypeScript experience to Golang, and this project is one of my early projects using this language.

I chose it because it is more performant and efficient and can compile the packages and their dependencies into a single executable binary.

I'm not confident about my coding approaches, maybe I'm stuck with JS strategies in my head, but I'm excited about how it will end up.

Feel free to join my journey by contributing to this project.

## FAQ

**Can I contribute to make IMDb-API better?**

[Please do!](https://github.com/Scrip7/imdb-api/blob/main/CONTRIBUTING.md) We are looking for any kind of contribution to improve IMDb-API's core functionality and documentation. When in doubt, make a PR!

## Disclaimer

> The "[IMDb-API](https://github.com/Scrip7/imdb-api)" repository is not affiliated, authorized, maintained, or endorsed by the [IMDb](https://en.wikipedia.org/wiki/IMDb) or any of its affiliates or subsidiaries (including [Amazon](<https://en.wikipedia.org/wiki/Amazon_(company)>)).
>
> This unofficial independent project is published as it is.
>
> Therefore, use it at your own risk and respect the copyrights of their content.

## License

See the [`LICENSE`](https://github.com/Scrip7/imdb-api/blob/main/LICENSE) file for more information.
