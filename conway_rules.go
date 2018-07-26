package conway

func ConwayRule() Rule {
	return AliveMappingRule{
		Alive: map[int]CellValue{
			0: Dead,
			1: Dead,
			4: Dead,
			5: Dead,
			6: Dead,
			7: Dead,
			8: Dead,
			9: Dead,
		},
		Dead: map[int]CellValue{
			3: Alive,
		},
	}
}
