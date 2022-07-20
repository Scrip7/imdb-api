package index

type schemaOrgData struct {
	Keywords string `json:"keywords"` // list of keywords as a string joined by the "," character
	Image    string `json:"image"`    // poster image URL
}

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
	PrimaryImage       titlePrimaryImage    `json:"primaryImage"`
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
	Series             titleSeries          `json:"series"` // only when viewing an episode of a series
	Certificate        titleCertificate     `json:"certificate"`
	// TODO: improve typings
	Metacritic interface{} `json:"metacritic"`
}

type titleCertificate struct {
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

type titlePrimaryImage struct {
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
	ID           string         `json:"id"`
	NameText     withText       `json:"nameText"`
	PrimaryImage titleThumbnail `json:"primaryImage"`
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
	Day     int64         `json:"day"`
	Month   int64         `json:"month"`
	Year    int64         `json:"year"`
	Country withTextAndID `json:"country,omitempty"`
}

type aboveTheFoldDataReleaseYear struct {
	Year    int64 `json:"year"`
	EndYear int64 `json:"endYear"`
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
	VideoStrip              videoStrip                      `json:"videoStrip"`
	TitleMainImages         mainImages                      `json:"titleMainImages"`
	ProductionStatus        productionStatus                `json:"productionStatus"`
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
	Soundtrack              titleSoundtrack                 `json:"soundtrack"`
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
	AKAs                    akasWrapper                     `json:"akas"`
	FilmingLocations        titleKeywords                   `json:"filmingLocations"`
	Production              titleProduction                 `json:"production"`
	Companies               totalWrapper                    `json:"companies"`
	Runtime                 aboveTheFoldDataRuntime         `json:"runtime"`
	LifetimeGross           titleLifetimeGross              `json:"lifetimeGross"`
	OpeningWeekendGross     titleOpeningWeekendGross        `json:"openingWeekendGross"`
	WorldwideGross          titleWorldwideGross             `json:"worldwideGross"`
	PrestigiousAwardSummary prestigiousAwardSummary         `json:"prestigiousAwardSummary"`
	Creators                []titleCreator                  `json:"creators"`
	Episodes                titleEpisodesWrapper            `json:"episodes"` // only when viewing the parent series ID
	// TODO: improve typings
	Writers          []interface{} `json:"writers"`
	ProductionBudget interface{}   `json:"productionBudget"`
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
	WeekendEndDate string     `json:"weekendEndDate"`
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

type titleSeries struct {
	EpisodeNumber   episodeNumber `json:"episodeNumber"`
	NextEpisode     nextEpisode   `json:"nextEpisode"`
	PreviousEpisode nextEpisode   `json:"previousEpisode"`
	Series          seriesClass   `json:"series"`
}

type prestigiousAwardSummary struct {
	Award       awardSummary `json:"award"`
	Nominations int64        `json:"nominations"`
	WINS        int64        `json:"wins"`
}

type awardSummary struct {
	Event eventIDWrapper `json:"event"`
	ID    string         `json:"id"`
	Text  string         `json:"text"`
}

type eventIDWrapper struct {
	ID string `json:"id"`
}

type titleCreator struct {
	Category     creatorCategory `json:"category"`
	Credits      []creatorCredit `json:"credits"`
	TotalCredits int64           `json:"totalCredits"`
}

type creatorCategory struct {
	Text string `json:"text"`
}

type creatorCredit struct {
	Name       nameWrapper `json:"name"`
	Attributes interface{} `json:"attributes"`
}

type nameWrapper struct {
	ID       string          `json:"id"`
	NameText creatorCategory `json:"nameText"`
}

type titleEpisodesWrapper struct {
	Episodes      titleEpisodes      `json:"episodes"`
	Seasons       []seasonNumWrapper `json:"seasons"`
	TopRated      topRated           `json:"topRated"`
	TotalEpisodes titleEpisodes      `json:"totalEpisodes"`
	Years         []yearWrapper      `json:"years"`
}

type titleEpisodes struct {
	Total int64 `json:"total"`
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
	OriginalText plaidHTMLWrapper `json:"originalText"`
}

type plaidHTMLWrapper struct {
	PlaidHTML string `json:"plaidHtml"`
}

type moreLikeThisTitles struct {
	Edges []moreLikeThisTitlesEdge `json:"edges"`
}

type moreLikeThisTitlesEdge struct {
	Node cunningNode `json:"node"`
}

type cunningNode struct {
	ID                string                         `json:"id"`
	TitleText         withText                       `json:"titleText"`
	OriginalTitleText withText                       `json:"originalTitleText"`
	TitleType         withTextAndID                  `json:"titleType"`
	PrimaryImage      imageNode                      `json:"primaryImage"`
	ReleaseYear       aboveTheFoldDataReleaseYear    `json:"releaseYear"`
	RatingsSummary    aboveTheFoldDataRatingsSummary `json:"ratingsSummary"`
	Runtime           aboveTheFoldDataRuntime        `json:"runtime"`
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
	StageDirection interface{}     `json:"stageDirection"`
}

type lineCharacter struct {
	Character string    `json:"character"`
	Name      IDWrapper `json:"name"`
}

type mainColumnDataRatingsSummary struct {
	TopRanking interface{} `json:"topRanking"`
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
	Text plaidHTMLWrapper `json:"text"`
}
