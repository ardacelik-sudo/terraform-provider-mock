package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Config struct {
	SimulateDelayMs int
	FailureRate     float64
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"simulate_delay_ms": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Simulated artificial network latency in milliseconds",
			},
			"failure_rate": {
				Type:        schema.TypeFloat,
				Optional:    true,
				Default:     0.0,
				Description: "Simulated probability (0.0 to 1.0) of API calls failing",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"mock_network":  resourceNetwork(),
			"mock_instance": resourceInstance(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := &Config{
		SimulateDelayMs: d.Get("simulate_delay_ms").(int),
		FailureRate:     d.Get("failure_rate").(float64),
	}
	return config, nil
}
