/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package wuid

import (
	"database/sql"
	"fmt"
	"sort"
	"strconv"

	"github.com/edwingeng/wuid/mysql/wuid"
)

var w *wuid.WUID

var powersOf10 = map[int]uint64{
	1:  10,
	2:  100,
	3:  1000,
	4:  10000,
	5:  100000,
	6:  1000000,
	7:  10000000,
	8:  100000000,
	9:  1000000000,
	10: 10000000000,
	11: 100000000000,
	12: 1000000000000,
	13: 10000000000000,
	14: 100000000000000,
	15: 1000000000000000,
	16: 10000000000000000,
	17: 100000000000000000,
	18: 1000000000000000000,
	19: 10000000000000000000,
}

func GenUGid(dsn string, digits int) string {
	if w == nil {
		Init(dsn)
	}

	if digits <= 0 || digits > 19 {
		return ""
	}

	id := w.Next()

	modulus, exists := powersOf10[digits]
	if !exists {
		return ""
	}

	digitID := uint64(id) % modulus

	return fmt.Sprintf("%d", digitID)
}

func Init(dsn string) {

	newDB := func() (*sql.DB, bool, error) {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return nil, false, err
		}
		return db, true, nil
	}

	w = wuid.NewWUID("default", nil)
	_ = w.LoadH28FromMysql(newDB, "wuid")
}

func GenUid(dsn string) string {
	if w == nil {
		Init(dsn)
	}

	return fmt.Sprintf("%v", w.Next())
}

func CombineId(aid, bid string) string {
	ids := []string{aid, bid}

	sort.Slice(ids, func(i, j int) bool {
		a, _ := strconv.ParseUint(ids[i], 0, 64)
		b, _ := strconv.ParseUint(ids[j], 0, 64)
		return a < b
	})

	return fmt.Sprintf("%s_%s", ids[0], ids[1])
}
