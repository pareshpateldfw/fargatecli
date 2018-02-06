package cmd

import (
	"errors"

	"github.com/jpignata/fargate/elbv2"
	"github.com/spf13/cobra"
)

const defaultTargetGroupFormat = "%s-default"

type lbOperation struct {
	elbv2 elbv2.Client
}

func (o lbOperation) findLb(lbName string) (elbv2.LoadBalancer, error) {
	output.Debug("Finding load balancer[API=elb2 Action=DescribeLoadBalancers]")
	loadBalancers, err := o.elbv2.DescribeLoadBalancersByName([]string{lbName})

	if err != nil {
		return elbv2.LoadBalancer{}, err
	}

	switch {
	case len(loadBalancers) == 0:
		return elbv2.LoadBalancer{}, ErrLbNotFound
	case len(loadBalancers) > 1:
		return elbv2.LoadBalancer{}, ErrLbTooManyFound
	}

	return loadBalancers[0], nil
}

var lbCmd = &cobra.Command{
	Use:   "lb",
	Short: "Manage load balancers",
	Long: `Manage load balancers

Load balancers distribute incoming traffic between the tasks within a service
for HTTP/HTTPS and TCP applications. HTTP/HTTPS load balancers can route to
multiple services based upon rules you specify when you create a new service.`,
}

var (
	ErrLbNotFound     = errors.New("Load balancer not found")
	ErrLbTooManyFound = errors.New("Too many load balancers found")
)

func init() {
	rootCmd.AddCommand(lbCmd)
}
