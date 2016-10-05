package main

import (
	"flag"
	"fmt"
	"github.com/arthurandres/sklib"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	country       = flag.String("country", "GB", "Country of the user")
	currency      = flag.String("currency", "GBP", "Currency of the user")
	language      = flag.String("language", "en-GB", "Language of the user")
	keyFile       = flag.String("keyFile", "key", "API key provided by skyscanner")
	origin        = flag.String("from", "LON", "Origin Town/Airport")
	departureDate = flag.String("out", "20161101", "Date of departure/outbound flight")
	returnDate    = flag.String("in", "20161103", "Date of return/inbound flight")
	noCache       = flag.Bool("noCache", false, "Do not read from cache")
	delay         = flag.Bool("delay", false, "Simulate random delay in query")
	destinations  = flag.String("to", "", "selected destinations")
	departAfter   = flag.String("departAfter", "", "Minimum departure time")
	returnAfter   = flag.String("returnAfter", "", "Minimum return time")
	directOnly    = flag.Bool("direct", false, "Direct flights only")
)

type ApplicationParameters struct {
	Localisation  sklib.Localisation
	KeyFile       string
	Key           string
	Origin        string
	DepartureDate string
	ReturnDate    string
	NoCache       bool
	Delay         bool
	Destinations  []string
	DepartAfter   *time.Duration
	ReturnAfter   *time.Duration
	DirectOnly    bool
}

func ParseDestinations(input string) []string {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return make([]string, 0)
	} else {
		split := strings.Split(input, ",")
		results := make([]string, 0, len(split))
		for _, s := range split {
			s := strings.TrimSpace(s)
			if len(s) != 0 {
				results = append(results, s)
			}
		}
		return results
	}

}

func ReadArguments() ApplicationParameters {
	flag.Parse()
	key := sklib.ReadKey(*keyFile)
	return ApplicationParameters{
		Localisation:  ReadLocalisationArguments(),
		KeyFile:       *keyFile,
		Key:           key,
		Origin:        *origin,
		DepartureDate: *departureDate,
		ReturnDate:    *returnDate,
		NoCache:       *noCache,
		Delay:         *delay,
		Destinations:  ParseDestinations(*destinations),
		DepartAfter:   ReadDurationArgument(departAfter),
		ReturnAfter:   ReadDurationArgument(returnAfter),
		DirectOnly:    *directOnly}

}

func ReadLocalisationArguments() sklib.Localisation {
	return sklib.Localisation{
		Country:  *country,
		Currency: *currency,
		Language: *language}
}

func ReadDurationArgument(input *string) *time.Duration {
	if input == nil || len(*input) == 0 {
		return nil
	}
	duration, err := ParseTimeOfDay(*input)
	if err != nil {
		panic(err)
	}
	return &duration
}

func DisplayFullQuotes(request sklib.BrowseRoutesRequest, quotes sklib.FullQuotes) {

	sort.Sort(sort.Reverse(quotes))
	for _, v := range quotes {
		link := sklib.GetLink(request.Origin, v.Destination.SkyscannerCode, request.DepartureDate, request.ReturnDate)
		fmt.Printf("%s %.0f %s %s\n", v.Destination.SkyscannerCode, v.Quote.MinPrice, v.Destination.Name, link)
	}
	fmt.Printf("%d results\n", len(quotes))
}

func main() {
	arguments := ReadArguments()
	engine := &sklib.LiveEngine{Key: arguments.Key}
	db := sklib.CreateDB()
	defer db.Close()
	var cache sklib.CacheStore = sklib.CreateCache(db)
	if arguments.NoCache {
		cache = &sklib.WriteOnlyStore{Store: cache}
	}
	var ce sklib.RequestEngine = &sklib.CachedEngine{engine, cache}

	if arguments.Delay {
		ce = &sklib.SlowEngine{ce}
	}
	if len(arguments.Destinations) == 0 {
		browse(ce, arguments)
	} else {
		search(ce, arguments)
	}
}

func search(engine sklib.RequestEngine, arguments ApplicationParameters) {
	itineraries, err := sklib.Search(engine, arguments.ToSearchRequest())
	if err != nil {
		panic(err)
	}
	filtered := sklib.ApplyFilter(itineraries, arguments.ToFilter())
	DisplayItineraries(filtered)
}

func browse(engine sklib.RequestEngine, arguments ApplicationParameters) {
	results, err := sklib.Browse(engine, arguments.ToBrowseRoutesRequest())
	if err != nil {
		panic(err)
	}
	if arguments.DirectOnly {
		results = sklib.FilterDirects(results)
	}
	DisplayFullQuotes(arguments.ToBrowseRoutesRequest(), results)
}

func DisplayItineraries(input sklib.Itineraries) {
	sort.Sort(sort.Reverse(input))
	for _, itinerary := range input {
		fmt.Println(
			itinerary.OutboundLeg.Display(),
			itinerary.InboundLeg.Display(),
			itinerary.GetPrice())
	}
	fmt.Println(len(input), "results")
}

func ParseTimeOfDay(input string) (time.Duration, error) {
	display, err := strconv.ParseUint(input, 10, 32)
	if err != nil {
		return time.Minute, err
	}
	minutes := display % 100
	hours := display / 100
	if minutes >= 60 || hours >= 24 {
		panic(fmt.Errorf("Wrong time of day %s", input))
	}
	return time.Minute*time.Duration(minutes) + time.Hour*time.Duration(hours), nil
}

func (m *ApplicationParameters) ToFilter() sklib.ItineraryFilter {
	filters := make(sklib.CompositeFilter, 0)
	filters = append(filters, &sklib.DirectFilter{m.DirectOnly})
	filters = sklib.AppendTimeFilter(filters, m.DepartAfter, false, true)
	filters = sklib.AppendTimeFilter(filters, m.ReturnAfter, false, false)
	return filters
}

func (m *ApplicationParameters) ToBrowseRoutesRequest() sklib.BrowseRoutesRequest {
	return sklib.BrowseRoutesRequest{
		Localisation:  m.Localisation,
		Origin:        m.Origin,
		Destination:   "",
		DepartureDate: m.DepartureDate,
		ReturnDate:    m.ReturnDate}
}

func (m *ApplicationParameters) ToSearchRequest() sklib.SearchRequest {
	return sklib.SearchRequest{
		Localisation:  m.Localisation,
		Origin:        m.Origin,
		Destinations:  m.Destinations,
		DepartureDate: m.DepartureDate,
		ReturnDate:    m.ReturnDate}
}
