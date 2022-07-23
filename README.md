# IMDb-API

[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)][1]
[![tests](https://github.com/Scrip7/imdb-api/actions/workflows/tests.yml/badge.svg)][2]
[![CodeFactor](https://www.codefactor.io/repository/github/scrip7/imdb-api/badge)][3]
[![Go Report Card](https://goreportcard.com/badge/github.com/Scrip7/imdb-api)][4]
[![License](https://img.shields.io/github/license/Scrip7/imdb-api?color=orange)][5]
[![Contribute](https://img.shields.io/badge/PRs-welcome-blue.svg?color=d9ecde)][6]

<a href="https://imdb.com" target="_blank">
  <img align="right" src="https://ia.media-imdb.com/images/M/MV5BMTk3ODA4Mjc0NF5BMl5BcG5nXkFtZTgwNDc1MzQ2OTE@._V1_.png" width="110" />
</a>
<a href="https://go.dev" target="_blank">
  <img align="right" src="https://go.dev/images/go-logo-blue.svg" width="140" />
</a>

A cross-platform [Go][7] microservice to scrape the [IMDb][8] website.

> **Warning**  
> Keep in mind this app does **not** have rate-limiting or authentication features to protect itself against spams.  
> The end-users should **not** have direct access to its interface unless you understand what you're doing.

## Features

- [ ] Proxy support
- [x] Swagger documentation
- [x] [Slug][9] generator
- [x] In-memory caching
- [ ] CLI tool
- [ ] Docker Hub image

## Installation

You can always download the latest binary version from the [releases page][15].

If you're going to build the project on your own, here are the pre-requirements:

- [Git][16]
- [Go][17]

Once you have installed and configured them _(if necessary)_, execute the following commands:

```bash
git clone https://github.com/Scrip7/imdb-api.git
cd imdb-api
cp .env.example .env
go run main.go
```

## Coverage

We have listed the IMDb website paths that can potentially use in our application to collect data.  
Because this is a reverse-engineered project, things might break unexpectedly!  
Therefore, we keep this list updated along with the new features we introduce or when we encounter new bugs to visualize our extensive coverage.

> ✅ Done.
>
> 🚧 Under development.
>
> 💭 The route is known, but we haven't thoroughly thought about how to implement it yet!
>
> 😱 It broke! _(IMDb changed its response structure, it needs minor changes or re-implementation)_
>
> ❌ Failed to implement.

| URL _(prefix: `imdb.com`)_          | Scope  | Second Scope | Module                            | Status | Tests |
| ----------------------------------- | ------ | ------------ | --------------------------------- | :----: | :---: |
| `/title/:id`                        | Title  | -            | Index                             |   🚧   |   -   |
| `/title/:id/mediaindex`             | Title  | -            | Photo Gallery                     |   💭   |   -   |
| `/title/:id/videogallery`           | Title  | -            | Video Gallery                     |   💭   |   -   |
| `/title/:id/news`                   | Title  | -            | News                              |   💭   |   -   |
| `/title/:id/fullcredits`            | Title  | Details      | Full Cast & Crew                  |   💭   |   -   |
| `/title/:id/releaseinfo`            | Title  | Details      | Release Info (Dates)              |   💭   |   -   |
| `/title/:id/externalsites`          | Title  | Details      | External Sites                    |   💭   |   -   |
| `/title/:id/companycredits`         | Title  | Details      | Company Credits                   |   💭   |   -   |
| `/title/:id/locations`              | Title  | Details      | Filming & Production              |   💭   |   -   |
| `/title/:id/technical`              | Title  | Details      | Technical Specifications          |   💭   |   -   |
| `/title/:id/taglines`               | Title  | Story        | Taglines                          |   💭   |   -   |
| `/title/:id/plotsummary`            | Title  | Story        | Plot Summary                      |   💭   |   -   |
| `/title/:id/keywords`               | Title  | Story        | Plot Keywords                     |   💭   |   -   |
| `/title/:id/parentalguide`          | Title  | Story        | Parents Guide                     |   💭   |   -   |
| `/title/:id/trivia`                 | Title  | Inform       | Trivia                            |   💭   |   -   |
| `/title/:id/trgoofsivia`            | Title  | Inform       | Goofs                             |   💭   |   -   |
| `/title/:id/crazycredits`           | Title  | Inform       | Crazy Credits                     |   💭   |   -   |
| `/title/:id/quotes`                 | Title  | Inform       | Quotes                            |   💭   |   -   |
| `/title/:id/alternateversions`      | Title  | Inform       | Alternate Versions                |   💭   |   -   |
| `/title/:id/movieconnections`       | Title  | Inform       | Connections                       |   💭   |   -   |
| `/title/:id/soundtrack`             | Title  | Inform       | Soundtracks                       |   💭   |   -   |
| `/title/:id/awards`                 | Title  | Review       | Awards                            |   💭   |   -   |
| `/title/:id/faq`                    | Title  | Review       | FAQ                               |   💭   |   -   |
| `/title/:id/reviews`                | Title  | Review       | User Reviews                      |   💭   |   -   |
| `/title/:id/ratings`                | Title  | Review       | User Ratings                      |   💭   |   -   |
| `/title/:id/externalreviews`        | Title  | Review       | External Reviews                  |   💭   |   -   |
| `/title/:id/episodes`               | Title  | TV           | Episode List                      |   💭   |   -   |
| `/title/:id/tvschedule`             | Title  | TV           | Schedule                          |   💭   |   -   |
| `/search/title`                     | Search | -            | Title                             |   💭   |   -   |
| `/search/name`                      | Search | -            | Name                              |   💭   |   -   |
| `/search/common`                    | Search | -            | Collaborations                    |   💭   |   -   |
| `/chart/boxoffice`                  | Chart  | -            | Box Office                        |   💭   |   -   |
| `/chart/moviemeter`                 | Chart  | -            | Most Popular Movies               |   💭   |   -   |
| `/chart/top`                        | Chart  | -            | Top 250 Movies                    |   💭   |   -   |
| `/chart/top-english-movies`         | Chart  | -            | Top Rated English Movies          |   💭   |   -   |
| `/chart/toptv`                      | Chart  | -            | Top Rated TV Shows                |   💭   |   -   |
| `/chart/bottom`                     | Chart  | -            | Lowest Rated Movies               |   💭   |   -   |
| `/news/top`                         | News   | -            | Top                               |   💭   |   -   |
| `/news/movie`                       | News   | -            | Movie                             |   💭   |   -   |
| `/news/tv`                          | News   | -            | TV                                |   💭   |   -   |
| `/news/celebrity`                   | News   | -            | Celebrity                         |   💭   |   -   |
| `/news/indie`                       | News   | -            | Indie                             |   💭   |   -   |
| `/india/upcoming`                   | Indian | -            | Most Anticipated New Movies/Shows |   💭   |   -   |
| `/india/top-rated-indian-movies`    | Indian | -            | Top Rated Movies                  |   💭   |   -   |
| `/india/top-rated-malayalam-movies` | Indian | -            | Top Rated Malayalam Movies        |   💭   |   -   |
| `/india/top-rated-tamil-movies`     | Indian | -            | Top Rated Tamil Movies            |   💭   |   -   |
| `/india/top-rated-telugu-movies`    | Indian | -            | Top Rated Telugu Movies           |   💭   |   -   |
| `/india/tamil`                      | Indian | -            | Trending Tamil Movies             |   💭   |   -   |
| `/india/telugu`                     | Indian | -            | Trending Telugu Movies            |   💭   |   -   |
| `/event/all`                        | Event  | -            | All Events Alphabetically         |   💭   |   -   |
| `/calendar`                         | -      | -            | Upcoming Releases                 |   💭   |   -   |

Is a route missing? Feel free to [open a new issue][10] to let us know!

> **Note** Please check that the issue has not already been reported to prevent duplication.

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

## FAQ

**Can I contribute to make IMDb-API better?**

[Please do!][11] We are looking for any kind of contribution to improve IMDb-API's core functionality and documentation. When in doubt, make a PR!

## Disclaimer

> The "[IMDb-API][12]" repository is not affiliated, authorized, maintained, or endorsed by the [IMDb][13] or any of its affiliates or subsidiaries (including [Amazon][14]).  
> This unofficial independent project is published as it is.  
> Therefore, use it at your own risk and respect the copyrights of their content.

## License

See the [`LICENSE`][5] file for more information.

[1]:  https://pkg.go.dev/github.com/Scrip7/imdb-api
[2]:  https://github.com/Scrip7/imdb-api/actions/workflows/tests.yml
[3]:  https://codefactor.io/repository/github/scrip7/imdb-api
[4]:  https://goreportcard.com/report/github.com/Scrip7/imdb-api
[5]:  https://github.com/Scrip7/imdb-api/blob/main/LICENSE
[6]:  https://github.com/Scrip7/imdb-api/pulls
[7]:  https://go.dev
[8]:  https://imdb.com
[9]:  https://en.wikipedia.org/wiki/Clean_URL#Slug
[10]: https://github.com/Scrip7/imdb-api/issues
[11]: https://github.com/Scrip7/imdb-api/blob/main/CONTRIBUTING.md
[12]: https://github.com/Scrip7/imdb-api
[13]: https://en.wikipedia.org/wiki/IMDb
[14]: <https://en.wikipedia.org/wiki/Amazon_(company)>
[15]: https://github.com/Scrip7/imdb-api/releases
[16]: https://linode.com/docs/guides/how-to-install-git-on-linux-mac-and-windows
[17]: https://go.dev/doc/install
