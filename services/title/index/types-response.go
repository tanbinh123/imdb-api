package index

type scriptResponse struct {
	Keywords string `json:"keywords"` // list of keywords as a string joined by the "," character
	Image    string `json:"image"`    // poster image URL
}

type titleIndex struct {
	Props props `json:"props"`
}

type props struct {
	PageProps pageProps `json:"pageProps"`
}

type pageProps struct {
	AboveTheFoldData aboveTheFoldData `json:"aboveTheFoldData"`
	MainColumnData   mainColumnData   `json:"mainColumnData"`
}

type aboveTheFoldData struct {
	MeterRanking       meterRanking         `json:"meterRanking"`
	PrimaryVideos      primaryVideos        `json:"primaryVideos"`
	ExternalLinks      totalWrapper         `json:"externalLinks"`
	Keywords           titleKeywords        `json:"keywords"`
	Genres             genresWrapper        `json:"genres"`
	Plot               plotWrapper          `json:"plot"`
	Credits            totalWrapper         `json:"credits"`
	PrincipalCredits   []principalCredit    `json:"principalCredits"`
	CriticReviewsTotal totalWrapper         `json:"criticReviewsTotal"` // total number of external reviews
	Meta               meta                 `json:"meta"`
	CastPageTitle      castPageTitle        `json:"castPageTitle"`
	DirectorsPageTitle []directorsPageTitle `json:"directorsPageTitle"`
	Certificate        interface{}          `json:"certificate"`
	Metacritic         interface{}          `json:"metacritic"`
	CreatorsPageTitle  []interface{}        `json:"creatorsPageTitle"`
	Series             titleSeries          `json:"series"`
}

type castPageTitle struct {
	Edges []castPageTitleEdge `json:"edges"`
}

type castPageTitleEdge struct {
	Node nodeWithNameText `json:"node"`
}

type nodeWithNameText struct {
	Name nameTextWrapper `json:"name"`
}

type nameTextWrapper struct {
	NameText withText `json:"nameText"`
}

type withText struct {
	Text string `json:"text"`
}

type IDWrapper struct {
	ID string `json:"id"`
}

type totalWrapper struct {
	Total int64 `json:"total"`
}

type directorsPageTitle struct {
	Credits []nodeWithNameText `json:"credits"`
}

type titleGoofs struct {
	Edges []goofsEdge `json:"edges"`
}

type goofsEdge struct {
	Node goofNode `json:"node"`
}

type goofNode struct {
	Author         nicknameWrapper `json:"author"`
	Summary        summary         `json:"summary"`
	Text           textWrapper     `json:"text"`
	AuthorRating   int64           `json:"authorRating"`
	SubmissionDate string          `json:"submissionDate"`
}

type nicknameWrapper struct {
	NickName string `json:"nickName"`
}

type summary struct {
	OriginalText string `json:"originalText"`
}

type textWrapper struct {
	OriginalText plotTextWrapper `json:"originalText"`
}

type plotTextWrapper struct {
	PlainText string `json:"plainText"`
}

type genresWrapper struct {
	Genres []withTextAndID `json:"genres"`
}

type withTextAndID struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type titleKeywords struct {
	Total int64 `json:"total"`
}

type meta struct {
	CanonicalID       string `json:"canonicalId"`
	PublicationStatus string `json:"publicationStatus"`
}

type meterRanking struct {
	CurrentRank int64      `json:"currentRank"`
	RankChange  rankChange `json:"rankChange"`
}

type rankChange struct {
	ChangeDirection string `json:"changeDirection"`
	Difference      int64  `json:"difference"`
}

type plotWrapper struct {
	PlotText plotTextWrapper `json:"plotText"`
}

type imageNode struct {
	ID      string           `json:"id"`
	Width   int64            `json:"width"`
	Height  int64            `json:"height"`
	URL     string           `json:"url"`
	Caption *plotTextWrapper `json:"caption,omitempty"`
}

type primaryVideos struct {
	Edges []primaryVideosEdge `json:"edges"`
}

type primaryVideosEdge struct {
	Node fluffyNode `json:"node"`
}

type fluffyNode struct {
	ID           string               `json:"id"`
	IsMature     bool                 `json:"isMature"`
	ContentType  PurpleContentType    `json:"contentType"`
	Thumbnail    titleThumbnail       `json:"thumbnail"`
	Runtime      runtimeWrapper       `json:"runtime"`
	Description  withValueAndLanguage `json:"description"`
	Name         withValueAndLanguage `json:"name"`
	PlaybackURLs []urlWrapper         `json:"playbackURLs"`
	PreviewURLs  []urlWrapper         `json:"previewURLs"`
}

type PurpleContentType struct {
	DisplayName DisplayNameClass `json:"displayName"`
}

type DisplayNameClass struct {
	Value string `json:"value"`
}

type withValueAndLanguage struct {
	Value    string `json:"value"`
	Language string `json:"language"`
}

type urlWrapper struct {
	DisplayName withValueAndLanguage `json:"displayName"`
	MIMEType    string               `json:"mimeType"`
	URL         string               `json:"url"`
}

type runtimeWrapper struct {
	Value int64 `json:"value"`
}

type titleThumbnail struct {
	URL    string `json:"url"`
	Height int64  `json:"height"`
	Width  int64  `json:"width"`
}

type principalCredit struct {
	TotalCredits int64         `json:"totalCredits"`
	Category     withTextAndID `json:"category"`
	Credits      []credit      `json:"credits"`
}

type credit struct {
	Name       fluffyName  `json:"name"`
	Attributes interface{} `json:"attributes"`
}

type fluffyName struct {
	ID           string          `json:"id"`
	NameText     withText        `json:"nameText"`
	PrimaryImage *titleThumbnail `json:"primaryImage"`
}

type titleProduction struct {
	Edges []productionEdge `json:"edges"`
}

type productionEdge struct {
	Node companyWrapper `json:"node"`
}

type companyWrapper struct {
	Company company `json:"company"`
}

type company struct {
	ID          string   `json:"id"`
	CompanyText withText `json:"companyText"`
}

type productionStatus struct {
	CurrentProductionStage  withTextAndID             `json:"currentProductionStage"`
	ProductionStatusHistory []productionStatusHistory `json:"productionStatusHistory"`
	Restriction             interface{}               `json:"restriction"`
}

type productionStatusHistory struct {
	Status withTextAndID `json:"status"`
}

type aboveTheFoldDataRatingsSummary struct {
	AggregateRating float64 `json:"aggregateRating"`
	VoteCount       int64   `json:"voteCount"`
}

type faqEdges struct {
	Edges []faqEdge `json:"edges"`
}

type faqEdge struct {
	Node faqNode `json:"node"`
}

type faqNode struct {
	ID       string      `json:"id"`
	Question faqQuestion `json:"question"`
}

type faqQuestion struct {
	PlainText string `json:"plainText"`
}

type releaseDate struct {
	Day     int64          `json:"day"`
	Month   int64          `json:"month"`
	Year    int64          `json:"year"`
	Country *withTextAndID `json:"country,omitempty"`
}

type aboveTheFoldDataReleaseYear struct {
	Year    int64       `json:"year"`
	EndYear interface{} `json:"endYear"`
}

type aboveTheFoldDataRuntime struct {
	Seconds int64 `json:"seconds"`
}

type mainColumnData struct {
	ID                      string                          `json:"id"`
	WINS                    totalWrapper                    `json:"wins"`
	Nominations             totalWrapper                    `json:"nominations"`
	RatingsSummary          mainColumnDataRatingsSummary    `json:"ratingsSummary"`
	Videos                  totalWrapper                    `json:"videos"`
	TitleMainImages         mainImages                      `json:"titleMainImages"`
	ProductionStatus        productionStatus                `json:"productionStatus"`
	PrimaryImage            IDWrapper                       `json:"primaryImage"`
	TitleType               IDWrapper                       `json:"titleType"`
	CanHaveEpisodes         bool                            `json:"canHaveEpisodes"`
	Cast                    castClass                       `json:"cast"`
	Directors               []director                      `json:"directors"`
	IsAdult                 bool                            `json:"isAdult"`
	MoreLikeThisTitles      moreLikeThisTitles              `json:"moreLikeThisTitles"`
	TriviaTotal             totalWrapper                    `json:"triviaTotal"`
	Trivia                  titleTrivia                     `json:"trivia"`
	GoofsTotal              totalWrapper                    `json:"goofsTotal"`
	Goofs                   titleGoofs                      `json:"goofs"`
	QuotesTotal             totalWrapper                    `json:"quotesTotal"`
	Quotes                  quotes                          `json:"quotes"`
	CrazyCredits            castPageTitle                   `json:"crazyCredits"`
	AlternateVersions       titleKeywords                   `json:"alternateVersions"`
	Connections             connections                     `json:"connections"`
	Soundtrack              soundtrack                      `json:"soundtrack"`
	TitleText               withText                        `json:"titleText"`
	OriginalTitleText       withText                        `json:"originalTitleText"`
	ReleaseYear             associatedTitleReleaseYear      `json:"releaseYear"`
	Reviews                 totalWrapper                    `json:"reviews"` // total number of users reviews
	FeaturedReviews         featuredReviews                 `json:"featuredReviews"`
	FaqsTotal               totalWrapper                    `json:"faqsTotal"`
	Faqs                    faqEdges                        `json:"faqs"`
	ReleaseDate             releaseDate                     `json:"releaseDate"`
	CountriesOfOrigin       mainColumnDataCountriesOfOrigin `json:"countriesOfOrigin"`
	DetailsExternalLinks    detailsExternalLinks            `json:"detailsExternalLinks"`
	SpokenLanguages         spokenLanguagesWrapper          `json:"spokenLanguages"`
	Akas                    akasWrapper                     `json:"akas"`
	FilmingLocations        titleKeywords                   `json:"filmingLocations"`
	Production              titleProduction                 `json:"production"`
	Companies               totalWrapper                    `json:"companies"`
	Runtime                 aboveTheFoldDataRuntime         `json:"runtime"`
	ProductionBudget        interface{}                     `json:"productionBudget"`
	LifetimeGross           interface{}                     `json:"lifetimeGross"`
	OpeningWeekendGross     interface{}                     `json:"openingWeekendGross"`
	WorldwideGross          interface{}                     `json:"worldwideGross"`
	PrestigiousAwardSummary interface{}                     `json:"prestigiousAwardSummary"`
	Episodes                interface{}                     `json:"episodes"`
	Creators                []interface{}                   `json:"creators"`
	Writers                 []interface{}                   `json:"writers"`
}

type titleSeries struct {
	EpisodeNumber   episodeNumber `json:"episodeNumber"`
	NextEpisode     nextEpisode   `json:"nextEpisode"`
	PreviousEpisode nextEpisode   `json:"previousEpisode"`
	Series          seriesClass   `json:"series"`
}

type episodeNumber struct {
	EpisodeNumber int64 `json:"episodeNumber"`
	SeasonNumber  int64 `json:"seasonNumber"`
}

type nextEpisode struct {
	ID string `json:"id"`
}

type seriesClass struct {
	ID                string           `json:"id"`
	TitleText         titleText        `json:"titleText"`
	OriginalTitleText titleText        `json:"originalTitleText"`
	TitleType         nextEpisode      `json:"titleType"`
	ReleaseYear       titleReleaseYear `json:"releaseYear"`
}

type titleText struct {
	Text string `json:"text"`
}

type titleReleaseYear struct {
	Year    int64 `json:"year"`
	EndYear int64 `json:"endYear"`
}

type akasWrapper struct {
	Edges []akaEdge `json:"edges"`
}

type akaEdge struct {
	Node withText `json:"node"`
}

type castClass struct {
	Edges []castEdge `json:"edges"`
}

type castEdge struct {
	Node stickyNode `json:"node"`
}

type stickyNode struct {
	Name           fluffyName      `json:"name"`
	Attributes     []withText      `json:"attributes"`
	Characters     []nodeCharacter `json:"characters"`
	EpisodeCredits episodeCredits  `json:"episodeCredits"`
}

type nodeCharacter struct {
	Name string `json:"name"`
}

type episodeCredits struct {
	Total     int64       `json:"total"`
	YearRange interface{} `json:"yearRange"`
}

type connections struct {
	Edges []connectionsEdge `json:"edges"`
}

type connectionsEdge struct {
	Node indigoNode `json:"node"`
}

type indigoNode struct {
	AssociatedTitle associatedTitle `json:"associatedTitle"`
	Category        withText        `json:"category"`
}

type associatedTitle struct {
	ID                string                     `json:"id"`
	ReleaseYear       associatedTitleReleaseYear `json:"releaseYear"`
	TitleText         withText                   `json:"titleText"`
	OriginalTitleText withText                   `json:"originalTitleText"`
	Series            associatedTitleSeries      `json:"series"`
}

type associatedTitleReleaseYear struct {
	Year int64 `json:"year"`
}

type associatedTitleSeries struct {
	Series seriesSeries `json:"series"`
}

type seriesSeries struct {
	TitleText         withText `json:"titleText"`
	OriginalTitleText withText `json:"originalTitleText"`
}

type mainColumnDataCountriesOfOrigin struct {
	Countries []withTextAndID `json:"countries"`
}

type detailsExternalLinks struct {
	Edges []detailsExternalLinksEdge `json:"edges"`
	Total int64                      `json:"total"`
}

type detailsExternalLinksEdge struct {
	Node hilariousNode `json:"node"`
}

type hilariousNode struct {
	URL                string      `json:"url"`
	Label              string      `json:"label"`
	ExternalLinkRegion interface{} `json:"externalLinkRegion"`
}

type director struct {
	TotalCredits int64    `json:"totalCredits"`
	Category     withText `json:"category"`
	Credits      []credit `json:"credits"`
}

type featuredReviews struct {
	Edges []featuredReviewEdge `json:"edges"`
}

type featuredReviewEdge struct {
	Node featuredReviewNode `json:"node"`
}

type featuredReviewNode struct {
	ID             string       `json:"id"`
	Author         fluffyAuthor `json:"author"`
	Summary        summary      `json:"summary"`
	Text           fluffyText   `json:"text"`
	AuthorRating   int64        `json:"authorRating"`
	SubmissionDate string       `json:"submissionDate"`
	Helpfulness    helpfulness  `json:"helpfulness"`
}

type fluffyAuthor struct {
	NickName string `json:"nickName"`
	UserID   string `json:"userId"`
}

type helpfulness struct {
	UpVotes   int64 `json:"upVotes"`
	DownVotes int64 `json:"downVotes"`
}

type fluffyText struct {
	OriginalText textElement `json:"originalText"`
}

type textElement struct {
	PlaidHTML string `json:"plaidHtml"`
}

type moreLikeThisTitles struct {
	Edges []moreLikeThisTitlesEdge `json:"edges"`
}

type moreLikeThisTitlesEdge struct {
	Node cunningNode `json:"node"`
}

type cunningNode struct {
	ID                 string                         `json:"id"`
	TitleText          withText                       `json:"titleText"`
	OriginalTitleText  withText                       `json:"originalTitleText"`
	TitleType          withTextAndID                  `json:"titleType"`
	PrimaryImage       imageNode                      `json:"primaryImage"`
	ReleaseYear        aboveTheFoldDataReleaseYear    `json:"releaseYear"`
	RatingsSummary     aboveTheFoldDataRatingsSummary `json:"ratingsSummary"`
	Runtime            *aboveTheFoldDataRuntime       `json:"runtime"`
	Certificate        *certificate                   `json:"certificate"`
	TitleCardGenres    titleCardGenres                `json:"titleCardGenres"`
	CanHaveEpisodes    bool                           `json:"canHaveEpisodes"`
	PrimaryWatchOption *primaryWatchOption            `json:"primaryWatchOption"`
}

type certificate struct {
	Rating string `json:"rating"`
}

type primaryWatchOption struct {
	AdditionalWatchOptionsCount int64 `json:"additionalWatchOptionsCount"`
}

type titleCardGenres struct {
	Genres []withText `json:"genres"`
}

type quotes struct {
	Edges []quotesEdge `json:"edges"`
}

type quotesEdge struct {
	Node magentaNode `json:"node"`
}

type magentaNode struct {
	Lines []line `json:"lines"`
}

type line struct {
	Characters     []lineCharacter `json:"characters"`
	Text           string          `json:"text"`
	StageDirection interface{}     `json:"stageDirection"`
}

type lineCharacter struct {
	Character string    `json:"character"`
	Name      IDWrapper `json:"name"`
}

type mainColumnDataRatingsSummary struct {
	TopRanking interface{} `json:"topRanking"`
}

type soundtrack struct {
	Edges []soundtrackEdge `json:"edges"`
}

type soundtrackEdge struct {
	Node friskyNode `json:"node"`
}

type friskyNode struct {
	Text     string        `json:"text"`
	Comments []textElement `json:"comments"`
}

type spokenLanguagesWrapper struct {
	SpokenLanguages []withTextAndID `json:"spokenLanguages"`
}

type mainImages struct {
	Total int64       `json:"total"`
	Edges []imageEdge `json:"edges"`
}

type imageEdge struct {
	Node imageNode `json:"node"`
}

type titleTrivia struct {
	Edges []triviaEdge `json:"edges"`
}

type triviaEdge struct {
	Node triviaNode `json:"node"`
}

type triviaNode struct {
	Text textElement `json:"text"`
}
