package main

import (
	"flag"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	AppVersion = "0.0.3"
)

var (
	argProject  = flag.String("project", "", "Specify a Project ID.")
	argZone     = flag.String("zone", "us-west1-b", "Specify a Zone Name.")
	argInstance = flag.String("instance", "", "Specify the instance ID.")
	argStart    = flag.Bool("start", false, "Start instances.")
	argStop     = flag.Bool("stop", false, "Stop instances.")
	argBatch    = flag.Bool("batch", false, "Enable batch mode.")
	argVersion  = flag.Bool("version", false, "Print Version.")
)

func lastElements(str string, sep string) string {
	splited := strings.Split(str, sep)
	element := splited[len(splited)-1]
	return element
}

func printInstances(instanceList *compute.InstanceList) {
	allInstances := [][]string{}
	for _, v := range instanceList.Items {
		var ipList []string
		for _, n := range v.NetworkInterfaces {
			if n.NetworkIP == "" {
				ipList = append(ipList, "N/A")
			} else {
				ipList = append(ipList, n.NetworkIP)
			}
		}
		instance := []string{
			strconv.FormatUint(v.Id, 10),
			v.Name,
			lastElements(v.MachineType, "/"),
			strings.Join(ipList, ","),
			v.Status,
		}
		allInstances = append(allInstances, instance)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "MachineType", "IP Address", "Status"})
	for _, value := range allInstances {
		table.Append(value)
	}
	table.Render()
}

func listInstances(client *http.Client, project_id string, zone string) (*compute.InstanceList, error) {
	service, err := compute.New(client)
	if err != nil {
		return nil, err
	}
	t := compute.NewInstancesService(service)
	list, err := t.List(project_id, zone).Do()
	if err != nil {
		return nil, err
	}
	return list, err
}

func stopInstance(client *http.Client, project_id string, zone string, instance string) bool {
	service, err := compute.New(client)
	if err != nil {
		return false
	}
	t := compute.NewInstancesService(service)
	fmt.Print("Do you want to stop an instance? (y/n): ")
	var stdin string
	fmt.Scan(&stdin)
	switch stdin {
	case "y", "Y":
		_, err = t.Stop(project_id, zone, instance).Do()
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	case "n", "N":
		return false
	default:
		return false
	}
}

func startInstance(client *http.Client, project_id string, zone string, instance string) bool {
	service, err := compute.New(client)
	if err != nil {
		return false
	}
	t := compute.NewInstancesService(service)
	fmt.Print("Do you want to start an instance? (y/n): ")
	var stdin string
	fmt.Scan(&stdin)
	switch stdin {
	case "y", "Y":
		_, err = t.Start(project_id, zone, instance).Do()
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	case "n", "N":
		return false
	default:
		return false
	}
}

func main() {
	flag.Parse()
	if *argVersion {
		fmt.Println(AppVersion)
		os.Exit(0)
	}

	var project string
	if *argProject == "" && os.Getenv("GCP_PROJECT") == "" {
		fmt.Println("Please set `project` option or environment variable `GCP_PROJECT`")
		os.Exit(1)
	} else {
		if os.Getenv("GCP_PROJECT") != "" {
			project = os.Getenv("GCP_PROJECT")
		} else {
			project = *argProject
		}
	}

	var zone string
	zone = *argZone

	ctx := context.Background()
	client, err := google.DefaultClient(ctx, compute.ComputeScope)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if *argInstance != "" {
		if *argStart {
			result := startInstance(client, project, zone, *argInstance)
			fmt.Println("Instance Starting...")
			if result == false {
				fmt.Println("Interrupted.")
				os.Exit(1)
			}
			fmt.Println("Instance Started.")
		} else if *argStop {
			result := stopInstance(client, project, zone, *argInstance)
			fmt.Println("Instance Stopping...")
			if result == false {
				fmt.Println("Interrupted.")
				os.Exit(1)
			}
			fmt.Println("Instance Stopped.")
		} else {
			fmt.Println("Please set `start` or `stop` option.")
			os.Exit(1)
		}
	}
	list, err := listInstances(client, project, zone)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	printInstances(list)
	os.Exit(0)
}
