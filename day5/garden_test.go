package main

import "testing"

func Test_findLowestLocationNumberForSeed(t *testing.T) {
	seedToSoilMapping := []gardenMap{
		{
			destinationRangeStart: 50,
			sourceRangeStart:      98,
			rangeLength:           2,
		},
		{
			destinationRangeStart: 52,
			sourceRangeStart:      50,
			rangeLength:           48,
		}}

	soilToFertilizerMapping := []gardenMap{
		{
			destinationRangeStart: 0,
			sourceRangeStart:      15,
			rangeLength:           37,
		},
		{
			destinationRangeStart: 37,
			sourceRangeStart:      52,
			rangeLength:           2,
		},
		{
			destinationRangeStart: 39,
			sourceRangeStart:      0,
			rangeLength:           15,
		},
	}

	fertilizerToWaterMapping := []gardenMap{
		{
			destinationRangeStart: 49,
			sourceRangeStart:      53,
			rangeLength:           8,
		},
		{
			destinationRangeStart: 0,
			sourceRangeStart:      11,
			rangeLength:           42,
		},
		{
			destinationRangeStart: 42,
			sourceRangeStart:      0,
			rangeLength:           7,
		},
		{
			destinationRangeStart: 57,
			sourceRangeStart:      7,
			rangeLength:           4,
		},
	}

	waterToLightMap := []gardenMap{
		{
			destinationRangeStart: 88,
			sourceRangeStart:      18,
			rangeLength:           7,
		},
		{
			destinationRangeStart: 18,
			sourceRangeStart:      25,
			rangeLength:           70,
		},
	}

	lightToTempMap := []gardenMap{
		{
			destinationRangeStart: 45,
			sourceRangeStart:      77,
			rangeLength:           23,
		},
		{
			destinationRangeStart: 81,
			sourceRangeStart:      45,
			rangeLength:           19,
		},
		{
			destinationRangeStart: 68,
			sourceRangeStart:      64,
			rangeLength:           13,
		},
	}

	tempToHumidityMap := []gardenMap{
		{
			destinationRangeStart: 0,
			sourceRangeStart:      69,
			rangeLength:           1,
		},
		{
			destinationRangeStart: 1,
			sourceRangeStart:      0,
			rangeLength:           69,
		},
	}

	humidityToLocationMap := []gardenMap{
		{
			destinationRangeStart: 60,
			sourceRangeStart:      56,
			rangeLength:           37,
		},
		{
			destinationRangeStart: 56,
			sourceRangeStart:      93,
			rangeLength:           4,
		},
	}

	type args struct {
		seed                     int
		seedToSoilMapping        []gardenMap
		soilToFertilizerMapping  []gardenMap
		fertilizerToWaterMapping []gardenMap
		waterToLightMap          []gardenMap
		lightToTemperatureMap    []gardenMap
		temperatureToHumidityMap []gardenMap
		humidityToLocationMap    []gardenMap
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// simple test case
		{name: "simple test case", args: args{
			seed:                     79,
			seedToSoilMapping:        seedToSoilMapping,
			soilToFertilizerMapping:  soilToFertilizerMapping,
			fertilizerToWaterMapping: fertilizerToWaterMapping,
			waterToLightMap:          waterToLightMap,
			lightToTemperatureMap:    lightToTempMap,
			temperatureToHumidityMap: tempToHumidityMap,
			humidityToLocationMap:    humidityToLocationMap,
		}, want: 82},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLowestLocationNumberForSeed(tt.args.seed, tt.args.seedToSoilMapping, tt.args.soilToFertilizerMapping, tt.args.fertilizerToWaterMapping, tt.args.waterToLightMap, tt.args.lightToTemperatureMap, tt.args.temperatureToHumidityMap, tt.args.humidityToLocationMap); got != tt.want {
				t.Errorf("findLowestLocationNumberForSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
