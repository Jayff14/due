package gate

import (
	"context"
	packets "github.com/dobyte/due/v2/packet"
	"github.com/dobyte/due/v2/session"
	"github.com/dobyte/due/v2/transport/drpc/internal/client"
	"github.com/dobyte/due/v2/transport/drpc/internal/codes"
	"github.com/dobyte/due/v2/transport/drpc/internal/packet"
	"sync/atomic"
)

type Client struct {
	seq              uint64
	client           *client.Client
	bindPacker       *packet.BindPacker
	unbindPacker     *packet.UnbindPacker
	getIPPacker      *packet.GetIPPacker
	statPacker       *packet.StatPacker
	disconnectPacker *packet.DisconnectPacker
	pushPacker       *packet.PushPacker
}

// Bind 绑定用户与连接
func (c *Client) Bind(ctx context.Context, cid, uid int64) (bool, error) {
	seq := atomic.AddUint64(&c.seq, 1)

	buf, err := c.bindPacker.PackReq(seq, cid, uid)
	if err != nil {
		return false, err
	}

	data, err := c.client.Call(ctx, seq, buf)
	if err != nil {
		return false, err
	}

	code, err := c.bindPacker.UnpackRes(data)
	if err != nil {
		return false, err
	}

	return code == codes.NotFoundSession, nil
}

// Unbind 解绑用户与连接
func (c *Client) Unbind(ctx context.Context, uid int64) (bool, error) {
	seq := atomic.AddUint64(&c.seq, 1)

	buf, err := c.unbindPacker.PackReq(seq, uid)
	if err != nil {
		return false, err
	}

	data, err := c.client.Call(ctx, seq, buf)
	if err != nil {
		return false, err
	}

	code, err := c.unbindPacker.UnpackRes(data)
	if err != nil {
		return false, err
	}

	return code == codes.NotFoundSession, nil
}

// GetIP 获取客户端IP
func (c *Client) GetIP(ctx context.Context, kind session.Kind, target int64) (string, bool, error) {
	seq := atomic.AddUint64(&c.seq, 1)

	buf, err := c.getIPPacker.PackReq(seq, kind, target)
	if err != nil {
		return "", false, err
	}

	data, err := c.client.Call(ctx, seq, buf)
	if err != nil {
		return "", false, err
	}

	code, ip, err := c.getIPPacker.UnpackRes(data)
	if err != nil {
		return "", false, err
	}

	return ip, code == codes.NotFoundSession, nil
}

// Stat 推送广播消息
func (c *Client) Stat(ctx context.Context, kind session.Kind) (int64, error) {
	seq := atomic.AddUint64(&c.seq, 1)

	buf, err := c.statPacker.PackReq(seq, kind)
	if err != nil {
		return 0, err
	}

	data, err := c.client.Call(ctx, seq, buf)
	if err != nil {
		return 0, err
	}

	return c.statPacker.UnpackRes(data)
}

// Disconnect 断开连接
func (c *Client) Disconnect(ctx context.Context, kind session.Kind, target int64, isForce bool) (bool, error) {
	seq := atomic.AddUint64(&c.seq, 1)

	buf, err := c.disconnectPacker.PackReq(seq, kind, target, isForce)
	if err != nil {
		return false, err
	}

	data, err := c.client.Call(ctx, seq, buf)
	if err != nil {
		return false, err
	}

	code, err := c.disconnectPacker.UnpackRes(data)
	if err != nil {
		return false, err
	}

	return code == codes.NotFoundSession, nil
}

// Push 推送消息
func (c *Client) Push(ctx context.Context, kind session.Kind, target int64, message *packets.Message) (bool, error) {
	seq := atomic.AddUint64(&c.seq, 1)

	buf, err := c.pushPacker.PackReq(seq, kind, target, message)
	if err != nil {
		return false, err
	}

}

// AsyncPush 异步推送消息
func (c *Client) AsyncPush(ctx context.Context, kind session.Kind, target int64, message *packets.Message) error {

}

// Multicast 推送组播消息
func (c *Client) Multicast(ctx context.Context, kind session.Kind, targets []int64, message *packets.Message) (total int64, err error) {

}

// AsyncMulticast 推送组播消息
func (c *Client) AsyncMulticast(ctx context.Context, kind session.Kind, targets []int64, message *packets.Message) error {

}

// Broadcast 推送广播消息
func (c *Client) Broadcast(ctx context.Context, kind session.Kind, message *packets.Message) (total int64, err error) {

}

// AsyncBroadcast 推送广播消息
func (c *Client) AsyncBroadcast(ctx context.Context, kind session.Kind, message *packets.Message) error {

}