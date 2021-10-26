package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AidanFogarty/go-f1/pkg/testutil"
)

const (
	successScheduleRespXML = `
	<?xml version="1.0" encoding="utf-8"?>
	<?xml-stylesheet type="text/xsl" href="/schemas/mrd-1.4.xsl"?>
	<MRData xmlns="http://ergast.com/mrd/1.4" series="f1" url="http://ergast.com/api/f1/current" limit="30" offset="0" total="22">
		<RaceTable season="2021">
			<Race season="2021" round="1" url="https://en.wikipedia.org/wiki/2021_Bahrain_Grand_Prix">
				<RaceName>Bahrain Grand Prix</RaceName>
				<Circuit circuitId="bahrain" url="http://en.wikipedia.org/wiki/Bahrain_International_Circuit">
					<CircuitName>Bahrain International Circuit</CircuitName>
					<Location lat="26.0325" long="50.5106">
						<Locality>Sakhir</Locality>
						<Country>Bahrain</Country>
					</Location>
				</Circuit>
				<Date>2021-03-28</Date>
				<Time>15:00:00Z</Time>
			</Race>
			<Race season="2021" round="2" url="http://en.wikipedia.org/wiki/2021_Emilia_Romagna_Grand_Prix">
				<RaceName>Emilia Romagna Grand Prix</RaceName>
				<Circuit circuitId="imola" url="http://en.wikipedia.org/wiki/Autodromo_Enzo_e_Dino_Ferrari">
					<CircuitName>Autodromo Enzo e Dino Ferrari</CircuitName>
					<Location lat="44.3439" long="11.7167">
						<Locality>Imola</Locality>
						<Country>Italy</Country>
					</Location>
				</Circuit>
				<Date>2021-04-18</Date>
				<Time>13:00:00Z</Time>
			</Race>
		</RaceTable>
	</MRData>
	`
)

func newTestServer(response string) *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(response))
	})
	return httptest.NewServer(handler)
}

func TestErgast_Schedule(t *testing.T) {
	type args struct {
		ctx  context.Context
		year int
	}
	tests := []struct {
		name     string
		args     args
		response string
		want     []Race
		wantErr  bool
	}{
		{
			name:     "Successful Schedule Response",
			response: successScheduleRespXML,
			args:     args{context.TODO(), 2021},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			srv := newTestServer(test.response)

			ergast := Ergast{
				BaseURL:    srv.URL,
				HTTPClient: srv.Client(),
			}

			_, err := ergast.Schedule(test.args.ctx, test.args.year)
			testutil.Ok(t, err)
		})
	}
}
