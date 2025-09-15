/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package ctxdata

import "context"

func GetUId(ctx context.Context) string {
	if u, ok := ctx.Value(Identify).(string); ok {
		return u
	}
	return ""
}

func IsAdmin(id string) bool {
	return id == "100000000000"
}
