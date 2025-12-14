package service

import (
	"fmt"
	"time"

	gotmdbapi "github.com/ralvarezdev/go-tmdb-api"
	v1 "github.com/ralvarezdev/proto-movies/gen/go/ralvarezdev/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// MapToOptionalFloat64 maps a gotmdbapi optional float32 to a pointer to float64
//
// Parameters:
//
//   - value: the gotmdbapi optional float32 to map
//
// Returns:
//
// - *float64: the mapped pointer to float64
func MapToOptionalFloat64(value *float32) *float64 {
	if value == nil {
		return nil
	}
	mappedValue := float64(*value)
	return &mappedValue
}

// MapDateStringToTimestamp maps a date string to a timestamppb.Timestamp
//
// Parameters:
//
// - dateString: the date string to map
//
// Returns:
//
// - *timestamppb.Timestamp: the mapped timestamppb.Timestamp
func MapDateStringToTimestamp(dateString string) *timestamppb.Timestamp {
	if dateString == "" {
		return nil
	}
	parsedTime, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return nil
	}
	return timestamppb.New(parsedTime)
}

// MapToGender maps an integer to a v1.Gender
//
// Parameters:
//
//   - gender: the integer representing the gender
//
// Returns:
//
// - v1.Gender: the mapped gender
func MapToGender(gender *int32) v1.Gender {
	if gender == nil {
		return v1.Gender_NOT_SET_OR_NOT_SPECIFIED
	}

	switch *gender {
	case int32(v1.Gender_MALE):
		return v1.Gender_MALE
	case int32(v1.Gender_FEMALE):
		return v1.Gender_FEMALE
	case int32(v1.Gender_NON_BINARY):
		return v1.Gender_NON_BINARY
	default:
		return v1.Gender_NOT_SET_OR_NOT_SPECIFIED
	}
}

// MapToCastMember maps a gotmdbapi.Cast to a v1.CastMember
//
// Parameters:
//
//   - castMember: the gotmdbapi.CastMember to map
//
// Returns:
//
// - *v1.CastMember: the mapped v1.CastMember
func MapToCastMember(castMember *gotmdbapi.Cast) *v1.CastMember {
	if castMember == nil {
		return &v1.CastMember{}
	}

	// Map popularity to float64 pointer
	var popularity *float64
	if castMember.Popularity != nil {
		popularity = new(float64)
		*popularity = float64(*castMember.Popularity)
	}

	// Parse profile path from relative to full URL
	var profileURL *string
	if castMember.ProfilePath != nil && *castMember.ProfilePath != "" {
		profileURL = new(string)
		*profileURL = fmt.Sprintf(
			gotmdbapi.ImageVariableQualityURL,
			CastMemberProfileImageWidthSize,
			(*castMember.ProfilePath)[1:],
		)
	}

	return &v1.CastMember{
		Adult:           castMember.Adult,
		Gender:          MapToGender(castMember.Gender),
		Id:              castMember.ID,
		KnownDepartment: castMember.KnownForDepartment,
		Name:            castMember.Name,
		OriginalName:    castMember.OriginalName,
		Popularity:      popularity,
		ProfileUrl:      profileURL,
		CastId:          string(castMember.ID),
		Character:       castMember.Character,
		CreditId:        castMember.CreditID,
		Order:           castMember.Order,
	}
}

// MapToCrewMember maps a gotmdbapi.Crew to a v1.CrewMember
//
// Parameters:
//
//   - crewMember: the gotmdbapi.CrewMember to map
//
// Returns:
//
// - *v1.CrewMember: the mapped v1.CrewMember
func MapToCrewMember(crewMember *gotmdbapi.Crew) *v1.CrewMember {
	if crewMember == nil {
		return &v1.CrewMember{}
	}

	// Parse profile path from relative to full URL
	var profileURL *string
	if crewMember.ProfilePath != nil && *crewMember.ProfilePath != "" {
		profileURL = new(string)
		*profileURL = fmt.Sprintf(
			gotmdbapi.ImageVariableQualityURL,
			CrewMemberProfileImageWidthSize,
			(*crewMember.ProfilePath)[1:],
		)
	}

	return &v1.CrewMember{
		Adult:           crewMember.Adult,
		Gender:          MapToGender(crewMember.Gender),
		Id:              crewMember.ID,
		KnownDepartment: crewMember.KnownForDepartment,
		Name:            crewMember.Name,
		OriginalName:    crewMember.OriginalName,
		Popularity:      MapToOptionalFloat64(crewMember.Popularity),
		ProfileUrl:      profileURL,
		CreditId:        crewMember.CreditID,
		Department:      crewMember.Department,
		Job:             crewMember.Job,
	}
}

// MapCastMembers maps a slice of gotmdbapi.Cast to a slice of v1.CastMember
//
// Parameters:
//
// - castMembers: the slice of gotmdbapi.Cast to map
//
// Returns:
//
// - []*v1.CastMember: the mapped slice of v1.CastMember
func MapCastMembers(castMembers []gotmdbapi.Cast) []*v1.CastMember {
	mappedCastMembers := make([]*v1.CastMember, len(castMembers))
	for i := range castMembers {
		mappedCastMembers[i] = MapToCastMember(&castMembers[i])
	}
	return mappedCastMembers
}

// MapCrewMembers maps a slice of gotmdbapi.Crew to a slice of v1.CrewMember
//
// Parameters:
//
//   - crewMembers: the slice of gotmdbapi.Crew to map
//
// Returns:
//
// - []*v1.CrewMember: the mapped slice of v1.CrewMember
func MapCrewMembers(crewMembers []gotmdbapi.Crew) []*v1.CrewMember {
	mappedCrewMembers := make([]*v1.CrewMember, len(crewMembers))
	for i := range crewMembers {
		mappedCrewMembers[i] = MapToCrewMember(&crewMembers[i])
	}
	return mappedCrewMembers
}

// MapToGetMovieCreditsResponse maps a gotmdbapi.MovieCreditsResponse to a v1.GetMovieCreditsResponse
//
// Parameters:
//
// - response: the gotmdbapi.MovieCreditsResponse to map
//
// Returns:
//
// - *v1.GetMovieCreditsResponse: the mapped v1.GetMovieCreditsResponse
func MapToGetMovieCreditsResponse(response *gotmdbapi.MovieCreditsResponse) *v1.GetMovieCreditsResponse {
	if response == nil {
		return &v1.GetMovieCreditsResponse{}
	}
	castMembers := MapCastMembers(response.Cast)
	crewMembers := MapCrewMembers(response.Crew)
	return &v1.GetMovieCreditsResponse{
		Cast: castMembers,
		Crew: crewMembers,
	}
}

// MapToSimpleMovie maps a TMDB API movie to a simple movie
//
// Parameters:
//
// - movie: the TMDB API movie to map
//
// Returns:
//
// - *v1.SimpleMovie: the mapped simple movie
func MapToSimpleMovie(movie *gotmdbapi.SimpleMovie) *v1.SimpleMovie {
	if movie == nil {
		return &v1.SimpleMovie{}
	}

	// Parse poster path from relative to full URL
	var posterURL string
	if movie.PosterPath != "" {
		posterURL = fmt.Sprintf(
			gotmdbapi.ImageVariableQualityURL,
			SimpleMoviePosterImageWidthSize,
			movie.PosterPath[1:],
		)
	}

	return &v1.SimpleMovie{
		Adult:                movie.Adult,
		GenreIds:             movie.GenreIDs,
		Id:                   movie.ID,
		OriginalLanguage:     movie.OriginalLanguage,
		OriginalTitle:        movie.OriginalTitle,
		Overview:             movie.Overview,
		Popularity:           MapToOptionalFloat64(movie.Popularity),
		PosterUrl:            posterURL,
		ReleaseDate:          MapDateStringToTimestamp(movie.ReleaseDate),
		Title:                movie.Title,
		RatingAverageCritics: MapToOptionalFloat64(movie.VoteAverage),
		RatingCountCritics:   movie.VoteCount,
	}
}

// MapToSimpleMovies maps a slice of TMDB API movies to a slice of simple movies
//
// Parameters:
//
// - movies: the slice of TMDB API movies to map
//
// Returns:
//
// - []*v1.SimpleMovie: the mapped slice of simple movies
func MapToSimpleMovies(movies []gotmdbapi.SimpleMovie) []*v1.SimpleMovie {
	mappedMovies := make([]*v1.SimpleMovie, len(movies))
	for i := range movies {
		mappedMovies[i] = MapToSimpleMovie(&movies[i])
	}
	return mappedMovies
}

// MapToDateRange maps a TMDB API date range to a gRPC date range
//
// Parameters:
//
// - dateRange: the TMDB API date range to map
//
// Returns:
//
// - *v1.DateRange: the mapped gRPC date range
func MapToDateRange(dateRange *gotmdbapi.DateRange) *v1.DateRange {
	if dateRange == nil {
		return &v1.DateRange{}
	}
	return &v1.DateRange{
		Maximum: MapDateStringToTimestamp(dateRange.Maximum),
		Minimum: MapDateStringToTimestamp(dateRange.Minimum),
	}
}

// MapToGetNowPlayingMoviesResponse maps a TMDB API movie list response to a gRPC now playing movies response
//
// Parameters:
//
// - response: the TMDB API movie list response to map
//
// Returns:
//
// - *v1.GetNowPlayingMoviesResponse: the mapped gRPC now playing movies response
func MapToGetNowPlayingMoviesResponse(response *gotmdbapi.DateMovieListResponse) *v1.GetNowPlayingMoviesResponse {
	if response == nil {
		return &v1.GetNowPlayingMoviesResponse{}
	}
	return &v1.GetNowPlayingMoviesResponse{
		Dates:        MapToDateRange(&response.Dates),
		Page:         response.Page,
		Results:      MapToSimpleMovies(response.Results),
		TotalPages:   response.TotalPages,
		TotalResults: response.TotalResults,
	}
}

// MapToGetTopRatedMoviesResponse maps a TMDB API movie list response to a gRPC top rated movies response
//
// Parameters:
//
// - response: the TMDB API movie list response to map
//
// Returns:
//
// - *v1.GetTopRatedMoviesResponse: the mapped gRPC top rated movies response
func MapToGetTopRatedMoviesResponse(response *gotmdbapi.MovieListResponse) *v1.GetTopRatedMoviesResponse {
	if response == nil {
		return &v1.GetTopRatedMoviesResponse{}
	}
	return &v1.GetTopRatedMoviesResponse{
		Page:         response.Page,
		Results:      MapToSimpleMovies(response.Results),
		TotalPages:   response.TotalPages,
		TotalResults: response.TotalResults,
	}
}

// MapToGetPopularMoviesResponse maps a TMDB API movie list response to a gRPC popular movies response
//
// Parameters:
//
// - response: the TMDB API movie list response to map
//
// Returns:
//
// - *v1.GetPopularMoviesResponse: the mapped gRPC popular movies response
func MapToGetPopularMoviesResponse(response *gotmdbapi.MovieListResponse) *v1.GetPopularMoviesResponse {
	if response == nil {
		return &v1.GetPopularMoviesResponse{}
	}
	return &v1.GetPopularMoviesResponse{
		Page:         response.Page,
		Results:      MapToSimpleMovies(response.Results),
		TotalPages:   response.TotalPages,
		TotalResults: response.TotalResults,
	}
}

// MapToGetUpcomingMoviesResponse maps a TMDB API movie list response to a gRPC upcoming movies response
//
// Parameters:
//
// - response: the TMDB API movie list response to map
//
// Returns:
//
// - *v1.GetUpcomingMoviesResponse: the mapped gRPC upcoming movies response
func MapToGetUpcomingMoviesResponse(response *gotmdbapi.DateMovieListResponse) *v1.GetUpcomingMoviesResponse {
	if response == nil {
		return &v1.GetUpcomingMoviesResponse{}
	}
	return &v1.GetUpcomingMoviesResponse{
		Dates:        MapToDateRange(&response.Dates),
		Page:         response.Page,
		Results:      MapToSimpleMovies(response.Results),
		TotalPages:   response.TotalPages,
		TotalResults: response.TotalResults,
	}
}

// MapToSimilarMoviesResponse maps a TMDB API movie list response to a gRPC similar movies response
//
// Parameters:
//
// - response: the TMDB API movie list response to map
//
// Returns:
//
// - *v1.SimilarMoviesResponse: the mapped gRPC similar movies response
func MapToSimilarMoviesResponse(response *gotmdbapi.MovieListResponse) *v1.SimilarMoviesResponse {
	if response == nil {
		return &v1.SimilarMoviesResponse{}
	}
	return &v1.SimilarMoviesResponse{
		Page:         response.Page,
		Results:      MapToSimpleMovies(response.Results),
		TotalPages:   response.TotalPages,
		TotalResults: response.TotalResults,
	}
}

// MapToSearchMoviesResponse maps a TMDB API movie list response to a gRPC search movies response
//
// Parameters:
//
// - response: the TMDB API movie list response to map
//
// Returns:
//
// - *v1.SearchMoviesResponse: the mapped gRPC search movies response
func MapToSearchMoviesResponse(response *gotmdbapi.MovieListResponse) *v1.SearchMoviesResponse {
	if response == nil {
		return &v1.SearchMoviesResponse{}
	}
	return &v1.SearchMoviesResponse{
		Page:         response.Page,
		Results:      MapToSimpleMovies(response.Results),
		TotalPages:   response.TotalPages,
		TotalResults: response.TotalResults,
	}
}

// MapToGenre maps a TMDB API genre to a gRPC genre
//
// Parameters:
//
// - genre: the TMDB API genre to map
//
// Returns:
//
// - *v1.Genre: the mapped gRPC genre
func MapToGenre(genre *gotmdbapi.Genre) *v1.Genre {
	if genre == nil {
		return &v1.Genre{}
	}
	return &v1.Genre{
		Id:   genre.ID,
		Name: genre.Name,
	}
}

// MapToGenres maps a slice of TMDB API genres to a slice of gRPC genres
//
// Parameters:
//
// - genres: the slice of TMDB API genres to map
//
// Returns:
//
// - []*v1.Genre: the mapped slice of gRPC genres
func MapToGenres(genres []gotmdbapi.Genre) []*v1.Genre {
	mappedGenres := make([]*v1.Genre, len(genres))
	for i, genre := range genres {
		mappedGenres[i] = MapToGenre(&genre)
	}
	return mappedGenres
}

// MapToProductionCompany maps a TMDB API production company to a gRPC production company
//
// Parameters:
//
// - company: the TMDB API production company to map
//
// Returns:
//
// - *v1.ProductionCompany: the mapped gRPC production company
func MapToProductionCompany(company *gotmdbapi.ProductionCompany) *v1.ProductionCompany {
	if company == nil {
		return &v1.ProductionCompany{}
	}

	// Parse logo path from relative to full URL
	var logoURL *string
	if company.LogoPath != nil && *company.LogoPath != "" {
		logoURL = new(string)
		*logoURL = fmt.Sprintf(
			gotmdbapi.ImageVariableQualityURL,
			ProductionCompanyLogoImageWidthSize,
			(*company.LogoPath)[1:],
		)
	}

	return &v1.ProductionCompany{
		Id:            company.ID,
		LogoUrl:       logoURL,
		Name:          company.Name,
		OriginCountry: company.OriginCountry,
	}
}

// MapToProductionCompanies maps a slice of TMDB API production companies to a slice of gRPC production companies
//
// Parameters:
//
// - companies: the slice of TMDB API production companies to map
//
// Returns:
//
// - []*v1.ProductionCompany: the mapped slice of gRPC production companies
func MapToProductionCompanies(companies []gotmdbapi.ProductionCompany) []*v1.ProductionCompany {
	mappedCompanies := make([]*v1.ProductionCompany, len(companies))
	for i, company := range companies {
		mappedCompanies[i] = MapToProductionCompany(&company)
	}
	return mappedCompanies
}

// MapToProductionCountry maps a TMDB API production country to a gRPC production country
//
// Parameters:
//
// - country: the TMDB API production country to map
//
// Returns:
//
// - *v1.ProductionCountry: the mapped gRPC production country
func MapToProductionCountry(country *gotmdbapi.ProductionCountry) *v1.ProductionCountry {
	if country == nil {
		return &v1.ProductionCountry{}
	}
	return &v1.ProductionCountry{
		Iso_3166_1: country.ISO3166_1,
		Name:       country.Name,
	}
}

// MapToProductionCountries maps a slice of TMDB API production countries to a slice of gRPC production countries
//
// Parameters:
//
// - countries: the slice of TMDB API production countries to map
//
// Returns:
//
// - []*v1.ProductionCountry: the mapped slice of gRPC production countries
func MapToProductionCountries(countries []gotmdbapi.ProductionCountry) []*v1.ProductionCountry {
	mappedCountries := make([]*v1.ProductionCountry, len(countries))
	for i, country := range countries {
		mappedCountries[i] = MapToProductionCountry(&country)
	}
	return mappedCountries
}

// MapToGetMovieDetailsResponse maps a TMDB API movie details response to a gRPC movie details response
//
// Parameters:
//
// - response: the TMDB API movie details response to map
//
// Returns:
//
// - *v1.GetMovieDetailsResponse: the mapped gRPC movie details response
func MapToGetMovieDetailsResponse(response *gotmdbapi.MovieDetailsResponse) *v1.GetMovieDetailsResponse {
	if response == nil {
		return &v1.GetMovieDetailsResponse{}
	}

	// Parse poster path from relative to full URL
	var posterURL string
	if response.PosterPath != "" {
		posterURL = fmt.Sprintf(
			gotmdbapi.ImageVariableQualityURL,
			MovieDetailsPosterImageWidthSize,
			response.PosterPath[1:],
		)
	}

	return &v1.GetMovieDetailsResponse{
		Adult:                response.Adult,
		Budget:               response.Budget,
		Genres:               MapToGenres(response.Genres),
		OriginalLanguage:     response.OriginalLanguage,
		Homepage:             response.Homepage,
		Id:                   response.ID,
		OriginalTitle:        response.OriginalTitle,
		Overview:             response.Overview,
		PosterUrl:            posterURL,
		Popularity:           MapToOptionalFloat64(response.Popularity),
		ProductionCompanies:  MapToProductionCompanies(response.ProductionCompanies),
		ProductionCountries:  MapToProductionCountries(response.ProductionCountries),
		ReleaseDate:          MapDateStringToTimestamp(response.ReleaseDate),
		Revenue:              response.Revenue,
		Runtime:              response.Runtime,
		Status:               response.Status,
		Tagline:              response.Tagline,
		Title:                response.Title,
		RatingAverageCritics: MapToOptionalFloat64(response.VoteAverage),
		RatingCountCritics:   response.VoteCount,
	}
}

// MapToCriticAuthorDetails maps a TMDB API author details to a gRPC critic author details
//
// Parameters:
//
// - authorDetails: the TMDB API author details to map
//
// Returns:
//
// - *v1.CriticAuthorDetails: the mapped gRPC author details
func MapToCriticAuthorDetails(authorDetails *gotmdbapi.AuthorDetails) *v1.CriticAuthorDetails {
	if authorDetails == nil {
		return &v1.CriticAuthorDetails{}
	}

	// Parse avatar path from relative to full URL
	var avatarURL *string
	if authorDetails.AvatarPath != nil && *authorDetails.AvatarPath != "" {
		avatarURL = new(string)
		*avatarURL = fmt.Sprintf(
			gotmdbapi.ImageVariableQualityURL,
			AvatarImageWidthSize,
			(*authorDetails.AvatarPath)[1:],
		)
	}

	return &v1.CriticAuthorDetails{
		Name:       authorDetails.Name,
		Username:   authorDetails.Username,
		AvatarPath: avatarURL,
		Rating:     authorDetails.Rating,
	}
}

// MapToMovieReview maps a TMDB API movie review to a gRPC movie review
//
// Parameters:
//
// - review: the TMDB API movie review to map
//
// Returns:
//
// - *v1.MovieReview: the mapped gRPC movie review
func MapToMovieReview(review *gotmdbapi.Review) *v1.MovieCriticReview {
	if review == nil {
		return &v1.MovieCriticReview{}
	}
	return &v1.MovieCriticReview{
		Id:            review.ID,
		Author:        review.Author,
		AuthorDetails: MapToCriticAuthorDetails(&review.AuthorDetails),
		Content:       review.Content,
		CreatedAt: func() *timestamppb.Timestamp {
			parsedTime, err := time.Parse(time.RFC3339, review.CreatedAt)
			if err != nil {
				return nil
			}
			return timestamppb.New(parsedTime)
		}(),
		UpdatedAt: func() *timestamppb.Timestamp {
			parsedTime, err := time.Parse(time.RFC3339, review.UpdatedAt)
			if err != nil {
				return nil
			}
			return timestamppb.New(parsedTime)
		}(),
		Url: review.URL,
	}
}

// MapToMovieReviews maps a slice of TMDB API movie reviews to a slice of gRPC movie reviews
//
// Parameters:
//
// - reviews: the slice of TMDB API movie reviews to map
//
// Returns:
//
// - []*v1.MovieReview: the mapped slice of gRPC movie reviews
func MapToMovieReviews(reviews []gotmdbapi.Review) []*v1.MovieCriticReview {
	mappedReviews := make([]*v1.MovieCriticReview, len(reviews))
	for i := range reviews {
		mappedReviews[i] = MapToMovieReview(&reviews[i])
	}
	return mappedReviews
}

// MapToGetMovieReviewsResponse maps a TMDB API movie reviews response to a gRPC movie reviews response
//
// Parameters:
//
// - response: the TMDB API movie reviews response to map
//
// Returns:
//
// - *v1.GetMovieReviewsResponse: the mapped gRPC movie reviews response
func MapToGetMovieReviewsResponse(response *gotmdbapi.MovieReviewsResponse) *v1.GetMovieReviewsResponse {
	if response == nil {
		return &v1.GetMovieReviewsResponse{}
	}
	return &v1.GetMovieReviewsResponse{
		CriticReviews: MapToMovieReviews(response.Results),
		Page:          response.Page,
		TotalPages:    response.TotalPages,
		TotalResults:  response.TotalResults,
	}
}

// MapToGetMovieGenresResponse maps a TMDB API genre list response to a gRPC get movie genres response
//
// Parameters:
//
// - response: the TMDB API genre list response to map
//
// Returns:
//
// - *v1.GetMovieGenresResponse: the mapped gRPC get movie genres response
func MapToGetMovieGenresResponse(response *gotmdbapi.GenreListResponse) *v1.GetMovieGenresResponse {
	if response == nil {
		return &v1.GetMovieGenresResponse{}
	}
	return &v1.GetMovieGenresResponse{
		Genres: MapToGenres(response.Genres),
	}
}

// MapToDiscoverMoviesResponse maps a TMDB API movie list response to a gRPC discover movies response
//
// Parameters:
//
// - response: the TMDB API movie list response to map
//
// Returns:
//
// - *v1.DiscoverMoviesResponse: the mapped gRPC discover movies response
func MapToDiscoverMoviesResponse(response *gotmdbapi.MovieListResponse) *v1.DiscoverMoviesResponse {
	if response == nil {
		return &v1.DiscoverMoviesResponse{}
	}
	return &v1.DiscoverMoviesResponse{
		Page:         response.Page,
		Results:      MapToSimpleMovies(response.Results),
		TotalPages:   response.TotalPages,
		TotalResults: response.TotalResults,
	}
}

// MapToSortBy maps a gRPC sort by to a TMDB API sort by
//
// Parameters:
//
// - sortBy: the gRPC sort by to map
//
// Returns:
//
// - gotmdbapi.SortByEnum: the mapped TMDB API sort by
func MapToSortBy(sortBy v1.SortBy) gotmdbapi.SortByEnum {
	switch sortBy {
	case v1.SortBy_SORT_BY_UNSPECIFIED:
		return ""
	case v1.SortBy_POPULARITY_ASC:
		return gotmdbapi.SortByPopularityAsc
	case v1.SortBy_POPULARITY_DESC:
		return gotmdbapi.SortByPopularityDesc
	case v1.SortBy_REVENUE_ASC:
		return gotmdbapi.SortByRevenueAsc
	case v1.SortBy_REVENUE_DESC:
		return gotmdbapi.SortByRevenueDesc
	case v1.SortBy_PRIMARY_RELEASE_DATE_ASC:
		return gotmdbapi.SortByPrimaryReleaseDateAsc
	case v1.SortBy_PRIMARY_RELEASE_DATE_DESC:
		return gotmdbapi.SortByPrimaryReleaseDateDesc
	case v1.SortBy_ORIGINAL_TITLE_ASC:
		return gotmdbapi.SortByOriginalTitleAsc
	case v1.SortBy_ORIGINAL_TITLE_DESC:
		return gotmdbapi.SortByOriginalTitleDesc
	case v1.SortBy_VOTE_AVERAGE_ASC:
		return gotmdbapi.SortByVoteAverageAsc
	case v1.SortBy_VOTE_AVERAGE_DESC:
		return gotmdbapi.SortByVoteAverageDesc
	case v1.SortBy_VOTE_COUNT_ASC:
		return gotmdbapi.SortByVoteCountAsc
	case v1.SortBy_VOTE_COUNT_DESC:
		return gotmdbapi.SortByVoteCountDesc
	default:
		return ""
	}
}

// MapToWatchMonetizationType maps a gRPC watch monetization type to a TMDB API watch monetization type
//
// Parameters:
//
// - t: the gRPC watch monetization type to map
//
// Returns:
//
// - gotmdbapi.WatchMonetizationTypeEnums: the mapped TMDB API watch monetization type
func MapToWatchMonetizationType(t v1.WatchMonetizationType) gotmdbapi.WatchMonetizationTypeEnums {
	switch t {
	case v1.WatchMonetizationType_WATCH_MONETIZATION_TYPE_UNSPECIFIED:
		return ""
	case v1.WatchMonetizationType_FLATRATE:
		return gotmdbapi.WatchMonetizationTypeFlatrate
	case v1.WatchMonetizationType_FREE:
		return gotmdbapi.WatchMonetizationTypeFree
	case v1.WatchMonetizationType_ADS:
		return gotmdbapi.WatchMonetizationTypeAds
	case v1.WatchMonetizationType_RENT:
		return gotmdbapi.WatchMonetizationTypeRent
	case v1.WatchMonetizationType_BUY:
		return gotmdbapi.WatchMonetizationTypeBuy
	default:
		return ""
	}
}

// MapToWatchMonetizationTypes maps a slice of gRPC watch monetization types to a slice of TMDB API watch monetization
// types
//
// Parameters:
//
//   - types: the slice of gRPC watch monetization types to map
//
// Returns:
//
// - []gotmdbapi.WatchMonetizationTypeEnums: the mapped slice of TMDB API watch monetization types
func MapToWatchMonetizationTypes(types []v1.WatchMonetizationType) []gotmdbapi.WatchMonetizationTypeEnums {
	mappedTypes := make([]gotmdbapi.WatchMonetizationTypeEnums, 0, len(types))
	for _, t := range types {
		mappedType := MapToWatchMonetizationType(t)
		if mappedType != "" {
			mappedTypes = append(mappedTypes, mappedType)
		}
	}
	return mappedTypes
}
