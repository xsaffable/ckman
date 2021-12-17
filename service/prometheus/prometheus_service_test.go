package prometheus

//import (
//	"fmt"
//	"github.com/housepower/ckman/model"
//	"testing"
//)
//
//func TestPrometheusService_QueryMetric(t *testing.T) {
//	// 配置, 应该从 cluster.jso 配置文件读取
//	promHost := "127.0.0.1"
//	promPort := 9090
//	var params model.MetricQueryReq
//	params.Metric = "xx"
//	params.Time = 12
//
//	promService := NewPrometheusService(fmt.Sprintf("%s:%d", promHost, promPort), 10)
//
//	value, err := promService.QueryRangeMetric(&params)
//}
