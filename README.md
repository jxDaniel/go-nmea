# go-nmea [![Build Status](https://travis-ci.org/adrianmo/go-nmea.svg?branch=master)](https://travis-ci.org/adrianmo/go-nmea) [![Go Report Card](https://goreportcard.com/badge/github.com/adrianmo/go-nmea)](https://goreportcard.com/report/github.com/adrianmo/go-nmea) [![Coverage Status](https://coveralls.io/repos/adrianmo/go-nmea/badge.svg?branch=master&service=github)](https://coveralls.io/github/adrianmo/go-nmea?branch=master) [![GoDoc](https://godoc.org/github.com/adrianmo/go-nmea?status.svg)](https://godoc.org/github.com/adrianmo/go-nmea)

This is a NMEA library for the Go programming language (http://golang.org).

## Installing

### Using `go get`

    go get github.com/adrianmo/go-nmea

After this command *go-nmea* is ready to use. Its source will be in:

    $GOPATH/src/github.com/adrianmo/go-nmea

## Supported sentences

At this moment, this library supports the following sentence types:

- [RMC](http://aprs.gids.nl/nmea/#rmc) - Recommended Minimum Specific GPS/Transit data
- [GGA](http://aprs.gids.nl/nmea/#gga) - GPS Positioning System Fix Data
- [GSA](http://aprs.gids.nl/nmea/#gsa) - GPS DOP and active satellites
- [GSV](http://aprs.gids.nl/nmea/#gsv) - GPS Satellites in view
- [GLL](http://aprs.gids.nl/nmea/#gll) - Geographic Position, Latitude / Longitude and time
- [VTG](http://aprs.gids.nl/nmea/#vtg) - Track Made Good and Ground Speed
- [ZDA](http://aprs.gids.nl/nmea/#zda) - Date & time data
- [HDT](http://aprs.gids.nl/nmea/#hdt) - Actual vessel heading in degrees True
- [GNS](https://www.trimble.com/oem_receiverhelp/v4.44/en/NMEA-0183messages_GNS.html) - Combined GPS fix for GPS, Glonass, Galileo, and BeiDou
- [PGRME](http://aprs.gids.nl/nmea/#rme) - Estimated Position Error (Garmin proprietary sentence)
- [THS](http://www.nuovamarea.net/pytheas_9.html) - Actual vessel heading in degrees True and status
- [VDM/VDO](http://catb.org/gpsd/AIVDM.html) - Encapsulated binary payload
- [WPL](http://aprs.gids.nl/nmea/#wpl) - Waypoint location
- [RTE](http://aprs.gids.nl/nmea/#rte) - Route

## Example

```go
package main

import (
	"fmt"
	"log"
	"github.com/adrianmo/go-nmea"
)

func main() {
	sentence := "$GPRMC,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*70"
	s, err := nmea.Parse(sentence)
	if err != nil {
		log.Fatal(err)
	}
	if s.DataType() == nmea.TypeRMC {
		m := s.(nmea.RMC)
		fmt.Printf("Raw sentence: %v\n", m)
		fmt.Printf("Time: %s\n", m.Time)
		fmt.Printf("Validity: %s\n", m.Validity)
		fmt.Printf("Latitude GPS: %s\n", nmea.FormatGPS(m.Latitude))
		fmt.Printf("Latitude DMS: %s\n", nmea.FormatDMS(m.Latitude))
		fmt.Printf("Longitude GPS: %s\n", nmea.FormatGPS(m.Longitude))
		fmt.Printf("Longitude DMS: %s\n", nmea.FormatDMS(m.Longitude))
		fmt.Printf("Speed: %f\n", m.Speed)
		fmt.Printf("Course: %f\n", m.Course)
		fmt.Printf("Date: %s\n", m.Date)
		fmt.Printf("Variation: %f\n", m.Variation)
	}
}
```

Output:

```
$ go run main/main.go

Raw sentence: $GPRMC,220516,A,5133.82,N,00042.24,W,173.8,231.8,130694,004.2,W*70
Time: 22:05:16.0000
Validity: A
Latitude GPS: 5133.8200
Latitude DMS: 51° 33' 49.200000"
Longitude GPS: 042.2400
Longitude DMS: 0° 42' 14.400000"
Speed: 173.800000
Course: 231.800000
Date: 13/06/94
Variation: -4.200000
```

## Contributions

Please, feel free to implement support for new sentences, fix bugs, refactor code, etc. and send a pull-request to update the library.
