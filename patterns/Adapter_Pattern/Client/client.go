package client

import (
	ad "adapter/Adapter"
)

type Client struct{}

func (c *Client) ChargeMobile(mob ad.Mobile) {
	mob.ChargeAppleMobile()
}
