package iou_test

import (
	"encoding/json"
	"testing"

	"github.com/khlipeng/iou"
	"github.com/stretchr/testify/require"
)

func BenchmarkIoU(b *testing.B) {
	b.Run("Polygon", func(b *testing.B) {
		str1 := `
		[[419,374],[539,251],[657,251],[823,183],[941,181],[1058,216],[1185,273],[1185,425],[1050,592],[793,592],[751,650],[834,794],[1174,782],[1311,525],[1328,344],[1489,438],[1489,598],[1398,896],[1075,928],[1543,920],[1555,736],[1579,538],[1707,486],[1707,784],[1650,1056],[1403,1064],[799,1047],[515,1017],[441,754],[322,555]]
		`

		str2 := `
		[[552,277],[654,259],[827,186],[974,159],[1040,203],[1206,304],[1194,425],[1052,597],[798,602],[753,654],[818,792],[933,781],[1165,776],[1260,629],[1303,525],[1319,358],[1382,359],[1477,408],[1503,576],[1446,798],[1409,893],[1220,924],[1516,907],[1555,894],[1571,548],[1691,487],[1711,749],[1648,1044],[1554,1074],[1222,1071],[525,1026],[455,816],[366,669],[329,578],[325,501],[477,293]]
		`

		p1 := iou.Polygon{}
		err := json.Unmarshal([]byte(str1), &p1)
		require.NoError(b, err)
		p2 := iou.Polygon{}
		err = json.Unmarshal([]byte(str2), &p2)
		require.NoError(b, err)

		for n := 0; n < b.N; n++ {
			require.Equal(b, p1.IoU(p2) > 0.93, true)
		}

	})

	b.Run("Box 0.25", func(b *testing.B) {
		str1 := `
		[[0,0],[500,0],[500,500],[0,500]]
		`

		str2 := `
		[[0,0],[250,0],[250,250],[0,250]]
		`

		p1 := iou.Polygon{}
		err := json.Unmarshal([]byte(str1), &p1)
		require.NoError(b, err)
		p2 := iou.Polygon{}
		err = json.Unmarshal([]byte(str2), &p2)
		require.NoError(b, err)

		for n := 0; n < b.N; n++ {
			require.Equal(b, p1.IoU(p2) == 0.25, true)
		}

	})
	b.Run("Box 1", func(b *testing.B) {

		str1 := `
		[[0,0],[500,0],[500,500],[0,500]]
		`

		str2 := `
		[[0,0],[500,0],[500,500],[0,500]]
		`

		p1 := iou.Polygon{}
		err := json.Unmarshal([]byte(str1), &p1)
		require.NoError(b, err)
		p2 := iou.Polygon{}
		err = json.Unmarshal([]byte(str2), &p2)
		require.NoError(b, err)
		for n := 0; n < b.N; n++ {

			require.Equal(b, p1.IoU(p2) == 1, true)
		}
	})

}
