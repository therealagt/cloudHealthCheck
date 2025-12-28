/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/spf13/cobra"
	"google.golang.org/api/iterator"
)

// instancesCmd represents the instances command
var instancesCmd = &cobra.Command{
	Use:   "instances",
	Short: "List all instances in the specified GCP project and zone",
	Long:  `List all instances in the specified GCP project and zone.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectID, _ := cmd.Flags().GetString("project")
		if projectID == "" {
			log.Fatal("Project ID is required; use --project flag to specify it")
		}

		listInstances(projectID)
	},
}

func init() {
	rootCmd.AddCommand(instancesCmd)

	// Define flags for the instances command
	instancesCmd.Flags().StringP("project", "p", "", "GCP Project ID (required)")
	instancesCmd.MarkFlagRequired("project")
}

func listInstances(projectID string) {
	ctx := context.Background()

	// Create the Instances client
	c, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create instances client: %v", err)
	}
	defer c.Close()

	req := &computepb.AggregatedListInstancesRequest{
		Project: projectID,
	}

	fmt.Printf("Listing instances for project: %s\n", projectID)

	// Über alle Zonen iterieren (Aggregated List)
	it := c.AggregatedList(ctx, req)
	for {
		pair, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error listing instances: %v", err)
		}

		// pair.Value ist eine Liste von Instanzen in einer Zone
		for _, instance := range pair.Value.Instances {
			fmt.Printf("Name: %s | Zone: %s | Status: %s\n", *instance.Name, *instance.Zone, *instance.Status)
			if len(instance.Labels) > 0 {
				fmt.Println("  Labels:")
				for k, v := range instance.Labels {
					fmt.Printf("    - %s: %s\n", k, v)
				}
			}
		}
	}
}
