package index

type scriptResponse struct {
	Keywords string `json:"keywords"` // list of keywords as a string joined by the "," character
	Image    string `json:"image"`    // poster image URL
}

type TitleIndex struct {
	Props Props `json:"props"`
}

type Props struct {
	PageProps PageProps `json:"pageProps"`
}

type PageProps struct {
	AboveTheFoldData AboveTheFoldData `json:"aboveTheFoldData"`
	MainColumnData   MainColumnData   `json:"mainColumnData"`
}

type AboveTheFoldData struct {
	MeterRanking       MeterRanking         `json:"meterRanking"`
	PrimaryVideos      PrimaryVideos        `json:"primaryVideos"`
	ExternalLinks      Total                `json:"externalLinks"`
	Keywords           Keywords             `json:"keywords"`
	Genres             GenresWrapper        `json:"genres"`
	Plot               Plot                 `json:"plot"`
	Credits            Total                `json:"credits"`
	PrincipalCredits   []PrincipalCredit    `json:"principalCredits"`
	CriticReviewsTotal Total                `json:"criticReviewsTotal"` // total number of external reviews
	Meta               Meta                 `json:"meta"`
	CastPageTitle      CastPageTitle        `json:"castPageTitle"`
	DirectorsPageTitle []DirectorsPageTitle `json:"directorsPageTitle"`
	Certificate        interface{}          `json:"certificate"`
	Metacritic         interface{}          `json:"metacritic"`
	CreatorsPageTitle  []interface{}        `json:"creatorsPageTitle"`
}

type CastPageTitle struct {
	Edges []CastPageTitleEdge `json:"edges"`
}

type CastPageTitleEdge struct {
	Node NodeElement `json:"node"`
}

type NodeElement struct {
	Name PurpleName `json:"name"`
}

type PurpleName struct {
	NameText withText `json:"nameText"`
}

type withText struct {
	Text string `json:"text"`
}

type AboveTheFoldDataCountriesOfOrigin struct {
	Countries []PrimaryImageElement `json:"countries"`
}

type PrimaryImageElement struct {
	ID string `json:"id"`
}

type Total struct {
	Total int64 `json:"total"`
}

type DirectorsPageTitle struct {
	Credits []NodeElement `json:"credits"`
}

type FeaturedReviews struct {
	Edges []GoofsEdge `json:"edges"`
}

type GoofsEdge struct {
	Node PurpleNode `json:"node"`
}

type PurpleNode struct {
	Author         PurpleAuthor `json:"author"`
	Summary        Summary      `json:"summary"`
	Text           PurpleText   `json:"text"`
	AuthorRating   int64        `json:"authorRating"`
	SubmissionDate string       `json:"submissionDate"`
}

type PurpleAuthor struct {
	NickName string `json:"nickName"`
}

type Summary struct {
	OriginalText string `json:"originalText"`
}

type PurpleText struct {
	OriginalText PlotText `json:"originalText"`
}

type PlotText struct {
	PlainText string `json:"plainText"`
}

type GenresWrapper struct {
	Genres []withTextAndID `json:"genres"`
}

type withTextAndID struct {
	Text string `json:"text"`
	ID   string `json:"id"`
}

type Keywords struct {
	Total int64 `json:"total"`
}

type Meta struct {
	CanonicalID       string `json:"canonicalId"`
	PublicationStatus string `json:"publicationStatus"`
}

type MeterRanking struct {
	CurrentRank int64      `json:"currentRank"`
	RankChange  RankChange `json:"rankChange"`
}

type RankChange struct {
	ChangeDirection string `json:"changeDirection"`
	Difference      int64  `json:"difference"`
}

type Plot struct {
	PlotText PlotText `json:"plotText"`
}

type NodeClass struct {
	ID      string    `json:"id"`
	Width   int64     `json:"width"`
	Height  int64     `json:"height"`
	URL     string    `json:"url"`
	Caption *PlotText `json:"caption,omitempty"`
}

type PrimaryVideos struct {
	Edges []PrimaryVideosEdge `json:"edges"`
}

type PrimaryVideosEdge struct {
	Node FluffyNode `json:"node"`
}

type FluffyNode struct {
	ID           string            `json:"id"`
	IsMature     bool              `json:"isMature"`
	ContentType  PurpleContentType `json:"contentType"`
	Thumbnail    Thumbnail         `json:"thumbnail"`
	Runtime      PurpleRuntime     `json:"runtime"`
	Description  Description       `json:"description"`
	Name         Description       `json:"name"`
	PlaybackURLs []URL             `json:"playbackURLs"`
	PreviewURLs  []URL             `json:"previewURLs"`
}

type PurpleContentType struct {
	DisplayName DisplayNameClass `json:"displayName"`
}

type DisplayNameClass struct {
	Value string `json:"value"`
}

type Description struct {
	Value    string `json:"value"`
	Language string `json:"language"`
}

type URL struct {
	DisplayName Description `json:"displayName"`
	MIMEType    string      `json:"mimeType"`
	URL         string      `json:"url"`
}

type PurpleRuntime struct {
	Value int64 `json:"value"`
}

type Thumbnail struct {
	URL    string `json:"url"`
	Height int64  `json:"height"`
	Width  int64  `json:"width"`
}

type PrincipalCredit struct {
	TotalCredits int64         `json:"totalCredits"`
	Category     withTextAndID `json:"category"`
	Credits      []Credit      `json:"credits"`
}

type Credit struct {
	Name       FluffyName  `json:"name"`
	Attributes interface{} `json:"attributes"`
}

type FluffyName struct {
	NameText     withText   `json:"nameText"`
	ID           string     `json:"id"`
	PrimaryImage *Thumbnail `json:"primaryImage"`
}

type Production struct {
	Edges []ProductionEdge `json:"edges"`
}

type ProductionEdge struct {
	Node TentacledNode `json:"node"`
}

type TentacledNode struct {
	Company Company `json:"company"`
}

type Company struct {
	ID          string   `json:"id"`
	CompanyText withText `json:"companyText"`
}

type ProductionStatus struct {
	CurrentProductionStage  withTextAndID             `json:"currentProductionStage"`
	ProductionStatusHistory []ProductionStatusHistory `json:"productionStatusHistory"`
	Restriction             interface{}               `json:"restriction"`
}

type ProductionStatusHistory struct {
	Status withTextAndID `json:"status"`
}

type AboveTheFoldDataRatingsSummary struct {
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

type ReleaseDate struct {
	Day     int64          `json:"day"`
	Month   int64          `json:"month"`
	Year    int64          `json:"year"`
	Country *withTextAndID `json:"country,omitempty"`
}

type AboveTheFoldDataReleaseYear struct {
	Year    int64       `json:"year"`
	EndYear interface{} `json:"endYear"`
}

type AboveTheFoldDataRuntime struct {
	Seconds int64 `json:"seconds"`
}

type TitleType struct {
	Text      string `json:"text"`
	ID        string `json:"id"`
	IsSeries  bool   `json:"isSeries"`
	IsEpisode bool   `json:"isEpisode"`
}

type ImageModel struct {
	URL       string `json:"url"`
	Caption   string `json:"caption"`
	MaxHeight int64  `json:"maxHeight"`
	MaxWidth  int64  `json:"maxWidth"`
}

type MainColumnData struct {
	ID                      string                          `json:"id"`
	WINS                    Total                           `json:"wins"`
	Nominations             Total                           `json:"nominations"`
	RatingsSummary          MainColumnDataRatingsSummary    `json:"ratingsSummary"`
	Videos                  Total                           `json:"videos"`
	TitleMainImages         TitleMainImages                 `json:"titleMainImages"`
	ProductionStatus        ProductionStatus                `json:"productionStatus"`
	PrimaryImage            PrimaryImageElement             `json:"primaryImage"`
	TitleType               PrimaryImageElement             `json:"titleType"`
	CanHaveEpisodes         bool                            `json:"canHaveEpisodes"`
	Cast                    CastClass                       `json:"cast"`
	Directors               []Director                      `json:"directors"`
	IsAdult                 bool                            `json:"isAdult"`
	MoreLikeThisTitles      MoreLikeThisTitles              `json:"moreLikeThisTitles"`
	TriviaTotal             Total                           `json:"triviaTotal"`
	Trivia                  Trivia                          `json:"trivia"`
	GoofsTotal              Total                           `json:"goofsTotal"`
	Goofs                   FeaturedReviews                 `json:"goofs"`
	QuotesTotal             Total                           `json:"quotesTotal"`
	Quotes                  Quotes                          `json:"quotes"`
	CrazyCredits            CastPageTitle                   `json:"crazyCredits"`
	AlternateVersions       Keywords                        `json:"alternateVersions"`
	Connections             Connections                     `json:"connections"`
	Soundtrack              Soundtrack                      `json:"soundtrack"`
	TitleText               withText                        `json:"titleText"`
	OriginalTitleText       withText                        `json:"originalTitleText"`
	ReleaseYear             AssociatedTitleReleaseYear      `json:"releaseYear"`
	Reviews                 Total                           `json:"reviews"` // total number of users reviews
	FeaturedReviews         PurpleFeaturedReviews           `json:"featuredReviews"`
	FaqsTotal               Total                           `json:"faqsTotal"`
	Faqs                    faqEdges                        `json:"faqs"`
	ReleaseDate             ReleaseDate                     `json:"releaseDate"`
	CountriesOfOrigin       MainColumnDataCountriesOfOrigin `json:"countriesOfOrigin"`
	DetailsExternalLinks    DetailsExternalLinks            `json:"detailsExternalLinks"`
	SpokenLanguages         SpokenLanguages                 `json:"spokenLanguages"`
	Akas                    Akas                            `json:"akas"`
	FilmingLocations        Keywords                        `json:"filmingLocations"`
	Production              Production                      `json:"production"`
	Companies               Total                           `json:"companies"`
	Runtime                 AboveTheFoldDataRuntime         `json:"runtime"`
	ProductionBudget        interface{}                     `json:"productionBudget"`
	LifetimeGross           interface{}                     `json:"lifetimeGross"`
	OpeningWeekendGross     interface{}                     `json:"openingWeekendGross"`
	WorldwideGross          interface{}                     `json:"worldwideGross"`
	Series                  interface{}                     `json:"series"`
	PrestigiousAwardSummary interface{}                     `json:"prestigiousAwardSummary"`
	Episodes                interface{}                     `json:"episodes"`
	Creators                []interface{}                   `json:"creators"`
	Writers                 []interface{}                   `json:"writers"`
}

type Akas struct {
	Edges []AKAEdge `json:"edges"`
}

type AKAEdge struct {
	Node withText `json:"node"`
}

type CastClass struct {
	Edges []CastEdge `json:"edges"`
}

type CastEdge struct {
	Node StickyNode `json:"node"`
}

type StickyNode struct {
	Name           FluffyName      `json:"name"`
	Attributes     []withText      `json:"attributes"`
	Characters     []NodeCharacter `json:"characters"`
	EpisodeCredits EpisodeCredits  `json:"episodeCredits"`
}

type NodeCharacter struct {
	Name string `json:"name"`
}

type EpisodeCredits struct {
	Total     int64       `json:"total"`
	YearRange interface{} `json:"yearRange"`
}

type Connections struct {
	Edges []ConnectionsEdge `json:"edges"`
}

type ConnectionsEdge struct {
	Node IndigoNode `json:"node"`
}

type IndigoNode struct {
	AssociatedTitle AssociatedTitle `json:"associatedTitle"`
	Category        withText        `json:"category"`
}

type AssociatedTitle struct {
	ID                string                     `json:"id"`
	ReleaseYear       AssociatedTitleReleaseYear `json:"releaseYear"`
	TitleText         withText                   `json:"titleText"`
	OriginalTitleText withText                   `json:"originalTitleText"`
	Series            AssociatedTitleSeries      `json:"series"`
}

type AssociatedTitleReleaseYear struct {
	Year int64 `json:"year"`
}

type AssociatedTitleSeries struct {
	Series SeriesSeries `json:"series"`
}

type SeriesSeries struct {
	TitleText         withText `json:"titleText"`
	OriginalTitleText withText `json:"originalTitleText"`
}

type MainColumnDataCountriesOfOrigin struct {
	Countries []withTextAndID `json:"countries"`
}

type DetailsExternalLinks struct {
	Edges []DetailsExternalLinksEdge `json:"edges"`
	Total int64                      `json:"total"`
}

type DetailsExternalLinksEdge struct {
	Node HilariousNode `json:"node"`
}

type HilariousNode struct {
	URL                string      `json:"url"`
	Label              string      `json:"label"`
	ExternalLinkRegion interface{} `json:"externalLinkRegion"`
}

type Director struct {
	TotalCredits int64    `json:"totalCredits"`
	Category     withText `json:"category"`
	Credits      []Credit `json:"credits"`
}

type PurpleFeaturedReviews struct {
	Edges []featuredReviewEdge `json:"edges"`
}

type featuredReviewEdge struct {
	Node featuredReviewNode `json:"node"`
}

type featuredReviewNode struct {
	ID             string       `json:"id"`
	Author         FluffyAuthor `json:"author"`
	Summary        Summary      `json:"summary"`
	Text           FluffyText   `json:"text"`
	AuthorRating   int64        `json:"authorRating"`
	SubmissionDate string       `json:"submissionDate"`
	Helpfulness    Helpfulness  `json:"helpfulness"`
}

type FluffyAuthor struct {
	NickName string `json:"nickName"`
	UserID   string `json:"userId"`
}

type Helpfulness struct {
	UpVotes   int64 `json:"upVotes"`
	DownVotes int64 `json:"downVotes"`
}

type FluffyText struct {
	OriginalText TextElement `json:"originalText"`
}

type TextElement struct {
	PlaidHTML string `json:"plaidHtml"`
}

type MoreLikeThisTitles struct {
	Edges []MoreLikeThisTitlesEdge `json:"edges"`
}

type MoreLikeThisTitlesEdge struct {
	Node CunningNode `json:"node"`
}

type CunningNode struct {
	ID                 string                         `json:"id"`
	TitleText          withText                       `json:"titleText"`
	OriginalTitleText  withText                       `json:"originalTitleText"`
	TitleType          withTextAndID                  `json:"titleType"`
	PrimaryImage       NodeClass                      `json:"primaryImage"`
	ReleaseYear        AboveTheFoldDataReleaseYear    `json:"releaseYear"`
	RatingsSummary     AboveTheFoldDataRatingsSummary `json:"ratingsSummary"`
	Runtime            *AboveTheFoldDataRuntime       `json:"runtime"`
	Certificate        *Certificate                   `json:"certificate"`
	TitleCardGenres    TitleCardGenres                `json:"titleCardGenres"`
	CanHaveEpisodes    bool                           `json:"canHaveEpisodes"`
	PrimaryWatchOption *PrimaryWatchOption            `json:"primaryWatchOption"`
}

type Certificate struct {
	Rating string `json:"rating"`
}

type PrimaryWatchOption struct {
	AdditionalWatchOptionsCount int64 `json:"additionalWatchOptionsCount"`
}

type TitleCardGenres struct {
	Genres []withText `json:"genres"`
}

type Quotes struct {
	Edges []QuotesEdge `json:"edges"`
}

type QuotesEdge struct {
	Node MagentaNode `json:"node"`
}

type MagentaNode struct {
	Lines []Line `json:"lines"`
}

type Line struct {
	Characters     []LineCharacter `json:"characters"`
	Text           string          `json:"text"`
	StageDirection interface{}     `json:"stageDirection"`
}

type LineCharacter struct {
	Character string              `json:"character"`
	Name      PrimaryImageElement `json:"name"`
}

type MainColumnDataRatingsSummary struct {
	TopRanking interface{} `json:"topRanking"`
}

type Soundtrack struct {
	Edges []SoundtrackEdge `json:"edges"`
}

type SoundtrackEdge struct {
	Node FriskyNode `json:"node"`
}

type FriskyNode struct {
	Text     string        `json:"text"`
	Comments []TextElement `json:"comments"`
}

type SpokenLanguages struct {
	SpokenLanguages []withTextAndID `json:"spokenLanguages"`
}

type TitleMainImages struct {
	Total int64                 `json:"total"`
	Edges []TitleMainImagesEdge `json:"edges"`
}

type TitleMainImagesEdge struct {
	Node NodeClass `json:"node"`
}

type Trivia struct {
	Edges []TriviaEdge `json:"edges"`
}

type TriviaEdge struct {
	Node MischievousNode `json:"node"`
}

type MischievousNode struct {
	Text         TextElement `json:"text"`
	Trademark    interface{} `json:"trademark"`
	RelatedNames interface{} `json:"relatedNames"`
}

type FluffyContentType struct {
	DisplayName DisplayNameClass `json:"displayName"`
}
