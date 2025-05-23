package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/hypervisors"
	"github.com/gophercloud/gophercloud/v2/testhelper"
	th "github.com/gophercloud/gophercloud/v2/testhelper"
	"github.com/gophercloud/gophercloud/v2/testhelper/client"
)

// HypervisorListBodyPre253 represents a raw hypervisor list from the Compute
// API with microversion older than 2.53.
// The first hypervisor represents what the specification says (~Newton)
// The second is exactly the same, but what you can get off a real system (~Kilo)
const HypervisorListBodyPre253 = `
{
    "hypervisors": [
        {
            "cpu_info": {
                "arch": "x86_64",
                "model": "Nehalem",
                "vendor": "Intel",
                "features": [
                    "pge",
                    "clflush"
                ],
                "topology": {
                    "cores": 1,
                    "threads": 1,
                    "sockets": 4
                }
            },
            "current_workload": 0,
            "status": "enabled",
            "state": "up",
            "disk_available_least": 0,
            "host_ip": "1.1.1.1",
            "free_disk_gb": 1028,
            "free_ram_mb": 7680,
            "hypervisor_hostname": "fake-mini",
            "hypervisor_type": "fake",
            "hypervisor_version": 2002000,
            "id": 1,
            "local_gb": 1028,
            "local_gb_used": 0,
            "memory_mb": 8192,
            "memory_mb_used": 512,
            "running_vms": 0,
            "service": {
                "host": "e6a37ee802d74863ab8b91ade8f12a67",
                "id": 2,
                "disabled_reason": null
            },
            "vcpus": 1,
            "vcpus_used": 0
        },
        {
            "cpu_info": "{\"arch\": \"x86_64\", \"model\": \"Nehalem\", \"vendor\": \"Intel\", \"features\": [\"pge\", \"clflush\"], \"topology\": {\"cores\": 1, \"threads\": 1, \"sockets\": 4}}",
            "current_workload": 0,
            "status": "enabled",
            "state": "up",
            "disk_available_least": 0,
            "host_ip": "1.1.1.1",
            "free_disk_gb": 1028,
            "free_ram_mb": 7680,
            "hypervisor_hostname": "fake-mini",
            "hypervisor_type": "fake",
            "hypervisor_version": 2.002e+06,
            "id": 1,
            "local_gb": 1028,
            "local_gb_used": 0,
            "memory_mb": 8192,
            "memory_mb_used": 512,
            "running_vms": 0,
            "service": {
                "host": "e6a37ee802d74863ab8b91ade8f12a67",
                "id": 2,
                "disabled_reason": null
            },
            "vcpus": 1,
            "vcpus_used": 0
        }
    ]
}`

// HypervisorListBodyPage1 represents page 1 of a raw hypervisor list result with Pike+ release.
const HypervisorListBodyPage1 = `
{
    "hypervisors": [
        {
            "cpu_info": {
                "arch": "x86_64",
                "model": "Nehalem",
                "vendor": "Intel",
                "features": [
                    "pge",
                    "clflush"
                ],
                "topology": {
                    "cores": 1,
                    "threads": 1,
                    "sockets": 4
                }
            },
            "current_workload": 0,
            "status": "enabled",
            "state": "up",
            "disk_available_least": 0,
            "host_ip": "1.1.1.1",
            "free_disk_gb": 1028,
            "free_ram_mb": 7680,
            "hypervisor_hostname": "fake-mini",
            "hypervisor_type": "fake",
            "hypervisor_version": 2002000,
            "id": "c48f6247-abe4-4a24-824e-ea39e108874f",
            "local_gb": 1028,
            "local_gb_used": 0,
            "memory_mb": 8192,
            "memory_mb_used": 512,
            "running_vms": 0,
            "service": {
                "host": "e6a37ee802d74863ab8b91ade8f12a67",
                "id": "9c2566e7-7a54-4777-a1ae-c2662f0c407c",
                "disabled_reason": null
            },
            "vcpus": 1,
            "vcpus_used": 0
        }
    ],
    "hypervisors_links": [
        {
            "href": "%s/os-hypervisors/detail?marker=c48f6247-abe4-4a24-824e-ea39e108874f",
            "rel": "next"
        }
    ]
}`

// HypervisorListBodyPage2 represents page 2 of a raw hypervisor list result with Pike+ release.
const HypervisorListBodyPage2 = `
{
    "hypervisors": [
        {
            "cpu_info": "{\"arch\": \"x86_64\", \"model\": \"Nehalem\", \"vendor\": \"Intel\", \"features\": [\"pge\", \"clflush\"], \"topology\": {\"cores\": 1, \"threads\": 1, \"sockets\": 4}}",
            "current_workload": 0,
            "status": "enabled",
            "state": "up",
            "disk_available_least": 0,
            "host_ip": "1.1.1.1",
            "free_disk_gb": 1028,
            "free_ram_mb": 7680,
            "hypervisor_hostname": "fake-mini",
            "hypervisor_type": "fake",
            "hypervisor_version": 2.002e+06,
            "id": "c48f6247-abe4-4a24-824e-ea39e108874f",
            "local_gb": 1028,
            "local_gb_used": 0,
            "memory_mb": 8192,
            "memory_mb_used": 512,
            "running_vms": 0,
            "service": {
                "host": "e6a37ee802d74863ab8b91ade8f12a67",
                "id": "9c2566e7-7a54-4777-a1ae-c2662f0c407c",
                "disabled_reason": null
            },
            "vcpus": 1,
            "vcpus_used": 0
        }
    ]
}`

// HypervisorListBodyEmpty represents an empty raw hypervisor list result, marking the end of pagination.
const HypervisorListBodyEmpty = `{ "hypervisors": [] }`

// HypervisorListWithParametersBody represents a raw hypervisor list result with Pike+ release.
const HypervisorListWithParametersBody = `
{
    "hypervisors": [
        {
            "cpu_info": {
                "arch": "x86_64",
                "model": "Nehalem",
                "vendor": "Intel",
                "features": [
                    "pge",
                    "clflush"
                ],
                "topology": {
                    "cores": 1,
                    "threads": 1,
                    "sockets": 4
                }
            },
            "current_workload": 0,
            "status": "enabled",
            "state": "up",
            "disk_available_least": 0,
            "host_ip": "1.1.1.1",
            "free_disk_gb": 1028,
            "free_ram_mb": 7680,
            "hypervisor_hostname": "fake-mini",
            "hypervisor_type": "fake",
            "hypervisor_version": 2002000,
            "id": "c48f6247-abe4-4a24-824e-ea39e108874f",
            "local_gb": 1028,
            "local_gb_used": 0,
            "memory_mb": 8192,
            "memory_mb_used": 512,
            "running_vms": 0,
            "service": {
                "host": "e6a37ee802d74863ab8b91ade8f12a67",
                "id": "9c2566e7-7a54-4777-a1ae-c2662f0c407c",
                "disabled_reason": null
            },
            "servers": [
                {
                    "name": "instance-00000001",
                    "uuid": "c42acc8d-eab3-4e4d-9d90-01b0791328f4"
                },
                {
                    "name": "instance-00000002",
                    "uuid": "8aaf2941-b774-41fc-921b-20c4757cc359"
                }
            ],
            "vcpus": 1,
            "vcpus_used": 0
        },
        {
            "cpu_info": "{\"arch\": \"x86_64\", \"model\": \"Nehalem\", \"vendor\": \"Intel\", \"features\": [\"pge\", \"clflush\"], \"topology\": {\"cores\": 1, \"threads\": 1, \"sockets\": 4}}",
            "current_workload": 0,
            "status": "enabled",
            "state": "up",
            "disk_available_least": 0,
            "host_ip": "1.1.1.1",
            "free_disk_gb": 1028,
            "free_ram_mb": 7680,
            "hypervisor_hostname": "fake-mini",
            "hypervisor_type": "fake",
            "hypervisor_version": 2.002e+06,
            "id": "c48f6247-abe4-4a24-824e-ea39e108874f",
            "local_gb": 1028,
            "local_gb_used": 0,
            "memory_mb": 8192,
            "memory_mb_used": 512,
            "running_vms": 0,
            "servers": [
                {
                    "name": "instance-00000001",
                    "uuid": "c42acc8d-eab3-4e4d-9d90-01b0791328f4"
                },
                {
                    "name": "instance-00000002",
                    "uuid": "8aaf2941-b774-41fc-921b-20c4757cc359"
                }
            ],
            "service": {
                "host": "e6a37ee802d74863ab8b91ade8f12a67",
                "id": "9c2566e7-7a54-4777-a1ae-c2662f0c407c",
                "disabled_reason": null
            },
            "vcpus": 1,
            "vcpus_used": 0
        }
    ]
}`

const HypervisorsStatisticsBody = `
{
    "hypervisor_statistics": {
        "count": 1,
        "current_workload": 0,
        "disk_available_least": 0,
        "free_disk_gb": 1028,
        "free_ram_mb": 7680,
        "local_gb": 1028,
        "local_gb_used": 0,
        "memory_mb": 8192,
        "memory_mb_used": 512,
        "running_vms": 0,
        "vcpus": 2,
        "vcpus_used": 0
    }
}
`

// HypervisorGetBody represents a raw hypervisor GET result with Pike+ release.
const HypervisorGetBody = `
{
    "hypervisor":{
        "cpu_info":{
            "arch":"x86_64",
            "model":"Nehalem",
            "vendor":"Intel",
            "features":[
                "pge",
                "clflush"
            ],
            "topology":{
                "cores":1,
                "threads":1,
                "sockets":4
            }
        },
        "current_workload":0,
        "status":"enabled",
        "state":"up",
        "disk_available_least":0,
        "host_ip":"1.1.1.1",
        "free_disk_gb":1028,
        "free_ram_mb":7680,
        "hypervisor_hostname":"fake-mini",
        "hypervisor_type":"fake",
        "hypervisor_version":2002000,
        "id":"c48f6247-abe4-4a24-824e-ea39e108874f",
        "local_gb":1028,
        "local_gb_used":0,
        "memory_mb":8192,
        "memory_mb_used":512,
        "running_vms":0,
        "service":{
            "host":"e6a37ee802d74863ab8b91ade8f12a67",
            "id":"9c2566e7-7a54-4777-a1ae-c2662f0c407c",
            "disabled_reason":null
        },
        "vcpus":1,
        "vcpus_used":0
    }
}
`

// HypervisorGetEmptyCPUInfoBody represents a raw hypervisor GET result with
// no cpu_info
const HypervisorGetEmptyCPUInfoBody = `
{
    "hypervisor":{
        "cpu_info": "",
        "current_workload":0,
        "status":"enabled",
        "state":"up",
        "disk_available_least":0,
        "host_ip":"1.1.1.1",
        "free_disk_gb":1028,
        "free_ram_mb":7680,
        "hypervisor_hostname":"fake-mini",
        "hypervisor_type":"fake",
        "hypervisor_version":2002000,
        "id":"c48f6247-abe4-4a24-824e-ea39e108874f",
        "local_gb":1028,
        "local_gb_used":0,
        "memory_mb":8192,
        "memory_mb_used":512,
        "running_vms":0,
        "service":{
            "host":"e6a37ee802d74863ab8b91ade8f12a67",
            "id":"9c2566e7-7a54-4777-a1ae-c2662f0c407c",
            "disabled_reason":null
        },
        "vcpus":1,
        "vcpus_used":0
    }
}
`

// HypervisorAfterV287ResponseBody represents a raw hypervisor GET result with
// missing cpu_info, free_disk_gb, local_gb as seen after v2.87
const HypervisorAfterV287ResponseBody = `
{
    "hypervisor":{
        "current_workload":0,
        "status":"enabled",
        "state":"up",
        "disk_available_least":0,
        "host_ip":"1.1.1.1",
        "free_ram_mb":7680,
        "hypervisor_hostname":"fake-mini",
        "hypervisor_type":"fake",
        "hypervisor_version":2002000,
        "id":"c48f6247-abe4-4a24-824e-ea39e108874f",
        "local_gb_used":0,
        "memory_mb":8192,
        "memory_mb_used":512,
        "running_vms":0,
        "service":{
            "host":"e6a37ee802d74863ab8b91ade8f12a67",
            "id":"9c2566e7-7a54-4777-a1ae-c2662f0c407c",
            "disabled_reason":null
        },
        "vcpus":1,
        "vcpus_used":0
    }
}
`

// HypervisorUptimeBody represents a raw hypervisor uptime request result with
// Pike+ release.
const HypervisorUptimeBody = `
{
    "hypervisor": {
        "hypervisor_hostname": "fake-mini",
        "id": "c48f6247-abe4-4a24-824e-ea39e108874f",
        "state": "up",
        "status": "enabled",
        "uptime": " 08:32:11 up 93 days, 18:25, 12 users,  load average: 0.20, 0.12, 0.14"
    }
}
`

var (
	HypervisorFakePre253 = hypervisors.Hypervisor{
		CPUInfo: hypervisors.CPUInfo{
			Arch:   "x86_64",
			Model:  "Nehalem",
			Vendor: "Intel",
			Features: []string{
				"pge",
				"clflush",
			},
			Topology: hypervisors.Topology{
				Cores:   1,
				Threads: 1,
				Sockets: 4,
			},
		},
		CurrentWorkload:    0,
		Status:             "enabled",
		State:              "up",
		DiskAvailableLeast: 0,
		HostIP:             "1.1.1.1",
		FreeDiskGB:         1028,
		FreeRamMB:          7680,
		HypervisorHostname: "fake-mini",
		HypervisorType:     "fake",
		HypervisorVersion:  2002000,
		ID:                 "1",
		LocalGB:            1028,
		LocalGBUsed:        0,
		MemoryMB:           8192,
		MemoryMBUsed:       512,
		RunningVMs:         0,
		Service: hypervisors.Service{
			Host:           "e6a37ee802d74863ab8b91ade8f12a67",
			ID:             "2",
			DisabledReason: "",
		},
		VCPUs:     1,
		VCPUsUsed: 0,
	}

	HypervisorFake = hypervisors.Hypervisor{
		CPUInfo: hypervisors.CPUInfo{
			Arch:   "x86_64",
			Model:  "Nehalem",
			Vendor: "Intel",
			Features: []string{
				"pge",
				"clflush",
			},
			Topology: hypervisors.Topology{
				Cores:   1,
				Threads: 1,
				Sockets: 4,
			},
		},
		CurrentWorkload:    0,
		Status:             "enabled",
		State:              "up",
		DiskAvailableLeast: 0,
		HostIP:             "1.1.1.1",
		FreeDiskGB:         1028,
		FreeRamMB:          7680,
		HypervisorHostname: "fake-mini",
		HypervisorType:     "fake",
		HypervisorVersion:  2002000,
		ID:                 "c48f6247-abe4-4a24-824e-ea39e108874f",
		LocalGB:            1028,
		LocalGBUsed:        0,
		MemoryMB:           8192,
		MemoryMBUsed:       512,
		RunningVMs:         0,
		Service: hypervisors.Service{
			Host:           "e6a37ee802d74863ab8b91ade8f12a67",
			ID:             "9c2566e7-7a54-4777-a1ae-c2662f0c407c",
			DisabledReason: "",
		},
		VCPUs:     1,
		VCPUsUsed: 0,
	}

	HypervisorFakeWithParameters = hypervisors.Hypervisor{
		CPUInfo: hypervisors.CPUInfo{
			Arch:   "x86_64",
			Model:  "Nehalem",
			Vendor: "Intel",
			Features: []string{
				"pge",
				"clflush",
			},
			Topology: hypervisors.Topology{
				Cores:   1,
				Threads: 1,
				Sockets: 4,
			},
		},
		CurrentWorkload:    0,
		Status:             "enabled",
		State:              "up",
		DiskAvailableLeast: 0,
		HostIP:             "1.1.1.1",
		FreeDiskGB:         1028,
		FreeRamMB:          7680,
		HypervisorHostname: "fake-mini",
		HypervisorType:     "fake",
		HypervisorVersion:  2002000,
		ID:                 "c48f6247-abe4-4a24-824e-ea39e108874f",
		LocalGB:            1028,
		LocalGBUsed:        0,
		MemoryMB:           8192,
		MemoryMBUsed:       512,
		RunningVMs:         0,
		Service: hypervisors.Service{
			Host:           "e6a37ee802d74863ab8b91ade8f12a67",
			ID:             "9c2566e7-7a54-4777-a1ae-c2662f0c407c",
			DisabledReason: "",
		},
		Servers: &[]hypervisors.Server{
			{
				Name: "instance-00000001",
				UUID: "c42acc8d-eab3-4e4d-9d90-01b0791328f4",
			},
			{
				Name: "instance-00000002",
				UUID: "8aaf2941-b774-41fc-921b-20c4757cc359",
			},
		},
		VCPUs:     1,
		VCPUsUsed: 0,
	}

	HypervisorEmptyCPUInfo = hypervisors.Hypervisor{
		CPUInfo:            hypervisors.CPUInfo{},
		CurrentWorkload:    0,
		Status:             "enabled",
		State:              "up",
		DiskAvailableLeast: 0,
		HostIP:             "1.1.1.1",
		FreeDiskGB:         1028,
		FreeRamMB:          7680,
		HypervisorHostname: "fake-mini",
		HypervisorType:     "fake",
		HypervisorVersion:  2002000,
		ID:                 "c48f6247-abe4-4a24-824e-ea39e108874f",
		LocalGB:            1028,
		LocalGBUsed:        0,
		MemoryMB:           8192,
		MemoryMBUsed:       512,
		RunningVMs:         0,
		Service: hypervisors.Service{
			Host:           "e6a37ee802d74863ab8b91ade8f12a67",
			ID:             "9c2566e7-7a54-4777-a1ae-c2662f0c407c",
			DisabledReason: "",
		},
		VCPUs:     1,
		VCPUsUsed: 0,
	}

	HypervisorAfterV287Response = hypervisors.Hypervisor{
		CPUInfo:            hypervisors.CPUInfo{},
		CurrentWorkload:    0,
		Status:             "enabled",
		State:              "up",
		DiskAvailableLeast: 0,
		HostIP:             "1.1.1.1",
		FreeDiskGB:         0,
		FreeRamMB:          7680,
		HypervisorHostname: "fake-mini",
		HypervisorType:     "fake",
		HypervisorVersion:  2002000,
		ID:                 "c48f6247-abe4-4a24-824e-ea39e108874f",
		LocalGB:            0,
		LocalGBUsed:        0,
		MemoryMB:           8192,
		MemoryMBUsed:       512,
		RunningVMs:         0,
		Service: hypervisors.Service{
			Host:           "e6a37ee802d74863ab8b91ade8f12a67",
			ID:             "9c2566e7-7a54-4777-a1ae-c2662f0c407c",
			DisabledReason: "",
		},
		VCPUs:     1,
		VCPUsUsed: 0,
	}

	HypervisorsStatisticsExpected = hypervisors.Statistics{
		Count:              1,
		CurrentWorkload:    0,
		DiskAvailableLeast: 0,
		FreeDiskGB:         1028,
		FreeRamMB:          7680,
		LocalGB:            1028,
		LocalGBUsed:        0,
		MemoryMB:           8192,
		MemoryMBUsed:       512,
		RunningVMs:         0,
		VCPUs:              2,
		VCPUsUsed:          0,
	}

	HypervisorUptimeExpected = hypervisors.Uptime{
		HypervisorHostname: "fake-mini",
		ID:                 "c48f6247-abe4-4a24-824e-ea39e108874f",
		State:              "up",
		Status:             "enabled",
		Uptime:             " 08:32:11 up 93 days, 18:25, 12 users,  load average: 0.20, 0.12, 0.14",
	}
)

func HandleHypervisorsStatisticsSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-hypervisors/statistics", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, HypervisorsStatisticsBody)
	})
}

func HandleHypervisorListPre253Successfully(t *testing.T) {
	th.Mux.HandleFunc("/os-hypervisors/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, HypervisorListBodyPre253)
	})
}

func HandleHypervisorListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-hypervisors/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		switch r.URL.Query().Get("marker") {
		case "":
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprintf(w, HypervisorListBodyPage1, testhelper.Server.URL)
		case "c48f6247-abe4-4a24-824e-ea39e108874f":
			w.Header().Add("Content-Type", "application/json")
			fmt.Fprint(w, HypervisorListBodyPage2)
		default:
			http.Error(w, "unexpected marker value", http.StatusInternalServerError)
		}
	})
}

func HandleHypervisorListWithParametersSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-hypervisors/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestFormValues(t, r, map[string]string{
			"with_servers": "true",
		})

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, HypervisorListWithParametersBody)
	})
}

func HandleHypervisorGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-hypervisors/"+HypervisorFake.ID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, HypervisorGetBody)
	})
}

func HandleHypervisorGetEmptyCPUInfoSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-hypervisors/"+HypervisorFake.ID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, HypervisorGetEmptyCPUInfoBody)
	})
}

func HandleHypervisorAfterV287ResponseSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-hypervisors/"+HypervisorFake.ID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, HypervisorAfterV287ResponseBody)
	})
}

func HandleHypervisorUptimeSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-hypervisors/"+HypervisorFake.ID+"/uptime", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, HypervisorUptimeBody)
	})
}
