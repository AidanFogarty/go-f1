package api

import (
	"context"
	"testing"

	"github.com/AidanFogarty/go-f1/pkg/testutil"
)

const (
	successLapsRespXML = `
	<?xml version="1.0" encoding="utf-8"?>
	<?xml-stylesheet type="text/xsl" href="/schemas/mrd-1.4.xsl"?>
	<MRData xmlns="http://ergast.com/mrd/1.4" series="f1" url="http://ergast.com/api/f1/2021/last/laps" limit="40" offset="0" total="40">
		<RaceTable season="2021" round="20">
			<Race season="2021" round="20" url="http://en.wikipedia.org/wiki/2021_Qatar_Grand_Prix">
				<RaceName>Qatar Grand Prix</RaceName>
				<Circuit circuitId="losail" url="http://en.wikipedia.org/wiki/Losail_International_Circuit">
					<CircuitName>Losail International Circuit</CircuitName>
					<Location lat="25.49" long="51.4542">
						<Locality>Al Daayen</Locality>
						<Country>Qatar</Country>
					</Location>
				</Circuit>
				<Date>2021-11-21</Date>
				<Time>14:00:00Z</Time>
				<LapsList>
					<Lap number="1">
						<Timing driverId="hamilton" lap="1" position="1" time="1:33.760"/>
						<Timing driverId="alonso" lap="1" position="2" time="1:35.613"/>
						<Timing driverId="gasly" lap="1" position="3" time="1:36.603"/>
						<Timing driverId="max_verstappen" lap="1" position="4" time="1:37.292"/>
						<Timing driverId="norris" lap="1" position="5" time="1:38.135"/>
						<Timing driverId="ocon" lap="1" position="6" time="1:38.931"/>
						<Timing driverId="sainz" lap="1" position="7" time="1:39.804"/>
						<Timing driverId="tsunoda" lap="1" position="8" time="1:40.382"/>
						<Timing driverId="perez" lap="1" position="9" time="1:41.182"/>
						<Timing driverId="stroll" lap="1" position="10" time="1:41.296"/>
						<Timing driverId="bottas" lap="1" position="11" time="1:41.898"/>
						<Timing driverId="raikkonen" lap="1" position="12" time="1:42.464"/>
						<Timing driverId="leclerc" lap="1" position="13" time="1:43.059"/>
						<Timing driverId="giovinazzi" lap="1" position="14" time="1:43.751"/>
						<Timing driverId="russell" lap="1" position="15" time="1:44.238"/>
						<Timing driverId="ricciardo" lap="1" position="16" time="1:44.471"/>
						<Timing driverId="vettel" lap="1" position="17" time="1:44.759"/>
						<Timing driverId="mick_schumacher" lap="1" position="18" time="1:45.606"/>
						<Timing driverId="latifi" lap="1" position="19" time="1:45.752"/>
						<Timing driverId="mazepin" lap="1" position="20" time="1:46.318"/>
					</Lap>
					<Lap number="2">
						<Timing driverId="hamilton" lap="2" position="1" time="1:30.198"/>
						<Timing driverId="alonso" lap="2" position="2" time="1:30.573"/>
						<Timing driverId="gasly" lap="2" position="3" time="1:30.457"/>
						<Timing driverId="max_verstappen" lap="2" position="4" time="1:30.572"/>
						<Timing driverId="norris" lap="2" position="5" time="1:30.972"/>
						<Timing driverId="ocon" lap="2" position="6" time="1:31.127"/>
						<Timing driverId="sainz" lap="2" position="7" time="1:31.392"/>
						<Timing driverId="tsunoda" lap="2" position="8" time="1:31.817"/>
						<Timing driverId="perez" lap="2" position="9" time="1:31.896"/>
						<Timing driverId="stroll" lap="2" position="10" time="1:32.272"/>
						<Timing driverId="bottas" lap="2" position="11" time="1:32.520"/>
						<Timing driverId="raikkonen" lap="2" position="12" time="1:32.653"/>
						<Timing driverId="leclerc" lap="2" position="13" time="1:32.498"/>
						<Timing driverId="giovinazzi" lap="2" position="14" time="1:32.713"/>
						<Timing driverId="russell" lap="2" position="15" time="1:33.009"/>
						<Timing driverId="ricciardo" lap="2" position="16" time="1:33.120"/>
						<Timing driverId="vettel" lap="2" position="17" time="1:33.623"/>
						<Timing driverId="mick_schumacher" lap="2" position="18" time="1:33.739"/>
						<Timing driverId="latifi" lap="2" position="19" time="1:33.655"/>
						<Timing driverId="mazepin" lap="2" position="20" time="1:34.272"/>
					</Lap>
				</LapsList>
			</Race>
		</RaceTable>
	</MRData>
	`
)

func TestErgast_Laps(t *testing.T) {
	type args struct {
		ctx   context.Context
		year  int
		round string
	}
	tests := []struct {
		name         string
		args         args
		response     string
		want         []Lap
		wantErr      bool
		expectedLaps int
	}{
		{
			name:         "Successful Laps Response",
			args:         args{context.TODO(), 2021, "last"},
			response:     successLapsRespXML,
			want:         []Lap{},
			wantErr:      false,
			expectedLaps: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			srv := newTestServer(successLapsRespXML)

			ergast := &Ergast{
				BaseURL:    srv.URL,
				HTTPClient: srv.Client(),
			}

			laps, err := ergast.Laps(test.args.ctx, test.args.year, test.args.round)

			testutil.Ok(t, err)
			testutil.Equals(t, test.expectedLaps, len(laps))
		})
	}
}
