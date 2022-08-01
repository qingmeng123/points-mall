/*******
* @Author:qingmeng
* @Description:
* @File:interceptor
* @Date:2022/7/30
 */

package auth

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func UnaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// pre-processing
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...) // invoking RPC method
	// post-processing
	end := time.Now()
	log.Printf("RPC: %s, req:%v start time: %s, end time: %s, err: %v", method, req, start.Format(time.RFC3339), end.Format(time.RFC3339), err)
	return err
}