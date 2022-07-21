package parser

// type schemaOrgData struct {
// 	Keywords string `json:"keywords"` // list of keywords as a string joined by the "," character
// 	Image    string `json:"image"`    // poster image URL
// }

type nextJSData struct {
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
	PrimaryImage  primaryImage  `json:"primaryImage"`
	MeterRanking  meterRanking  `json:"meterRanking"`
	PrimaryVideos primaryVideos `json:"primaryVideos"`
	Genres        genresWrapper `json:"genres"`
	Plot          plotWrapper   `json:"plot"`
	Meta          meta          `json:"meta"`
	// CastPageTitle      castPageTitle        `json:"castPageTitle"`
	// DirectorsPageTitle []directorsPageTitle `json:"directorsPageTitle"`
	Series      series            `json:"series"` // only when viewing an episode of a series
	Certificate certificateRating `json:"certificate"`
	// PrincipalCredits   []principalCredit    `json:"principalCredits"`
	Keywords           keywords     `json:"keywords"`
	ExternalLinks      totalWrapper `json:"externalLinks"`
	Credits            totalWrapper `json:"credits"`
	CriticReviewsTotal totalWrapper `json:"criticReviewsTotal"` // total number of external reviews
	Metacritic         interface{}  `json:"metacritic"`         // TODO: improve typings
}

type certificateRating struct {
	Rating string `json:"rating"`
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

type primaryImage struct {
	ID      string                     `json:"id"`
	Width   int64                      `json:"width"`
	Height  int64                      `json:"height"`
	URL     string                     `json:"url"`
	Caption primaryImageCaptionWrapper `json:"caption"`
}

type primaryImageCaptionWrapper struct {
	PlainText string `json:"plainText"`
}

type IDWrapper struct {
	ID string `json:"id"`
}

type totalWrapper struct {
	Total int64 `json:"total"`
}

type videoStrip struct {
	Edges []videoStripEdge `json:"edges"`
}

type videoStripEdge struct {
	Node videoStripNode `json:"node"`
}

type videoStripNode struct {
	ID          string              `json:"id"`
	ContentType videoContentType    `json:"contentType"`
	Name        stringValueWrapper  `json:"name"`
	Runtime     int64ValueWrapper   `json:"runtime"`
	Thumbnail   videoThumbnailClass `json:"thumbnail"`
}

type videoContentType struct {
	DisplayName stringValueWrapper `json:"displayName"`
}

type stringValueWrapper struct {
	Value string `json:"value"`
}

type int64ValueWrapper struct {
	Value int64 `json:"value"`
}

type videoThumbnailClass struct {
	Height int64  `json:"height"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
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

type keywords struct {
	Total int64         `json:"total"`
	Edges []keywordEdge `json:"edges"`
}

type keywordEdge struct {
	Node keywordNode `json:"node"`
}

type keywordNode struct {
	Text string `json:"text"`
}

type meta struct {
	CanonicalID       string `json:"canonicalId"`       // tt0123456
	PublicationStatus string `json:"publicationStatus"` // PUBLISHED
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
	ID      string          `json:"id"`
	Width   int64           `json:"width"`
	Height  int64           `json:"height"`
	URL     string          `json:"url"`
	Caption plotTextWrapper `json:"caption,omitempty"`
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

type creditNameWrapper struct {
	Name creditName `json:"name"`
}

type creditName struct {
	ID           string         `json:"id"`
	NameText     withText       `json:"nameText"`
	PrimaryImage titleThumbnail `json:"primaryImage"`
}

type titleProduction struct {
	Edges []productionEdge `json:"edges"`
}

type productionEdge struct {
	Node productionCompanyNode `json:"node"`
}

type productionCompanyNode struct {
	Company titleCompany `json:"company"`
}

type titleCompany struct {
	ID          string   `json:"id"`
	CompanyText withText `json:"companyText"`
}

type productionStatus struct {
	CurrentProductionStage  withTextAndID             `json:"currentProductionStage"`
	ProductionStatusHistory []productionStatusHistory `json:"productionStatusHistory"`
	Restriction             interface{}               `json:"restriction"` // TODO: improve typing
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
	Day     int64         `json:"day"`
	Month   int64         `json:"month"`
	Year    int64         `json:"year"`
	Country withTextAndID `json:"country,omitempty"`
}

type aboveTheFoldDataReleaseYear struct {
	Year    int64 `json:"year"`
	EndYear int64 `json:"endYear"`
}

type secondsWrapper struct {
	Seconds int64 `json:"seconds"`
}

type mainColumnData struct {
	ID                      string                          `json:"id"`
	TitleText               withText                        `json:"titleText"`
	OriginalTitleText       withText                        `json:"originalTitleText"`
	CanHaveEpisodes         bool                            `json:"canHaveEpisodes"`
	IsAdult                 bool                            `json:"isAdult"`
	TitleMainImages         mainImages                      `json:"titleMainImages"`
	VideoStrip              videoStrip                      `json:"videoStrip"`
	ProductionStatus        productionStatus                `json:"productionStatus"`
	TitleType               IDWrapper                       `json:"titleType"`
	Cast                    castEdgeWrapper                 `json:"cast"`
	MoreLikeThisTitles      moreLikeThisTitles              `json:"moreLikeThisTitles"`
	CrazyCredits            castPageTitle                   `json:"crazyCredits"`
	Soundtrack              titleSoundtrack                 `json:"soundtrack"`
	ReleaseYear             associatedTitleReleaseYear      `json:"releaseYear"`
	FeaturedReviews         featuredReviews                 `json:"featuredReviews"`
	Faqs                    faqEdges                        `json:"faqs"`
	ReleaseDate             releaseDate                     `json:"releaseDate"`
	CountriesOfOrigin       mainColumnDataCountriesOfOrigin `json:"countriesOfOrigin"`
	DetailsExternalLinks    detailsExternalLinks            `json:"detailsExternalLinks"`
	SpokenLanguages         spokenLanguagesWrapper          `json:"spokenLanguages"`
	AKAs                    akasWrapper                     `json:"akas"`
	Runtime                 secondsWrapper                  `json:"runtime"`
	LifetimeGross           titleLifetimeGross              `json:"lifetimeGross"`       // Boxoffice stats
	OpeningWeekendGross     titleOpeningWeekendGross        `json:"openingWeekendGross"` // Boxoffice stats
	WorldwideGross          titleWorldwideGross             `json:"worldwideGross"`      // Boxoffice stats
	Connections             connections                     `json:"connections"`
	PrestigiousAwardSummary prestigiousAwardSummary         `json:"prestigiousAwardSummary"`
	RatingsSummary          titleRatingsSummary             `json:"ratingsSummary"`
	Episodes                titleEpisodesWrapper            `json:"episodes"` // only when viewing the parent series ID
	Wins                    totalWrapper                    `json:"wins"`
	Nominations             totalWrapper                    `json:"nominations"`
	Videos                  totalWrapper                    `json:"videos"`
	Reviews                 totalWrapper                    `json:"reviews"` // total number of users reviews
	FaqsTotal               totalWrapper                    `json:"faqsTotal"`
	TriviaTotal             totalWrapper                    `json:"triviaTotal"`
	Trivia                  titleTrivia                     `json:"trivia"`
	GoofsTotal              totalWrapper                    `json:"goofsTotal"`
	Goofs                   titleGoofs                      `json:"goofs"`
	QuotesTotal             totalWrapper                    `json:"quotesTotal"`
	Quotes                  quotes                          `json:"quotes"`
	AlternateVersions       titleAlternateVersions          `json:"alternateVersions"`
	FilmingLocations        titleFilmingLocations           `json:"filmingLocations"`
	Companies               totalWrapper                    `json:"companies"`  // Companies total number
	Production              titleProduction                 `json:"production"` // Companies list
	Creators                []titleCreator                  `json:"creators"`
	Directors               []titleDirector                 `json:"directors"`
	TechnicalSpecifications technicalSpecifications         `json:"technicalSpecifications"`
	Writers                 []interface{}                   `json:"writers"` // TODO: better typing (it's probably same as "creators", and "directors")
	ProductionBudget        interface{}                     `json:"productionBudget"`
}

type technicalSpecifications struct {
	SoundMixes   soundMixes   `json:"soundMixes"`
	AspectRatios aspectRatios `json:"aspectRatios"`
	Colorations  colorations  `json:"colorations"`
}

type aspectRatios struct {
	Items []aspectRatiosItem `json:"items"`
}

type aspectRatiosItem struct {
	AspectRatio string             `json:"aspectRatio"`
	Attributes  []textValueWrapper `json:"attributes"`
}

type colorations struct {
	Items []colorationsItem `json:"items"`
}

type colorationsItem struct {
	ConceptID string `json:"conceptId"`
	Text      string `json:"text"`
}

type soundMixes struct {
	Items []soundMixesItem `json:"items"`
}

type soundMixesItem struct {
	ID         string        `json:"id"`
	Text       string        `json:"text"`
	Attributes []interface{} `json:"attributes"`
}

type titleFilmingLocations struct {
	Total int64                      `json:"total"`
	Edges []titleFilmingLocationEdge `json:"edges"`
}

type titleFilmingLocationEdge struct {
	Node textValueWrapper `json:"node"`
}

type titleAlternateVersions struct {
	Total int64                       `json:"total"`
	Edges []titleAlternateVersionEdge `json:"edges"`
}

type titleAlternateVersionEdge struct {
	Node titleAlternateVersionNode `json:"node"`
}

type titleAlternateVersionNode struct {
	Text plaidHTMLWrapper `json:"text"`
}

type titleLifetimeGross struct {
	Total lifetimeGrossTotal `json:"total"`
}

type lifetimeGrossTotal struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

type titleOpeningWeekendGross struct {
	Gross          titleGross `json:"gross"`
	WeekendEndDate string     `json:"weekendEndDate"` // "" "2019-11-24"
}

type titleGross struct {
	Total grossTotal `json:"total"`
}

type grossTotal struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

type titleWorldwideGross struct {
	Total worldwideGrossTotal `json:"total"`
}

type worldwideGrossTotal struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

type series struct {
	EpisodeNumber   episodeNumber `json:"episodeNumber"`
	NextEpisode     nextEpisode   `json:"nextEpisode"`
	PreviousEpisode nextEpisode   `json:"previousEpisode"`
	Series          seriesClass   `json:"series"`
}

type prestigiousAwardSummary struct {
	Nominations int64        `json:"nominations"` // Total nominations
	Wins        int64        `json:"wins"`        // Total wins
	Award       awardSummary `json:"award"`
}

type awardSummary struct {
	ID    string         `json:"id"` // an0123456
	Text  string         `json:"text"`
	Event eventIDWrapper `json:"event"`
}

type eventIDWrapper struct {
	ID string `json:"id"` // ev0123456
}

type titleCreator struct {
	TotalCredits int64           `json:"totalCredits"`
	Credits      []creatorCredit `json:"credits"`
}

type creatorCredit struct {
	Name creatorNameWrapper `json:"name"`
}

type creatorNameWrapper struct {
	ID       string           `json:"id"`
	NameText textValueWrapper `json:"nameText"`
}

type textValueWrapper struct {
	Text string `json:"text"`
}

type titleEpisodesWrapper struct {
	Episodes      totalWrapper       `json:"episodes"`
	TotalEpisodes totalWrapper       `json:"totalEpisodes"`
	Seasons       []seasonNumWrapper `json:"seasons"`
	Years         []yearWrapper      `json:"years"`
	TopRated      topRated           `json:"topRated"`
}

type seasonNumWrapper struct {
	Number int64 `json:"number"`
}

type topRated struct {
	Edges []ratingEdge `json:"edges"`
}

type ratingEdge struct {
	Node ratingNode `json:"node"`
}

type ratingNode struct {
	RatingsSummary ratingsSummary `json:"ratingsSummary"`
}

type ratingsSummary struct {
	AggregateRating float64 `json:"aggregateRating"`
}

type yearWrapper struct {
	Year int64 `json:"year"`
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

type castEdgeWrapper struct {
	Edges []castNodeWrapper `json:"edges"`
}

type castNodeWrapper struct {
	Node castNode `json:"node"`
}

type castNode struct {
	Name           castName           `json:"name"`
	Characters     []castCharacter    `json:"characters"`
	EpisodeCredits castEpisodeCredits `json:"episodeCredits"`
}

type castName struct {
	ID           string              `json:"id"`
	NameText     castNameTextWrapper `json:"nameText"`
	PrimaryImage castPrimaryImage    `json:"primaryImage"`
}

type castNameTextWrapper struct {
	Text string `json:"text"`
}

type castPrimaryImage struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

type castCharacter struct {
	Name string `json:"name"`
}

type castEpisodeCredits struct {
	Total     int64                      `json:"total"`
	YearRange castEpisodeCreditYearRange `json:"yearRange"`
}

type castEpisodeCreditYearRange struct {
	Year    int64 `json:"year"`
	EndYear int64 `json:"endYear"`
}

type connections struct {
	Edges []connectionsEdge `json:"edges"`
}

type connectionsEdge struct {
	Node connectionNode `json:"node"`
}

type connectionNode struct {
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
	URL                string           `json:"url"`
	Label              string           `json:"label"`
	ExternalLinkRegion textValueWrapper `json:"externalLinkRegion"`
}

type titleDirector struct {
	TotalCredits int64               `json:"totalCredits"`
	Credits      []creditNameWrapper `json:"credits"`
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
	OriginalText plaidHTMLWrapper `json:"originalText"`
}

type plaidHTMLWrapper struct {
	PlaidHTML string `json:"plaidHtml"`
}

type moreLikeThisTitles struct {
	Edges []moreLikeThisTitlesEdge `json:"edges"`
}

type moreLikeThisTitlesEdge struct {
	Node relatedTitleNode `json:"node"`
}

type relatedTitleNode struct {
	ID                string                         `json:"id"`
	TitleText         withText                       `json:"titleText"`
	OriginalTitleText withText                       `json:"originalTitleText"`
	TitleType         withTextAndID                  `json:"titleType"`
	PrimaryImage      imageNode                      `json:"primaryImage"`
	ReleaseYear       aboveTheFoldDataReleaseYear    `json:"releaseYear"`
	RatingsSummary    aboveTheFoldDataRatingsSummary `json:"ratingsSummary"`
	Runtime           secondsWrapper                 `json:"runtime"`
	Certificate       certificate                    `json:"certificate"`
	TitleCardGenres   titleCardGenres                `json:"titleCardGenres"`
	CanHaveEpisodes   bool                           `json:"canHaveEpisodes"`
}

type certificate struct {
	Rating string `json:"rating"`
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
	StageDirection interface{}     `json:"stageDirection"` // TODO: better type
}

type lineCharacter struct {
	Character string    `json:"character"`
	Name      IDWrapper `json:"name"`
}

type titleRatingsSummary struct {
	TopRanking ratingSummaryTopRaking `json:"topRanking"`
}

type ratingSummaryTopRaking struct {
	ID   string             `json:"id"` // topRatedTv:tt0108778:en_US
	Text stringValueWrapper `json:"text"`
	Rank int64              `json:"rank"`
}

type titleSoundtrack struct {
	Edges []soundtrackEdge `json:"edges"`
}

type soundtrackEdge struct {
	Node friskyNode `json:"node"`
}

type friskyNode struct {
	Text     string             `json:"text"`
	Comments []plaidHTMLWrapper `json:"comments"`
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
	Text         plaidHTMLWrapper `json:"text"`
	Trademark    interface{}      `json:"trademark"`    // TODO: better type
	RelatedNames interface{}      `json:"relatedNames"` // TODO: better type
}
